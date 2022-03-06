/*
 * File: \strings\tries\non-nodes\tries_map.go                                 *
 * Project: algorithms                                                         *
 * Created At: Tuesday, 2022/02/15 , 16:44:46                                  *
 * Author: elchn                                                               *
 * -----                                                                       *
 * Last Modified: Monday, 2022/03/7 , 00:05:57                                 *
 * Modified By: elchn                                                          *
 * -----                                                                       *
 * HISTORY:                                                                    *
 * Date      	By	Comments                                                   *
 * ----------	---	---------------------------------------------------------  *
 */

package tries

// tries_map map based tries
type tries_map struct {
	children map[rune]*tries_map
	val      any
	size     int
}

// NewTriesArr constructor
func NewTriesMap() *tries_map {
	return &tries_map{
		children: make(map[rune]*tries_map),
	}
}

// Put 插入键值对
func (t *tries_map) Put(key string, val any) {
	if key == "" {
		t.val = val
		t.size++
	}
	cur := t
	for i, c := range key {
		if _, ok := cur.children[c]; !ok {
			cur.children[c] = &tries_map{
				children: make(map[rune]*tries_map),
			}
		}

		cur = cur.children[c]

		if i == len(key)-1 {
			cur.val = val
			t.size++
		}
	}
}

// Get 查找字符串对应的值，如果不存在，则返回nil
func (t *tries_map) Get(key string) any {
	cur := t
	for _, c := range key {
		cur = cur.children[c]

		if cur.val != nil && byte(c) == key[len(key)-1] {
			return cur.val
		}
	}
	return t.val
}

// Size 获取键值对的数量
func (t tries_map) Size() int {
	return t.size
}

// Contains 表中是否存在键为key的值
func (t *tries_map) Contains(key string) bool {
	if key == "" {
		return t.val != nil
	}
	cur := t
	for i, c := range key {
		if k := cur.children[c]; k == nil {
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

// Keys 返回所有键的列表
func (t *tries_map) Keys() (res []string) {
	res = t.KeysWithPrefix("")

	return
}

// KeysWithPrefix returns all the keys having pre as prefix
func (t *tries_map) KeysWithPrefix(pre string) (res []string) {
	res = []string{}
	if t == nil {
		return
	}

	for i, c := range pre {
		if b := t.children[c]; b != nil {
			t = b
		}
		if t.val != nil && i == len(pre)-1 {
			res = append(res, pre)
		}
	}
	// collect
	var collect func(x *tries_map, key string) []string

	collect = func(x *tries_map, key string) []string {
		strs := []string{}

		if x == nil {
			return strs
		}

		if x.val != nil {
			strs = append(strs, key)
		}

		for i, c := range x.children {
			if c != nil {
				strs = append(strs, collect(c, key+string(i))...)
			}
		}
		return strs
	}

	res = append(res, collect(t, pre)...)

	return
}

// KeysThatMatch returns all the keys that matches pattern,where '.' matches any byte
func (t *tries_map) KeysThatMatch(pattern string) (res []string) {
	res = []string{}

	if t == nil {
		return
	}

	// collect
	var collect func(x *tries_map, key, pattern string) []string

	collect = func(x *tries_map, key, pattern string) []string {
		strs := []string{}

		if x == nil {
			return res
		}

		if kl, pl := len(key), len(pattern); kl == pl && x.val != nil {
			strs = append(strs, key)
		} else if kl < pl {
			for i, c := range x.children {
				if pattern[kl] == '.' || rune(pattern[kl]) == i {
					strs = append(strs, collect(c, key+string(rune(i)), pattern)...)
				}
			}
		}

		return strs
	}

	res = append(res, collect(t, "", pattern)...)

	return
}

// LongestPrefixOf returns the longest key that has a prefix of pre
func (t *tries_map) LongestPrefixOf(pre string) string {
	//Todo
	var search func(x *tries_map, s string, d, length int) int

	search = func(x *tries_map, s string, d, length int) int {
		if x == nil {
			return length
		}

		if x.val != nil {
			length = d
		}

		if d == len(s) {
			return length
		}

		return search(x.children[rune(s[d])], s, d+1, length)
	}
	return pre[0:search(t, pre, 0, 0)]
}
