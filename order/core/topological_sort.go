package core

// TopologicalSort 对冲突图拓扑排序
func (cg *ConflictGraph) TopologicalSort() {
	var queue []string

	//找到入度矩阵中入度为0的节点，将其加入到队列中
	for vertex, degree := range cg.InDegree {
		if degree == 0 {
			queue = append(queue, vertex)
		}
	}

	for len(queue) > 0 {
		vertex := queue[0]
		queue = queue[1:]
		cg.Queue = append(cg.Queue, vertex)

		//与该节点相邻的节点入度减一
		for _, neighbor := range cg.Graph[vertex] {
			cg.InDegree[neighbor]--
			//如果节点入度为0，则将其加入到queue中
			if cg.InDegree[neighbor] == 0 {
				queue = append(queue, neighbor)
			}
		}
	}

	if len(cg.Queue) != len(cg.InDegree) {
		cg.HasLoop = true
		cg.Queue = nil
	}
}
