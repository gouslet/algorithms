package pq

type Value interface {
	Less(b Value) bool
}

type HeapPQ struct {
	pq []Value
	N  int
}

func NewHeapPQWithSize(maxN int) *HeapPQ {
	if maxN >= 0 {
		return &HeapPQ{make([]Value, maxN+1, maxN+1), 0}
	} else {
		panic("size of HeapPQ can't be negative")
	}
}

func NewHeapPQ() *HeapPQ {
	return &HeapPQ{make([]Value, 0), 0}
}

func NewHeapPQFrom(a []Value) *HeapPQ {
	return &HeapPQ{a, 0}
}

func (this HeapPQ) Size() int {
	return this.N
}

func (this HeapPQ) IsEmpty() bool {
	return this.N == 0
}

func (this *HeapPQ) Insert(value Value) {
	this.N++
	this.pq[this.N] = value
	this.swim(this.N)
}

func (this *HeapPQ) DelExtre() Value {
	max := this.pq[1]
	this.swap(1, this.N)
	this.N--
	this.pq[this.N+1] = nil
	this.sink(1)
	return max
}

// swim 如果堆的有序状态因为某个结点变得比它的父结点更大而被打破，那么就需要将该结点不断上浮
func (this *HeapPQ) swim(k int) {
	for k > 1 && this.pq[k/2].Less(this.pq[k]) {
		this.swap(k/2, k)
		k /= 2
	}
}

// sink 如果堆的有序状态因为某个结点变得比它的两个子结点之一小而被打破，那么就需要将该结点不断下沉
func (this *HeapPQ) sink(k int) {
	for 2*k <= this.N {
		j := 2 * k
		if j < this.N && this.pq[j].Less(this.pq[j+1]) {
			j++
		}
		if !this.pq[k].Less(this.pq[j]) {
			break
		}
		this.swap(k, j)
		k = j
	}

}

func (this *HeapPQ) swap(i, j int) {
	temp := this.pq[i]
	this.pq[i] = this.pq[j]
	this.pq[j] = temp
}
