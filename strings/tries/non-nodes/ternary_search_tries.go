// CreatedAt: 2022-02-13 18:29:23
// LastEditors: Elon Chen
// LastEditTime: 2022-02-15 18:58:25
// RelativePath: \algorithms\strings\tries\non-nodes\tstay.go

// package tries includes all kinds of implementation of tries
package tries

// ternary search tries
type tst struct {
	left, mid, right *tst
	char             rune
	val              any
	size             int
}

// NewTriesArr constructor
func NewTST() *tst {
	return &tst{}
}

// Put inserts a pair of key and value into the tries
func (t *tst) Put(key string, val any) {
	t.put(key, val, 0)
	t.size++
}

func (t *tst) put(key string, val any, d int) *tst {
	if key == "" {
		t.val = val
		return t
	}
	ch := rune(key[d])
	if t == nil {
		t = &tst{
			char: ch,
		}
	}
	if ch < t.char {
		t.left = t.left.put(key, val, d)
	} else if ch > t.char {
		t.right = t.right.put(key, val, d)
	} else if d < len(key)-1 {
		t.mid = t.mid.put(key, val, d+1)
	} else {
		t.val = val
	}

	return t
}

// Get returns the value paired with a key,or nil if absent
func (t *tst) Get(key string) any {
	x := t.get(key, 0)
	if x == nil {
		return nil
	}
	return x.val
}

func (t *tst) get(key string, d int) *tst {
	if key == "" {
		return t
	}

	if t == nil {
		return nil
	}

	ch := rune(key[d])
	if ch > t.char {
		return t.right.get(key, d)
	} else if ch < t.char {
		return t.left.get(key, d)
	} else if d < len(key)-1 {
		return t.mid.get(key, d+1)
	} else {
		return t
	}
}

// Size returns the number of key-value pairs in the tries
func (t tst) Size() int {
	return t.size
}

// Delete removes key and its value from the tries
func (t *tst) Delete(key string) {

}

// Contains returns if their is a value paired with key in the tries
func (t *tst) Contains(key string) bool {
	if key == "" {
		return t.val != nil
	}
	cur := t
	for i := 0; i < len(key); {
		c := rune(key[i])
		if ch := cur.char; ch < c {
			if r := cur.right; r == nil {
				return false
			}
			cur = cur.right
		} else if ch > c {
			if l := cur.left; l == nil {
				return false
			}
			cur = cur.left
		} else {
			if i == len(key)-1 && cur.val != nil {
				return true
			}
			if m := cur.mid; m == nil {
				return false
			}
			cur = cur.mid
			i++
		}

	}
	return false
}

// Keys gets all strings in the tries
func (t *tst) Keys() (res []string) {
	res = t.KeysWithPrefix("")

	return
}

// KeysWithPrefix returns all the keys having pre as prefix
func (t *tst) KeysWithPrefix(pre string) []string {
	res := []string{}
	if t == nil {
		return res
	}
	if pre == "" {
		if t.char == 0 && t.val != nil {
			res = append(res, "")
		}
	}

	for i, l := 0, len(pre); i < l; {
		b := rune(pre[i])
		if t == nil {
			break
		}

		if t.char > b {
			t = t.left
		} else if t.char < b {
			t = t.right
		} else {
			if t.val != nil && i == len(pre)-1 {
				res = append(res, pre)
			}
			t = t.mid
			i++
		}

	}
	// collect
	var collect func(x *tst, key string) []string

	collect = func(x *tst, key string) []string {
		strs := []string{}

		if x == nil {
			return strs
		}
		s := key
		if ch := x.char; ch != 0 {
			s += string(ch)
			if x.val != nil {
				strs = append(strs, s)
			}
		}
		strs = append(strs, collect(x.mid, s)...)
		strs = append(strs, collect(x.left, key)...)
		strs = append(strs, collect(x.right, key)...)
		return strs
	}

	res = append(res, collect(t, pre)...)

	return res
}

// KeysThatMatch returns all the keys that matches pattern,where '.' matches any byte
func (t *tst) KeysThatMatch(pattern string) (res []string) {
	res = []string{}

	if t == nil {
		return
	}

	// collect
	var collect func(x *tst, key, pattern string) []string

	collect = func(x *tst, key, pattern string) []string {
		strs := []string{}

		if x == nil {
			return res
		}

		if kl, pl := len(key), len(pattern); kl == pl && x.val != nil {
			strs = append(strs, key)
		} else if kl < pl {
			if pattern[kl] == '.' || int(pattern[kl]) == 1 {
				strs = append(strs, collect(x, key, pattern)...)
			}
		}

		return strs
	}

	res = append(res, collect(t, "", pattern)...)

	return
}

// LongestPrefixOf returns the longest key that has a prefix of pre
func (t *tst) LongestPrefixOf(pre string) string {
	//Todo
	return ""
}
