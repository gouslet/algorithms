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

func main() {
	/*********************1.*******************/
	// exampleGraph()
	/*********************2.*******************/
	// exampleSymbolGraph()
	/*********************3.*******************/
	// exampleDFSPaths()
	/*********************4.*******************/
	exampleBFSPaths()
}

func exampleGraph() {
	_, crtfile, _, ok := runtime.Caller(0)
	if !ok {
		panic(errors.New("Can not get current file info"))
	}
	filename := filepath.Join(filepath.Dir(crtfile), "../TinyCG.txt")

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
}

func exampleSymbolGraph() {
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

func exampleDFSPaths() {
	_, crtfile, _, ok := runtime.Caller(0)
	if !ok {
		panic(errors.New("Can not get current file info"))
	}
	filename := filepath.Join(filepath.Dir(crtfile), "../tinyCG.txt")
	file, err := os.Open(filename)
	defer file.Close()

	if err != nil {
		log.Fatalln(err)
	}
	g := graph.NewGraphFrom(file)

	search := graph.NewDepthFirsPaths(g, 0)
	for i := 0; i < g.V(); i++ {
		fmt.Printf("0 to %d: ", i)
		if search.HasPathTo(i) {
			paths := search.PathTo(i)
			for j := 0; j < len(paths); j++ {
				if paths[j] == 0 {
					fmt.Print(paths[j])
				} else {
					fmt.Printf("-%d", paths[j])
				}
			}
		}
		fmt.Println()
	}
	// Output:
	//
	// 0 to 0: 0
	// 0 to 1: 0-1-2-3-5
	// 0 to 2: 0-2-3-5
	// 0 to 3: 0-3-5
	// 0 to 4: 0-4-2-3-5
	// 0 to 5: 0-5

}

func exampleBFSPaths() {
	_, crtfile, _, ok := runtime.Caller(0)
	if !ok {
		panic(errors.New("Can not get current file info"))
	}
	filename := filepath.Join(filepath.Dir(crtfile), "../tinyCG.txt")
	file, err := os.Open(filename)
	defer file.Close()

	if err != nil {
		log.Fatalln(err)
	}
	g := graph.NewGraphFrom(file)

	search := graph.NewBreadthFirsPaths(g, 0)
	for i := 0; i < g.V(); i++ {
		fmt.Printf("0 to %d: ", i)
		if search.HasPathTo(i) {
			paths := search.PathTo(i)
			for j := 0; j < len(paths); j++ {
				if paths[j] == 0 {
					fmt.Print(paths[j])
				} else {
					fmt.Printf("-%d", paths[j])
				}
			}
		}
		fmt.Println()
	}
	// Output:
	//
	// 0 to 0: 0
	// 0 to 1: 0-1-2-3-5
	// 0 to 2: 0-2-3-5
	// 0 to 3: 0-3-5
	// 0 to 4: 0-4-2-3-5
	// 0 to 5: 0-5

}
