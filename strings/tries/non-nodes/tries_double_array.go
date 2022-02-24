// CreatedAt: 2022-02-13 18:29:23
// LastEditors: Elon Chen
// LastEditTime: 2022-02-15 18:58:25
// RelativePath: \algorithms\strings\tries\non-nodes\tries_double_array.go

// package tries includes all kinds of implementation of tries
package tries

// find returns a sub-tries whose root matches ch
func (t *tries_double_array) find(ch rune) *tries_double_array {
	if t == nil || t.children == nil {
		return nil
	}

	for _, x := range t.children {
		if x.char == ch {
			return x
		}
	}
	return nil
}

// tries_double_array slice based tries
type tries_double_array struct {
	children []*tries_double_array
	char     rune
	val      any
	size     int
}

// NewTriesArr constructor
func NewTriesDoubleArray() *tries_double_array {
	return &tries_double_array{
		children: make([]*tries_double_array, 0),
	}
}

// Put inserts a pair of key and value into the tries
func (t *tries_double_array) Put(key string, val any) {
	if key == "" {
		t.val = val
		t.size++
	}
	cur := t
	for i, c := range key {
		if child := cur.find(c); child == nil {
			child = &tries_double_array{
				children: make([]*tries_double_array, 0),
				char:     c,
			}
			cur.children = append(cur.children, child)
		} else {
			cur = child
		}

		if i == len(key)-1 {
			cur.val = val
			t.size++
		}
	}
}

// Get returns the value paired with a key,or nil if absent
func (t *tries_double_array) Get(key string) any {
	cur := t
	for _, c := range key {
		cur = cur.find(c)

		if cur.val != nil && byte(c) == key[len(key)-1] {
			return cur.val
		}
	}
	return t.val
}

// Size returns the number of key-value pairs in the tries
func (t tries_double_array) Size() int {
	return t.size
}

// Delete removes key and its value from the tries
func (t *tries_double_array) Delete(key string) {

}

// Contains returns if their is a value paired with key in the tries
func (t *tries_double_array) Contains(key string) bool {
	if key == "" {
		return t.val != nil
	}
	cur := t
	for _, c := range key {
		if k := cur.find(c); k == nil {
			return false
		} else {
			cur = k
		}
		if cur.val != nil {
			return true
		}
	}
	return key == t.val
}

// Keys gets all strings in the tries
func (t *tries_double_array) Keys() (res []string) {
	res = t.KeysWithPrefix("")

	return
}

// KeysWithPrefix returns all the keys having pre as prefix
func (t *tries_double_array) KeysWithPrefix(pre string) (res []string) {
	res = []string{}
	if t == nil {
		return
	}

	for i, c := range pre {
		if b := t.find(c); b != nil {
			t = b
		}
		if t.val != nil && i == len(pre)-1 {
			res = append(res, pre)
		}
	}
	// collect
	var collect func(x *tries_double_array, key string) []string

	collect = func(x *tries_double_array, key string) []string {
		strs := []string{}

		if x == nil {
			return strs
		}

		if x.val != nil {
			strs = append(strs, key)
		}

		for _, c := range x.children {
			if c != nil {
				strs = append(strs, collect(c, key+string(c.char))...)
			}
		}
		return strs
	}

	res = append(res, collect(t, pre)...)

	return
}

// KeysThatMatch returns all the keys that matches pattern,where '.' matches any byte
func (t *tries_double_array) KeysThatMatch(pattern string) (res []string) {
	res = []string{}

	if t == nil {
		return
	}

	// collect
	var collect func(x *tries_double_array, key, pattern string) []string

	collect = func(x *tries_double_array, key, pattern string) []string {
		strs := []string{}

		if x == nil {
			return res
		}

		if kl, pl := len(key), len(pattern); kl == pl && x.val != nil {
			strs = append(strs, key)
		} else if kl < pl {
			for _, c := range x.children {
				if pattern[kl] == '.' || rune(pattern[kl]) == c.char {
					strs = append(strs, collect(c, key+string(c.char), pattern)...)
				}
			}
		}

		return strs
	}

	res = append(res, collect(t, "", pattern)...)

	return
}

// LongestPrefixOf returns the longest key that has a prefix of pre
func (t *tries_double_array) LongestPrefixOf(pre string) string {
	//Todo
	return ""
}
