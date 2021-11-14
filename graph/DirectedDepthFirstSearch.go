package graph

type DirectedDFS struct {
	marked []bool
}

func NewDirectedDFS(g Graph, s int) *DirectedDFS {
	marked := make([]bool, g.V())
	DDFS := &DirectedDFS{marked}
	DDFS.dfs(g, s)
	return DDFS
}

func (this *DirectedDFS) dfs(g Graph, v int) {
	this.marked[v] = true
	for _, w := range g.Adj(v) {
		if !this.marked[w] {
			this.dfs(g, w)
		}
	}
}

func (this DirectedDFS) Marked(w int) bool {
	return this.marked[w]
}
