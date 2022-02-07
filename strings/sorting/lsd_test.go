package sorting

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

func TestLSDSort(t *testing.T) {
	_, crtfile, _, ok := runtime.Caller(0)
	if !ok {
		panic(errors.New("Can not get current file info"))
	}

	tests := []struct {
		w          int
		items_file string
		res_file   string
	}{
		{7, "./test_data/7Characters.txt", "./test_data/7Characters_sorted.txt"},
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
			s     string
			items []string
			res   []string
		)
		for _, err := fmt.Fscanf(file1, "%s\n", &s); err != io.EOF; _, err = fmt.Fscanf(file1, "%s\n", &s) {
			items = append(items, s)
		}
		for _, err := fmt.Fscanf(file2, "%s\n", &s); err != io.EOF; _, err = fmt.Fscanf(file2, "%s\n", &s) {
			res = append(res, s)
		}

		if items = NewLSD(test.w, items).Sort(); !reflect.DeepEqual(items, res) {
			t.Fatalf("got %v,want %v\n", items, res)
		}
	}
}

func ExampleLSDSort() {
	_, crtfile, _, ok := runtime.Caller(0)
	if !ok {
		panic(errors.New("Can not get current file info"))
	}
	filename := filepath.Join(filepath.Dir(crtfile), "./test_data/7Characters.txt")

	file, err := os.Open(filename)
	if err != nil {
		log.Fatalln(err)
	}
	defer file.Close()

	var (
		s string
		a []string
	)
	for _, err := fmt.Fscanf(file, "%s\n", &s); err != io.EOF; _, err = fmt.Fscanf(file, "%s\n", &s) {
		a = append(a, s)
	}

	lsd := NewLSD(7, a)
	for _, t := range lsd.Sort() {
		fmt.Println(t)
	}
	// Output:
	// 10HV845
	// 10HV845
	// 10HV845
	// 1ICK750
	// 1ICK750
	// 2IYE230
	// 2RLA629
	// 2RLA629
	// 3ATW723
	// 3CIO720
	// 3CIO720
	// 4JZY524
	// 4PGC938
}
