package main

import (
	"container/heap"
	"fmt"
)

// Heap implementation
// An IntHeap is a min-heap of ints.
type IntHeap []int

func (h IntHeap) Len() int           { return len(h) }
func (h IntHeap) Less(i, j int) bool { return h[i] < h[j] }
func (h IntHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *IntHeap) Push(x interface{}) {
	// Push and Pop use pointer receivers because they modify the slice's length,
	// not just its contents.
	*h = append(*h, x.(int))
}

func (h *IntHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

// Actual ZipMin function
func ZipMin(numbers []int, zip int, ntop int) [][]int {
	// range window array. we original make it from index 0
	window := numbers[0:zip]
	// end result, array of (int)array
	var result [][]int

	for i := 0; i < len(numbers)-zip+1; i++ {
		// local result of ntop in current window
		var r []int

		// now we build a heap from window elements
		h := &IntHeap{}
		for z := 0; z < len(window); z++ {
			heap.Push(h, window[z])
		}
		heap.Init(h)

		// pop 3 mins elements and add them to local result
		for z := 0; z < ntop; z++ {
			top := heap.Pop(h)
			r = append(r, top.(int))
		}

		// append local result to global result
		result = append(result, r)

		// if we are not at the last step, move window by 1 to the left
		if i+zip < len(numbers) {
			window = append(window[:0], window[1:]...)
			window = append(window, numbers[i+zip])
		}
	}

	return result
}

func main() {
	fmt.Printf("Welcome to ZipTop!\n")
	numbers := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 14}
	result := ZipMin(numbers, 4, 2)
	fmt.Printf("%v\n", result)
}
