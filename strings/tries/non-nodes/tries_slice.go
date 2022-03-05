/*
 * File: \strings\tries\non-nodes\tries_slice.go                               *
 * Project: algorithms                                                         *
 * Created At: Wednesday, 2022/02/16 , 01:29:08                                *
 * Author: elchn                                                               *
 * -----                                                                       *
 * Last Modified: Saturday, 2022/03/5 , 00:18:34                               *
 * Modified By: elchn                                                          *
 * -----                                                                       *
 * HISTORY:                                                                    *
 * Date      	By	Comments                                                   *
 * ----------	---	---------------------------------------------------------  *
 */

package tries

// find returns a sub-tries whose root matches ch
func (t *tries_slice) find(ch rune) *tries_slice {
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

// tries_slice slice based tries
type tries_slice struct {
	children []*tries_slice
	char     rune
	val      any
	size     int
}

// NewTriesArr constructor
func NewTriesSlice() *tries_slice {
	return &tries_slice{
		children: make([]*tries_slice, 0),
	}
}

// Put inserts a pair of key and value into the tries
func (t *tries_slice) Put(key string, val any) {
	if key == "" {
		t.val = val
		t.size++
	}
	cur := t
	for i, c := range key {
		if child := cur.find(c); child == nil {
			child = &tries_slice{
				children: make([]*tries_slice, 0),
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
func (t *tries_slice) Get(key string) any {
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
func (t tries_slice) Size() int {
	return t.size
}

// Delete removes key and its value from the tries
func (t *tries_slice) Delete(key string) {

}

// Contains returns if their is a value paired with key in the tries
func (t *tries_slice) Contains(key string) bool {
	if key == "" {
		return t.val != nil
	}
	cur := t
	for i, c := range key {
		if k := cur.find(c); k == nil {
			return false
		} else {
			cur = k
		}
		if i == len(key)-1 && cur.val != nil {
			return true
		}
	}
	return false
}

// Keys gets all strings in the tries
func (t *tries_slice) Keys() (res []string) {
	res = t.KeysWithPrefix("")

	return
}

// KeysWithPrefix returns all the keys having pre as prefix
func (t *tries_slice) KeysWithPrefix(pre string) (res []string) {
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
	var collect func(x *tries_slice, key string) []string

	collect = func(x *tries_slice, key string) []string {
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
func (t *tries_slice) KeysThatMatch(pattern string) (res []string) {
	res = []string{}

	if t == nil {
		return
	}

	// collect
	var collect func(x *tries_slice, key, pattern string) []string

	collect = func(x *tries_slice, key, pattern string) []string {
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
func (t *tries_slice) LongestPrefixOf(pre string) string {
	//Todo
	return ""
}
