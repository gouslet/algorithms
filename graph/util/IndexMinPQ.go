package util

type Item interface {
}

type IndexMinPQ struct {
}

func NewIndexMinPQ(n int) *IndexMinPQ {
	return nil
}

func (this *IndexMinPQ) Insert(k int, item Item) {

}

func (this *IndexMinPQ) Set(k int, item Item) {

}

func (this IndexMinPQ) Contains(k int) bool {
	return false
}

func (this IndexMinPQ) IsEmpty() bool {
	return false
}

func (this IndexMinPQ) DelMin() int {
	return 0
}
