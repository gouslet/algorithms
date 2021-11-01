package main

import (
	uf "algorithms/union-find/quick-find"
	"errors"
	"fmt"
	"io"
	"log"
	"os"
	"path/filepath"
	"runtime"
)

func main() {
	_, crtfile, _, ok := runtime.Caller(0)
	if !ok {
		panic(errors.New("Can not get current file info"))
	}
	filename := filepath.Join(filepath.Dir(crtfile), "../../TinyUF.txt")

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
