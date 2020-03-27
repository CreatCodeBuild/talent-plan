package main

import (
	"container/heap"
	"sort"
)

// SortableInt64 implements 
//   sort.Interface and 
//   heap.Interface for []int64
type SortableInt64 []int64

// Len .
func (is *SortableInt64) Len() int {
	return len(*is)
}

// Less .
func (is *SortableInt64) Less(i, j int) bool {
	return (*is)[i] < (*is)[j]
}

func (is *SortableInt64) Swap(i, j int) {
	t := (*is)[i]
	(*is)[i] = (*is)[j]
	(*is)[j] = t
}


// Push .
func (is *SortableInt64) Push(x interface{}) {
	// Push and Pop use pointer receivers because they modify the slice's length,
	// not just its contents.
	*is = append(*is, x.(int64))
}

// Pop .
func (is *SortableInt64) Pop() interface{} {
	old := *is
	n := len(old)
	x := old[n-1]
	*is = old[0 : n-1]
	return x
}
// MergeSort performs the merge sort algorithm.
// Please supplement this function to accomplish the home work.
func MergeSort(src []int64) {
	// s := SortableInt64(src)
	// heap.Init(&s)
	// sort.Sort(&s)

	l, r := split(src)
	MergeSort(l)
	MergeSort(r)
	merge(l, r)
}

func split(src []int64) (l, r []int64) {
	m := len(src) / 2
	return src[0:m], src[m:]
}

func merge(l, r []int64) {

}
