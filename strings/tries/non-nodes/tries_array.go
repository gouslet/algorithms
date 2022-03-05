/*
 * File: \strings\tries\non-nodes\tries_array.go                               *
 * Project: algorithms                                                         *
 * Created At: Sunday, 2022/02/13 , 18:29:23                                   *
 * Author: elchn                                                               *
 * -----                                                                       *
 * Last Modified: Saturday, 2022/03/5 , 00:16:32                               *
 * Modified By: elchn                                                          *
 * -----                                                                       *
 * HISTORY:                                                                    *
 * Date      	By	Comments                                                   *
 * ----------	---	---------------------------------------------------------  *
 */

package tries

import "algorithms/strings"

// tries_arr array based tries
type tries_arr struct {
	children [strings.R]*tries_arr
	val      any
	size     int
}

// NewTriesArr constructor
func NewTriesArr() *tries_arr {
	return &tries_arr{}
}

// Put inserts a pair of key and value into the tries
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

// Get returns the value paired with a key,or nil if absent
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

// Size returns the number of key-value pairs in the tries
func (t tries_arr) Size() int {
	return t.size
}

// Delete removes key and its value from the tries
func (t *tries_arr) Delete(key string) {

}

// Contains returns if their is a value paired with key in the tries
func (t *tries_arr) Contains(key string) bool {
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

// Keys gets all strings in the tries
func (t *tries_arr) Keys() (res []string) {
	res = t.KeysWithPrefix("")

	return
}

// KeysWithPrefix returns all the keys having pre as prefix
func (t *tries_arr) KeysWithPrefix(pre string) (res []string) {
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
	var collect func(x *tries_arr, key string) []string

	collect = func(x *tries_arr, key string) []string {
		strs := []string{}

		if x == nil {
			return strs
		}

		if x.val != nil {
			strs = append(strs, key)
		}

		for i, c := range x.children {
			if c != nil {
				strs = append(strs, collect(c, key+string(rune(i)))...)
			}
		}
		return strs
	}

	res = append(res, collect(t, pre)...)

	return
}

// KeysThatMatch returns all the keys that matches pattern,where '.' matches any byte
func (t *tries_arr) KeysThatMatch(pattern string) (res []string) {
	res = []string{}

	if t == nil {
		return
	}

	// collect
	var collect func(x *tries_arr, key, pattern string) []string

	collect = func(x *tries_arr, key, pattern string) []string {
		strs := []string{}

		if x == nil {
			return res
		}

		if kl, pl := len(key), len(pattern); kl == pl && x.val != nil {
			strs = append(strs, key)
		} else if kl < pl {
			for i, c := range x.children {
				if pattern[kl] == '.' || int(pattern[kl]) == i {
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
func (t *tries_arr) LongestPrefixOf(pre string) string {
	//Todo
	return ""
}
