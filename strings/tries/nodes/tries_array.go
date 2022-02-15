// CreatedAt: 2022-02-13 18:29:23
// LastEditors: Elon Chen
// LastEditTime: 2022-02-15 18:58:25
// RelativePath: \algorithms\strings\tries\non-nodes\tries_array.go

// package tries includes all kinds of implementation of tries
package tries

import "algorithms/strings"

// tries_arr 基于数组的字典树结构
type tries_arr struct {
	children [strings.R]*tries_arr
	val      any
	size     int
}

// NewTriesArr 构造函数
func NewTriesArr() *tries_arr {
	return &tries_arr{}
}

// Put 插入键值对
func (t *tries_arr) Put(key string, val any) {
	if key == "" {
		t.val = val
		t.size++
	}
	cur := t
	for i, c := range key {
		if cur.children[c] == nil {
			cur.children[c] = new(tries_arr)
		}

		cur = cur.children[c]

		if i == len(key)-1 {
			cur.val = val
			t.size++
		}
	}
}

// Get 查找字符串对应的值，如果不存在，则返回nil
func (t *tries_arr) Get(key string) any {
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
func (t tries_arr) Size() int {
	return t.size
}

// Contains 表中是否存在键为key的值
func (t *tries_arr) Contains(key string) bool {
	if key == "" {
		return t.val != nil
	}
	cur := t
	for _, c := range key {
		if cur.children[c] == nil {
			return false
		}
		cur = cur.children[c]
		if cur.val != nil {
			return true
		}
	}
	return key == t.val
}

// Keys 返回所有键的列表
func (t *tries_arr) Keys() (res []string) {
	res = t.KeysWithPrefix("")

	return
}

// KeysWithPrefix
func (t *tries_arr) KeysWithPrefix(pre string) (res []string) {
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
	res = append(res, t.collect(pre, ".")...)

	return
}

// KeysThatMatch
func (t *tries_arr) KeysThatMatch(pattern string) (res []string) {
	res = append(res, t.collect("", pattern)...)
	return
}

// collect
func (t *tries_arr) collect(key, pattern string) []string {
	res := []string{}

	if t == nil {
		return res
	}

	d := len(key)
	f := len(pattern)
	if f == d {
		if t.val != nil {
			res = append(res, key)
			return res
		}
	} else {
		if f < d {
			return res
		}
	}

	for i, c := range t.children {
		if pattern[d] == '.' || int(pattern[d]) == i {
			res = append(res, c.collect(key+string(rune(i)), pattern)...)
		}
	}

	return res
}
