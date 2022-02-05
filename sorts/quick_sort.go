package sorts

import (
	"sort"
)

// QuickSort 快速排序
func QuickSort(data sort.Interface) {
	quickSort(data, 0, data.Len()-1)
}

func quickSort(data sort.Interface, lo, hi int) {
	if hi <= lo {
		return
	}

	j := partition(data, lo, hi)
	quickSort(data, lo, j-1)
	quickSort(data, j+1, hi)
}

func partition(data sort.Interface, lo, hi int) int {
	i, j := lo, hi+1

	for true {

		for i += 1; data.Less(i, lo); i++ {
			if i == hi {
				break
			}
		}

		for j -= 1; data.Less(lo, j); j-- {
			if j == lo {
				break
			}
		}
		if i >= j {
			break
		}
		data.Swap(i, j)
	}
	data.Swap(lo, j)
	return j
}

// Quick3Way 三向快速切分
func Quick3Way(data sort.Interface) {
	quick3Way(data, 0, data.Len()-1)
}

func quick3Way(data sort.Interface, lo, hi int) {
	if hi <= lo {
		return
	}
	lt, i, gt := lo, lo+1, hi
	th := lo //th为切分值的索引
	for i <= gt {
		a := data.Less(i, th)
		b := data.Less(th, i)
		if a {
			data.Swap(lt, i)
			th = i
			lt++
			i++
		} else if b {
			data.Swap(i, gt)
			gt--
		} else {
			i++
		}
	} //现在data[lo..lt-1] < data[th] = data[lt..gt] < data[gt+1..h]成立
	quick3Way(data, lo, lt-1)
	quick3Way(data, gt+1, hi)
}
