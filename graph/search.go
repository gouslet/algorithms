package graph

type Search interface {
	marked(v int) bool // 与顶点v是连通的吗？
	count() int        // 与本顶点连通的顶点总数
}
