package finding

const R = 256

// BoyerMooreSearch BM字符串查找算法
func BoyerMooreSearch(txt, pattern string) int {
	// 预处理
	right := [R]int{-1}
	M, N, i := len(pattern), len(txt), 0
	for ; i < M; i++ {
		right[pattern[i]] = i // 包含在模式字符串中的字符的值为其出现的最右位置
	}

	// 在txt中查找字符串
	skip := 0
	for i = 0; i <= N-M; i += skip { //文本字符串和模式在位置i匹配吗？
		skip = 0
		for j := M - 1; j > -1; j-- {
			if pattern[j] != txt[i+j] {
				skip = j - right[txt[i+j]]
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
