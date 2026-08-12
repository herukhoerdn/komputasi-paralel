package main

import (
	"runtime"
	"sync"
)

func parallelGaussJordan(A [][]float64, b []float64) []float64 {
	n := len(A)
	numWorkers := runtime.NumCPU()
	var wg sync.WaitGroup

	for k := 0; k < n; k++ {
		pivot := A[k][k]
		for j := k; j < n; j++ {
			A[k][j] /= pivot
		}
		b[k] /= pivot

		wg.Add(numWorkers)
		for w := 0; w < numWorkers; w++ {
			go func(worker int) {
				defer wg.Done()
				for i := worker; i < n; i += numWorkers {
					if i != k {
						factor := A[i][k]
						for j := k; j < n; j++ {
							A[i][j] -= factor * A[k][j]
						}
						b[i] -= factor * b[k]
					}
				}
			}(w)
		}
		wg.Wait()
	}

	return b
}

// func main() {
// 	var n int
// 	fmt.Print("Masukkan jumlah variabel (n): ")
// 	fmt.Scan(&n)

// 	A := make([][]float64, n)
// 	for i := range A {
// 		A[i] = make([]float64, n)
// 	}

// 	b := make([]float64, n)

// 	fmt.Println("Masukkan elemen matriks A:")
// 	for i := 0; i < n; i++ {
// 		for j := 0; j < n; j++ {
// 			fmt.Scan(&A[i][j])
// 		}
// 	}

// 	fmt.Println("Masukkan elemen vektor b:")
// 	for i := 0; i < n; i++ {
// 		fmt.Scan(&b[i])
// 	}

// 	result := parallelGaussJordan(A, b)

// 	fmt.Println("Solusi sistem persamaan:")
// 	for i := 0; i < n; i++ {
// 		fmt.Printf("x%d = %.6f\n", i+1, result[i])
// 	}
// }
