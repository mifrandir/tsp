package main

import (
	"fmt"
	"os"
	"strconv"
	"time"

	"github.com/miltfra/tools/ds/graph"
)

func main() {
	var n int
	if len(os.Args) > 1 {
		n, _ = strconv.Atoi(os.Args[1])
	} else {
		n = 10
	}
	fmt.Printf("[INF] Starting random TSP with %s nodes\n", strconv.Itoa(n))
	start := time.Now()
	g := graph.NewPolygonGraph(n, float64(100), true)
	fmt.Println("[INF] Created test Graph in", time.Since(start))
	start = time.Now()
	cost, path := TSPBB(g.Edges)
	fmt.Println("[INF] Completed TSP in", time.Since(start))
	fmt.Println("[OUT]", cost, path)
	fmt.Println("[INF] Actual Cost", actualCost(path, g.Edges))
}
