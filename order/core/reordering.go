package core

// TxReorder 交易重排序
// 输入：txId 交易ID  g 交易分组情况
// 输出：bool 交易是否需要提前中止
func TxReorder(txId string, g *TransactionGroup) bool {

	num := g.Rely[txId]

	if g.Group[num] != nil {
		cg := CreateConflictGraph(g.Group[num])
		cg.TopologicalSort()
		if cg.HasLoop == true {
			return true
		}
	}

	return false
}
