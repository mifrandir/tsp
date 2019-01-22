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
	cost, path := TSPBB(g.Edges)
	fmt.Println("[INF] TSP-Time:", time.Since(start))
	fmt.Println("[OUT] Path:", path)
	fmt.Println("[OUT] Predicted Cost:", cost)
	fmt.Println("[INF] Acutal Cost:", actualCost(path, g.Edges))
}

//func tspDefault() {
//	var n int
//	if len(os.Args) > 1 {
//		n, _ = strconv.Atoi(os.Args[1])
//	} else {
//		n = 10
//	}
//	fmt.Printf("[INF] Starting random TSP with %s nodes\n", strconv.Itoa(n))
//	start := time.Now()
//	g := graph.NewPolygonGraph(n, float64(100), false)
//	fmt.Println("[INF] Created test Graph in", time.Since(start))
//	start = time.Now()
//	cost, path := TSPBB(g.Edges)
//	fmt.Println("[INF] Completed TSP in", time.Since(start))
//	fmt.Println("[OUT]", cost, path)
//	fmt.Println("[INF] Actual Cost", actualCost(path, g.Edges))
//}
