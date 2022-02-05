package sorts

import (
	"errors"
	"fmt"
	"io"
	"log"
	"os"
	"path/filepath"
	"reflect"
	"runtime"
	"testing"
)

func ExampleInsertionSort() {
	_, crtfile, _, ok := runtime.Caller(0)
	if !ok {
		panic(errors.New("Can not get current file info"))
	}
	filename := filepath.Join(filepath.Dir(crtfile), "./test_data/Tiny.txt")

	file, err := os.Open(filename)
	if err != nil {
		log.Fatalln(err)
	}
	defer file.Close()

	var (
		d int
		s string
		a []string
	)
	_, err = fmt.Fscanf(file, "%d\n", &d)
	if err != nil {
		panic("int value of string length needed")
	}
	for _, err = fmt.Fscanf(file, "%s\n", &s); err != io.EOF; _, err = fmt.Fscanf(file, "%s\n", &s) {
		a = append(a, s)
	}
	l := len(a) - 1
	NewInsertion(a, 0, l, d).Sort()
	for _, t := range a {
		fmt.Println(t)
	}
	// Output:
	// A
	// E
	// E
	// L
	// M
	// O
	// P
	// R
	// S
	// T
	// X

}

func TestInsertionSort(t *testing.T) {
	_, crtfile, _, ok := runtime.Caller(0)
	if !ok {
		panic(errors.New("Can not get current file info"))
	}

	tests := []struct {
		items_file string
		res_file   string
	}{
		{"./test_data/7Characters.txt", "./test_data/7Characters_sorted.txt"},
		{"./test_data/Tiny.txt", "./test_data/Tiny_sorted.txt"},
	}

	for _, test := range tests {
		filename1 := filepath.Join(filepath.Dir(crtfile), test.items_file)
		filename2 := filepath.Join(filepath.Dir(crtfile), test.res_file)

		file1, err := os.Open(filename1)
		if err != nil {
			log.Fatalln(err)
		}
		defer file1.Close()

		file2, err := os.Open(filename2)
		if err != nil {
			log.Fatalln(err)
		}
		defer file2.Close()

		var (
			d, dd int
			s     string
			items []string
			res   []string
		)
		_, err = fmt.Fscanf(file1, "%d\n", &d)
		if err != nil {
			panic("int value of string length needed")
		}
		_, err = fmt.Fscanf(file2, "%d\n", &dd)
		if err != nil {
			panic("int value of string length needed")
		}
		for _, err = fmt.Fscanf(file1, "%s\n", &s); err != io.EOF; _, err = fmt.Fscanf(file1, "%s\n", &s) {
			items = append(items, s)
		}
		for _, err := fmt.Fscanf(file2, "%s\n", &s); err != io.EOF; _, err = fmt.Fscanf(file2, "%s\n", &s) {
			res = append(res, s)
		}
		l := len(items) - 1

		if NewInsertion(items, 0, l, d).Sort(); !reflect.DeepEqual(items, res) {
			t.Fatalf("got %v,want %v\n", items, res)
		}
	}
}
