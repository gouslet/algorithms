package uf

type UF struct {
	id    []int //父链接数组
	sz    []int //各个根节点对应的分量大小
	count int   //连通分量数量
}

func NewUF(N int) *UF {
	id := make([]int, N)
	sz := make([]int, N)
	for i := 0; i < N; i++ {
		id[i] = i
		sz[i] = 1
	}

	return &UF{id, sz, N}
}

func (this *UF) Count() int {
	return this.count
}

func (this *UF) Connected(p, q int) bool {
	return this.Find(p) == this.Find(q)
}

func (this *UF) Find(p int) int {
	for p != this.id[p] {
		p = this.id[p]
	}
	return p
}

func (this *UF) Union(p, q int) {
	pRoot := this.Find(p)
	qRoot := this.Find(q)

	if pRoot == qRoot {
		return
	}
	// 将小树的根节点链接到大树的根节点
	if this.sz[pRoot] < this.sz[qRoot] {
		this.id[pRoot] = qRoot
		this.sz[qRoot] += this.sz[pRoot]
	} else {
		this.id[qRoot] = pRoot
		this.sz[pRoot] += this.sz[qRoot]
	}

	this.count--
}
