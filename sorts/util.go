package sorts

import "constraints"

type Comparable[T any] struct {
	d   []T
	cmp func(T, T) bool
}

func (data Comparable[T]) Less(i, j int) bool {
	return data.cmp(data.d[i], data.d[j])
}

func (data Comparable[T]) Len() int {
	return len(data.d)
}

func (data Comparable[T]) Swap(i, j int) {
	data.d[i], data.d[j] = data.d[j], data.d[i]
}

type Slice[T constraints.Ordered] []T

type Sortable[T any] interface {
	[]T
}

func (data Slice[T]) Less(i, j int) bool {
	return data[i] < data[j]
}

// func (data Slice[T]) Less(i, j int) bool {
// 	switch data.(type) {
// 	case constraints.Ordered:
// 		return data[i] < data[j]
// 	default:
// 		return data.d.cmp(i, j)
// 	}
// }

func (data Slice[T]) Len() int {
	return len(data)
}

func (data Slice[T]) Swap(i, j int) {
	data[i], data[j] = data[j], data[i]
}
