package sorting

type Item struct {
	K int
	V string
}

type indexCountingSort struct {
	n int // 分组最大值+1
	a []Item
}

func NewIndexCountingSort(a []Item) *indexCountingSort {
	var n int
	for _, t := range a {
		if t.K > n {
			n = t.K
		}
	}
	return &indexCountingSort{n + 1, a} //键索引计数法中键为[0,n)之间的一个整数
}

// Sort 用键索引计数法对a排序
func (ics *indexCountingSort) Sort() []Item {
	l := len(ics.a)

	count := make([]int, ics.n+1)
	aux := make([]Item, l)
	// 计算出现频率
	for _, t := range ics.a {
		count[t.K+1]++
	}
	// 将频率转化为索引
	for i := 0; i < ics.n; i++ {
		count[i+1] += count[i]
	}

	// 将元素分类
	for i := 0; i < l; i++ {
		aux[count[ics.a[i].K]] = ics.a[i]
		count[ics.a[i].K]++
	}

	// 回写
	for i, t := range aux {
		ics.a[i] = t
	}

	return ics.a
}
