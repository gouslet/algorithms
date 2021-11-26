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
	defer file.Close()

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

// func ExampleDirectedCycle() {
// 	_, crtfile, _, ok := runtime.Caller(0)
// 	if !ok {
// 		panic(errors.New("Can not get current file info"))
// 	}
// 	filename := filepath.Join(filepath.Dir(crtfile), "../jobs.txt")
// 	file, err := os.Open(filename)
// 	defer file.Close()

// 	if err != nil {
// 		log.Fatalln(err)
// 	}
// 	g := graph.NewDigraphFrom(file)

// }

func ExampleTopological() {
	_, crtfile, _, ok := runtime.Caller(0)
	if !ok {
		panic(errors.New("Can not get current file info"))
	}
	filename := filepath.Join(filepath.Dir(crtfile), "../jobs.txt")
	file, err := os.Open(filename)
	defer file.Close()

	if err != nil {
		log.Fatalln(err)
	}
	sg := graph.NewSymbolDigraphFrom(file, "/")

	top := graph.NewTopological(*sg.G())
	top.Order().Map(func(v int) {
		fmt.Println(sg.Name(v))
	})
	// Output:
	//
	//Calculus
	//Linear Algebra
	//Introduction to CS
	//Advanced Programming
	//Algorithms
	//Scientific Computing
	//Databases
	//Theoretical CS
	//Artificial Intelligence
	//Machine Learning
	//Robotics
	//Neural Networks
	//Computational Biology
}

func ExampleSymbolGraph() {
	_, crtfile, _, ok := runtime.Caller(0)
	if !ok {
		panic(errors.New("Can not get current file info"))
	}
	filename := filepath.Join(filepath.Dir(crtfile), "../routes.txt")
	file, err := os.Open(filename)
	defer file.Close()

	if err != nil {
		log.Fatalln(err)
	}
	g := graph.NewSymbolGraphFrom(file, " ")

	start := []string{
		"JFK",
		"LAX",
	}
	for i := 0; i < len(start); i++ {
		fmt.Println(start[i])
		ends := g.G().Adj(g.Index(start[i]))
		for j := 0; j < len(ends); j++ {
			fmt.Printf("\t%s\n", g.Name(ends[j]))
		}
	}
	// Output:
	//JFK
	//	MCO
	//	ATL
	//	ORD
	//LAX
	//	PHX
	//	LAS
}
