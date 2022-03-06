/*
 * File: \tree\bst_test.go                                                     *
 * Project: algorithms                                                         *
 * Created At: Saturday, 2022/03/5 , 14:21:45                                  *
 * Author: elchn                                                               *
 * -----                                                                       *
 * Last Modified: Sunday, 2022/03/6 , 22:54:07                                 *
 * Modified By: elchn                                                          *
 * -----                                                                       *
 * HISTORY:                                                                    *
 * Date      	By	Comments                                                   *
 * ----------	---	---------------------------------------------------------  *
 */

package tree

import (
	"reflect"
	"sort"
	"testing"
)

func TestPutAndGet(t *testing.T) {

	tests := []element[int]{
		{1, "a"},
		{2, "b"},
		{3, "c"},
		{4, "d"},
		{5, "e"},
	}

	bst := NewBST(func(a, b int) int {
		return a - b
	})

	for _, test := range tests {
		bst.Put(test.key, test.val)
		t.Logf("Put(%d, %s)", test.key, test.val)

	}

	for _, test := range tests {
		if v := bst.Get(test.key); v != test.val {
			t.Errorf("Get(%d) = %s,want %s", test.key, v, test.val)
		}
	}

}

func TestMinAndMax(t *testing.T) {

	tests := []struct {
		eles []element[int]
		min  int
		max  int
	}{
		{
			[]element[int]{
				{1, "a"},
				{2, "b"},
				{3, "c"},
				{4, "d"},
				{5, "e"},
				{6, "f"},
			},
			1,
			6,
		},
		{
			[]element[int]{},
			0,
			0,
		},
	}

	for _, test := range tests {
		bst := NewBST(
			func(a, b int) int {
				return a - b
			})
		for _, kv := range test.eles {
			bst.Put(kv.key, kv.val)
			t.Logf("Put(%d, %s)", kv.key, kv.val)
		}

		if min := bst.Min(); min != test.min {
			t.Errorf("Min() = %d,want %d", min, test.min)
		}

		if max := bst.Max(); max != test.max {
			t.Errorf("Max() = %d,want %d", max, test.max)
		}
	}
}

func TestFloorAndCeil(t *testing.T) {

	tests := []struct {
		eles      []element[int]
		key_floor map[int]int
		key_ceil  map[int]int
	}{
		{
			[]element[int]{
				{1, "a"},
				{2, "b"},
				{3, "c"},
				{4, "d"},
				{5, "e"},
				{6, "f"},
			},
			map[int]int{
				0: 0,
				1: 1,
				2: 2,
				3: 3,
				4: 4,
				6: 6,
				7: 6,
			},
			map[int]int{
				0: 1,
				1: 1,
				2: 2,
				3: 3,
				4: 4,
				6: 6,
				7: 0,
			},
		},
		{
			[]element[int]{},
			map[int]int{
				0: 0,
				1: 0,
				2: 0,
				3: 0,
				4: 0,
				6: 0,
				7: 0,
			},
			map[int]int{
				0: 0,
				1: 0,
				2: 0,
				3: 0,
				4: 0,
				6: 0,
				7: 0,
			},
		},
	}

	for _, test := range tests {
		bst := NewBST(
			func(a, b int) int {
				return a - b
			})
		for _, kv := range test.eles {
			bst.Put(kv.key, kv.val)
			t.Logf("Put(%d, %s)", kv.key, kv.val)
		}
		for k, f := range test.key_floor {
			if v := bst.Floor(k); v != f {
				t.Errorf("Floor(%d) = %d,want %d", k, v, f)
			}
		}
		for k, f := range test.key_ceil {
			if v := bst.Ceil(k); v != f {
				t.Errorf("Ceil(%d) = %d,want %d", k, v, f)
			}
		}
	}
}

func TestSelect(t *testing.T) {

	tests := []struct {
		eles []element[int]
		r_k  map[int]int
	}{
		{
			[]element[int]{
				{1, "a"},
				{2, "b"},
				{3, "c"},
				{4, "d"},
				{5, "e"},
				{6, "f"},
			},
			map[int]int{
				0:  1,
				1:  2,
				2:  3,
				3:  4,
				4:  5,
				5:  6,
				6:  0,
				-1: 0,
			},
		},
		{
			[]element[int]{},
			map[int]int{
				0:  0,
				-1: 0,
				1:  0,
			},
		},
	}

	for _, test := range tests {
		bst := NewBST(
			func(a, b int) int {
				return a - b
			})
		for _, kv := range test.eles {
			bst.Put(kv.key, kv.val)
			t.Logf("Put(%d, %s)", kv.key, kv.val)
		}
		for r, k := range test.r_k {
			if key := bst.Select(r); key != k {
				t.Errorf("Select(%d) = %d,want %d", r, key, k)
			}
		}

	}
}

func TestRank(t *testing.T) {

	tests := []struct {
		eles []element[int]
		k_r  map[int]int
	}{
		{
			[]element[int]{
				{1, "a"},
				{2, "b"},
				{3, "c"},
				{4, "d"},
				{5, "e"},
				{6, "f"},
			},
			map[int]int{
				1:  0,
				2:  1,
				3:  2,
				4:  3,
				5:  4,
				6:  5,
				-1: -1,
			},
		},
		{
			[]element[int]{},
			map[int]int{
				0:  -1,
				1:  -1,
				-1: -1,
			},
		},
	}

	for _, test := range tests {
		bst := NewBST(
			func(a, b int) int {
				return a - b
			})
		for _, kv := range test.eles {
			bst.Put(kv.key, kv.val)
			t.Logf("Put(%d, %s)", kv.key, kv.val)
		}
		for k, v := range test.k_r {
			if r := bst.Rank(k); r != v {
				t.Errorf("Rank(%d) = %d,want %d", k, r, v)
			}
		}

	}
}

func TestDeleteMinAndDeleteMax(t *testing.T) {

	tests := []struct {
		eles []element[int]
		min  int
		max  int
	}{
		{
			[]element[int]{
				{1, "a"},
				{2, "b"},
				{3, "c"},
				{4, "d"},
				{5, "e"},
				{6, "f"},
			},
			1,
			6,
		},
		{
			[]element[int]{},
			0,
			0,
		},
	}

	for _, test := range tests {
		bst := NewBST(
			func(a, b int) int {
				return a - b
			})
		for _, kv := range test.eles {
			bst.Put(kv.key, kv.val)
			t.Logf("Put(%d, %s)", kv.key, kv.val)
		}

		if max := bst.DeleteMax(); max != test.max {
			t.Errorf("Max() = %d,want %d", max, test.max)
		}

		if min := bst.DeleteMin(); min != test.min {
			t.Errorf("Min() = %d,want %d", min, test.min)
		}
	}
}

func TestDelete(t *testing.T) {

	tests := []struct {
		eles []element[int]
		kv   map[int]string
	}{
		{
			[]element[int]{
				{1, "a"},
				{2, "b"},
				{3, "c"},
				{4, "d"},
				{5, "e"},
				{6, "f"},
			},
			map[int]string{
				1: "a",
				2: "b",
				3: "c",
				4: "d",
				5: "e",
				6: "f",
			},
		},
		{
			[]element[int]{},
			nil,
		},
	}

	for _, test := range tests {
		bst := NewBST(
			func(a, b int) int {
				return a - b
			})
		for _, kv := range test.eles {
			bst.Put(kv.key, kv.val)
			t.Logf("Put(%d, %s)", kv.key, kv.val)
		}
		for _, kv := range test.eles {
			if v := bst.Delete(kv.key); v != kv.val {
				t.Errorf("Delete(%d) = %s,want %s", kv.key, v, kv.val)
			}
		}
	}
}

func TestKeys(t *testing.T) {

	tests := []struct {
		eles []element[int]
		keys []int
	}{
		{
			[]element[int]{
				{1, "a"},
				{2, "b"},
				{3, "c"},
				{4, "d"},
				{5, "e"},
				{6, "f"},
			},
			[]int{
				1, 2, 3, 4, 5, 6,
			},
		},
		{
			[]element[int]{},
			nil,
		},
	}

	for _, test := range tests {
		bst := NewBST(
			func(a, b int) int {
				return a - b
			})
		for _, kv := range test.eles {
			bst.Put(kv.key, kv.val)
			t.Logf("Put(%d, %s)", kv.key, kv.val)
		}
		keys := bst.Keys()
		sort.Ints(test.keys)
		sort.Ints(keys)
		if reflect.DeepEqual(keys, test.keys) {
			t.Errorf("Keys() = %v,want %v", keys, test.keys)
		}
	}
}
