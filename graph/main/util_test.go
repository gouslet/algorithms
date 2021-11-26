package main

import (
	"algorithms/graph/util"
	"fmt"
)

func ExampleStack() {
	nums := []int{
		1, 2, 3, 4, 5, 6, 7,
	}
	stack := util.NewStack()
	for _, v := range nums {
		stack.Push(v)
	}
	f := func(v int) {
		fmt.Println(v)
	}
	stack.Map(f)
	// Output:
	//7
	//6
	//5
	//4
	//3
	//2
	//1
}
