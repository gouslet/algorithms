package finding

// ViolentSearch1 暴力字符串查找（实现一），返回模式字符串pattern在文本字符串txt中第一次出现的位置
func ViolentSearch1(txt, pattern string) int {
	for N, M, i := len(txt), len(pattern), 0; i < N-M+1; i++ {
		var j int
		for j = 0; j < M; j++ {
			if txt[i+j] != pattern[j] {
				break
			}
		}
		if j == M {
			return i // 找到匹配
		}
	}
	return -1 // 未找到匹配
}

// ViolentSearch2 暴力字符串查找（实现二），返回模式字符串pattern在文本字符串txt中第一次出现的位置
func ViolentSearch2(txt, pattern string) int {
	M := len(pattern)
	i, j := 0, 0
	for N := len(txt); i < N && j < M; i++ {
		if txt[i] == pattern[j] {
			j++
		} else {
			i -= j
			j = 0
		}
	}
	if j == M {
		return i - M // 找到匹配
	}
	return -1 // 未找到匹配
}
