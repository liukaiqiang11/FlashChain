package core

const minNLen = 16

// TransactionGroup 事务分组
type TransactionGroup struct {
	N     int            //总共分了多少个组
	Group [][][2]string  //分组的结果
	Rely  map[string]int //存储依赖关系
}

// NewGroup 新建分组
func NewGroup() *TransactionGroup {
	buf1 := make([][][2]string, minNLen)
	buf2 := make(map[string]int)
	return &TransactionGroup{
		N:     1,
		Group: buf1,
		Rely:  buf2,
	}
}

// resizeGroup 对分组大小进行扩容
func (g *TransactionGroup) resizeGroup() {
	newBuf := make([][][2]string, g.N<<1)
	copy(newBuf, g.Group[:])
	g.Group = newBuf

}

func (g *TransactionGroup) GroupTransactionsByDeps(table [][2]string) {
	for _, t := range table {
		g.AddEdge(t[0], t[1])
	}
}

// AddEdge 新加分组，t1、t2是存在依赖关系的两个事务
func (g *TransactionGroup) AddEdge(t1, t2 string) {
	// 新元素入队之前，当队列长度等于缓存区长度时，缓存区长度重设为两个队列长度
	if g.N >= len(g.Group)*8/10 {
		g.resizeGroup()
	}
	//当事务t1和事务t2都没进行分组时，将t1和t2加入到一个新的分组中
	if g.Rely[t1] == 0 && g.Rely[t2] == 0 {
		g.Rely[t1], g.Rely[t2] = g.N, g.N
		g.Group[g.N] = append(g.Group[g.N], [2]string{t1, t2})
		g.N++
	} else if g.Rely[t1] > 0 && g.Rely[t2] == 0 { //如果t1已经分组，但是t2还没分组，则将t2加入到t1分组
		g.Rely[t2] = g.Rely[t1]
		g.Group[g.Rely[t1]] = append(g.Group[g.Rely[t1]], [2]string{t1, t2})
	} else if g.Rely[t1] == 0 && g.Rely[t2] > 0 { //如果t2已经分组，但是t1还没分组，则将t1加入到t2分组
		g.Rely[t1] = g.Rely[t2]
		g.Group[g.Rely[t2]] = append(g.Group[g.Rely[t2]], [2]string{t1, t2})
	} else if g.Rely[t1] == g.Rely[t2] { //事务t1和t2都已经分组,并且分到了同一个组中
		g.Group[g.Rely[t1]] = append(g.Group[g.Rely[t1]], [2]string{t1, t2})
	} else { //事务t1和t2都已经分组，但是分到了不同的组中
		//记录t2的分组
		temp := g.Rely[t2]
		//将t2组中的事务和t1组中的事务合并
		g.Group[g.Rely[t1]] = append(g.Group[g.Rely[t1]], g.Group[g.Rely[t2]]...)
		g.Group[g.Rely[t1]] = append(g.Group[g.Rely[t1]], [2]string{t1, t2})
		//将t2的分组改为t1的分组
		for _, txs := range g.Group[temp] {
			for _, tx := range txs {
				g.Rely[tx] = g.Rely[t1]
			}
		}
		g.Group[temp] = nil
	}
}

func DeepCopyGroup(group *TransactionGroup) *TransactionGroup {
	buf := make([][][2]string, len(group.Group))
	for i, arr := range group.Group {
		buf[i] = make([][2]string, len(arr))
		for j, subArr := range arr {
			buf[i][j] = [2]string{subArr[0], subArr[1]}
		}
	}
	copyMap := make(map[string]int)
	for key, value := range group.Rely {
		copyMap[key] = value
	}

	return &TransactionGroup{N: group.N, Group: buf, Rely: copyMap}
}
