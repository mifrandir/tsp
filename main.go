package main

import (
	"fmt"
	"github.com/miltfra/tools"
	"os"
	"strconv"
	"time"
)

func main() {
	n, _ := strconv.Atoi(os.Args[1])
	fmt.Printf("[INF] Starting random TSP with %s nodes\n", os.Args[1])
	start := time.Now()
	rnd := tools.RndGraph(n)
	fmt.Println("[INF] Created test Graph in", time.Since(start))
	start = time.Now()
	cost, path := TSPBB(rnd)
	fmt.Println("[INF] Completed TSP in", time.Since(start))
	fmt.Println("[OUT]", cost, path)
}
