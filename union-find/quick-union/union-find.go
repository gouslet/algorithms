package uf

type UF struct {
	id    []int //分量id
	count int   //分量数量
}

func NewUF(N int) *UF {
	id := make([]int, N)
	for i := 0; i < N; i++ {
		id[i] = i
	}

	return &UF{id, N}
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

	this.id[pRoot] = qRoot

	this.count--
}
