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

// func main() {
// 	_, crtfile, _, ok := runtime.Caller(0)
// 	if !ok {
// 		panic(errors.New("Can not get current file info"))
// 	}
// 	filename := filepath.Join(filepath.Dir(crtfile), "../TinyG.txt")

// 	file, err := os.Open(filename)
// 	if err != nil {
// 		log.Fatalln(err)
// 	}
// 	g := graph.NewGraphFrom(file)
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
// }

func main() {
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
