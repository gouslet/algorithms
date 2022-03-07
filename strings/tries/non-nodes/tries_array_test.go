/*
 * File: \strings\tries\non-nodes\tries_array_test.go                          *
 * Project: algorithms                                                         *
 * Created At: Tuesday, 2022/02/15 , 16:07:07                                  *
 * Author: elchn                                                               *
 * -----                                                                       *
 * Last Modified: Monday, 2022/03/7 , 21:57:20                                 *
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

func TestTriesArr(t *testing.T) {
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
		tries_arr := NewTriesArr()
		var keys []string
		for k, v := range test.pairs {
			tries_arr.Put(k, v)
			t.Logf("Put(%q,%v),size = %d\n", k, v, tries_arr.Size())
			if b := tries_arr.Contains(k); !b {
				t.Errorf("Contains(%q) = %v,want true", k, b)
			}
			if b := tries_arr.Contains(k + "#"); b {
				t.Errorf("Contains(%q) = %v,want false", k, b)
			}
			if b := tries_arr.Get(k); b != v {
				t.Errorf("Get(%q) = %v,want %v", k, b, v)
			}
			keys = append(keys, k)
		}
		sort.Strings(keys)
		if b := tries_arr.Keys(); !reflect.DeepEqual(b, keys) {
			t.Errorf("Keys() = %v,want %v", b, keys)
		}
		for pre, strs := range test.prefix_pairs {
			sort.Strings(strs)
			if b := tries_arr.KeysWithPrefix(pre); !reflect.DeepEqual(b, strs) {
				t.Errorf("KeysWithPrefix(%q) = %v,want %v", pre, b, strs)
			}
		}

		for wpre, strs := range test.wildcard_pairs {
			sort.Strings(strs)
			if b := tries_arr.KeysThatMatch(wpre); !reflect.DeepEqual(b, strs) {
				t.Errorf("KeysThatMatch(%q) = %v,want %v", wpre, b, strs)
			}
		}

		for pre, fix := range test.longestPrefixOf {
			if lpf := tries_arr.LongestPrefixOf(pre); lpf != fix {
				t.Errorf("LongestPrefixOf(%q) = %v,want %v", pre, lpf, fix)
			}
		}

		for k, v := range test.pairs {
			if del := tries_arr.Delete(k); del != v {
				t.Errorf("Delete(%q) = %v,want %v", k, del, v)
			}
		}
	}
}

// func Example() {
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
// 		tries_arr := NewTriesArr()
// 		var keys []string
// 		for k, v := range test.pairs {
// 			tries_arr.Put(k, v)
// 			fmt.Printf("Put(%q,%v),size = %d\n", k, v, tries_arr.Size())
// 			fmt.Printf("Contains(%q) = %v,want true\n", k, tries_arr.Contains(k))
// 			keys = append(keys, k)
// 		}

// 		b := tries_arr.Keys()
// 		fmt.Printf("Keys() = %v,want %v\n", b, keys)

// 		for pre, strs := range test.prefix_pairs {
// 			fmt.Printf("KeysWithPrefix(%q) = %v,want %v\n", pre, b, strs)
// 		}
// 		fmt.Println("------------------------------")
// 	}
// 	// Output:
// 	// 20
// }
