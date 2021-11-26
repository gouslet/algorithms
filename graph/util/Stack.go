package util

type Stack struct {
	val  []int
	size int
}

func NewStack() *Stack {
	return &Stack{nil, 0}
}

func (this *Stack) IsEmpty() bool {
	return this.size == 0
}

func (this *Stack) Push(v int) {
	this.val = append(this.val, v)
	this.size++
}

func (this *Stack) Pop() int {
	if !this.IsEmpty() {
		v := this.val[len(this.val)-1]
		this.val = this.val[:len(this.val)-1]
		this.size--
		return v
	} else {
		panic("Stack is empty,can't pop")
	}
}

func (this Stack) Map(fn func(v int)) {
	for i := len(this.val) - 1; i > -1; i-- {
		fn(int(this.val[i]))
	}
}

func (this *Stack) Size() int {
	return len(this.val)
}
