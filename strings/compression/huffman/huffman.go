package huffman

type Node struct {
	ch          byte
	freq        int
	left, right *Node
}

func (n Node) IsLeaf() bool {
	return n.left == nil && n.right == nil
}

func (n Node) CompareTo(that Node) int {
	return n.freq - that.freq
}
