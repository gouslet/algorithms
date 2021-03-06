package graph

import (
	"bufio"
	"io"
	"os"
	"strings"
)

type SymbolGraph struct {
	st   map[string]int // 符号名->索引
	keys []string       // 索引->符号名
	g    *Graph
}

// NewSymbolGraphFrom 根据filename指定的文件构造图，使用delim来分隔顶点名
func NewSymbolGraphFrom(in io.ReadSeekCloser, sep string) *SymbolGraph {
	st := make(map[string]int)

	scanner := bufio.NewScanner(in)
	for scanner.Scan() {
		line := scanner.Text()
		strs := strings.Split(line, sep)

		for _, s := range strs {
			if _, ok := st[s]; !ok {
				st[s] = len(st)
			}
		}
	}

	keys := make([]string, len(st))
	for name, _ := range st {
		keys[st[name]] = name
	}

	g := NewGraph(len(st))

	in.Seek(0, os.SEEK_SET)
	scanner = bufio.NewScanner(in)
	for scanner.Scan() {
		line := scanner.Text()
		strs := strings.Split(line, sep)

		v := st[strs[0]]
		for i := 1; i < len(strs); i++ {
			g.AddEdge(v, st[strs[i]])
		}
	}

	return &SymbolGraph{st, keys, g}
}

// Contains key是一个顶点吗？
func (this *SymbolGraph) Cotains(key string) bool {
	_, ok := this.st[key]
	return ok
}

// Index key的索引
func (this *SymbolGraph) Index(key string) int {
	return this.st[key]
}

// Name 索引v的顶点名
func (this *SymbolGraph) Name(v int) string {
	return this.keys[v]
}

// G 隐藏的Graph对象
func (this *SymbolGraph) G() *Graph {
	return this.g
}
