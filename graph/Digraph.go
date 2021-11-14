package graph

import (
	"fmt"
	"io"
)

type Digraph struct {
	v int // 顶点数目
	e int // 边的数目

	adj []Bag // 	邻接表数组
}

func NewDigraph(v int) *Digraph {
	adj := make([]Bag, v)
	return &Digraph{v, 0, adj}
}

func NewDigraphFrom(r io.Reader) *Digraph {
	var v, e int
	fmt.Fscanf(r, "%d\n%d\n", &v, &e)
	graph := NewDigraph(v)
	var m, n int
	for i := 0; i < e; i++ {
		fmt.Fscanf(r, "%d%d\n", &m, &n)
		graph.AddEdge(m, n)
	}
	return graph
}

func (this *Digraph) V() int {
	return this.v
}

func (this *Digraph) E() int {
	return this.e
}

func (this *Digraph) AddEdge(v, w int) {
	this.adj[v] = append(this.adj[v], w)
	this.e++
}

// 给出由顶点v指出的边所连接的所有顶点
func (this *Digraph) Adj(v int) []int {
	return this.adj[v]
}

// 返回该有向图的一个副本，但将其中所有边的方向反转
func (this *Digraph) Reverse() *Digraph {
	rdg := NewDigraph(this.v)
	for i := 0; i < this.v; i++ {
		for _, j := range this.adj[i] {
			rdg.AddEdge(j, i)
		}
	}
	return rdg
}

func (this *Digraph) String() string {
	s := fmt.Sprintf("%d vertices, %d edges\n", this.v, this.e)
	for i := 0; i < this.v; i++ {
		s += fmt.Sprintf("%d: ", i)
		for _, j := range this.adj[i] {
			s += fmt.Sprintf("%d", j)
		}
		s += fmt.Sprintln()
	}
	return s
}
