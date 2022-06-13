package graph

type Node struct {
	Name     string
	Duration uint64
	Float    uint64
	Prev     []*Node
	Next     []*Node
	es       uint64
	ef       uint64
	ls       uint64
	lf       uint64
}

func NewNode(n string, d uint64) *Node {
	return &Node{
		Name:     n,
		Duration: d,
	}
}
