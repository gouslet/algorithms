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

func TestSort(t *testing.T) {
	_, crtfile, _, ok := runtime.Caller(0)
	if !ok {
		panic(errors.New("Can not get current file info"))
	}

	tests := []struct {
		items_file string
		res_file   string
	}{
		{"./test_data/Students_Group(name-indexed).txt", "./test_data/Students_Group(group-indexed).txt"},
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
			r     int
			items []Item
			res   []Item
		)
		for _, err := fmt.Fscanf(file1, "%s %d\n", &s, &r); err != io.EOF; _, err = fmt.Fscanf(file1, "%s %d\n", &s, &r) {
			items = append(items, Item{r, s})
		}
		for _, err := fmt.Fscanf(file2, "%s %d\n", &s, &r); err != io.EOF; _, err = fmt.Fscanf(file2, "%s %d\n", &s, &r) {
			res = append(res, Item{r, s})
		}

		if items = NewIndexCountingSort(items).Sort(); !reflect.DeepEqual(items, res) {
			t.Fatalf("got %v,want %v\n", items, res)
		}
	}
}
