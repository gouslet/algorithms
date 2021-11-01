package graph

import (
	"fmt"
	"io"
)

type Bag []int

type Graph struct {
	v   int   // 顶点数目
	e   int   // 边的数
	adj []Bag // 邻接表
}

// NewGraph 创建一个含有个顶点但不含有边的图
func NewGraph(v int) *Graph {
	adj := make([]Bag, v)
	return &Graph{v, 0, adj}
}

func NewGraphFromReader(r io.Reader) *Graph {
	var v, e int
	fmt.Fscanf(r, "%d\n%d\n", &v, &e)
	graph := NewGraph(v)
	var m, n int
	for i := 0; i < e; i++ {
		fmt.Fscanf(r, "%d%d\n", &m, &n)
		graph.AddEdge(m, n)
	}
	return graph
}

// V 返回图的顶点数
func (this *Graph) V() int {
	return this.v
}

// V 返回图的边数
func (this Graph) E() int {
	return this.e
}

//AddEdge 向图中添一条边v-w
func (this *Graph) AddEdge(v, w int) {
	this.adj[v] = append(this.adj[v], w)
	this.adj[w] = append(this.adj[w], v)
	this.e++
}

func (this *Graph) Adj(v int) Bag {
	return this.adj[v]
}

//Degree 计算v的度数
func Degree(g Graph, v int) int {
	return len(g.Adj(v))
}

//MaxDegree 计算所有顶点的最大数
func MaxDegree(g Graph) int {
	var max int = 0
	var i int = 0
	for d := g.V(); i < d; i++ {
		if dg := Degree(g, d); dg > max {
			max = dg
		}
	}
	return max
}

//AvgDegree计算所有顶点的平均度数
func AvgDegree(g Graph) float32 {
	return 2 * float32(g.E()) / float32(g.V())
}

//NumOfSelfLoops 计算自环的个数
func NumOfSelfoops(g Graph) int {
	var count int
	var i int
	for v := g.V(); i < v; i++ {
		for _, w := range g.Adj(i) {
			if i == w {
				count++
			}

		}
	}
	return count / 2
}
