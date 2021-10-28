package main

import (
	uf "algorithms/union-find"
	"fmt"
	"io"
	"log"
	"os"
)

func main() {
	filename := "TinyUF.txt"
	if len(os.Args) > 1 {
		filename = os.Args[1]
	}
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

}
