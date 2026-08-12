package main

import (
	"math"
	"runtime"
	"sync"
)

func parallelSimpson(f func(float64) float64, a, b float64, n int) float64 {
	if n%2 != 0 {
		panic("n harus genap untuk aturan Simpson")
	}

	h := (b - a) / float64(n)
	total := f(a) + f(b)

	numWorkers := runtime.NumCPU()

	var sumOdd, sumEven float64
	var wg sync.WaitGroup
	mu := &sync.Mutex{}

	worker := func(start, step int, isOdd bool) {
		defer wg.Done()
		localSum := 0.0
		for i := start; i < n; i += step {
			x := a + float64(i)*h
			localSum += f(x)
		}

		mu.Lock() // ✅ no shadowing
		if isOdd {
			sumOdd += localSum
		} else {
			sumEven += localSum
		}
		mu.Unlock()
	}

	// Odd indices
	for w := 0; w < numWorkers; w++ {
		start := 1 + 2*w
		if start < n {
			wg.Add(1)
			go worker(start, numWorkers*2, true)
		}
	}

	// Even indices
	for w := 0; w < numWorkers; w++ {
		start := 2 + 2*w
		if start < n {
			wg.Add(1)
			go worker(start, numWorkers*2, false)
		}
	}

	wg.Wait()

	total = total + 4*sumOdd + 2*sumEven
	return (h / 3) * total
}

func f(x float64) float64 { return math.Sin(x) }

// func main() {
// 	var a, b float64
// 	var n int

// 	fmt.Print("Masukkan batas bawah (a): ")
// 	fmt.Scan(&a)
// 	fmt.Print("Masukkan batas atas (b): ")
// 	fmt.Scan(&b)
// 	fmt.Print("Masukkan jumlah partisi (genap): ")
// 	fmt.Scan(&n)

// 	result := parallelSimpson(f, a, b, n)
// 	fmt.Println("Hasil integrasi:", result)
// }
