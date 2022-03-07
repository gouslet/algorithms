/*
 * File: \strings\tries\non-nodes\tries_slice_test.go                          *
 * Project: algorithms                                                         *
 * Created At: Wednesday, 2022/02/16 , 01:56:01                                *
 * Author: elchn                                                               *
 * -----                                                                       *
 * Last Modified: Monday, 2022/03/7 , 22:01:23                                 *
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

func TestTriesSlice(t *testing.T) {
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
				"by the sea": "by",
				"shell":      "she",
				"shells":     "shells",
				"":           "",
			},
		},
		{
			map[string]int{
				"":     1,
				" ":    2,
				"sea":  3,
				"_":    4,
				"'\"":  5,
				"\\`~": 6,
			},
			nil,
			nil,
			nil,
		},
	}

	for _, test := range tests {
		tries_slice := NewTriesMap()
		var keys []string
		for k, v := range test.pairs {
			tries_slice.Put(k, v)
			t.Logf("Put(%q,%v),size = %d\n", k, v, tries_slice.Size())
			if b := tries_slice.Contains(k); !b {
				t.Errorf("Contains(%q) = %v,want true", k, b)
			}
			if b := tries_slice.Contains(k + "#"); b {
				t.Errorf("Contains(%q) = %v,want false", k, b)
			}
			if b := tries_slice.Get(k); b != v {
				t.Errorf("Get(%q) = %v,want %v", k, b, v)
			}
			keys = append(keys, k)
		}
		sort.Strings(keys)
		b := tries_slice.Keys()
		sort.Strings(b)
		if !reflect.DeepEqual(b, keys) {
			t.Errorf("Keys() = %v,want %v", b, keys)
		}
		for pre, strs := range test.prefix_pairs {
			sort.Strings(strs)
			b := tries_slice.KeysWithPrefix(pre)
			sort.Strings(b)
			if !reflect.DeepEqual(b, strs) {
				t.Errorf("KeysWithPrefix(%q) = %v,want %v", pre, b, strs)
			}
		}

		for wpre, strs := range test.wildcard_pairs {
			sort.Strings(strs)
			b := tries_slice.KeysThatMatch(wpre)
			sort.Strings(b)
			if !reflect.DeepEqual(b, strs) {
				t.Errorf("KeysThatMatch(%q) = %v,want %v", wpre, b, strs)
			}
		}

		for pre, fix := range test.longestPrefixOf {
			if lpf := tries_slice.LongestPrefixOf(pre); lpf != fix {
				t.Errorf("LongestPrefixOf(%q) = %v,want %v", pre, lpf, fix)
			}
		}

		for k, v := range test.pairs {
			if del := tries_slice.Delete(k); del != v {
				t.Errorf("Delete(%q) = %v,want %v", k, del, v)
			}
		}
	}
}

// func ExampleTrriesslice() {
// 	tests := []struct {
// 		pairs       map[string]int
// 		prefix_pars map[string][]string
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
// 		tries_slice := NewTriesSlice()
// 		var keys []string
// 		for k, v := range test.pairs {
// 			tries_slice.Put(k, v)
// 			fmt.Printf("Put(%q,%v),size = %d\n", k, v, tries_slice.Size())
// 			fmt.Printf("Contains(%q) = %v,want true\n", k, tries_slice.Contains(k))
// 			keys = append(keys, k)
// 		}

// 		b := tries_slice.Keys()
// 		fmt.Printf("Keys() = %v,want %v\n", b, keys)

// 		for pre, strs := range test.prefix_pars {
// 			fmt.Printf("KeysWithPrefix(%q) = %v,want %v\n", pre, b, strs)
// 		}
// 		fmt.Println("------------------------------")
// 	}
// 	// Output:
// 	// 20
// }
