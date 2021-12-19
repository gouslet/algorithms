package pq

import (
	"fmt"
	"testing"
)

type Number int

func (this Number) Less(n Value) bool {
	return this < n.(Number)
}

func TestNewIndexHeapPQ(t *testing.T) {

	tests := []struct {
		maxN int
		cap  int
	}{
		{5, 5},
		{0, 0},
		{1000, 1000},
	}
	for _, test := range tests {
		pq := NewIndexHeapPQ(test.maxN)

		if pq.N != 0 {
			t.Fatalf("want PQ with 0 elements,got %d\n", pq.N)
		}
		if l := len(pq.values); l != test.cap+1 {
			t.Fatalf("want PQ with %d capacity,got %d\n", test.cap, l)
		}
	}

}

func TestIsEmpty(t *testing.T) {

	tests := []struct {
		maxN int
		len  int
		vals []int
	}{
		{5, 0, []int{}},
		{1000, 3, []int{1, 2, 3}},
		{5, 5, []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 7, 4, 5, 6, 3, 2, 1, 4, 5, 6}},
	}
	for _, test := range tests {
		pq := NewIndexHeapPQ(test.maxN)

		if !pq.IsEmpty() {
			t.Fatalf("want PQ with 0 elements,got %d\n", pq.N)
		}
		for i, v := range test.vals {
			pq.Insert(i%test.maxN, Number(v))
		}
		if pq.IsEmpty() && test.len != 0 {
			t.Fatalf("want PQ with %d elements,got %d\n", test.len, pq.N)
		}
	}

}

func TestContains(t *testing.T) {
	tests := []struct {
		maxN int
		vals []int
	}{
		{5, []int{}},
		{1000, []int{1, 2, 3}},
		{5, []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 7, 4, 5, 6, 3, 2, 1, 4, 5, 6}},
	}
	for _, test := range tests {
		pq := NewIndexHeapPQ(test.maxN)
		for i := 0; i < test.maxN; i++ {
			if pq.Contains(i) {
				t.Fatalf("want PQ that does not contain No.%d elements", i)
			}
		}
		for i, v := range test.vals {
			pq.Insert(i%test.maxN, Number(v))
			if !pq.Contains(i % test.maxN) {
				t.Fatalf("want PQ that contains No.%d elements", i%test.maxN)
			}
		}
	}

}

func TestSwim(t *testing.T) {
	tests := []struct {
		maxN int
		vals []int
		heap []int
	}{
		{5, []int{}, []int{}},
		{10, []int{1, 2, 3}, []int{3, 1, 2}},
		{5, []int{1, 2}, []int{2, 1}},
		{5, []int{1, 2, 3}, []int{3, 1, 2}},
		{5, []int{1, 2, 3, 4}, []int{4, 3, 2, 1}},
		{5, []int{1, 2, 3, 4, 5}, []int{5, 4, 2, 1, 3}},
	}
	for _, test := range tests {
		pq := NewIndexHeapPQ(test.maxN)
		for i, v := range test.vals {
			k := i%test.maxN + 1
			pq.values[k] = Number(v)
			if pq.qp[k] == -1 {

				pq.N += 1

				pq.pq[pq.N] = k
				pq.qp[k] = pq.N
			}

			pq.swim(pq.qp[k])
		}

		if s1, s2 := fmt.Sprintf("%v", pq), fmt.Sprintf("%v", test.heap); s1 != s2 {
			t.Fatalf("want %s,got %s", s2, s1)
		}
	}
}

func TestSink(t *testing.T) {
	tests := []struct {
		maxN int
		vals []int
		heap []int
	}{
		{5, []int{}, []int{}},
		{10, []int{1, 2, 3}, []int{3, 1, 2}},
		{2, []int{1, 2}, []int{2, 1}},
		{3, []int{1, 2, 3}, []int{3, 1, 2}},
		{4, []int{1, 2, 3, 4}, []int{3, 1, 2, 4}},
		{5, []int{1, 2, 3, 4, 5}, []int{3, 1, 2, 4, 5}},
		{6, []int{1, 2, 3, 4, 5, 6}, []int{3, 1, 2, 4, 5, 6}},
		{7, []int{1, 2, 3, 4, 5, 6, 7}, []int{3, 1, 2, 4, 5, 6, 7}},
	}
	for _, test := range tests {
		pq := NewIndexHeapPQ(test.maxN)
		for i, v := range test.vals {
			k := i%test.maxN + 1
			pq.values[k] = Number(v)
			if pq.qp[k] == -1 {

				pq.N += 1

				pq.pq[pq.N] = k
				pq.qp[k] = pq.N
			}
			pq.sink(1)
		}

		if s1, s2 := fmt.Sprintf("%v", pq), fmt.Sprintf("%v", test.heap); s1 != s2 {
			t.Fatalf("want %s,got %s", s2, s1)
		}
	}
}

func TestInsert(t *testing.T) {
	tests := []struct {
		maxN int
		vals []int
		heap []int
	}{
		{5, []int{}, []int{}},
		{1000, []int{1, 2, 3}, []int{3, 1, 2}},
		{2, []int{1, 2}, []int{2, 1}},
		{3, []int{1, 2, 3}, []int{3, 1, 2}},
		{4, []int{1, 2, 3, 4}, []int{4, 3, 2, 1}},
		{5, []int{1, 2, 3, 4, 5}, []int{5, 4, 2, 1, 3}},
		{6, []int{1, 2, 3, 4, 5, 6}, []int{6, 4, 5, 1, 3, 2}},
		{7, []int{1, 2, 3, 4, 5, 6, 7}, []int{7, 4, 6, 1, 3, 2, 5}},
	}
	for _, test := range tests {
		pq := NewIndexHeapPQ(test.maxN)
		for i, v := range test.vals {
			pq.Insert(i%test.maxN, Number(v))
		}
		if s1, s2 := fmt.Sprintf("%v", pq), fmt.Sprintf("%v", test.heap); s1 != s2 {
			t.Fatalf("want %s,got %s", s1, s2)
		}
	}
}

func TestDelExtre(t *testing.T) {
	tests := []struct {
		maxN   int
		vals   []int
		extres []int
	}{
		{5, []int{}, []int{}},
		{10, []int{1, 2, 3}, []int{3, 2, 1}},
		{2, []int{1, 2}, []int{2, 1}},
		{3, []int{1, 2, 3}, []int{3, 2, 1}},
		{4, []int{1, 2, 3, 4}, []int{4, 3, 2, 1}},
		{5, []int{1, 2, 3, 4, 5}, []int{5, 4, 3, 2, 1}},
		{6, []int{1, 2, 3, 4, 5, 6}, []int{6, 5, 4, 3, 2, 1}},
		{7, []int{1, 2, 3, 4, 5, 6, 7}, []int{7, 6, 5, 4, 3, 2, 1}},
	}
	for _, test := range tests {
		pq := NewIndexHeapPQ(test.maxN)
		for i, v := range test.vals {
			pq.Insert(i, Number(v))
		}
		larges := make([]Value, 0)
		for !pq.IsEmpty() {
			larges = append(larges, pq.Extre())
			pq.DelExtre()
		}
		if s1, s2 := fmt.Sprintf("%v", larges), fmt.Sprintf("%v", test.extres); s1 != s2 {
			t.Fatalf("want %s,got %s", s2, s1)
		}
	}
}

func ExampleString() {
	tests := []struct {
		maxN int
		vals []int
		heap []int
	}{
		{5, []int{}, []int{}},
		{10, []int{1, 2, 3}, []int{1, 2, 3}},
		{5, []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 7, 4, 5, 6, 3, 2, 1, 4, 5, 6}, []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 7, 4, 5, 6, 3, 2, 1, 4, 5, 6}},
	}
	for _, test := range tests {
		pq := NewIndexHeapPQ(test.maxN)
		for i, v := range test.vals {
			pq.Insert(i%test.maxN, Number(v))
		}
		fmt.Printf("%v\n", pq)
	}
	// Output:
	// []
	// [3 1 2]
	// [6 5 1 4 2]
}
