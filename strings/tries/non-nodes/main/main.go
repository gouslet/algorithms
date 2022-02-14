package main

import (
	tries "algorithms/strings/tries/non-nodes"
	"fmt"
)

func main() {
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
		tries_arr := tries.NewTriesArr()
		var keys []string
		for k, v := range test.pairs {
			tries_arr.Put(k, v)
			fmt.Printf("Put(\"%s\",%v),size = %d\n", k, v, tries_arr.Size())
			fmt.Printf("Contains(\"%s\") = %v,want true\n", k, tries_arr.Contains(k))
			keys = append(keys, k)
		}

		b := tries_arr.Keys()
		fmt.Printf("Keys() = %v,want %v\n", b, keys)

		for pre, strs := range test.prefix_pars {
			fmt.Printf("KeysWithPrefix(\"%s\") = %v,want %v\n", pre, b, strs)
		}
		fmt.Println("------------------------------")
	}
	// Output:
	// 20
}
