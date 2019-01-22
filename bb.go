package main

import (
	"github.com/miltfra/tsp/status"
)

// AdjMatrix is the matrix of the current TSP calculation
var AdjMatrix [][]uint

// TSPBB calculates the Traveling Salesman Problem on a given
// edge matrix and returns the best value and the best path
func TSPBB(mtrx [][]uint) (uint, []int8) {
	AdjMatrix = mtrx
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
		if candidate.Count == l+1 {
			return candidate.Boundary, fwdToPath(candidate.FwdPath)
		} else if candidate.FwdPath[candidate.LastVertex] == -1 {
			// 0 has been visited in every path so we don't have to consider it
			for i = 0; i < l; i++ {
				if candidate.BckPath[i] == -1 &&
					AdjMatrix[candidate.LastVertex][i] != 0 {
					stat.Put(getNewElement(candidate, i))
				}
			}
		}
	}
	return 2147483647, make([]int8, 0)
}

// UpdateBoundary updates the boundary of the Status Element
// TODO: Use more PQs to manage the edges to update more quickly
func UpdateBoundary(e *status.Element) {
	l := len(AdjMatrix)
	// Declaring variables so we don't need to allocate space multiple times
	var min, v uint
	var j, i int
	// Outgoing edges
	var out uint
	for i = 0; i < l; i++ {
		if e.FwdPath[i] != -1 {
			// If there is a path we can add it's value immediately
			out += AdjMatrix[i][e.FwdPath[i]]
		} else {
			// Else we have to cycle through the matrix to find the lowest value
			min = ^uint(0)
			for j = 0; j < l; j++ {
				if v = AdjMatrix[i][j]; v != 0 && v < min {
					min = v
				}
			}
			out += min
		}
	}
	// Incoming edges
	var in uint
	for i = 0; i < l; i++ {
		if e.BckPath[i] != -1 {
			// If there is a path we can add it's value immediately
			in += AdjMatrix[e.BckPath[i]][i]
		} else {
			// Else we have to cycle through the matrix to find the lowest value
			min = ^uint(0)
			for j = 0; j < l; j++ {
				if v = AdjMatrix[j][i]; v != 0 && v < min {
					min = v
				}
			}
			in += min
		}
	}
	//if (in == 38408 || out == 38408) && e.Count == 20 {
	//	fmt.Println("38408")
	//}
	if in > out {
		e.Boundary = in
	} else {
		e.Boundary = out
	}
}

// Adds a vertex to the paths of a candidate
func getNewElement(candidate *status.Element, i int8) *status.Element {
	l := len(AdjMatrix)
	fwd := make([]int8, l)
	bck := make([]int8, l)
	copy(fwd, candidate.FwdPath)
	copy(bck, candidate.BckPath)
	fwd[candidate.LastVertex] = i
	bck[i] = candidate.LastVertex
	e := &status.Element{fwd, bck, i, candidate.Count + 1, 0}
	UpdateBoundary(e)
	return e
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
