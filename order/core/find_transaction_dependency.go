package core

import (
	pb "fchain/proto"
)

type ReadWriteTransaction struct {
	ReadTransactions  []string
	WriteTransactions []string
}

// FindTransactionDependency 发现事务之间的依赖关系
// 输入： kvRwSet 事务的读写集
// 输出： bool 该事务是否需要中止
func FindTransactionDependency(kvRwSet *pb.KVRWSet, rwSet map[string]ReadWriteTransaction) [][2]string {

	txSet := ReadWriteTransaction{}
	var deps [][2]string

	for _, read := range kvRwSet.Reads {
		tx, _ := rwSet[read.Key]
		for _, txID := range tx.WriteTransactions {
			if kvRwSet.TxID != txID {
				//fmt.Println("RW依赖", kvRwSet.TxID, "依赖于", txID)
				deps = append(deps, [2]string{kvRwSet.TxID, txID})
			}
		}
		txSet.ReadTransactions = append(tx.ReadTransactions, kvRwSet.TxID)
		txSet.WriteTransactions = tx.WriteTransactions
		rwSet[read.Key] = txSet
	}

	for _, write := range kvRwSet.Writes {
		tx, _ := rwSet[write.Key]
		for _, txID := range tx.ReadTransactions {
			if kvRwSet.TxID != txID {
				//fmt.Println("RW依赖", txID, "依赖于", kvRwSet.TxID)
				deps = append(deps, [2]string{txID, kvRwSet.TxID})
			}
		}
		txSet.ReadTransactions = tx.ReadTransactions
		txSet.WriteTransactions = append(tx.WriteTransactions, kvRwSet.TxID)
		rwSet[write.Key] = txSet
	}

	return deps
}
