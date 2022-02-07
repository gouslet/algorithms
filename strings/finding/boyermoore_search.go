package finding

const R = 256

type boyerMooreSearch struct {
	right   [R]int
	m       int
	pattern string
}

func NewBoyerMooreSearch(pattern string) boyerMooreSearch {
	// 预处理
	right := [R]int{-1}
	M, i := len(pattern), 0
	for ; i < M; i++ {
		right[pattern[i]] = i // 包含在模式字符串中的字符的值为其出现的最右位置
	}
	return boyerMooreSearch{right, M, pattern}
}

// Search BM字符串查找算法
func (b boyerMooreSearch) Search(txt string) int {
	N := len(txt)
	// 在txt中查找字符串
	skip := 0
	for i := 0; i <= N-b.m; i += skip { //文本字符串和模式在位置i匹配吗？
		skip = 0
		for j := b.m - 1; j > -1; j-- {
			if b.pattern[j] != txt[i+j] {
				skip = j - b.right[txt[i+j]]
				if skip < 1 {
					skip = 1
				}
				break
			}
		}
		if skip == 0 {
			return i
		}
	}
	return -1
}
