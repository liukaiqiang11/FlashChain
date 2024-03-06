# FlashChain
## 1、需求
运行FlashChain，你至少需要golang (version >= 1.19).
## 2、如何去使用
### 运行代码
`go run main.go`
## 3、修改配置文件
```yml
block:
#  区块大小
  blksize: 50
#  区块的生成时间（单位ms）
  createBlockTime: 300
#  是否查看状态数据库
  showStateDB: true
#  是否查看区块链账本数据库
  showBlockchainDB: true

config:
#  地址数量
  addrNum: 1000
#  发送的事务数量
  txNum: 1000
#  发送事务时的偏斜度
  skewness: 1
#  发送事务的读占比
  ratio: 0.5
#  是否限制发送速率
  isLimited: false
#  发送速率（单位：个/每秒）
  rate: 35000
#  合约名（目前只支持 SmallBank 和 KvStore 两个合约）
  contractName: "SmallBank"

organizations:
  - name: org1
    ports:
      - ":1308"

client:
  MSPID: "*.lkq.com"
  IdentityPath: "cert/client.pem"
  KeyPath: "cert/client.key"

peer:
  address: ":1308"
  MSPID: "*.lkq.com"
  IdentityPath: "cert/peer.pem"
  KeyPath: "cert/peer.key"

order:
  address: ":1309"
  MSPID: "*.lkq.com"
  IdentityPath: "cert/order.pem"
  KeyPath: "cert/order.key"

ca:
  crt: "cert/ca.crt"
```
## 4、证明满足原子性和一致性
打开main.go中用于测试事务的原子性的代码
```golang
    // 将以下代码注释打开
    //from := source.Intn(int(config.AddrNum))
		//to := source.Intn(int(config.AddrNum))
		//args = [][]byte{
		//	[]byte("sendPayment"),
		//	[]byte(fmt.Sprintf("%x", from)),
		//	[]byte(fmt.Sprintf("%x", to)),
		//	[]byte("10"),
		//}
```

