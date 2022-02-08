package finding

type kmpSearch struct {
	dfa [R][]int
	m   int // 模式字符串长度
}

func NewKMPSearch(pattern string) kmpSearch {
	// 预处理
	M := len(pattern)
	dfa := [R][]int{}

	if M > 0 {
		for i, _ := range dfa {
			dfa[i] = make([]int, M)
		}
		dfa[pattern[0]][0] = 1
	}
	for x, i := 0, 1; i < M; i++ {
		for c := 0; c < R; c++ {
			dfa[c][i] = dfa[c][x]
		}
		dfa[pattern[i]][i] = i + 1
		x = dfa[pattern[i]][x]
	}

	return kmpSearch{dfa, M}
}

// Search KMP字符串查找算法
func (k kmpSearch) Search(txt string) int {
	N := len(txt)
	var (
		i int
		j int
	)
	// 在txt中查找字符串
	for ; i < N && j < k.m; i++ { //文本字符串和模式在位置i匹配吗？
		j = k.dfa[txt[i]][j]
	}
	if j == k.m {
		return i - k.m
	}
	return -1
}
