package graph

import "io"

type SymbolGraph struct {
	st   map[string]int // 符号名->索引
	keys []string       // 索引->符号名
	g    *Graph
}

func NewSymbolGraphFrom(in io.Reader, sep string)

func (this *SymbolGraph) Cotains(key string) bool

func (this *SymbolGraph) Index(key string) int

func (this *SymbolGraph) Name(v int) string

func (this *SymbolGraph) G() *Graph
