package main

import (
	"fmt"

	"github.com/logrusorgru/aurora"

	"lab_06/graph"
)

func main() {
	fmt.Printf("%v\n\n", aurora.Magenta("Ant algorythm"))

	graph.Compare("data/data.txt")
}
