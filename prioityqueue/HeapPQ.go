package pq

type Key interface {
	Less(b Key) bool
}

type HeapMaxPQ struct {
	pq []Key
	N  int
}

func NewHeapMaxPQWithSize(maxN int) *HeapMaxPQ {
	if maxN >= 0 {
		return &HeapMaxPQ{make([]Key, maxN+1, maxN+1), 0}
	} else {
		panic("size of HeapMaxPQ can't be negative")
	}
}

func NewHeapMaxPQ() *HeapMaxPQ {
	return &HeapMaxPQ{make([]Key, 0), 0}
}

func NewHeapMaxPQFrom(a []Key) *HeapMaxPQ {
	return &HeapMaxPQ{a, 0}
}

func (this HeapMaxPQ) Size() int {
	return this.N
}

func (this HeapMaxPQ) IsEmpty() bool {
	return this.N == 0
}

func (this *HeapMaxPQ) Insert(key Key) {
	this.N++
	this.pq[this.N] = key
	this.swim(this.N)
}

func (this *HeapMaxPQ) DelMax() Key {
	max := this.pq[1]
	this.swap(1, this.N)
	this.N--
	this.pq[this.N+1] = nil
	this.sink(1)
	return max
}

// swim 如果堆的有序状态因为某个结点变得比它的父结点更大而被打破，那么就需要将该结点不断上浮
func (this *HeapMaxPQ) swim(k int) {
	for k > 1 && this.pq[k/2].Less(this.pq[k]) {
		this.swap(k/2, k)
		k /= 2
	}
}

// sink 如果堆的有序状态因为某个结点变得比它的两个子结点之一小而被打破，那么就需要将该结点不断下沉
func (this *HeapMaxPQ) sink(k int) {
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

func (this *HeapMaxPQ) swap(i, j int) {
	temp := this.pq[i]
	this.pq[i] = this.pq[j]
	this.pq[j] = temp
}
