package main

import (
	"fmt"
	"os"
	"strconv"
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
		file = "/home/miltfra/wild20"
	}
	start := time.Now()
	g := graph.FromFile(file, 1)
	fmt.Println("[INF] Read Graph in", time.Since(start))
	start = time.Now()
	cost, path := TSPBB(g.Edges)
	fmt.Println("[INF] Completed TSP in", time.Since(start))
	fmt.Println("[OUT]", cost, path)
}

func tspDefault() {
	var n int
	if len(os.Args) > 1 {
		n, _ = strconv.Atoi(os.Args[1])
	} else {
		n = 10
	}
	fmt.Printf("[INF] Starting random TSP with %s nodes\n", strconv.Itoa(n))
	start := time.Now()
	g := graph.NewPolygonGraph(n, float64(100), false)
	fmt.Println("[INF] Created test Graph in", time.Since(start))
	start = time.Now()
	cost, path := TSPBB(g.Edges)
	fmt.Println("[INF] Completed TSP in", time.Since(start))
	fmt.Println("[OUT]", cost, path)
}
