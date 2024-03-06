package shim

import (
	"fchain/config"
	"fchain/peer/ledger"
	pb "fchain/proto"
	"fchain/proto/util"
	"fmt"
	cmap "github.com/orcaman/concurrent-map/v2"
	"strconv"
	"strings"
)

type StateDB struct {
	Value   []byte
	Version *pb.Version
}

type timer struct {
	startTime uint64
	endTime   uint64
}

var TxLatency = cmap.New[*timer]()

var State = cmap.New[*StateDB]()

var detection cmap.ConcurrentMap[string, struct{}]

var BlockNum uint64

var Blockchain *ledger.BlockStore

// GetStateDB 查看状态数据库
func GetStateDB() {
	valueNum := 0
	num := 0
	var maxColWidths = []int{18, 8, 12, 12, 7}

	fmt.Println("------------------------------状态数据库------------------------------")
	State.IterCb(func(key string, value *StateDB) {
		num++
		BlockNumber := value.Version.BlockNum
		GroupNum := value.Version.GroupNum
		TxNum := value.Version.TxNum
		if config.ContractName == "SmallBank" {
			val, _ := strconv.Atoi(string(value.Value))
			valueNum += val
			var line string
			line = fmt.Sprintf("账号:%s, 余额:%s, BlockNum:%d, GroupNum:%d, TxNum:%d\n",
				key, value.Value, BlockNumber, GroupNum, TxNum)
			fields := strings.Fields(line)
			if val != 100 && BlockNumber != 0 || GroupNum != 0 || TxNum != 0 {
				for i, field := range fields {
					// 根据最大宽度对齐字段
					format := fmt.Sprintf("%%-%ds", maxColWidths[i])
					fmt.Printf(format, field)

					// 打印字段之间的分隔符
					if i < len(fields)-1 {
						fmt.Print(" | ")
					}
				}
				fmt.Println()
			}
		} else if config.ContractName == "KvStore" {
			val, _ := strconv.Atoi(string(value.Value))
			if val != 100 && BlockNumber != 0 || GroupNum != 0 || TxNum != 0 {
				fmt.Printf("key:%s, value:%v, BlockNum:%d, GroupNum:%d, TxNum:%d\n",
					key, value.Value, value.Version.BlockNum, value.Version.GroupNum, value.Version.TxNum)
			}
		}
	})

	if config.ContractName == "SmallBank" {
		fmt.Println("账号数量为：", num)
		fmt.Println("所有账户的总余额为：", valueNum)
	}
}

func GetBlockchainDB() {
	fmt.Println("----------------------------------------区块链账本----------------------------------------")
	ito, _ := Blockchain.RetrieveBlocks(0)
	for ito.Next() {
		blockByte := ito.Value()
		util.StringBlockByte(blockByte)
	}
}

func GetTotalLatency() uint64 {
	var totalTime uint64 = 0
	TxLatency.IterCb(func(key string, t *timer) {
		if t.endTime != 0 {
			totalTime += t.endTime - t.startTime
		}
	})
	return totalTime
}
