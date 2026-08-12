package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

type BubbleResponse struct {
	Before string `json:"before"`
	After  string `json:"after"`
	Time   string `json:"time"`
}

func BubbleSortSequential(arr []int) {
	n := len(arr)
	for i := 0; i < n-1; i++ {
		for j := 0; j < n-i-1; j++ {
			if arr[j] > arr[j+1] {
				arr[j], arr[j+1] = arr[j+1], arr[j]
			}
		}
	}
}

func BubbleSortParallel(arr []int, workers int) {
	n := len(arr)
	var wg sync.WaitGroup

	for i := 0; i < n; i++ {
		startIndex := i % 2

		wg.Add(workers)
		for w := 0; w < workers; w++ {
			go func(w int) {
				defer wg.Done()
				for j := startIndex + w*2; j < n-1; j += workers * 2 {
					if arr[j] > arr[j+1] {
						arr[j], arr[j+1] = arr[j+1], arr[j]
					}
				}
			}(w)
		}
		wg.Wait()
	}
}

func GenerateRandomArray(size int) []int {
	arr := make([]int, size)
	for i := 0; i < size; i++ {
		arr[i] = rand.Intn(1000)
	}
	return arr
}

func ArrayToString(arr []int) string {
	limit := 200
	var text string
	for i := 0; i < len(arr) && i < limit; i++ {
		text += fmt.Sprintf("%d, ", arr[i])
	}
	return text + "... (truncated)"
}

func RunBubbleSort(size int, parallel bool) BubbleResponse {
	workers := 8
	original := GenerateRandomArray(size)

	copyArr := make([]int, size)
	copy(copyArr, original)

	start := time.Now()
	if parallel {
		BubbleSortParallel(copyArr, workers)
	} else {
		BubbleSortSequential(copyArr)
	}
	// duration := time.Since(start).String()
	duration := fmt.Sprintf("%.3f ms", float64(time.Since(start).Microseconds())/1000)

	return BubbleResponse{
		Before: ArrayToString(original),
		After:  ArrayToString(copyArr),
		Time:   duration,
	}
}
