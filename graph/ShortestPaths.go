package graph

import (
	"algorithms/graph/util"
	"math"
)

type SP struct {
	distTo []float64 // 从s到点的已知最短路径长度
	edgeTo []int     // 父链接数组
}

func NewSP(g EdgeWeightedDigraph, s int) {
	return
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
func (this SP) PathTo(v int) []*DirectedEdge {
	if !this.HasPathTo(v) {
		return nil
	}
	paths := util.NewStack()
}
