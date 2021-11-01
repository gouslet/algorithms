package graph

type DepthFirstGraph struct {
	marked []bool
	count  int
}

func NewDepthFirstGraph(g Graph, s int) *DepthFirstGraph {
	marked := make([]bool, g.V())
	DFS := &DepthFirstGraph{marked, s}
	DFS.dfs(g, s)
	return DFS
}

func (this *DepthFirstGraph) dfs(g Graph, v int) {
	this.marked[v] = true
	this.count++
	for _, w := range g.Adj(v) {
		if !this.marked[w] {
			this.dfs(g, w)
		}
	}
}

func (this DepthFirstGraph) Marked(w int) bool {
	return this.marked[w]
}

func (this DepthFirstGraph) Count() int {
	return this.count
}
