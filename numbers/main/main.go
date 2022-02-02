package main

import (
	"algorithms/numbers"
	"fmt"
)

func main() {
	fmt.Println(numbers.Convert(0))
	fmt.Println(numbers.Convert(1))
	fmt.Println(numbers.Convert(123))
	fmt.Println(numbers.Convert(1023))
	fmt.Println(numbers.Convert(1203))
	fmt.Println(numbers.Convert(1003))
	fmt.Println(numbers.Convert(10230))
	fmt.Println(numbers.Convert(10000))
}
