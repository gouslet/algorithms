/*
 * File: \strings\tries\non-nodes\ternary_search_tries_test.go                 *
 * Project: algorithms                                                         *
 * Created At: Wednesday, 2022/02/16 , 22:25:07                                *
 * Author: elchn                                                               *
 * -----                                                                       *
 * Last Modified: Monday, 2022/03/7 , 00:58:30                                 *
 * Modified By: elchn                                                          *
 * -----                                                                       *
 * HISTORY:                                                                    *
 * Date      	By	Comments                                                   *
 * ----------	---	---------------------------------------------------------  *
 */

package tries

import (
	"reflect"
	"sort"
	"testing"
)

func TestTernarySearchTries(t *testing.T) {
	tests := []struct {
		pairs           map[string]int
		prefix_pairs    map[string][]string
		wildcard_pairs  map[string][]string
		longestPrefixOf map[string]string
	}{
		{
			map[string]int{
				"sells":  1,
				"shells": 2,
				"sea":    3,
				"shore":  4,
				"by":     5,
				"she":    6,
				"the":    7,
			},
			map[string][]string{
				"": {
					"sells", "shells", "she", "sea", "shore", "by", "the",
				},
				"b": {
					"by",
				},
				"s": {
					"sells", "shells", "she", "sea", "shore",
				},
				"sh": {
					"shells", "she", "shore",
				},
			},
			map[string][]string{
				".": {},
				"b.": {
					"by",
				},
				".he": {
					"the", "she",
				},
				"s..": {
					"she", "sea",
				},
			},
			map[string]string{
				"by the sea":   "by",
				"shell":        "she",
				"shells":       "shells",
				"shoresdfdgfh": "shore",
				"":             "",
			},
		},
		{
			map[string]int{
				"":         1,
				" ":        2,
				"sea":      3,
				"_":        4,
				"'\"":      5,
				"\\`~":     6,
				"./&_+-*%": 7,
			},
			map[string][]string{
				"": {
					"",
					" ",
					"sea",
					"_",
					"'\"",
					"\\`~",
					"./&_+-*%",
				},
				" ": {
					" ",
				},
				"s": {
					"sea",
				},
				"_": {
					"_",
				},
				"'": {
					"'\"",
				},
			},
			map[string][]string{
				"": {
					"",
				},
				".": {
					" ",
					"_",
				},
				"..": {
					"'\"",
				},
				"...": {
					"\\`~",
					"sea",
				},
				"_": {
					"_",
				},
				"s..": {
					"sea",
				},
				"'.": {
					"'\"",
				},
			},
			map[string]string{
				"sea":    "sea",
				"'\"'\"": "'\"",
				"":       "",
			},
		},
	}

	for _, test := range tests {
		tst := NewTST()
		var keys []string
		for k, v := range test.pairs {
			tst.Put(k, v)
			t.Logf("Put(%q,%v),size = %d\n", k, v, tst.Size())
			if b := tst.Contains(k); !b {
				t.Errorf("Contains(%q) = %v,want true", k, b)
			}
			if b := tst.Contains(k + "#"); b {
				t.Errorf("Contains(%q) = %v,want false", k+"#", b)
			}
			if b := tst.Get(k); b != v {
				t.Errorf("Get(%q) = %v,want %v", k, b, v)
			}
			keys = append(keys, k)

			t.Run("Keys", func(t *testing.T) {
				sort.Strings(keys)
				b := tst.Keys()
				sort.Strings(b)
				if !reflect.DeepEqual(b, keys) {
					t.Errorf("Keys() = %v,want %v", b, keys)
				}
			})

		}
		t.Run("KeysWithPrefix", func(t *testing.T) {
			for pre, strs := range test.prefix_pairs {
				sort.Strings(strs)
				b := tst.KeysWithPrefix(pre)
				sort.Strings(b)
				if !reflect.DeepEqual(b, strs) {
					t.Errorf("KeysWithPrefix(%q) = %v,want %v", pre, b, strs)
				}
			}
		})
		t.Run("KeysThatMatch", func(t *testing.T) {
			for wpre, strs := range test.wildcard_pairs {
				sort.Strings(strs)
				b := tst.KeysThatMatch(wpre)
				sort.Strings(b)
				if !reflect.DeepEqual(b, strs) {
					t.Errorf("KeysThatMatch(%q) = %v,want %v", wpre, b, strs)
				}
			}
		})
		t.Run("LongestPrefixOf", func(t *testing.T) {
			for pre, fix := range test.longestPrefixOf {
				if lpf := tst.LongestPrefixOf(pre); lpf != fix {
					t.Errorf("LongestPrefixOf(%q) = %v,want %v", pre, lpf, fix)
				}
			}
		})
	}
}

// func ExampleTST() {
// 	tests := []struct {
// 		pairs        map[string]int
// 		prefix_pairs map[string][]string
// 	}{
// 		{
// 			map[string]int{
// 				"sells":  1,
// 				"shells": 2,
// 				"sea":    3,
// 				"shore":  4,
// 				"by":     5,
// 				"she":    6,
// 			},
// 			map[string][]string{
// 				"": {
// 					"sells", "shells", "she", "sea", "shore", "by",
// 				},
// 				"b": {
// 					"by",
// 				},
// 				"s": {
// 					"sells", "shells", "she", "sea", "shore",
// 				},
// 				"sh": {
// 					"shells", "she", "shore",
// 				},
// 			},
// 		},
// 		{
// 			map[string]int{
// 				"":     1,
// 				" ":    2,
// 				"sea":  3,
// 				"_":    4,
// 				"'\"":  5,
// 				"\\`~": 6,
// 			},
// 			nil,
// 		},
// 	}

// 	for _, test := range tests {
// 		tst := NewTST()
// 		var keys []string
// 		for k, v := range test.pairs {
// 			tst.Put(k, v)
// 			fmt.Printf("Put(%q,%v),size = %d\n", k, v, tst.Size())
// 			fmt.Printf("Contains(%q) = %v,want true\n", k, tst.Contains(k))
// 			keys = append(keys, k)
// 		}

// 		b := tst.Keys()
// 		fmt.Printf("Keys() = %v,want %v\n", b, keys)

// 		for pre, strs := range test.prefix_pairs {
// 			fmt.Printf("KeysWithPrefix(%q) = %v,want %v\n", pre, b, strs)
// 		}
// 		fmt.Println("------------------------------")
// 	}
// 	// Output:
// 	// 20
// }
