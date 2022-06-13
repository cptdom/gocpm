package graph

type InputNode struct {
	Name         string
	Duration     uint64
	Dependencies []string
}
