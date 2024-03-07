package config

import (
	"fchain/common/signer"
	"fmt"
	"github.com/pkg/errors"
	"github.com/spf13/viper"
	"os"
)

// ClientConf 客户端签名配置
var ClientConf *signer.Config

// PeerAddress Peer节点地址
var PeerAddress string

// PeerConf Peer节点签名配置
var PeerConf *signer.Config

// OrderAddress Order节点地址
var OrderAddress string

// OrderConf Order节点签名配置
var OrderConf *signer.Config

// CaCrt 根证书地址
var CaCrt string

type Organization struct {
	Name  string
	Ports []string
}

var Organizations []Organization

var PeerAddressArr []string

// BlockSize 区块大小
var BlockSize uint64 = 500
var CreateBlockTime = 300
var AddrNum uint64 = 10000
var TxNum = 10000
var Skewness float64 = 0
var Ratio = 0.5
var Rate uint64 = 15000
var IsLimited = false
var ContractName = "SmallBank"
var ProjectPath string
var configFilePath string
var ShowStateDB bool
var ShowBlockchainDB bool

// PathExists 判断一个文件是否存在
func PathExists(path string) (bool, error) {
	_, err := os.Stat(path)
	if err == nil {
		return true, nil
	}
	if os.IsNotExist(err) {
		return false, nil
	}
	return false, err
}

func init() {
	var err error

	pwd, err := os.Getwd()
	if err != nil {
		panic(err)
	}
	ProjectPath = pwd + "/"
	configFilePath = pwd + "/config/"

	ClientConf, err = GetClientConf()
	if err != nil {
		panic(errors.New("client conf acquisition failed"))
	}

	PeerConf, err = GetPeerConf()
	if err != nil {
		panic(errors.New("peer conf acquisition failed"))
	}

	OrderConf, err = GetOrderConf()
	if err != nil {
		panic(errors.New("order conf acquisition failed"))
	}

	fmt.Println(configFilePath)
	if confFileExists, _ := PathExists(configFilePath); confFileExists != true {
		panic(errors.New("配置文件不存在"))
	}

	viper.AddConfigPath(configFilePath)
	viper.SetConfigName("config")
	if err := viper.ReadInConfig(); err != nil {
		panic(err)
	}

	// GetPeerAddress 获取Peer节点的地址
	PeerAddress = viper.GetString("peer.address")

	// GetOrderAddress 获取Order节点的地址
	OrderAddress = viper.GetString("order.address")

	// GetCaCrt 获取ca节点的ca.crt
	CaCrt = ProjectPath + viper.GetString("ca.crt")

	Organizations, err = GetOrganizations()
	if err != nil {
		panic(errors.New("organizations acquisition failed"))
	}

	for _, organization := range Organizations {
		PeerAddressArr = append(PeerAddressArr, organization.Ports...)
	}

	isDuplicate := hasDuplicate(PeerAddressArr)
	if isDuplicate {
		panic(errors.New("organization peer node address error!"))
	}

	// GetBlksize 从配置中获取区块大小
	BlockSize = viper.GetUint64("block.blksize")

	// GetCreateBlockTime 从配置文件中获取区块生成时间
	CreateBlockTime = viper.GetInt("block.createBlockTime")

	// ShowStateDB 是否查看状态数据库
	ShowStateDB = viper.GetBool("block.showStateDB")

	// ShowBlockchainDB 是否查看区块链账本数据库
	ShowBlockchainDB = viper.GetBool("block.showBlockchainDB")

	// GetAddrNum 从配置文件中获取地址数量
	AddrNum = viper.GetUint64("config.addrNum")

	// GetTxNum 从配置文件中获取交易数量
	TxNum = viper.GetInt("config.txNum")

	// GetSkewness 从配置文件中获取偏斜度
	Skewness = viper.GetFloat64("config.skewness")

	// GetRatio 从配置文件中获取读占比
	Ratio = viper.GetFloat64("config.ratio")

	// GetRate 从配置文件中获取发送速率
	Rate = viper.GetUint64("config.rate")

	// GetIsLimited 从配置文件中获取，是否限制发送速率
	IsLimited = viper.GetBool("config.isLimited")

	// GetContractName 获取合约名
	ContractName = viper.GetString("config.contractName")
}

func hasDuplicate(strArr []string) bool {
	// 遍历数组中的每个元素
	for i := 0; i < len(strArr); i++ {
		// 检查当前元素是否与后面的元素重复
		for j := i + 1; j < len(strArr); j++ {
			if strArr[i] == strArr[j] {
				return true
			}
		}
	}

	return false
}
