package graph

type Edge struct {
	v, w   int     // 顶点
	weight float64 //边的权重
}

// Weight 边的权重
func (this Edge) Weight() float64 {
	return this.weight
}

// Either 边的两个端点之一
func (this Edge) Either() int {
	return this.v
}

// Other 边的另一个端点
func (this Edge) Other(vertex int) int {
	if vertex == this.v {
		return this.w
	} else {
		return this.v
	}
}

type WeightedEdgeGraph struct {
	v, e int       // 顶点的数量、边的数量
	adj  [][]*Edge // 邻接表
}

func NewWeightedEdgeGraph(v int) *WeightedEdgeGraph {
	adj := make([][]*Edge, v)
	for i := 0; i < v; i++ {
		adj[i] = make([]*Edge, 0)
	}
	return &WeightedEdgeGraph{v, 0, adj}
}

func (this WeightedEdgeGraph) V() int {
	return this.v
}

func (this WeightedEdgeGraph) E() int {
	return this.e
}

func (this *WeightedEdgeGraph) AddEdge(e *Edge) {
	v := e.Either()
	w := e.Other(v)
	this.adj[v] = append(this.adj[v], e)
	this.adj[w] = append(this.adj[w], e)
	this.e++
}

func (this WeightedEdgeGraph) Adj(v int) []*Edge {
	return this.adj[v]
}
func (this *WeightedEdgeGraph) Edges() []*Edge {
	edges := make([]*Edge, this.e)
	for i := 0; i < this.v; i++ {
		for _, e := range this.adj[i] {
			if e.Other(i) > i {
				edges = append(edges, e)
			}
		}
	}
	return edges
}
