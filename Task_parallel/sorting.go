package main

import (
	"runtime"
)

var maxDepth = runtime.NumCPU() * 2

func merge(left, right []int) []int {
	result := make([]int, 0, len(left)+len(right))
	i, j := 0, 0

	for i < len(left) && j < len(right) {
		if left[i] <= right[j] {
			result = append(result, left[i])
			i++
		} else {
			result = append(result, right[j])
			j++
		}
	}

	result = append(result, left[i:]...)
	result = append(result, right[j:]...)
	return result
}

func parallelMergeSort(arr []int, depth int) []int {
	if len(arr) <= 1 {
		return arr
	}

	mid := len(arr) / 2
	left := arr[:mid]
	right := arr[mid:]

	if depth < maxDepth {
		ch := make(chan []int)
		go func() { ch <- parallelMergeSort(left, depth+1) }()
		rightSorted := parallelMergeSort(right, depth+1)
		leftSorted := <-ch
		return merge(leftSorted, rightSorted)
	}

	leftSorted := parallelMergeSort(left, depth+1)
	rightSorted := parallelMergeSort(right, depth+1)
	return merge(leftSorted, rightSorted)
}

// func main() {
// 	var n int
// 	fmt.Print("Masukkan jumlah data: ")
// 	fmt.Scan(&n)

// 	arr := make([]int, n)
// 	fmt.Println("Masukkan data:")
// 	for i := 0; i < n; i++ {
// 		fmt.Scan(&arr[i])
// 	}

// 	fmt.Println("\nSorting... ")
// 	result := parallelMergeSort(arr, 0)

// 	fmt.Println("Hasil sorted:")
// 	fmt.Println(result)
// }
