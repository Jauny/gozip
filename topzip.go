package main

import "fmt"
import "sort"

func ZipTop(numbers []int, zip int, ntop int) {
	var result [][]int

	for i := 0; i < len(numbers)-zip+1; i++ {
		window := numbers[i : i+zip]
		heap := make([]int, zip)
		copy(heap, window)
		sort.Sort(sort.Reverse(sort.IntSlice(heap)))
		result = append(result, heap[0:ntop])
	}
	fmt.Printf("%v\n", result)
}

func main() {
	fmt.Printf("Welcome to ZipTop!\n")
	numbers := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	ZipTop(numbers, 3, 2)
}
