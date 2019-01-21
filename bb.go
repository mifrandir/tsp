package main

import (
	"math"

	"github.com/miltfra/tsp/status"
)

// TSPBB calculates the Traveling Salesman Problem on a given
// edge matrix and returns the best value and the best path
func TSPBB(mtrx [][]float64) (float64, []int) {
	l := len(mtrx)
	stat := status.New(int(math.Pow(float64(l), 6)))
	overlay := func() [][]bool {
		l := len(mtrx)
		ol := make([][]bool, l)
		for i := 0; i < l; i++ {
			ol[i] = make([]bool, l)
			for j := 0; j < l; j++ {
				if mtrx[i][j] != 0 {
					ol[i][j] = true
				} else {
					// I don't think we need this, technically
					// TODO Find out!
					ol[i][j] = false
				}
			}
		}
		return ol
	}()
	v := make([]bool, l)
	v[0] = true
	stat.Put(status.NewElement(mtrx, overlay, v, 0, 1))
	for stat.Length > 0 {
		candidate := stat.Get()
		if candidate.Count == l {
			return candidate.Boundary, overlayToPath(candidate.Overlay)
		}
		for i := 0; i < l; i++ {
			if !candidate.Visited[i] && candidate.Overlay[candidate.LastVertex][i] {
				newOverlay := getNewOverlay(candidate.Overlay, candidate.LastVertex, i)
				newVisited := make([]bool, l)
				copy(newVisited, candidate.Visited)
				newVisited[i] = true
				stat.Put(status.NewElement(mtrx, newOverlay, newVisited, i, candidate.Count+1))
			}
		}
	}
	return 2147483647, make([]int, 0)
}

func getNewOverlay(overlay [][]bool, start, target int) [][]bool {
	n := make([][]bool, len(overlay))
	for i := 0; i < len(n); i++ {
		n[i] = make([]bool, len(overlay))
		copy(n[i], overlay[i])
	}
	for i := 0; i < len(n); i++ {
		n[start][i] = false
		n[i][target] = false
	}
	n[start][target] = true
	return n
}

func overlayToPath(overlay [][]bool) []int {
	path := make([]int, len(overlay))
	last := 0
	for i := 0; i < len(overlay); i++ {
		path[i] = last
		for j := 0; j < len(overlay); j++ {
			if overlay[last][j] {
				last = j
				break
			}
		}
	}
	return path
}

func actualCost(path []int, adjMatrix [][]float64) float64 {
	j := path[len(path)-1]
	sum := float64(0)
	for i := 0; i < len(path); i++ {
		sum += adjMatrix[j][path[i]]
		j = path[i]
	}
	return sum
}

func factorial(i int) int {
	if i == 0 {
		return 1
	}
	f := i
	for j := i - 1; j > 0; j-- {
		f *= j
	}
	return f
}
