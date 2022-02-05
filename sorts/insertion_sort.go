package sorts

type insertion[T Item] struct {
	a         []T
	lo, hi, d int
}

type Item interface {
	~string
}

//
func NewInsertion[T Item](a []T, lo, hi, d int) *insertion[T] {
	return &insertion[T]{a, lo, hi, d}
}

func (ins *insertion[T]) Sort() {
	for i := ins.lo; i <= ins.hi; i++ {
		for j := i; j > ins.lo && less(string(ins.a[j]), string(ins.a[j-1]), ins.d) == true; j-- {
			swap(ins.a, j, j-1)
		}
	}
}

func swap[T any](a []T, i, j int) {
	a[i], a[j] = a[j], a[i]
}

// less 比较a和b字符串第d位的大小
func less(a, b string, d int) bool {
	return a[d] < b[d]
}
