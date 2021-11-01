package graph

type DepthFirstPaths struct {
	marked []bool // 在这个定点上调用过dfs了吗？
	edgeTo []int  // 从起点到一个顶点的已知路径上的最后一个顶点
	s      int    // 起点
}

func NewDepthFirsPaths(g Graph, s int) *DepthFirstPaths {
	marked := make([]bool, g.V())
	edgeTo := make([]int, g.V())
	paths := &DepthFirstPaths{marked, edgeTo, s}
	paths.dfs(g, s)
	return paths
}

func (this *DepthFirstPaths) dfs(g Graph, v int) {
	this.marked[v] = true
	for _, w := range g.Adj(v) {
		if !this.marked[w] {
			this.edgeTo[w] = v
			this.dfs(g, w)
		}
	}
}

func (this *DepthFirstPaths) HasPathTo(v int) bool {
	return this.marked[v]
}

func (this *DepthFirstPaths) PathTo(v int) []int {
	if !this.HasPathTo(v) {
		return nil
	}
	var paths []int
	for x := v; x != this.s; x = this.edgeTo[x] {
		paths = append(paths, x)
	}
	return paths
}

type BreadthFirstPaths struct {
	marked []bool // 在这个定点上调用过dfs了吗？
	edgeTo []int  // 到达该顶点的已知路径上的最后一个顶点
	s      int    // 起点
}

func NewBreadthFirsPaths(g Graph, s int) *BreadthFirstPaths {
	marked := make([]bool, g.V())
	edgeTo := make([]int, g.V())
	paths := &BreadthFirstPaths{marked, edgeTo, s}
	paths.bfs(g, s)
	return paths
}

func (this *BreadthFirstPaths) bfs(g Graph, v int) {
	this.marked[v] = true
	queue := NewQueue()
	queue.Enqueue(v)
	for !queue.IsEmpty() {
		s := queue.Dequeue()
		for _, w := range g.Adj(s) {
			if !this.marked[w] {
				this.edgeTo[w] = s
				this.marked[w] = true
				queue.Enqueue(w)
			}
		}
	}
}

func (this *BreadthFirstPaths) HasPathTo(v int) bool {
	return this.marked[v]
}

func (this *BreadthFirstPaths) PathTo(v int) []int {
	if !this.HasPathTo(v) {
		return nil
	}
	var paths []int
	for x := v; x != this.s; x = this.edgeTo[x] {
		paths = append(paths, x)
	}
	return paths
}

type Queue struct {
	val  []int
	size int
}

func NewQueue() *Queue {
	return &Queue{nil, 0}
}

func (this *Queue) IsEmpty() bool {
	return this.size == 0
}

func (this *Queue) Enqueue(v int) {
	this.val = append(this.val, v)
	this.size++
}

func (this *Queue) Dequeue() int {
	if !this.IsEmpty() {
		v := this.val[0]
		this.val = this.val[1:]
		this.size--
		return v
	} else {
		panic("queue is empty,can't deque")
	}
}
