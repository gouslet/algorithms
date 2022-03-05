/*
 * File: \tree\node.go                                                          *
 * Project: algorithms                                                         *
 * Created At: Saturday, 2022/03/5 , 00:26:50                                  *
 * Author: elchn                                                               *
 * -----                                                                       *
 * Last Modified: Saturday, 2022/03/5 , 21:54:39                               *
 * Modified By: elchn                                                          *
 * -----                                                                       *
 * HISTORY:                                                                    *
 * Date      	By	Comments                                                   *
 * ----------	---	---------------------------------------------------------  *
 */

package tree

type Key interface {
	// compareTo(key Key) int
	comparable
}
type BST[K Key] struct {
	root    *node[K]
	compare func(K, K) int
}

type Value interface {
}

type element[K Key] struct {
	key K
	val Value
}

type node[K Key] struct {
	element[K]
	left, right *node[K] //left and right subtrees
	n           int      // 子树的结点总数
}

func NewBST[K comparable](cmp func(K, K) int) *BST[K] {
	return &BST[K]{compare: cmp, root: nil}
}

// Size returns the number of elements in a binary search tree
func (b *BST[K]) Size() int {
	if b == nil {
		return 0
	}
	return b.root.size()
}

func (n *node[K]) size() int {
	if n == nil {
		return 0
	}
	return n.n
}

// Put an element into the binary search tree
func (b *BST[K]) Put(key K, val Value) {
	var put func(x *node[K], key K, val Value) *node[K]

	put = func(x *node[K], key K, val Value) *node[K] {
		if x == nil {
			x = &node[K]{
				element[K]{
					key,
					val,
				}, nil, nil, 1,
			}

			return x
		}

		if cmp := b.compare(key, x.key); cmp < 0 {
			x.left = put(x.left, key, val)
		} else if cmp > 0 {
			x.right = put(x.right, key, val)
		} else {
			x.val = val
		}

		x.n = x.left.size() + x.right.size() + 1

		return x
	}

	b.root = put(b.root, key, val)

	return
}

// Get the value paired with key,returning nil if absent
func (b *BST[K]) Get(key K) Value {
	var get func(x *node[K], key K) Value

	get = func(x *node[K], key K) Value {
		if x == nil {
			return nil
		}

		if cmp := b.compare(key, x.key); cmp < 0 {
			return get(x.left, key)
		} else if cmp > 0 {
			return get(x.right, key)
		} else {
			return x.val
		}
	}

	return get(b.root, key)
}

func (b *BST[K]) Contains(key K) bool {
	return b.Get(key) != nil
}

func (b *BST[K]) Delete(key string) {
}
