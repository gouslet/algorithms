package pq

type IndexHeapMaxPQ struct {
	N    int   // 元素的数量
	pq   []int // 索引二叉堆，从1开始
	qp   []int // 逆序qp[pq[i]] = pq[qp[i]] = i
	keys []Key // 有优先级之分的元素
}

func NewIndexHeapMaxPQ(maxN int) *IndexHeapMaxPQ {
	if maxN >= 0 {
		pq := make([]int, maxN+1, maxN+1)
		qp := make([]int, maxN+1, maxN+1)
		for i, _ := range qp {
			qp[i] = -1
		}
		return &IndexHeapMaxPQ{0, pq, qp, make([]Key, maxN+1, maxN+1)}
	} else {
		panic("size of HeapMaxPQ can't be negative")
	}
}

func (this IndexHeapMaxPQ) IsEmpty() bool {
	return this.N == 0
}

func (this IndexHeapMaxPQ) Contains(k int) bool {
	return this.qp[k] != -1
}

func (this *IndexHeapMaxPQ) Insert(k int, key Key) {
	this.N++
	k++ // 堆内数组从位置1开始，堆外序号从0开始
	this.qp[k] = this.N
	this.pq[this.N] = k
	this.keys[k] = key
	this.swim(this.N)
}

func (this *IndexHeapMaxPQ) swim(k int) {
	for k > 1 && this.keys[this.pq[k/2]].Less(this.keys[this.pq[k]]) {
		this.swap(k/2, k)
		k /= 2
	}
}

func (this *IndexHeapMaxPQ) sink(k int) {
	for 2*k <= this.N {
		j := 2 * k
		if j < this.N && this.keys[this.pq[j]].Less(this.keys[this.pq[j+1]]) {
			j++
		}
		if !this.keys[this.pq[k]].Less(this.keys[this.pq[j]]) {
			break
		}
		this.swap(k, j)
		k = j
	}
}

func (this IndexHeapMaxPQ) Max() Key {
	return this.keys[this.pq[1]]
}

func (this *IndexHeapMaxPQ) DelMax() int {
	indexOfMax := this.pq[1]
	this.swap(1, this.N)
	this.N--
	this.sink(1)
	this.keys[this.pq[this.N+1]] = nil
	this.qp[this.pq[this.N+1]] = -1
	return indexOfMax - 1 // 堆内数组从位置1开始，堆外序号从0开始
}

func (this *IndexHeapMaxPQ) swap(i, j int) {
	this.qp[this.pq[i]] = j
	this.qp[this.pq[j]] = i
	this.pq[i], this.pq[j] = this.pq[j], this.pq[i]
}
