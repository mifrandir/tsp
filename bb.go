package main

import (
	"fmt"
	"math"
	"runtime"
	"sync"

	"github.com/miltfra/tools/ds/matrix"
)

// Element implements a branch in the TSP-Tree
type Element struct {
	FBPath   []int8 // fwd + bck
	LstVtx   int8
	Count    int8
	Boundary uint
}

// NewElement returns a new Element for the Status heap
func NewElement(fbpath []int8, lv, cnt int8) *Element {
	return &Element{fbpath, lv, cnt, 0}
}

// Status a queue of status elements
type Status struct {
	// TSP Stuff
	adj      []uint
	solution *Element
	solved   bool
	vc       uint16
	// Heap Stuff
	arr     []*Element
	segSize int
	curSize int
	// Sync Stuff
	lckr sync.Mutex
	wg   sync.WaitGroup
}

// NewStatus returns a new Status heap of segment size N
func NewStatus(adjMtrx []uint, segSize int) *Status {
	return &Status{
		// TSP Stuff
		adjMtrx, nil, false, uint16(math.Sqrt(float64(len(adjMtrx)))),
		// Heap Stuff
		make([]*Element, segSize), segSize, 0,
		// Sync Stuff
		sync.Mutex{}, sync.WaitGroup{},
	}
}

// Put inserts an element into the priority queue
func (stat *Status) Put(e *Element) {
	stat.lckr.Lock()
	stat.curSize++
	if stat.curSize%stat.segSize == 0 {
		fmt.Println("[INF] Resizing heap to", stat.curSize+stat.segSize, "elements")
		arr := make([]*Element, stat.curSize+stat.segSize)
		copy(arr, stat.arr)
		stat.arr = arr
	}
	stat.arr[stat.curSize-1] = e
	stat.up(stat.curSize - 1)
	stat.lckr.Unlock()
}

// Get returns the first element of the priority queue
func (stat *Status) Get() *Element {
	stat.lckr.Lock()
	if stat.curSize == 0 {
		return nil
	}
	v := stat.arr[0]
	stat.curSize--
	stat.arr[0] = stat.arr[stat.curSize]
	defer stat.down(0)
	return v
}

// less compares an Element to another Element
func (e *Element) less(other *Element) bool {
	if e.Boundary == other.Boundary {
		return e.Count < other.Count
	}
	return e.Boundary > other.Boundary
}

// greater compares an Element to another Element
func (e *Element) greater(other *Element) bool {
	if e.Boundary == other.Boundary {
		return e.Count > other.Count
	}
	return e.Boundary < other.Boundary
}

func (stat *Status) down(i int) {
	v := stat.arr[i]
	child := (i << 1) + 1
	for l := stat.curSize; child+1 < l; child = (i << 1) + 1 {
		if stat.arr[child].less(stat.arr[child+1]) {
			child++
		}
		cv := stat.arr[child]
		b := v.less(cv)
		if b {
			stat.arr[i] = cv
			i = child
		} else {
			break
		}
	}
	if child < stat.curSize {
		cv := stat.arr[child]
		if v.less(cv) {
			stat.arr[i] = cv
			i = child
		}
	}
	stat.arr[i] = v
	stat.lckr.Unlock()
}

func (stat *Status) up(i int) {
	v := stat.arr[i]
	parent := (i - 1) >> 1
	for ; i > 0; parent = (i - 1) >> 1 {
		pv := stat.arr[parent]
		if v.greater(pv) {
			stat.arr[i] = pv
			i = parent
		} else {
			break
		}
	}
	stat.arr[i] = v
}

// TSPBB calculates the Traveling Salesman Problem on a given
// edge matrix and returns the best value and the best path while
// utilizing goroutines
func TSPBB(mtrx [][]uint, maxProcs, segSize int, grCnt uint16) (uint, []int8) {
	runtime.GOMAXPROCS(maxProcs)
	status := NewStatus(matrix.Flatten(mtrx), segSize)
	var i uint16
	rootFBPath := make([]int8, status.vc<<1)
	for i = 0; i < status.vc; i++ {
		rootFBPath[i] = -1
		rootFBPath[status.vc+i] = -1
	}
	status.Put(NewElement(rootFBPath, 0, 1))
	for i = 0; i < grCnt; i++ {
		status.wg.Add(1)
		go extend(status)
	}
	status.wg.Wait()
	if status.solved {
		return status.solution.Boundary, elemToPath(status)
	}
	return 2147483647, make([]int8, 0)
}

func extend(status *Status) {
	var candidate *Element
	var i uint16
	for status.curSize > 0 {
		if status.solved {
			break
		}
		candidate = status.Get()
		if uint16(candidate.Count) == status.vc+1 {
			status.solution = candidate
			status.solved = true
		} else {
			if uint16(candidate.Count) == status.vc {
				i = 0
			} else {
				i = 1
			}
			for ; i < status.vc; i++ {
				if candidate.FBPath[status.vc+i] == -1 &&
					status.adj[uint16(candidate.LstVtx)*status.vc+i] != 0 {
					status.Put(getNewElement(status, candidate, i))
				}
			}
		}
	}
	status.wg.Done()
}

// UpdateBoundary updates the boundary of the Status Element
// TODO: Use more PQs to manage the edges to update more quickly
func UpdateBoundary(status *Status, e *Element) {
	// Declaring variables so we don't need to allocate space multiple times
	var min, v uint
	var j, i uint16
	// Outgoing edges
	var out uint
	for i = 0; i < status.vc; i++ {
		if e.FBPath[i] != -1 {
			// If there is a path we can add it's value immediately
			out += status.adj[i*status.vc+uint16(e.FBPath[i])]
		} else {
			// Else we have to cycle through the matrix to find the lowest value
			min = ^uint(0)
			for j = 0; j < status.vc; j++ {
				if v = status.adj[i*status.vc+j]; v != 0 && v < min {
					min = v
				}
			}
			out += min
		}
	}
	// Incoming edges
	var in uint
	for i = 0; i < status.vc; i++ {
		if e.FBPath[status.vc+i] != -1 {
			// If there is a path we can add it's value immediately
			in += status.adj[uint16(e.FBPath[status.vc+i])*status.vc+i]
		} else {
			// Else we have to cycle through the matrix to find the lowest value
			min = ^uint(0)
			for j = 0; j < status.vc; j++ {
				if v = status.adj[j*status.vc+i]; v != 0 && v < min {
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
func getNewElement(status *Status, candidate *Element, i uint16) *Element {
	fbPath := make([]int8, status.vc<<1)
	copy(fbPath, candidate.FBPath)
	fbPath[candidate.LstVtx] = int8(i)
	fbPath[status.vc+i] = candidate.LstVtx
	e := &Element{fbPath, int8(i), candidate.Count + 1, 0}
	UpdateBoundary(status, e)
	return e
}

func elemToPath(status *Status) []int8 {
	path := make([]int8, status.vc)
	var i uint16 // starts with 0 anyways
	var next int8
	for i = 0; i < status.vc; i++ {
		path[i] = next
		next = status.solution.FBPath[next]
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
