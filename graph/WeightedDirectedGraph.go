package graph

import (
	"fmt"
)

type DirectedEdge struct {
	v, w   int     // 边的起点和终点
	weight float64 // 边的权重
}

// Weight 边的权重
func (this DirectedEdge) Weight() float64 {
	return this.weight
}

// From 从这条边指出的顶点
func (this DirectedEdge) From() int {
	return this.v
}

// From 这条边指向的顶点
func (this DirectedEdge) To() int {
	return this.w
}

func (this DirectedEdge) String() string {
	return fmt.Sprintf("%d->%d %.2f", this.v, this.w, this.weight)
}

type EdgeWeightedDigraph struct {
	v, e int               // 顶点总数、边的总数
	adj  [][]*DirectedEdge // 邻接表
}

// V 顶点总数
func (this EdgeWeightedDigraph) V() int {
	return this.v
}

// E 边的总数
func (this EdgeWeightedDigraph) E() int {
	return this.e
}

// AddEdge 添加边e到该有向图中
func (this *EdgeWeightedDigraph) AddEdge(e *DirectedEdge) {
	this.adj[e.From()] = append(this.adj[e.From()], e)
	this.e++
}

// Adj 从v指出边
func (this EdgeWeightedDigraph) Adj(v int) []*DirectedEdge {
	return this.adj[v]
}

// Edges 该有向边中的所有边
func (this EdgeWeightedDigraph) Edges() []*DirectedEdge {
	edges := make([]*DirectedEdge, this.e)
	for _, e := range this.adj {
		edges = append(edges, e...)
	}
	return edges
}

func String() string {
	return ""
}
