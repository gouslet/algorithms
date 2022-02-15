// CreatedAt: 2022-02-15 16:44:46
// LastEditors: Elon Chen
// LastEditTime: 2022-02-15 19:40:37
// RelativePath: \algorithms\strings\tries\non-nodes\tries_map.go

package tries

import "algorithms/strings"

// tries_map 基于数组的字典树结构
type tries_map struct {
	children [strings.R]*tries_map
	val      any
	size     int
}

// NewTriesArr 构造函数
func NewTrieMap() *tries_map {
	return &tries_map{}
}

// Put 插入键值对
func (t *tries_map) Put(key string, val any) {
	if key == "" {
		t.val = val
		t.size++
	}
	cur := t
	for i, c := range key {
		if cur.children[c] == nil {
			cur.children[c] = new(tries_map)
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
func (t *tries_map) Keys() (res []string) {
	res = t.KeysWithPrefix("")

	return
}

func (t *tries_map) KeysWithPrefix(pre string) (res []string) {
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

func (t *tries_map) KeysThatMatch(pattern string) (res []string) {
	res = append(res, t.collect("", pattern)...)
	return
}

func (t *tries_map) collect(key, pattern string) []string {
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
