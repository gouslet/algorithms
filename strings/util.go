// Package strings includes some algorithms about manipulating strings
package strings

// R ASCII表的元素个数
const R = 256

// charAt 将字符串中字符索引转换为数组索引
func CharAt(s string, d int) int {
	if d < len(s) {
		return int(s[d])
	} else { //当指定的位置超过了字符串的末尾时，返回-1
		return -1
	}
}

// swap 交换a[i]和a[j]
func Swap(a []string, i, j int) {
	a[i], a[j] = a[j], a[i]
}
