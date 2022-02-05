package strings

import (
	"fmt"
)

const R = 256 // Ascii字符表长度

type lsd struct {
	w int // 通过前w个字符将a排序
	a []string
}

func NewLSD(w int, a []string) *lsd {
	if a == nil {
		panic(fmt.Errorf("string slice can not be nil"))
	}
	if len(a) == 0 {
		panic(fmt.Errorf("string slice can not be of 0 length"))
	}
	l := len(a[0])
	for _, s := range a {
		if len(s) != l {
			panic(fmt.Errorf("strings in slice must have the same length"))
		}
	}
	if w <= 0 {
		panic(fmt.Errorf("w must be a positive integer"))
	}
	if w > l {
		panic(fmt.Errorf("w can not be larger than the length of a string"))
	}
	return &lsd{w, a}
}

// Sort 通过前w个字符将a排序
func (l *lsd) Sort() []string {
	length := len(l.a)
	aux := make([]string, length)

	for d := l.w - 1; d >= 0; d-- { // 根据第d个字符用键索引计数法排序
		count := make([]int, R+1)

		for i := 0; i < length; i++ {
			count[l.a[i][d]+1]++
		}

		for r := 0; r < R; r++ {
			count[r+1] += count[r]
		}

		for i := 0; i < length; i++ {
			aux[count[l.a[i][d]]] = l.a[i]
			count[l.a[i][d]]++
		}

		for i := 0; i < length; i++ {
			l.a[i] = aux[i]
		}
	}
	return l.a
}
