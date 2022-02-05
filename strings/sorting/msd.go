package strings

import sorts "algorithms/sorts"

const M = 10 // 小数组的切换阈值
var aux []string

type msd struct {
	a []string
}

func NewMSD(a []string) *msd {
	return &msd{a}
}

func (m *msd) Sort() {
	length := len(m.a)
	aux = make([]string, length)
	m.sort(0, len(m.a)-1, 0)
}

func (m *msd) sort(lo, hi, d int) {
	if hi <= lo+M {
		// 对第d个字符排序，从a[lo]到a[hi]
		sorts.NewInsertion(m.a, lo, hi, d).Sort()
		return
	}

	count := make([]int, R+2)

	// 计算频率
	for i := lo; i <= hi; i++ {
		count[charAt(m.a[i], d)+2]++
	}

	// 将频率转换为索引
	for r := 0; r < R+1; r++ {
		count[r+1] += count[r]
	}

	// 数据分类
	for i := lo; i <= hi; i++ {
		aux[count[charAt(m.a[i], d)+1]] = m.a[i]
		count[charAt(m.a[i], d)+1]++
	}

	// 回写
	for i := lo; i <= hi; i++ {
		m.a[i] = aux[i-lo]
	}

	// 递归地以每个字符为键进行排序
	for r := 0; r < R; r++ {
		m.sort(lo+count[r], lo+count[r+1]-1, d+1)
	}
}
