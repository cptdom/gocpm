package cmd

import (
	"fmt"
	"os"
	"strconv"

	g "cptdom/gocpm/internal/graph"
)

type UserInput struct {
	Nodes []*g.InputNode
}

func (u *UserInput) ProcessInput(name string, duration uint64, dependencies []string) {
	Node := &g.InputNode{
		Name:         name,
		Duration:     duration,
		Dependencies: dependencies,
	}
	u.Nodes = append(u.Nodes, Node)
}

func GenerateGraph(input *UserInput) *g.Graph {
	G := &g.Graph{}
	G.AddNodes(input.Nodes)
	if err := G.AddEdges(input.Nodes); err != nil {
		fmt.Printf("An error occured: %v\n", err)
		os.Exit(1)
	}
	return G
}

func ProcessGraph(G *g.Graph) {
	Starts := G.FindStarts()
	g.DoEarly(Starts)
	Ends := G.FindEnds()
	g.DoLatest(Ends)
}

func PrintSolution(G *g.Graph) {
	var Seq string
	var TotalDuration uint64
	criticals := G.GetCritical()
	for i, n := range criticals {
		TotalDuration += n.Duration
		Seq = Seq + n.Name + " (" + strconv.Itoa(int(n.Duration)) + ")"
		if i != len(criticals)-1 {
			Seq = Seq + ", "
		}
	}
	fmt.Println(Seq)
	fmt.Printf("Total minimal duration of the project: %d\n", TotalDuration)
}
