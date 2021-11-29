package graph

import "algorithms/graph/util"

type DepthFirstPaths struct {
	marked []bool // 在这个定点上调用过dfs了吗？
	edgeTo []int  // 从起点到一个顶点的已知路径上的最后一个顶点
	s      int    // 起点
}

func NewDepthFirsPaths(g *Graph, s int) *DepthFirstPaths {
	marked := make([]bool, g.V())
	edgeTo := make([]int, g.V())
	paths := &DepthFirstPaths{marked, edgeTo, s}
	paths.dfs(g, s)
	return paths
}

func (this *DepthFirstPaths) dfs(g *Graph, v int) {
	this.marked[v] = true
	for _, w := range g.Adj(v) {
		if !this.marked[w] {
			this.edgeTo[w] = v
			this.dfs(g, w)
		}
	}
}

func (this *DepthFirstPaths) HasPathTo(v int) bool {
	return this.marked[v]
}

func (this *DepthFirstPaths) PathTo(v int) []int {
	if !this.HasPathTo(v) {
		return nil
	}
	paths := []int{0}
	for x := v; x != this.s; x = this.edgeTo[x] {
		paths = append(paths, x)
	}

	return paths
}

type BreadthFirstPaths struct {
	marked []bool // 在这个定点上调用过dfs了吗？
	edgeTo []int  // 到达该顶点的已知路径上的最后一个顶点
	s      int    // 起点
}

func NewBreadthFirsPaths(g *Graph, s int) *BreadthFirstPaths {
	marked := make([]bool, g.V())
	edgeTo := make([]int, g.V())
	paths := &BreadthFirstPaths{marked, edgeTo, s}
	paths.bfs(g, s)
	return paths
}

func (this *BreadthFirstPaths) bfs(g *Graph, v int) {
	this.marked[v] = true
	queue := util.NewQueue()
	queue.Enqueue(v)
	for !queue.IsEmpty() {
		s := queue.Dequeue()
		for _, w := range g.Adj(s) {
			if !this.marked[w] {
				this.edgeTo[w] = s
				this.marked[w] = true
				queue.Enqueue(w)
			}
		}
	}
}

func (this *BreadthFirstPaths) HasPathTo(v int) bool {
	return this.marked[v]
}

func (this *BreadthFirstPaths) PathTo(v int) []int {
	if !this.HasPathTo(v) {
		return nil
	}
	paths := util.NewStack()
	for x := v; x != this.s; x = this.edgeTo[x] {
		paths.Push(x)
	}
	paths.Push(this.s)
	p := make([]int, 0)
	for !paths.IsEmpty() {
		p = append(p, paths.Pop())
	}
	return p
}
