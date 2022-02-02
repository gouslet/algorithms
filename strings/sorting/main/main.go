package main

import (
	strings "algorithms/strings/sorting"
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
	filename := filepath.Join(filepath.Dir(crtfile), "../test_data/Students_Group(name-indexed).txt")

	file, err := os.Open(filename)
	if err != nil {
		log.Fatalln(err)
	}
	defer file.Close()

	var (
		s string
		r int
		a []strings.Item
	)
	for _, err := fmt.Fscanf(file, "%s %d\n", &s, &r); err != io.EOF; _, err = fmt.Fscanf(file, "%s %d\n", &s, &r) {
		fmt.Printf("s -> %s, r -> %d\n", s, r)
		a = append(a, strings.Item{r, s})
	}

	a = strings.NewIndexCountingSort(a).Sort()
	fmt.Println("---------------------------")
	for i, l := 0, len(a); i < l; i++ {
		fmt.Printf("%s %d\n", a[i].V, a[i].K)
	}
}
