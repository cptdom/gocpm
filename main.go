package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	api "cptdom/gocpm/cmd"
	v "cptdom/gocpm/version"
)

func main() {
	fmt.Printf("gocpm %s\n", v.Version)
	reader := bufio.NewReader(os.Stdin)
	i := 1
	tasks := &api.UserInput{}
L:
	for {
		fmt.Printf("\nTask n. %d\n", i)
		// name
		fmt.Println("Enter task's shortname or sign: ")
		input, _, _ := reader.ReadLine()
		Name := strings.Trim(string(input), " ")
		// duration
		fmt.Println("Enter task's duration (a number): ")
		var Duration uint64
		for {
			_, err := fmt.Scanf("%d", &Duration)
			if err != nil {
				fmt.Println("Invalid input, try again?")
			} else {
				break
			}
		}
		// dependencies
		fmt.Println("Enter the names of this task's dependencies, divided by a comma: ")
		input, _, _ = reader.ReadLine()
		strInput := strings.Trim(string(input), " ")
		var Dependencies []string
		if strInput != "" {
			Dependencies = strings.Split(strInput, ",")
		}
		tasks.ProcessInput(Name, Duration, Dependencies)
		// proceed
		fmt.Println("Task successfully added. Add another one? (y/n)")
		input, _, _ = reader.ReadLine()
		strInput = string(input)
		switch strInput {
		case "Y", "y":
			i++
		case "N", "n":
			break L
		default:
			i++
		}
	}
	fmt.Printf("Received %d tasks.\n", i)
	Graph := api.GenerateGraph(tasks)
	api.ProcessGraph(Graph)
	api.PrintSolution(Graph)
}
