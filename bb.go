package main

import (
	"github.com/miltfra/tsp/status"
)

// TSPBB calculates the Traveling Salesman Problem on a given
// edge matrix and returns the best value and the best path
func TSPBB(mtrx [][]uint) (uint, []int8) {
	var i int8
	l := int8(len(mtrx))
	stat := status.New(1 << 32)
	fwd := make([]int8, l)
	bck := make([]int8, l)
	for i := int8(0); i < l; i++ {
		fwd[i] = -1
		bck[i] = -1
	}
	stat.Put(status.NewElement(mtrx, fwd, bck, 0, 1))
	var candidate *status.Element
	for stat.Length > 0 {
		candidate = stat.Get()
		if candidate.Count == l {
			if candidate.AdjMatrix[candidate.LastVertex][0] != 0 {
				return candidate.Boundary, fwdToPath(candidate.FwdPath)
			}
		} else {
			// 0 has been visited in every path so we don't have to consider it
			for i = 1; i < l; i++ {
				if candidate.BckPath[i] == -1 &&
					candidate.AdjMatrix[candidate.LastVertex][i] != 0 {
					stat.Put(getNewElement(candidate, i))
				}
			}
		}
	}
	return 2147483647, make([]int8, 0)
}

// Adds a vertex to the paths of a candidate
func getNewElement(candidate *status.Element, i int8) *status.Element {
	l := len(candidate.AdjMatrix)
	fwd := make([]int8, l)
	bck := make([]int8, l)
	copy(fwd, candidate.FwdPath)
	copy(bck, candidate.BckPath)
	fwd[candidate.LastVertex] = i
	bck[i] = candidate.LastVertex
	return status.NewElement(candidate.AdjMatrix, fwd, bck, i, candidate.Count+1)
}

func fwdToPath(fwd []int8) []int8 {
	path := make([]int8, len(fwd))
	var next int8 // starts with 0 anyways
	for i := 0; i < len(fwd); i++ {
		path[i] = next
		next = fwd[next]
	}
	return path
}

func actualCost(path []int8, adjMatrix [][]uint) uint {
	j := path[len(path)-1]
	var sum uint
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
