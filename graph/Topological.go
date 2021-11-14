package graph

import (
	"algorithms/graph/util"
)

type Topological struct {
	order util.Iterable
}

func NewTopological(g Digraph) *Topological {
	cycleFinder := NewDirectedCycle(g)
	if !cycleFinder.HasCycle() {
		dfs := NewDepthFirstOrder(g)

		order := dfs.ReversePost()
		return &Topological{order}
	}

	return &Topological{nil}
}

func (this *Topological) Order() util.Iterable {
	return this.order
}

func (this *Topological) IsDAG() bool {
	return this.order == nil
}
