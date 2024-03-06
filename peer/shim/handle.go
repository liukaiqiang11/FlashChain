package shim

import (
	"fchain/common/signer"
	"fchain/config"
	pb "fchain/proto"
	"fchain/proto/util"
	"fmt"
	cmap "github.com/orcaman/concurrent-map/v2"
	"google.golang.org/protobuf/proto"
	"log"
	"sync"
	"sync/atomic"
)

type PeerChaincodeStream interface {
	Send(*pb.ChaincodeMessage) error
	Recv() (*pb.ChaincodeMessage, error)
}

// ClientStream supports the (original) chaincode-as-client interaction pattern
type ClientStream interface {
	PeerChaincodeStream
	CloseSend() error
}

type Handler struct {
	serialLock   sync.Mutex
	cc           Chaincode
	chatStream   PeerChaincodeStream
	readWriteSet cmap.ConcurrentMap[string, *pb.KVRWSet]
	signer       *signer.Signer
}

// serialSend 序列化发送
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

func newChaincodeHandler(peerChatStream PeerChaincodeStream) *Handler {

	sig, err := signer.NewSigner(config.PeerConf)
	if err != nil {
		log.Fatal(err)
	}
	return &Handler{
		chatStream:   peerChatStream,
		readWriteSet: cmap.New[*pb.KVRWSet](),
		signer:       sig,
	}
}

type stubHandlerFunc func(*pb.ChaincodeMessage) (*pb.ChaincodeMessage, error)

func (h *Handler) handleStubInteraction(handler stubHandlerFunc, msg *pb.ChaincodeMessage, errc chan<- error) {
	resp, err := handler(msg)
	if err != nil {
		resp = &pb.ChaincodeMessage{Type: pb.ChaincodeMessage_Type_ERROR, Payload: []byte(err.Error()), TxID: msg.TxID}
	}
	h.serialSendAsync(resp, errc)
}

// 查看读集或者写集钟是否有重复的元素
func hasDuplicate(rwSet *pb.KVRWSet) bool {

	for i := 0; i < len(rwSet.Reads); i++ {
		for j := i + 1; j < len(rwSet.Reads); j++ {
			if rwSet.Reads[i].Key == rwSet.Reads[j].Key {
				return true
			}
		}
	}
	for i := 0; i < len(rwSet.Writes); i++ {
		for j := i + 1; j < len(rwSet.Writes); j++ {
			if rwSet.Writes[i].Key == rwSet.Writes[j].Key {
				return true
			}
		}
	}
	return false
}

// detectionKVRwSet 检测事务的读写集，让不可序列化的事务提前中止
func detectionKVRwSet(kvRwSet *pb.KVRWSet, txID string) *pb.ChaincodeMessage {

	for _, write := range kvRwSet.Writes {
		if _, ok := detection.Get(write.Key); ok {
			err := fmt.Errorf("this is a Non-serializable transaction")
			return &pb.ChaincodeMessage{Type: pb.ChaincodeMessage_Type_ERROR, Payload: []byte(err.Error()), TxID: txID}
		}
	}

	for _, write := range kvRwSet.Writes {
		detection.Set(write.Key, struct{}{})
	}
	return nil
}

func (h *Handler) handleTransaction(msg *pb.ChaincodeMessage) (*pb.ChaincodeMessage, error) {

	input := &pb.ChaincodeInput{}
	err := proto.Unmarshal(msg.Payload, input)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal input: %s", err)
	}

	stub, err := newChaincodeStub(h, msg.TxID, input, msg.Proposal)

	if err != nil {
		return nil, fmt.Errorf("failed to create new ChaincodeStub: %s", err)
	}

	name, err := stub.getChainCodeName()
	if err != nil {
		return nil, fmt.Errorf("failed to get Chaincode Name: %s", err)
	}

	if name == "SmallBank" {
		h.cc = new(SmallBank)
	} else if name == "KvStore" {
		h.cc = new(KvStore)
	}

	//调用链码
	res := h.cc.Invoke(stub)
	if res.Status != 200 {
		err = fmt.Errorf("transaction invoke fail")
		return &pb.ChaincodeMessage{Type: pb.ChaincodeMessage_Type_ERROR, Payload: []byte(err.Error()), TxID: msg.TxID}, nil
	}

	t := &timer{startTime: msg.Timestamp}
	TxLatency.Set(msg.TxID, t)

	// 获取事务的读写集
	kvRwSet, ok := h.readWriteSet.Get(msg.TxID)
	if !ok {
		err = fmt.Errorf("obtain the transaction read/write set fail")
		return &pb.ChaincodeMessage{Type: pb.ChaincodeMessage_Type_ERROR, Payload: []byte(err.Error()), TxID: msg.TxID}, nil
	}

	// 检测事务读集或写集是否存在重复地址，如果存在则说明事务错误
	if hasDuplicate(kvRwSet) {
		err := fmt.Errorf("this is a Non-serializable transaction")
		return &pb.ChaincodeMessage{Type: pb.ChaincodeMessage_Type_ERROR, Payload: []byte(err.Error())}, nil
	}

	//检测该事务的读写集，判断它是否为一个不可序列化的事务，如果不可序列化则直接将该事务中止
	chaincodeMsg := detectionKVRwSet(kvRwSet, msg.TxID)
	if chaincodeMsg != nil {
		return chaincodeMsg, nil
	}

	//创建背书，将背书结果返回给客户端
	result := marshalOrPanic(kvRwSet)
	shdr := stub.proposal.GetHeader()
	payload := stub.proposal.GetPayload()
	pr, err := util.CreateProposalResponse(shdr, payload, &res, result, h.signer)
	if err != nil {
		err = fmt.Errorf("failed to CreateProposalResponse: %s", err)
		return &pb.ChaincodeMessage{Type: pb.ChaincodeMessage_Type_ERROR, Payload: []byte(err.Error()), TxID: msg.TxID}, nil
	}

	resBytes, err := proto.Marshal(pr)
	if err != nil {
		err = fmt.Errorf("failed to marshal ProposalResponse: %s", err)
		return &pb.ChaincodeMessage{Type: pb.ChaincodeMessage_Type_ERROR, Payload: []byte(err.Error()), TxID: msg.TxID}, nil
	}

	var msgType pb.ChaincodeMessage_Type

	if kvRwSet.Writes == nil {
		msgType = pb.ChaincodeMessage_Type_ONLY_READ_TRANSACTION
		atomic.AddInt32(&SuccessNum, 1)
	} else {
		msgType = pb.ChaincodeMessage_Type_COMPLETED
	}
	return &pb.ChaincodeMessage{Type: msgType, Payload: resBytes, Proposal: msg.Proposal, TxID: msg.TxID}, nil

}

// handleGetState 获取状态数据库中的数据
func (h *Handler) handleGetState(key string, txID string) ([]byte, error) {

	s, ok := State.Get(key)
	if !ok {
		s = &StateDB{}
	}

	kVRead := &pb.KVRead{Key: key, Version: s.Version}
	readWriteSet, ok := h.readWriteSet.Get(txID)
	if !ok {
		readWriteSet = &pb.KVRWSet{}
	}
	readWriteSet.Reads = append(readWriteSet.Reads, kVRead)
	readWriteSet.TxID = txID
	h.readWriteSet.Set(txID, readWriteSet)
	return s.Value, nil
}

// handlePutState 更改状态数据库中的数据
func (h *Handler) handlePutState(key string, value []byte, txID string) error {

	kVWrite := &pb.KVWrite{Key: key, Value: value}
	readWriteSet, ok := h.readWriteSet.Get(txID)
	if !ok {
		readWriteSet = &pb.KVRWSet{}
	}

	readWriteSet.Writes = append(readWriteSet.Writes, kVWrite)
	readWriteSet.TxID = txID
	h.readWriteSet.Set(txID, readWriteSet)

	return nil
}

// handleDelState 删除状态数据库中的数据
func (h *Handler) handleDelState(key string, txID string) error {

	kVWrite := &pb.KVWrite{Key: key, IsDelete: true}
	readWriteSet, _ := h.readWriteSet.Get(txID)
	if readWriteSet == nil {
		readWriteSet = &pb.KVRWSet{}
	}
	readWriteSet.Writes = append(readWriteSet.Writes, kVWrite)
	readWriteSet.TxID = txID
	h.readWriteSet.Set(txID, readWriteSet)

	return nil
}

// handleReady 处理已经准备好的消息
func (h *Handler) handleReady(msg *pb.ChaincodeMessage, errc chan error) error {
	switch msg.Type {

	case pb.ChaincodeMessage_Type_TRANSACTION:
		go h.handleStubInteraction(h.handleTransaction, msg, errc)
		return nil

	default:
		return fmt.Errorf("[%s] Chaincode h cannot core message (%s)", msg.TxID, msg.Type)
	}
}

// handleMessage 处理发送过来的消息
func (h *Handler) handleMessage(msg *pb.ChaincodeMessage, errc chan error) error {

	err := h.handleReady(msg, errc)

	if err != nil {
		payload := []byte(err.Error())
		errorMsg := &pb.ChaincodeMessage{Type: pb.ChaincodeMessage_Type_ERROR, Payload: payload, TxID: msg.TxID}
		return h.serialSend(errorMsg)
	}

	return nil
}

// marshalOrPanic 对proto消息进行编码
func marshalOrPanic(msg proto.Message) []byte {
	bytes, err := proto.Marshal(msg)
	if err != nil {
		panic(fmt.Sprintf("failed to marshal message: %s", err))
	}
	return bytes
}
