package finding

const q = 997 // 一个很大的素数Q

// RabinKarpSearch RK字符串查找算法
type rabinKarpSearch struct {
	m           int   // 模式字符串的长度
	rm          int64 // R^(M-1)%Q
	patternHash int64 // 模式字符串的散列值
}

// NewRabinKarpSearch 根据
func NewRabinKarpSearch(pattern string) rabinKarpSearch {
	M := len(pattern)

	var (
		RM int64 = 1
	)
	for i := 1; i < M; i++ { //计算R^(M-1)%Q，用于减去第一个数字时的计算
		RM = (R * RM) % q
	}

	return rabinKarpSearch{M, RM, hash(pattern, M)}
}

// hash 计算key[0..M-1]的散列值
func hash(key string, M int) int64 {
	var h int64
	for j := 0; j < M; j++ { // Horner法除k取余
		h = (R*h + int64(key[j])) % q
	}
	return h
}

// Search 在txt中查找字符串
func (r rabinKarpSearch) Search(txt string) int {
	N := len(txt)
	if r.m == 0 {
		return 0
	} else if N == 0 || N < r.m {
		return -1
	}

	//check 验证正确性（蒙特卡洛法）
	var check = func(i int) bool {
		return true
	}

	txtHash := hash(txt, r.m)
	if r.patternHash == txtHash && check(0) { // 一开始就匹配成功
		return 0
	}

	for i := r.m; i < N; i++ {
		txtHash = (txtHash + q - r.rm*int64(txt[i-r.m])%q) % q // 减去第一个数字，加上一个Q以保证所得值为正
		txtHash = (txtHash*R + int64(txt[i])) % q              // 加上最后一个数字

		if index := i - r.m + 1; r.patternHash == txtHash && check(index) { // 找到匹配
			return index
		}
	}

	return -1
}
