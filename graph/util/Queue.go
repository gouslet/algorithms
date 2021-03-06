package util

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

func (this *Queue) Map(fn func(v int)) {
	for i := 0; i < len(this.val); i++ {
		fn(this.val[i])
	}
}

func (this *Queue) Size() int {
	return len(this.val)
}
