package graph

import (
	"fmt"
	"io"
)

type WeightedDirectedEdge struct {
	v, w   int     // 边的起点和终点
	weight float64 // 边的权重
}

// Weight 边的权重
func (this WeightedDirectedEdge) Weight() float64 {
	return this.weight
}

// From 从这条边指出的顶点
func (this WeightedDirectedEdge) From() int {
	return this.v
}

// From 这条边指向的顶点
func (this WeightedDirectedEdge) To() int {
	return this.w
}

func (this WeightedDirectedEdge) String() string {
	return fmt.Sprintf("%d->%d %.2f", this.v, this.w, this.weight)
}

type WeightedEdgeDigraph struct {
	v, e int                       // 顶点总数、边的总数
	adj  [][]*WeightedDirectedEdge // 邻接表
}

func NewWeightedEdgeDigraphFrom(r io.Reader) *WeightedEdgeDigraph {
	var v, e int
	fmt.Fscanf(r, "%d\n%d\n", &v, &e)
	adj := make([][]*WeightedDirectedEdge, 0)
	for i := 0; i < v; i++ {
		adj = append(adj, make([]*WeightedDirectedEdge, 0))
	}
	graph := &WeightedEdgeDigraph{v, e, adj}
	var m, n int
	var w float64
	for i := 0; i < e; i++ {
		fmt.Fscanf(r, "%d%d%f\n", &m, &n, &w)
		graph.AddEdge(&WeightedDirectedEdge{m, n, w})
	}
	return graph
}

// V 顶点总数
func (this WeightedEdgeDigraph) V() int {
	return this.v
}

// E 边的总数
func (this WeightedEdgeDigraph) E() int {
	return this.e
}

// AddEdge 添加边e到该有向图中
func (this *WeightedEdgeDigraph) AddEdge(e *WeightedDirectedEdge) {
	this.adj[e.From()] = append(this.adj[e.From()], e)
	this.e++
}

// Adj 从v指出边
func (this WeightedEdgeDigraph) Adj(v int) []*WeightedDirectedEdge {
	return this.adj[v]
}

// Edges 该有向边中的所有边
func (this WeightedEdgeDigraph) Edges() []*WeightedDirectedEdge {
	edges := make([]*WeightedDirectedEdge, this.e)
	for _, e := range this.adj {
		edges = append(edges, e...)
	}
	return edges
}

func String() string {
	return ""
}
