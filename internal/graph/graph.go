package graph

import (
	"errors"
	"sort"
)

type Graph struct {
	Nodes []*Node
}

func (g *Graph) AddNodes(input []*InputNode) {
	for _, i := range input {
		newNode := NewNode(i.Name, i.Duration)
		g.Nodes = append(g.Nodes, newNode)
	}
}

func (g *Graph) getNodeByName(name string) (*Node, error) {
	for _, n := range g.Nodes {
		if n.Name == name {
			return n, nil
		}
	}
	return nil, errors.New("Node not found")
}

// going over the slice twice ensures all the desired
// connections exist
func (g *Graph) AddEdges(input []*InputNode) error {
	for _, i := range input {
		for _, d := range i.Dependencies {
			dep, err := g.getNodeByName(d)
			if err != nil {
				return err
			}
			current, err := g.getNodeByName(i.Name)
			if err != nil {
				return err
			}
			dep.Next = append(dep.Next, current)
			current.Prev = append(current.Prev, dep)
		}
	}
	return nil
}

func (g *Graph) FindStarts() []*Node {
	var Starts []*Node
	for _, n := range g.Nodes {
		if len(n.Prev) == 0 {
			Starts = append(Starts, n)
		}
	}
	return Starts
}

func (g *Graph) FindEnds() []*Node {
	var Ends []*Node
	var maxEf uint64
	for _, n := range g.Nodes {
		if len(n.Next) == 0 && n.ef > maxEf {
			maxEf = n.ef
		}
	}
	for _, n := range g.Nodes {
		if len(n.Next) == 0 {
			n.lf = maxEf
			n.ls = maxEf - n.Duration
			n.Float = n.lf - n.ef
			Ends = append(Ends, n)
		}
	}
	return Ends
}

func (g *Graph) GetCritical() []*Node {
	var criticals []*Node
	for _, n := range g.Nodes {
		if n.Float == 0 {
			criticals = append(criticals, n)
		}
	}
	// sort by start
	sort.Slice(criticals, func(i, j int) bool {
		return criticals[i].es < criticals[j].es
	})
	return criticals
}
