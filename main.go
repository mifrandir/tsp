package main

import (
	"fmt"
	"os"
	"time"

	"github.com/miltfra/tools/ds/graph"
)

func main() {
	tspCLI()
}

func tspCLI() {
	var file string
	if len(os.Args) > 1 {
		file = os.Args[1]
	} else {
		file = "/home/miltfra/projects/example_files/graphs/tsp1.txt"
	}
	start := time.Now()
	g := graph.FromFile(file, 1)
	fmt.Println("[INF] Read-Time:", time.Since(start))
	start = time.Now()
	cost, path := TSPBB(g.Edges, 8, 1<<28, 7)
	fmt.Println("[INF] TSP-Time:", time.Since(start))
	fmt.Println("[OUT] Path:", path)
	fmt.Println("[OUT] Predicted Cost:", cost)
	fmt.Println("[INF] Acutal Cost:", actualCost(path, g.Edges))
}
