package main

import (
	pq "algorithms/prioityqueue"
	"bufio"
	"bytes"
	"errors"
	"fmt"
	"io"
	"log"
	"os"
	"path/filepath"
	"runtime"
)

func main() {
	// TopM 打印输入行中的最大M行
	// TopM()

	// MultiwayMerge 多向归并：将多个有序的输入流归并成一个有序的输入流
	MultiwayMerge()

}

type Item int

func (this Item) Less(b pq.Key) bool {
	v, _ := b.(Item)
	return this > v
}

// TopM 打印输入行中的最大M行
func TopM() {
	_, crtfile, _, ok := runtime.Caller(0)
	if !ok {
		panic(errors.New("Can not get current file info"))
	}
	filename := filepath.Join(filepath.Dir(crtfile), "../numbers1.txt")

	file, err := os.Open(filename)
	if err != nil {
		log.Fatalln(err)
	}
	var M int = 5
	pq := pq.NewHeapMaxPQWithSize(M + 1)

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		var i int
		fmt.Sscanf(line, "%d", &i)
		pq.Insert(Item(i))
		if pq.Size() > M {
			pq.DelMax()
		}
	}

	for !pq.IsEmpty() {
		fmt.Printf("%d ", pq.DelMax())
	}
	fmt.Println()
}

type String string

func (this String) Less(b pq.Key) bool {
	v, _ := b.(String)
	return this[0] > v[0]
}

func MultiwayMerge() {
	_, crtfile, _, ok := runtime.Caller(0)
	if !ok {
		panic(errors.New("Can not get current file info"))
	}
	files := []string{
		"../m1.txt",
		"../m2.txt",
		"../m3.txt",
	}

	pq := pq.NewIndexHeapMaxPQ(len(files))
	streams := []*bufio.Reader{}
	finished := make([]bool, len(files))

	for i, f := range files {
		filename := filepath.Join(filepath.Dir(crtfile), f)
		file, err := os.Open(filename)
		defer file.Close()

		if err != nil {
			log.Fatalln(err)
		}

		reader := bufio.NewReader(file)

		s, _ := reader.ReadBytes(' ')
		pq.Insert(i, String(bytes.Trim(s, " ")))
		streams = append(streams, reader)
	}

	for !pq.IsEmpty() {
		fmt.Printf("%s ", pq.Max())

		i := pq.DelMax()

		if !finished[i] {
			if s, err := streams[i].ReadBytes(' '); err == nil || err == io.EOF {
				pq.Insert(i, String(bytes.Trim(s, " ")))
				if err == io.EOF {
					finished[i] = true
				}
			}
		}

	}

}
