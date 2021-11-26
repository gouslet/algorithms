package graph

import "algorithms/graph/util"

type DirectedCycle struct {
	marked  []bool
	edgeTo  []int
	cycle   *util.Stack // 有向环中的所有顶点（如果存在）
	onStack []bool      // 递归调用栈上的所有顶点
}

func NewDirectedCycle(g Digraph) *DirectedCycle {
	dc := &DirectedCycle{make([]bool, g.V()), make([]int, g.V()), nil, make([]bool, g.V())}
	for i := 0; i < g.V(); i++ {
		if !dc.marked[i] {
			dc.dfs(g, i)
		}
	}
	return dc
}

func (this *DirectedCycle) dfs(g Digraph, v int) {
	this.onStack[v] = true
	this.marked[v] = true

	for _, w := range g.Adj(v) {
		if this.HasCycle() {
			return
		} else if !this.marked[w] {
			this.edgeTo[w] = v
			this.dfs(g, w)
		} else if this.onStack[w] {
			this.cycle = util.NewStack()
			for x := v; x != w; x = this.edgeTo[x] {
				this.cycle.Push(x)
			}
			this.cycle.Push(w)
			this.cycle.Push(v)
		}
	}
	this.onStack[v] = false
}

// HasCycle 有向图中是否含有环
func (this DirectedCycle) HasCycle() bool {
	return this.cycle != nil && this.cycle.Size() != 0
}

// Cycle 返回有向环中的所有顶点，如果存在的话
func (this DirectedCycle) Cycle() util.Iterable {
	return this.cycle
}
