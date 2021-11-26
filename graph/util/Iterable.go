package util

type Iterable interface {
	Map(fn func(v int))
}
