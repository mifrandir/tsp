package status

import (
	"math/rand"
	"sync"
)

// Element implements a branch in the TSP-Tree
type Element struct {
	FwdPath    []int8
	BckPath    []int8
	LastVertex int8
	Count      int8
	Boundary   uint
}

// Status a queue of status elements
type Status struct {
	Arr    []*Element
	N      int
	Length int
	m      sync.Mutex
}

// New returns a new Status heap of segment size N
func New(N int) *Status {
	s := Status{make([]*Element, N), N, 0, sync.Mutex{}}
	return &s
}

// NewElement returns a new Element for the Status heap
func NewElement(AdjMatrix [][]uint, fwd, bck []int8, lastVertex, count int8) *Element {
	return &Element{fwd, bck, lastVertex, count, 0}
}

// Put inserts an element into the priority queue
func (stat *Status) Put(e *Element) {
	stat.m.Lock()
	stat.Length++
	if stat.Length%stat.N == 0 {
		arr := make([]*Element, stat.Length+stat.N)
		copy(arr, stat.Arr)
		stat.Arr = arr
	}
	stat.Arr[stat.Length-1] = e
	stat.up(stat.Length - 1)
	stat.m.Unlock()
}

// Get returns the first element of the priority queue
func (stat *Status) Get() *Element {
	stat.m.Lock()
	if stat.Length == 0 {
		return nil
	}
	v := stat.Arr[0]
	stat.Length--
	stat.Arr[0] = stat.Arr[stat.Length]
	stat.down(0)
	stat.m.Unlock()
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

// heapify converts a given array to a heap
func (stat *Status) heapify() {
	for i := stat.getStart(); i >= 0; i-- {
		stat.down(i)
	}
}

// Check checks a given array for the heap property
func (stat *Status) Check() bool {
	for i := 0; i < len(stat.Arr); i++ {
		c := (i << 1) + 1
		if c >= stat.Length {
			break
		}
		if stat.Arr[i].less(stat.Arr[c]) {
			return false
		}
		if c+1 >= stat.Length {
			break
		}
		if stat.Arr[i].less(stat.Arr[c+1]) {
			return false
		}
	}
	return true
}

// RandArr returns an array of size n with all integers in [0, 100) in random order
func RandArr(n int) []int {
	var arr = make([]int, n)
	for i := range arr {
		arr[i] = i
	}
	for i := range arr {
		j := rand.Intn(n)
		v := arr[i]
		arr[i] = arr[j]
		arr[j] = v
	}
	return arr
}

// Gets the first index to down() on the array
func (stat *Status) getStart() int {
	var i = 0
	for (i<<1)+1 < stat.Length {
		i = (i << 1) + 1
	}
	return i
}

func (stat *Status) down(i int) {
	v := stat.Arr[i]
	child := (i << 1) + 1
	for l := stat.Length; child+1 < l; child = (i << 1) + 1 {
		if stat.Arr[child].less(stat.Arr[child+1]) {
			child++
		}
		cv := stat.Arr[child]
		b := v.less(cv)
		if b {
			stat.Arr[i] = cv
			i = child
		} else {
			break
		}
	}
	if child < stat.Length {
		cv := stat.Arr[child]
		if v.less(cv) {
			stat.Arr[i] = cv
			i = child
		}
	}
	stat.Arr[i] = v
}

func (stat *Status) up(i int) {
	v := stat.Arr[i]
	parent := (i - 1) >> 1
	for ; i > 0; parent = (i - 1) >> 1 {
		pv := stat.Arr[parent]
		if v.greater(pv) {
			stat.Arr[i] = pv
			i = parent
		} else {
			break
		}
	}
	stat.Arr[i] = v
}
