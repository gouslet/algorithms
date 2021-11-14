package main

import (
	"algorithms/graph"
	"errors"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"runtime"
)

func ExampleGraph() {
	_, crtfile, _, ok := runtime.Caller(0)
	if !ok {
		panic(errors.New("Can not get current file info"))
	}
	filename := filepath.Join(filepath.Dir(crtfile), "../TinyG.txt")

	file, err := os.Open(filename)
	if err != nil {
		log.Fatalln(err)
	}
	g := graph.NewGraphFrom(file)
	s := 0
	search := graph.NewDepthFirstGraph(*g, s)

	for i := 0; i < g.V(); i++ {
		if search.Marked(i) {
			fmt.Printf("%d ", i)
		}
	}
	fmt.Println()
	if search.Count() != g.V() {
		fmt.Print("Not c")
	} else {
		fmt.Print("C")
	}
	fmt.Println("onnected")
	// Output:
	//
	// 0 1 2 3 4 5 6
	// Not connected
}

// func ExampleDigraph() {
// 	_, crtfile, _, ok := runtime.Caller(0)
// 	if !ok {
// 		panic(errors.New("Can not get current file info"))
// 	}
// 	filename := filepath.Join(filepath.Dir(crtfile), "../TinyDG.txt")

// 	file, err := os.Open(filename)
// 	if err != nil {
// 		log.Fatalln(err)
// 	}
// 	g := graph.NewDigraphFrom(file)
// 	s := 0
// 	search := graph.NewDepthFirstGraph(*g, s)

// 	for i := 0; i < g.V(); i++ {
// 		if search.Marked(i) {
// 			fmt.Printf("%d ", i)
// 		}
// 	}
// 	fmt.Println()
// 	if search.Count() != g.V() {
// 		fmt.Print("Not c")
// 	} else {
// 		fmt.Print("C")
// 	}
// 	fmt.Println("onnected")
// 	// Output:
// 	//
// 	// 0 1 2 3 4 5 6
// 	// Not connected
// }

// func ExampleTopological() {
// 	_, crtfile, _, ok := runtime.Caller(0)
// 	if !ok {
// 		panic(errors.New("Can not get current file info"))
// 	}
// 	filename := filepath.Join(filepath.Dir(crtfile), "../jobs.txt")
// 	file, err := os.Open(filename)
// 	if err != nil {
// 		log.Fatalln(err)
// 	}
// 	// g := graph.NewDigraphFrom(file)

// }

func ExampleSymbolGraph() {
	_, crtfile, _, ok := runtime.Caller(0)
	if !ok {
		panic(errors.New("Can not get current file info"))
	}
	filename := filepath.Join(filepath.Dir(crtfile), "../movies.txt")
	file, err := os.Open(filename)
	if err != nil {
		log.Fatalln(err)
	}
	// g := graph.NewDigraphFrom(file)
}
