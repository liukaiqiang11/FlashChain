package order

import (
	"context"
	"fchain/peer/shim"
	"fchain/proto/util"
	"fmt"
	"log"
	"sync"
	"sync/atomic"
	"time"

	"fchain/common/signer"
	"fchain/config"
	"fchain/order/core"
	pb "fchain/proto"
)

type Handler struct {
	serialLock  sync.Mutex
	chatStream  OrderChaincodeStream
	peerClients []pb.ChaincodeSupportClient
	start       time.Time
	rwSet       map[string]core.ReadWriteTransaction
	envelopes   []*pb.Envelope
	envGroup    []*pb.Envelopes
	signer      *signer.Signer
	group       *core.TransactionGroup
	num         uint64
}

func NewOrderHandler(ServerStream OrderChaincodeStream) *Handler {

	conf := config.OrderConf
	sig, err := signer.NewSigner(conf)
	if err != nil {
		log.Fatal(err)
	}

	clients, err := shim.GetPeerClient()
	if err != nil {
		log.Fatal(err)
	}

	return &Handler{
		chatStream:  ServerStream,
		peerClients: clients,
		rwSet:       make(map[string]core.ReadWriteTransaction),
		envelopes:   []*pb.Envelope{},
		envGroup:    []*pb.Envelopes{},
		signer:      sig,
		group:       core.NewGroup(),
		num:         0,
	}
}

func (h *Handler) serialSend(msg *pb.ChaincodeMessage) error {
	h.serialLock.Lock()
	defer h.serialLock.Unlock()
	return h.chatStream.Send(msg)
}

func (h *Handler) serialSendAsync(msg *pb.ChaincodeMessage, errc chan<- error) {
	go func() {
		errc <- h.serialSend(msg)
	}()
}

func (h *Handler) initBLock() {
	h.group = core.NewGroup()
	h.num = 0
	h.start = time.Now()
	h.rwSet = make(map[string]core.ReadWriteTransaction)
	h.envelopes = nil
	h.envGroup = nil
}

func (h *Handler) HandleMessage(msg *pb.ChaincodeMessage, errc chan error) error {
	if msg.Type == pb.ChaincodeMessage_Type_TRANSACTION {

		envelope, err := util.UnmarshalEnvelope(msg.Payload)
		if err != nil {
			log.Println(err)
		}

		kvRwSet, err := util.GetKVRWSetKwFromEnvelope(envelope)
		if err != nil {
			log.Println(err)
		}

		//拷贝交易的分组信息
		cloneGroup := core.DeepCopyGroup(h.group)

		// 分析事务直接的依赖关系
		deps := core.FindTransactionDependency(kvRwSet, h.rwSet)

		var isAbort bool
		if deps != nil {
			h.group.GroupTransactionsByDeps(deps)
			isAbort = core.TxReorder(kvRwSet.TxID, h.group)
		}

		// 如果该事务不可序列化，则把该事务删除
		if isAbort {
			err = fmt.Errorf("transaction reorder fail")
			resp := &pb.ChaincodeMessage{Type: pb.ChaincodeMessage_Type_ERROR, Payload: []byte(err.Error()), TxID: kvRwSet.TxID}
			h.serialSendAsync(resp, errc)
			for _, read := range kvRwSet.Reads {
				txs, _ := h.rwSet[read.Key]
				txs.ReadTransactions = core.DeleteTx(txs.ReadTransactions, kvRwSet.TxID)
				h.rwSet[read.Key] = txs
			}
			for _, write := range kvRwSet.Writes {
				txs, _ := h.rwSet[write.Key]
				txs.WriteTransactions = core.DeleteTx(txs.WriteTransactions, kvRwSet.TxID)
				h.rwSet[write.Key] = txs
			}
			h.group = cloneGroup
		} else {
			atomic.AddUint64(&h.num, 1)
			h.envelopes = append(h.envelopes, envelope)
			resp := &pb.ChaincodeMessage{Type: pb.ChaincodeMessage_Type_COMPLETED, TxID: kvRwSet.TxID}
			h.serialSendAsync(resp, errc)
		}
	}

	// 判断是否满足成块条件
	if h.num == config.BlockSize || (int(time.Since(h.start)) > config.CreateBlockTime*1e6 && h.start.UnixMilli() > 0) || msg.Type == pb.ChaincodeMessage_Type_COMPLETED {

		//将写写冲突的事务删除
		for _, rwTransaction := range h.rwSet {
			for i, wt := range rwTransaction.WriteTransactions {
				if i > 0 {
					_, h.envelopes = core.GetEnvelope(wt, h.envelopes)
					err := fmt.Errorf("this is a transaction that writes conflict")
					resp := &pb.ChaincodeMessage{Type: pb.ChaincodeMessage_Type_ERROR, Payload: []byte(err.Error()), TxID: wt}
					h.serialSendAsync(resp, errc)
				}
			}
		}

		h.envGroup = core.CreateEnvelopeGroup(h.group, h.envelopes)
		//h.envGroup = append(envGroup, &pb.Envelopes{Envelope: h.envelopes})

		info, err := h.peerClients[0].GetBlockInfo(context.Background(), &pb.Empty{})
		if err != nil {
			log.Fatal(err)
		}

		newBlock := util.CreateNewBlock(info.Height, info.CurrentBlockHash, h.envGroup)

		blockHeaderBytes := util.BlockHeaderBytes(newBlock.Header)

		sig, err := h.signer.Sign(blockHeaderBytes)
		if err != nil {
			log.Fatal(err)
		}
		shdr, err := util.CreateSignatureHeader(h.signer)
		if err != nil {
			log.Fatal(err)
		}

		block := util.NewBlockSign(newBlock, shdr, sig)

		blockBytes := util.BlockBytes(block)
		var blockMsg *pb.ValidateMessage

		if msg.Type == pb.ChaincodeMessage_Type_COMPLETED && h.num != 0 {
			blockMsg = &pb.ValidateMessage{Type: pb.ValidateMessage_Type_VALIDATE_COMPLETED, Payload: blockBytes}
		} else if h.num == 0 {
			blockMsg = &pb.ValidateMessage{Type: pb.ValidateMessage_Type_VALIDATE_COMPLETED}
		} else {
			blockMsg = &pb.ValidateMessage{Type: pb.ValidateMessage_Type_VALIDATE, Payload: blockBytes}
		}

		var wg sync.WaitGroup

		for _, peerClient := range h.peerClients {
			wg.Add(1)
			go func(client pb.ChaincodeSupportClient) {
				defer wg.Done()
				_, err = client.ValidateBlock(context.Background(), blockMsg)
				if err != nil {
					log.Fatal(err)
				}
			}(peerClient)
		}
		wg.Wait()

		h.initBLock()
	}
	return nil
}
