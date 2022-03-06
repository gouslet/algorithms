/*
 * File: \tree\node.go                                                          *
 * Project: algorithms                                                         *
 * Created At: Saturday, 2022/03/5 , 00:26:50                                  *
 * Author: elchn                                                               *
 * -----                                                                       *
 * Last Modified: Sunday, 2022/03/6 , 22:46:52                                 *
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
	return size(b.root)
}

func size[K Key](n *node[K]) int {
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

		x.n = size(x.left) + size(x.right) + 1

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

// Contains
func (b *BST[K]) Contains(key K) bool {
	return b.Get(key) != nil
}

func min[K Key](n *node[K]) *node[K] {
	if n == nil || n.left == nil {
		return n
	}
	return min(n.left)
}

// Min
func (b *BST[K]) Min() K {

	if m := min(b.root); m != nil {
		return m.key
	}
	return *new(K)
}

// Max
func (b *BST[K]) Max() K {
	var max func(n *node[K]) *node[K]

	max = func(n *node[K]) *node[K] {
		if n == nil || n.right == nil {
			return n
		}
		return max(n.right)
	}
	if t := max(b.root); t != nil {
		return t.key
	}

	return *new(K)
}

// Floor
func (b *BST[K]) Floor(key K) K {

	var floor func(n *node[K], key K) *node[K]

	floor = func(n *node[K], key K) *node[K] {
		if n == nil {
			return nil
		}
		if cmp := b.compare(key, n.key); cmp < 0 {
			return floor(n.left, key)
		} else if cmp > 0 {
			if n.right == nil {
				return n
			}
			return floor(n.right, key)
		} else {
			return n
		}
	}

	t := floor(b.root, key)
	if t == nil {
		return *new(K)
	} else {
		return t.key
	}

}

// Ceil
func (b *BST[K]) Ceil(key K) K {

	var ceil func(n *node[K], key K) *node[K]

	ceil = func(n *node[K], key K) *node[K] {
		if n == nil {
			return nil
		}
		if cmp := b.compare(key, n.key); cmp > 0 {
			return ceil(n.right, key)
		} else if cmp < 0 {
			if n.left == nil {
				return n
			}
			return ceil(n.left, key)
		} else {
			return n
		}
	}

	t := ceil(b.root, key)
	if t == nil {
		return *new(K)
	} else {
		return t.key
	}

}

func (b *BST[K]) Select(k int) K {
	if k < 0 || k > b.Size() {
		return *new(K)
	}

	var selectk func(n *node[K], k int) *node[K]

	selectk = func(n *node[K], k int) *node[K] {
		if n == nil {
			return nil
		}
		if t := size(n.left); t > k {
			return selectk(n.left, k)
		} else if t < k {
			return selectk(n.right, k-t-1)
		} else {
			return n
		}
	}

	if t := selectk(b.root, k); t != nil {
		return t.key
	}

	return *new(K)
}

func (b *BST[K]) Rank(key K) int {
	if b == nil {
		return -1
	}

	var rank func(n *node[K], key K) int
	rank = func(n *node[K], key K) int {
		if n == nil {
			return -1
		}
		if cmp := b.compare(key, n.key); cmp < 0 {
			return rank(n.left, key)
		} else if cmp > 0 {
			return rank(n.right, key) + size(n.left) + 1
		} else {
			return size(n.left)
		}
	}

	return rank(b.root, key)
}

func (b *BST[K]) DeleteMax() K {
	var res K
	var deleteMax func(n *node[K]) *node[K]

	deleteMax = func(n *node[K]) *node[K] {
		if n == nil {
			return nil
		}

		if n.right == nil {
			res = n.key
			return n.left
		}

		n.right = deleteMax(n.right)
		n.n = size(n.left) + size(n.right) + 1
		return n
	}

	b.root = deleteMax(b.root)
	return res
}

func deleteMin[K Key](n *node[K]) (res *node[K], min element[K]) {

	if n == nil {
		return
	}

	if n.left == nil {
		res = n.right
		min = n.element
		return
	}

	n.left, _ = deleteMin(n.left)
	n.n = size(n.left) + size(n.right) + 1

	res = n
	return
}

func (b *BST[K]) DeleteMin() K {
	var min element[K]
	b.root, min = deleteMin(b.root)

	return min.key
}

func (b *BST[K]) Delete(key K) (res any) {
	if b == nil {
		return
	}

	var delete func(n *node[K], key K) (x *node[K], res element[K])

	delete = func(n *node[K], key K) (x *node[K], res element[K]) {
		if n == nil {
			return
		}
		if cmp := b.compare(key, n.key); cmp < 0 {
			n.left, _ = delete(n.left, key)
		} else if cmp > 0 {
			n.right, _ = delete(n.right, key)
		} else {
			res = n.element

			if n.right == nil {
				x = n.left
				return
			}
			if n.left == nil {
				x = n.right
				return
			}

			x = min(n.right)
			x.right, _ = deleteMin(n.right)
			x.left = n.left
		}
		x.n = size(x.left) + size(x.right) + 1

		return
	}
	var res_e element[K]
	b.root, res_e = delete(b.root, key)
	res = res_e.val
	return
}

func (b *BST[K]) Keys() []K {

	if b == nil || b.Size() == 0 {
		return make([]K, 0)
	}

	var keys func(n *node[K], lo, hi K) []K
	keys = func(n *node[K], lo, hi K) []K {
		res := []K{}

		if n == nil {
			return res
		}
		cmplo := b.compare(lo, n.key)
		cmphi := b.compare(hi, n.key)
		if cmplo < 0 {
			res = append(res, keys(n.left, lo, hi)...)
		}

		if cmplo <= 0 && cmphi >= 0 {
			res = append(res, n.key)
		}

		if cmphi < 0 {
			res = append(res, keys(n.right, lo, hi)...)
		}

		return res
	}

	return keys(b.root, b.Min(), b.Max())
}
