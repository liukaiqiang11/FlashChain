package core

// ConflictGraph 冲突图
type ConflictGraph struct {
	N        int                 // 节点数
	HasLoop  bool                //是否存在环
	Graph    map[string][]string //记录节点间的依赖关系
	InDegree map[string]int      //记录每个节点的入度
	Queue    []string            //记录排序后的结果
}

// NewConflictGraph 新建一个冲突图
func NewConflictGraph() *ConflictGraph {
	buf1 := make(map[string][]string)
	buf2 := make(map[string]int)
	buf3 := make([]string, 0)
	return &ConflictGraph{
		N:        0,
		HasLoop:  false,
		Graph:    buf1,
		InDegree: buf2,
		Queue:    buf3,
	}
}

// AddEdge 添加边: t1,t2为事务id
func (cg *ConflictGraph) AddEdge(t1, t2 string) {

	//将依赖关系添加到冲突图
	cg.Graph[t1] = append(cg.Graph[t1], t2)

	//根据依赖关系计算入度表
	if cg.InDegree[t1] == 0 {
		cg.InDegree[t1] = 0
	}
	cg.InDegree[t2]++

	//计算节点数量
	cg.N = len(cg.InDegree)

}

// CreateConflictGraph 通过依赖关系表来创建冲突图
func CreateConflictGraph(table [][2]string) *ConflictGraph {
	s := NewConflictGraph()
	for _, t := range table {
		s.AddEdge(t[0], t[1])
	}
	return s
}
