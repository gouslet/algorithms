package graph

import "algorithms/graph/util"

type DepthFirstOrder struct {
	marked      []bool
	pre, post   *util.Queue // 所有顶点的前、后序排列
	reversePost *util.Stack // 所有顶点的逆后序排列
}

func NewDepthFirstOrder(g Digraph) *DepthFirstOrder {
	pre := util.NewQueue()
	post := util.NewQueue()
	reversePost := util.NewStack()

	marked := make([]bool, g.V())

	dfo := &DepthFirstOrder{marked, pre, post, reversePost}
	for v := 0; v < g.V(); v++ {
		if !dfo.marked[v] {
			dfo.dfs(g, v)
		}
	}

	return dfo
}

func (this *DepthFirstOrder) dfs(g Digraph, v int) {
	this.pre.Enqueue(v)

	this.marked[v] = true
	for _, w := range g.Adj(v) {
		if !this.marked[w] {
			this.dfs(g, w)
		}
	}

	this.post.Enqueue(v)
	this.reversePost.Push(v)
}

func (this *DepthFirstOrder) Pre() util.Iterable {
	return this.pre
}

func (this *DepthFirstOrder) Post() util.Iterable {
	return this.post
}

func (this *DepthFirstOrder) ReversePost() util.Iterable {
	return this.reversePost
}
