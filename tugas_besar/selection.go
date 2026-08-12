package main

import (
	"fmt"
	"math"
	"runtime"
	"sync"
	"time"
)

type MinResult struct {
	Value int
	Index int
}

func IsPrime(n int) bool {
	if n <= 1 {
		return false
	}
	if n == 2 || n == 3 {
		return true
	}
	if n%2 == 0 {
		return false
	}

	limit := int(math.Sqrt(float64(n)))
	for i := 3; i <= limit; i += 2 {
		if n%i == 0 {
			return false
		}
	}
	return true
}

func PrimesSequential(low, high int) []int {
	var primes []int
	for n := low; n <= high; n++ {
		if IsPrime(n) {
			primes = append(primes, n)
		}
	}
	return primes
}
func SelectionSort(arr []int) {
	for i := 0; i < len(arr)-1; i++ {
		minIdx := i
		for j := i + 1; j < len(arr); j++ {
			if arr[j] < arr[minIdx] {
				minIdx = j
			}
		}
		arr[i], arr[minIdx] = arr[minIdx], arr[i]
	}
}

func PrimesParallel(low, high int) []int {
	workers := runtime.NumCPU() * 2
	chunk := (high - low + 1) / workers

	var wg sync.WaitGroup
	results := make([][]int, workers)

	for w := 0; w < workers; w++ {
		wg.Add(1)
		go func(w int) {
			defer wg.Done()
			start := low + w*chunk
			end := start + chunk - 1
			if w == workers-1 {
				end = high
			}

			var local []int
			for n := start; n <= end; n++ {
				if IsPrime(n) {
					local = append(local, n)
				}
			}
			results[w] = local
		}(w)
	}
	wg.Wait()

	var merged []int
	for _, p := range results {
		merged = append(merged, p...)
	}
	return merged
}

func findMinParallel(arr []int, start, end int) MinResult {
	workers := runtime.NumCPU()
	chunk := (end - start + 1) / workers

	var wg sync.WaitGroup
	localMin := make([]MinResult, workers)

	for w := 0; w < workers; w++ {
		wg.Add(1)
		go func(w int) {
			defer wg.Done()
			s := start + w*chunk
			e := s + chunk - 1
			if w == workers-1 {
				e = end
			}

			minVal := arr[s]
			minIdx := s
			for i := s; i <= e; i++ {
				if arr[i] < minVal {
					minVal = arr[i]
					minIdx = i
				}
			}
			localMin[w] = MinResult{minVal, minIdx}
		}(w)
	}
	wg.Wait()

	result := localMin[0]
	for _, r := range localMin {
		if r.Value < result.Value {
			result = r
		}
	}
	return result
}

func SelectionSortParallel(arr []int) {
	for i := 0; i < len(arr)-1; i++ {
		min := findMinParallel(arr, i, len(arr)-1)
		arr[i], arr[min.Index] = arr[min.Index], arr[i]
	}
}

// func RunSelectionSort(size int, parallel bool) map[string]interface{} {
// 	arr := make([]int, size)
// 	for i := 0; i < size; i++ {
// 		arr[i] = size - i
// 	}

// 	before := make([]int, size)
// 	copy(before, arr)

// 	var start = now()
// 	if parallel {
// 		SelectionSortParallel(arr)
// 	} else {
// 		SelectionSort(arr)
// 	}
// 	duration := now() - start

// 	return map[string]interface{}{
// 		"time":  duration,
// 		"before": before,
// 		"after": arr,
// 	}
// }

// ...

func RunSelectionSort(size int, parallel bool) map[string]interface{} {
	arr := make([]int, size)
	for i := 0; i < size; i++ {
		arr[i] = size - i
	}

	before := make([]int, size)
	copy(before, arr)

	start := time.Now()
	if parallel {
		SelectionSortParallel(arr)
	} else {
		SelectionSort(arr)
	}
	// duration := time.Since(start).Milliseconds()
	duration := fmt.Sprintf("%.3f ms", float64(time.Since(start).Microseconds())/1000)

	return map[string]interface{}{
		"time":   duration,
		"before": before,
		"after":  arr,
	}
}
