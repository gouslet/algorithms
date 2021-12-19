package graph

import (
	pq "algorithms/prioityqueue"
	"math"
)

type SP struct {
	distTo []float64               // 从s到点的已知最短路径长度
	edgeTo []*WeightedDirectedEdge // 父链接数组
}

type distance float64

func (this distance) Less(b pq.Value) bool {
	v, _ := b.(distance)
	return this > v
}

func NewDijkstraSP(g *WeightedEdgeDigraph, s int) *SP {
	edgeTo := make([]*WeightedDirectedEdge, g.V())
	distTo := make([]float64, g.V())

	pq := pq.NewIndexHeapPQ(g.V())

	for v := 0; v < g.V(); v++ {
		distTo[v] = math.MaxFloat64
	}
	distTo[s] = 0.0

	pq.Insert(s, distance(0.0))

	relax := func(g WeightedEdgeDigraph, v int) {
		for _, e := range g.Adj(v) {
			w := e.To()
			if l := distTo[v] + e.Weight(); distTo[w] > l {
				distTo[w] = l
				edgeTo[w] = e
			}
			if pq.Contains(w) {
				pq.Set(w, distance(distTo[w]))
			} else {
				pq.Insert(w, distance(distTo[w]))
			}
		}
	}
	for !pq.IsEmpty() {
		relax(*g, pq.DelExtre())
	}
	return &SP{distTo, edgeTo}
}

// DistTo 从顶点s到v的距离，如果不存在则路径为无穷大
func (this SP) DistTo(v int) float64 {
	return this.distTo[v]
}

// HasPathTo 是否存在从顶点s到v的路径
func (this SP) HasPathTo(v int) bool {
	return this.distTo[v] < math.Inf(1)
}

// 从顶点s到v的路径，如果不存在则为nil
func (this SP) PathTo(v int) []*WeightedDirectedEdge {
	if !this.HasPathTo(v) {
		return nil
	}
	paths := []*WeightedDirectedEdge{}
	for e := this.edgeTo[v]; e != nil; e = this.edgeTo[e.From()] {
		paths = append(paths, e)
	}
	return paths
}
