package main

import (
	"runtime"
	"sync"
)

func parallelSearch(arr []int, target int) []int {
	numWorkers := runtime.NumCPU()
	chunkSize := len(arr) / numWorkers

	resultsChan := make(chan []int, numWorkers)
	var wg sync.WaitGroup

	for w := 0; w < numWorkers; w++ {
		start := w * chunkSize
		end := start + chunkSize

		if w == numWorkers-1 {
			end = len(arr)
		}

		wg.Add(1)
		go func(s, e int) {
			defer wg.Done()
			local := []int{}
			for i := s; i < e; i++ {
				if arr[i] == target {
					local = append(local, i)
				}
			}
			resultsChan <- local
		}(start, end)
	}

	go func() {
		wg.Wait()
		close(resultsChan)
	}()

	var results []int
	for local := range resultsChan {
		results = append(results, local...)
	}

	return results
}

// func main() {
// 	var n, target int
// 	fmt.Print("Masukkan jumlah elemen: ")
// 	fmt.Scan(&n)

// 	arr := make([]int, n)
// 	fmt.Println("Masukkan elemen array:")
// 	for i := 0; i < n; i++ {
// 		fmt.Scan(&arr[i])
// 	}

// 	fmt.Print("Masukkan target pencarian: ")
// 	fmt.Scan(&target)

// 	result := parallelSearch(arr, target)

// 	if len(result) == 0 {
// 		fmt.Println("Target tidak ditemukan")
// 	} else {
// 		fmt.Println("Target ditemukan pada indeks:", result)
// 	}
// }
