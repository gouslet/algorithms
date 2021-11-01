package main

import (
	uf "algorithms/union-find/quick-find"
	"fmt"
	"io"
	"log"
	"os"
)

func ExampleUnionFind() {
	filename := "../../TinyUF.txt"

	file, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}
	var N int
	fmt.Fscanf(file, "%d", &N)

	UF := uf.NewUF(N)
	var p, q int

	for _, err = fmt.Fscanf(file, "%d%d", &p, &q); err != io.EOF; _, err = fmt.Fscanf(file, "%d%d", &p, &q) {
		if !UF.Connected(p, q) {
			UF.Union(p, q)
			fmt.Printf("%d %d\n", p, q)
		}
	}
	fmt.Printf("%d components\n", UF.Count())
	// Output:
	// 4 3
	// 3 8
	// 6 5
	// 9 4
	// 2 1
	// 5 0
	// 7 2
	// 6 1
	// 2 components
}
