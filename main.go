package main

import (
	"context"
	crand "crypto/rand"
	"fmt"
	"io"
	"math/rand"
	"sync"
	"sync/atomic"
	"time"

	"fchain/common/signer"
	"fchain/config"
	"fchain/order"
	"fchain/peer"
	"fchain/peer/shim"
	pb "fchain/proto"
	"fchain/proto/util"

	"github.com/chinuy/zipf"
	log "github.com/corgi-kx/logcustom"
	cmap "github.com/orcaman/concurrent-map/v2"
	"google.golang.org/protobuf/proto"
)

// 随机数种子
var source = rand.New(rand.NewSource(time.Now().Unix()))

// 用于记录接收的事务数量
var num int32

// 信封通道，客户端生成的事务全都放入该通道，然后转发给order节点
var envelopeChan = make(chan *pb.Envelope, config.TxNum)

var proposalResponse = cmap.New[[]*pb.ProposalResponse]()

// getRandomArgs 生成随机的交易提案参数
func getRandomArgs() [][]byte {
	var args [][]byte

	if config.ContractName == "SmallBank" {
		selectFunc := []string{"almagate", "updateBalance", "updateSaving", "sendPayment", "writeCheck"}
		//控制事务的偏斜度
		z := zipf.NewZipf(source, config.Skewness, config.AddrNum)
		randNum := source.Intn(5)
		//控制事务的读写比例
		if source.Float64() <= config.Ratio {
			args = [][]byte{
				[]byte("getBalance"),
				[]byte(fmt.Sprintf("%d", z.Uint64())),
			}
		} else {
			switch randNum {
			case 0:
				args = [][]byte{
					[]byte(selectFunc[randNum]),
					[]byte(fmt.Sprintf("%d", z.Uint64())),
					[]byte(fmt.Sprintf("%d", z.Uint64())),
				}
			case 1, 2, 4:
				args = [][]byte{
					[]byte(selectFunc[randNum]),
					[]byte(fmt.Sprintf("%d", z.Uint64())),
					[]byte("10"),
				}
			case 3:
				args = [][]byte{
					[]byte(selectFunc[randNum]),
					[]byte(fmt.Sprintf("%d", z.Uint64())),
					[]byte(fmt.Sprintf("%d", z.Uint64())),
					[]byte("10"),
				}
			default:
				log.Debug("错误！")
			}
		}
		// 用于测试事务的原子性，如果遍历状态数据库，发现转账交易前后状态数据库的总余额不便，则可以证明该框架满足原子性和一致性。
		//from := source.Intn(int(config.AddrNum))
		//to := source.Intn(int(config.AddrNum))
		//args = [][]byte{
		//	[]byte("sendPayment"),
		//	[]byte(fmt.Sprintf("%d", from)),
		//	[]byte(fmt.Sprintf("%d", to)),
		//	[]byte("10"),
		//}
	} else if config.ContractName == "KvStore" {
		//控制事务的偏斜度
		z := zipf.NewZipf(source, config.Skewness, config.AddrNum)

		if source.Float64() <= config.Ratio {
			args = [][]byte{
				[]byte("read"),
				[]byte(fmt.Sprintf("%d", z.Uint64())),
			}
		} else {
			byteArray := make([]byte, 100) // 创建一个长度为 100 的字节数组

			_, err := crand.Read(byteArray) // 从随机源填充字节数组
			if err != nil {
				fmt.Println("Error generating random bytes:", err)
			}
			args = [][]byte{
				[]byte("write"),
				[]byte(fmt.Sprintf("%d", z.Uint64())),
				byteArray,
			}
		}
	}
	return args
}

// getClientMsg 生成客户端发送给Peer节点的消息
func getClientMsg(creator []byte, sign *signer.Signer) (msg *pb.ChaincodeMessage, err error) {
	args := getRandomArgs()
	//获取链码的调用信息
	in := &pb.ChaincodeInput{Args: args}
	inData, err := proto.Marshal(in)
	if err != nil {
		return nil, err
	}
	chaincodeID := &pb.ChaincodeID{Version: "v1", Path: ""}
	if config.ContractName == "SmallBank" {
		chaincodeID.Name = "SmallBank"
	} else if config.ContractName == "KvStore" {
		chaincodeID.Name = "KvStore"
	}
	chaincodeSpec := &pb.ChaincodeSpec{Input: in, ChaincodeID: chaincodeID}

	proposal, txID, err := util.CreateChaincodeProposal(chaincodeSpec, creator)
	if err != nil {
		return nil, err
	}

	signedProp, err := util.GetSignedProposal(proposal, sign)
	if err != nil {
		return nil, err
	}

	return &pb.ChaincodeMessage{Type: pb.ChaincodeMessage_Type_TRANSACTION, Timestamp: uint64(time.Now().UnixMilli()),
		Payload: inData, Proposal: signedProp, TxID: txID}, nil
}

// collectEndorsementsAndCreateTransaction 接收Peer节点返回的背书结果生成事务，并将事务打包进信封。
func collectEndorsementsAndCreateTransaction(clientStream shim.ClientStream, sign *signer.Signer) {

	for {
		// 接收Peer节点返回的背书结果
		msg, err := clientStream.Recv()
		if err == io.EOF {
			break
		}

		if msg == nil {
			break
		}

		atomic.AddInt32(&num, 1)

		if msg.Type == pb.ChaincodeMessage_Type_ERROR {
			//log.Debugf("transaction execute fail.txID is: %s,and the reason for the error is: %s", msg.TxID, msg.Payload)
		} else {
			//log.Infof("transaction execute success.txID is: %s", msg.TxID)
			pr, err := util.UnmarshalProposalResponse(msg.Payload)
			if err != nil {
				log.Warnf("UnmarshalProposalResponse error", err)
			}

			var prs []*pb.ProposalResponse

			if _, ok := proposalResponse.Get(msg.TxID); !ok {
				proposalResponse.Set(msg.TxID, prs)
			}
			prs, _ = proposalResponse.Get(msg.TxID)
			prs = append(prs, pr)
			proposalResponse.Set(msg.TxID, prs)

			if len(prs) == len(config.Organizations) {
				//解析提案请求
				proposal, err := util.UnmarshalProposal(msg.Proposal.ProposalBytes)
				if err != nil {
					log.Debug(err)
				}
				//将提案请求和背书结果一起打包成信封
				envelope, err := util.CreateSignedTx(msg.TxID, proposal, sign, prs...)
				if err != nil {
					log.Warnf("client CreateSignedTx fail!", err)
					continue
				}
				if msg.Type == pb.ChaincodeMessage_Type_ONLY_READ_TRANSACTION {
					//kvRwSet, err := util.GetKVRWSetKwFromEnvelope(envelope)
					//if err != nil {
					//	log.Warn(err)
					//}
					//fmt.Printf("Transaction information: TxID = %s, readSet = %s \n", kvRwSet.TxID, kvRwSet.Reads)
				} else {
					// 将信封放入通道
					envelopeChan <- envelope
				}
			}
		}
		if num == int32(config.TxNum) {
			close(envelopeChan)
		}
	}
}

// sendToOrder 把封装好的事务发送给order节点
func sendToOrder(ch <-chan *pb.Envelope, orderStream order.OrderStream) {
	for {
		rmsg, ok := <-ch
		if !ok {
			//如果事务发送完成，则给order节点发送一个完成的消息
			msg := &pb.ChaincodeMessage{Type: pb.ChaincodeMessage_Type_COMPLETED}
			err := orderStream.Send(msg)
			if err != nil {
				log.Fatal(err)
			}
			//如果发送完毕，则关闭流
			err = orderStream.CloseSend()
			if err != nil {
				log.Warn(err)
			}
			break
		}

		//将事务发送给order节点
		envelopeBytes := util.MarshalOrPanic(rmsg)
		msg := &pb.ChaincodeMessage{Type: pb.ChaincodeMessage_Type_TRANSACTION, Payload: envelopeBytes}
		err := orderStream.Send(msg)
		if err != nil {
			log.Warn(err)
		}
	}
}

// receiveOrderMsg 接收从order节点返回的消息
func receiveOrderMsg(orderStream order.OrderStream) {
	for {
		msg, err := orderStream.Recv()
		if err == io.EOF {
			break
		}

		if msg == nil {
			break
		}
		//if msg.Type == pb.ChaincodeMessage_Type_ERROR {
		//	log.Debugf("transaction order fail.txID is: %s,and the reason for the error is: %s", msg.TxID, msg.Payload)
		//} else {
		//	log.Infof("transaction order success.txID is: %s", msg.TxID)
		//}
	}
}

func main() {

	// msgChan 用于存放发送给Peer节点的消息
	var msgChan = make(chan *pb.ChaincodeMessage, config.TxNum)

	// 开启peer节点服务
	go peer.Start()

	// 开启order节点服务
	go order.Start()

	// 获取客户端的签名信息
	conf := config.ClientConf
	sig, err := signer.NewSigner(conf)
	if err != nil {
		log.Fatal(err)
	}
	creator, err := sig.Serialize()
	if err != nil {
		log.Fatal(err)
	}

	//建立与Peer节点的连接
	clientStreams, err := shim.GetPeerClientStream()
	if err != nil {
		log.Warn(err)
	}

	//建立与Order节点的连接
	orderStream, err := order.GetOrderStream()
	if err != nil {
		log.Warn(err)
	}

	//获取所有发送给Peer节点的交易信息
	for i := 0; i < config.TxNum; i++ {
		msg, err := getClientMsg(creator, sig)
		if err != nil {
			log.Warn(err)
		}
		msgChan <- msg
	}

	// 设置每秒发送的事务数量
	// 创建一个Ticker，每10毫秒触发一次
	var ticker *time.Ticker
	ticker = time.NewTicker(10 * time.Millisecond)
	defer ticker.Stop()

	// 创建一个通道，用于控制事务发送
	tokenBucket := make(chan struct{}, config.Rate)

	// 启动一个goroutine来发送令牌
	go func() {
		for range ticker.C {
			for i := 0; i < int(config.Rate)/100; i++ {
				tokenBucket <- struct{}{}
			}
		}
	}()

	// 创建一个与Peer节点的客户端连接，用于告诉客户端交易开始执行（与上面的连接不同，上面的是流连接）
	peerClients, err := shim.GetPeerClient()
	if err != nil {
		log.Warn(err)
	}

	var wg sync.WaitGroup
	for _, client := range peerClients {
		wg.Add(1)
		go func(client pb.ChaincodeSupportClient) {
			defer wg.Done()
			_, err = client.StartPeer(context.Background(), &pb.Empty{})
			if err != nil {
				log.Warnf("init peer fail", err)
			}
		}(client)
	}
	wg.Wait()

	// 发起提案请求
	go func() {
		for i := 0; i < config.TxNum; i++ {

			if config.IsLimited {
				//限制事务发送的速率，如果令牌桶中没有令牌，则阻塞事务的发送
				<-tokenBucket
			}

			msg := <-msgChan

			//将提案请求发送给Peer节点
			var totalNodesCount = 0
			for orgIdx := 0; orgIdx < len(config.Organizations); orgIdx++ {
				org := config.Organizations[orgIdx]
				nodesCount := len(org.Ports)
				if nodesCount == 0 {
					log.Fatalf("org%d peer node Address is null!", orgIdx+1)
				}
				randomNodeIdx := totalNodesCount + source.Intn(nodesCount)
				totalNodesCount += nodesCount

				wg.Add(1)
				go func(randomNodeIdx int) {
					defer wg.Done()
					err = clientStreams[randomNodeIdx].Send(msg)
					if err != nil {
						log.Warn(err)
					}
				}(randomNodeIdx)
			}
			wg.Wait()
		}
		//主动关闭流
		for _, clientStream := range clientStreams {
			err = clientStream.CloseSend()
			if err != nil {
				log.Debug(err)
			}
		}
	}()

	// 接收Peer节点返回的消息
	for _, clientStream := range clientStreams {
		go collectEndorsementsAndCreateTransaction(clientStream, sig)
	}

	//将打包好的事务发送给Order节点
	go sendToOrder(envelopeChan, orderStream)

	//接收Order节点返回的消息
	receiveOrderMsg(orderStream)

	for _, client := range peerClients {
		wg.Add(1)
		go func(client pb.ChaincodeSupportClient) {
			defer wg.Done()
			_, err = client.EndPeer(context.Background(), &pb.Empty{})
			if err != nil {
				log.Warnf("end peer fail", err)
			}
		}(client)
	}
	wg.Wait()

}
