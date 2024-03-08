package shim

import (
	"context"
	"fchain/config"
	"fchain/peer/ledger"
	"fchain/peer/shim/internal"
	pb "fchain/proto"
	"fmt"
	log "github.com/corgi-kx/logcustom"
	cmap "github.com/orcaman/concurrent-map/v2"
	"github.com/pkg/errors"
	"google.golang.org/protobuf/types/known/timestamppb"
	"strings"
	"sync/atomic"
	"time"
)

var StartTime time.Time

type ChaincodeServer struct {
	Address string
	CC      Chaincode
	pb.UnimplementedChaincodeSupportServer
}

func NewChaincodeServer(addr string) *ChaincodeServer {

	return &ChaincodeServer{Address: addr}
}

func (cs *ChaincodeServer) Register(ServerStream pb.ChaincodeSupport_RegisterServer) error {

	return chatWithPeer(ServerStream)
}

func (cs *ChaincodeServer) StartPeer(ctx context.Context, e *pb.Empty) (*pb.Empty, error) {

	// 获取Peer节点的端口地址
	dbPath := "blockDB" + config.PeerAddress
	dbPath = strings.Replace(dbPath, ":", "_", -1)
	//创建区块链账本
	Blockchain, _ = ledger.NewBlockStore(config.ProjectPath+dbPath, "blockchain")

	//设置冲突检测表的缓存大小
	detection = cmap.New[struct{}]()
	TxLatency = cmap.New[*timer]()

	// 给每个账号赋初值
	for i := 0; i < int(config.AddrNum); i++ {
		if config.ContractName == "SmallBank" {
			value := &StateDB{Value: []byte("100"), Version: &pb.Version{}}
			State.Set(fmt.Sprintf("saving_%d", i), value)
			State.Set(fmt.Sprintf("checking_%d", i), value)
		} else if config.ContractName == "KvStore" {
			value := &StateDB{Value: []byte{}, Version: &pb.Version{}}
			State.Set(fmt.Sprintf("%d", i), value)
		}
	}
	StartTime = time.Now()
	return e, nil
}

func (cs *ChaincodeServer) EndPeer(ctx context.Context, e *pb.Empty) (*pb.Empty, error) {

	<-EndChannel
	duration := time.Since(StartTime)

	//查看区块账本
	if config.ShowBlockchainDB {
		GetBlockchainDB()
	}

	//查看状态数据库
	if config.ShowStateDB {
		GetStateDB()
	}

	fmt.Printf("运行时间：%s\n", duration)
	successRatio := float64(SuccessNum) / float64(config.TxNum)
	throughput := float64(config.TxNum) * 1e9 / float64(duration)
	totalLatency := GetTotalLatency()
	latency := float64(totalLatency) / float64(SuccessNum)
	fmt.Printf("吞吐量为：%f\n", throughput*successRatio)
	fmt.Printf("延迟为： %f ms\n", latency)
	fmt.Printf("成功率为： %f%%\n", successRatio*100)

	BlockNum = 0
	SuccessNum = 0
	Blockchain.Shutdown()
	return e, nil
}

func (cs *ChaincodeServer) GetBlockInfo(ctx context.Context, e *pb.Empty) (*pb.BlockchainInfo, error) {

	info, err := Blockchain.GetBlockchainInfo()
	if err != nil {
		log.Debugf("get BlockchainInfo fail", err)
	}
	for BlockNum != info.Height {
		info, err = Blockchain.GetBlockchainInfo()
	}
	atomic.AddUint64(&BlockNum, 1)
	return info, nil
}

func (cs *ChaincodeServer) ValidateBlock(ctx context.Context, msg *pb.ValidateMessage) (*pb.ValidateMessage, error) {
	err := ValidateBlock(msg)
	var res pb.Response
	if err != nil {
		res = Error("block validate fail")
		return &pb.ValidateMessage{
			Type:      pb.ValidateMessage_Type_VALIDATE_FAIL,
			Timestamp: timestamppb.Now(),
			Response:  &res,
		}, err
	}
	res = Success(nil)
	return &pb.ValidateMessage{
		Type:      pb.ValidateMessage_Type_VALIDATE_SUCCESS,
		Timestamp: timestamppb.Now(),
		Response:  &res,
	}, nil
}

// Start 开启这个服务
func (cs *ChaincodeServer) Start() error {

	if cs.Address == "" {
		return errors.New("address must be specified")
	}

	var err error

	tlsCfg, err := internal.LoadTLSConfig()
	if err != nil {
		return err
	}

	server, err := internal.NewServer(cs.Address, tlsCfg)
	if err != nil {
		return err
	}

	defer server.Stop()

	pb.RegisterChaincodeSupportServer(server.Server, cs)

	return server.Start()
}
