package tries

import (
	"fmt"
	"reflect"
	"sort"
	"testing"
)

func TestTriesMap(t *testing.T) {
	tests := []struct {
		pairs          map[string]int
		prefix_pairs   map[string][]string
		wildcard_pairs map[string][]string
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
		},
	}

	for _, test := range tests {
		t.Run("A=1", func(t *testing.T) {
			tries_map := NewTriesMap()
			var keys []string
			for k, v := range test.pairs {
				tries_map.Put(k, v)
				t.Logf("Put(\"%s\",%v),size = %d\n", k, v, tries_map.Size())
				if b := tries_map.Contains(k); !b {
					t.Errorf("Contains(\"%s\") = %v,want true", k, b)
				}
				if b := tries_map.Get(k); b != v {
					t.Errorf("Get(\"%s\") = %v,want %v", k, b, v)
				}
				keys = append(keys, k)
			}
			sort.Strings(keys)
			b := tries_map.Keys()
			sort.Strings(b)
			if !reflect.DeepEqual(b, keys) {
				t.Errorf("Keys() = %v,want %v", b, keys)
			}
			for pre, strs := range test.prefix_pairs {
				sort.Strings(strs)
				b := tries_map.KeysWithPrefix(pre)
				sort.Strings(b)
				if !reflect.DeepEqual(b, strs) {
					t.Errorf("KeysWithPrefix(\"%s\") = %v,want %v", pre, b, strs)
				}
			}

			for wpre, strs := range test.wildcard_pairs {
				sort.Strings(strs)
				b := tries_map.KeysThatMatch(wpre)
				sort.Strings(b)
				if !reflect.DeepEqual(b, strs) {
					t.Errorf("KeysThatMatch(\"%s\") = %v,want %v", wpre, b, strs)
				}
			}
		})
	}
}

func ExampleTrriesMap() {
	tests := []struct {
		pairs       map[string]int
		prefix_pars map[string][]string
	}{
		{
			map[string]int{
				"sells":  1,
				"shells": 2,
				"sea":    3,
				"shore":  4,
				"by":     5,
				"she":    6,
			},
			map[string][]string{
				"": {
					"sells", "shells", "she", "sea", "shore", "by",
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
		},
	}

	for _, test := range tests {
		tries_map := NewTriesArr()
		var keys []string
		for k, v := range test.pairs {
			tries_map.Put(k, v)
			fmt.Printf("Put(\"%s\",%v),size = %d\n", k, v, tries_map.Size())
			fmt.Printf("Contains(\"%s\") = %v,want true\n", k, tries_map.Contains(k))
			keys = append(keys, k)
		}

		b := tries_map.Keys()
		fmt.Printf("Keys() = %v,want %v\n", b, keys)

		for pre, strs := range test.prefix_pars {
			fmt.Printf("KeysWithPrefix(\"%s\") = %v,want %v\n", pre, b, strs)
		}
		fmt.Println("------------------------------")
	}
	// Output:
	// 20
}
