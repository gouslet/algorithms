package pq

type MaxPQ interface {
	Insert(k Value) // 向优先队列中插入一个元素
	Max() Value     // 返回最大元素
	DelMax() *Value // 删除并返回最大元素
	IsEmpty() bool  // 返回队列是否为空
	Size() int      // 返回优先队列中的元素个数
}

type IndexPQ interface {
	Insert(k int, key Value)
	Change(k int, key Value)
	Contains(k int) bool
	Max()
}
