package pq

import "fmt"

type IndexHeapPQ struct {
	N      int     // 元素的数量
	pq     []int   // 索引二叉堆，从1开始
	qp     []int   // 逆序qp[pq[i]] = pq[qp[i]] = i
	values []Value // 有优先级之分的元素
}

func NewIndexHeapPQ(maxN int) *IndexHeapPQ {
	if maxN >= 0 {
		pq := make([]int, maxN+1, maxN+1)
		qp := make([]int, maxN+1, maxN+1)
		for i, _ := range qp {
			qp[i] = -1
		}
		return &IndexHeapPQ{0, pq, qp, make([]Value, maxN+1, maxN+1)}
	} else {
		panic("size of HeapExtrePQ can't be negative")
	}
}

func (this IndexHeapPQ) IsEmpty() bool {
	return this.N == 0
}

func (this IndexHeapPQ) Contains(k int) bool {
	k++ // 堆内数组从位置1开始，堆外序号从0开始
	return this.qp[k] != -1
}

func (this *IndexHeapPQ) Insert(k int, value Value) {
	if l := len(this.values); l < k+1 {
		msg := fmt.Sprintf("failed: index %d out of capacity %d\n", k, l)
		panic(msg)
	}
	if this.Contains(k) {
		this.Set(k, value)
		this.swim(k + 1)
		this.sink(k + 1)
	} else {
		this.N++
		k++ // 堆内数组从位置1开始，堆外序号从0开始
		this.qp[k] = this.N
		this.pq[this.N] = k
		this.values[k] = value
		this.swim(this.N)
	}

}

func (this *IndexHeapPQ) swim(k int) {
	for k > 1 && this.qp[k] != -1 && this.values[this.pq[k/2]].Less(this.values[this.pq[k]]) {
		this.swap(k/2, k)
		k /= 2
	}
}

func (this *IndexHeapPQ) sink(k int) {
	for 2*k <= this.N {
		j := 2 * k
		if j < this.N && this.values[this.pq[j]].Less(this.values[this.pq[j+1]]) {
			j++
		}
		if !this.values[this.pq[k]].Less(this.values[this.pq[j]]) {
			break
		}
		this.swap(k, j)
		k = j
	}
}

func (this IndexHeapPQ) Extre() Value {
	return this.values[this.pq[1]]
}

func (this *IndexHeapPQ) DelExtre() int {
	indexOfExtre := this.pq[1]
	this.swap(1, this.N)
	this.N--
	this.sink(1)
	this.values[this.pq[this.N+1]] = nil
	this.qp[this.pq[this.N+1]] = -1
	return indexOfExtre - 1 // 堆内数组从位置1开始，堆外序号从0开始
}

func (this *IndexHeapPQ) swap(i, j int) {
	this.qp[this.pq[i]] = j
	this.qp[this.pq[j]] = i
	this.pq[i], this.pq[j] = this.pq[j], this.pq[i]
}

func (this *IndexHeapPQ) Set(k int, value Value) {
	k++ // 堆内数组从位置1开始，堆外序号从0开始
	this.values[k] = value
	this.swim(k)
	this.sink(k)
}

func (this IndexHeapPQ) String() string {
	heap := make([]Value, this.N)
	for i := 1; i < this.N+1; i++ {
		if this.qp[i] != -1 {
			heap[i-1] = this.values[this.pq[i]]
		}
	}
	return fmt.Sprintf("%v", heap)
}
