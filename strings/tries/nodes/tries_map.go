package tries

import "algorithms/strings"

type node struct {
	size int
	next [strings.R]*node
	val  any
}

type tries struct {
	root *node
}

func NewTries() *tries {
	var root *node

	return &tries{root}
}

func (t tries) Size() int {
	return t.size(t.root)
}

// 单词查找树的延时递归size方法
func (t tries) size(x *node) int {
	if x == nil { //
		return 0
	}

	cnt := 0
	if x.val == nil {
		cnt++
	}

	for c := 0; c < R; c++ {
		cnt += t.size(x.next[c])
	}

	return cnt
}

// // 单词查找树的即时计算size方法
// func (t tries) size(x *node) int {
// 	if x == nil { //
// 		return 0
// 	}

// 	return t.root.size
// }

func (t tries) Get(key string) any {
	x := t.get(t.root, key, 0)
	if x == nil {
		return nil
	}
	return x.val
}

// get 返回以x作为根结点的子单词查找树中与key相关联的值
func (t tries) get(x *node, key string, d int) *node {
	if x == nil || d == len(key) {
		return x
	}

	return t.get(x.next[key[d]], key, d+1)
}

func (t *tries) Put(key string, val any) {
	t.root = t.put(t.root, key, val, 0)
}

// put 如果key存在于以x为根结点的子单词查找树中则更新与它相关联的值
func (t *tries) put(x *node, key string, val any, d int) *node {
	if x == nil {
		x = new(node)
		x.size++
	}

	if d == len(key) {
		x.val = val
		return x
	}

	x.next[key[d]] = t.put(x.next[key[d]], key, val, d+1)

	return x
}

type Iterator[K any] struct {
}

func (t tries) keys() Iterator[string] {
	return t.keysWithPrefix("")
}

func (t tries) keysWithPrefix(prefix string) Iterator[string] {
	var ret Iterator[string]
	return ret
}

func (t *tries) Delete(key string) {
	t.root = t.delete(t.root, key, 0)
}

func (t *tries) delete(x *node, key string, d int) *node {
	var r *node
	if x == nil {
		return r
	}

	if d == len(key) {
		x.val = nil
	} else {
		x.next[key[d]] = t.delete(x.next[key[d]], key, d+1)
	}

	if x.val != nil {
		return x
	}

	for c := 0; c < R; c++ {
		if x.next[c] != r {
			return x
		}
	}
	return r
}
