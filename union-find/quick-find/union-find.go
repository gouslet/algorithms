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
	return this.id[p]
}

func (this *UF) Union(p, q int) {
	pId := this.Find(p)
	qId := this.Find(q)

	if pId == qId {
		return
	}

	for i, id := range this.id {
		if id == pId {
			this.id[i] = qId
		}
	}
	this.count--
}
