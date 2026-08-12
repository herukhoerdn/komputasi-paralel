package main

import (
	"math"
	"runtime"
	"sync"
)

func parallelSieve(n int) []int {
	sieve := make([]bool, n+1)
	for i := range sieve {
		sieve[i] = true
	}
	sieve[0], sieve[1] = false, false

	limit := int(math.Sqrt(float64(n)))
	numWorkers := runtime.NumCPU()

	var wg sync.WaitGroup
	for x := 2; x <= limit; x++ {
		if sieve[x] {
			wg.Add(1)
			go func(i int) {
				defer wg.Done()
				for j := i * i; j <= n; j += i {
					sieve[j] = false
				}
			}(x)
		}
	}
	wg.Wait()

	chunkSize := n / numWorkers
	primesChan := make(chan int, n)

	var wg2 sync.WaitGroup
	for t := 0; t < numWorkers; t++ {
		start := t * chunkSize
		end := start + chunkSize
		if t == numWorkers-1 {
			end = n
		}

		wg2.Add(1)
		go func(s, e int) {
			defer wg2.Done()
			for i := s; i <= e; i++ {
				if sieve[i] {
					primesChan <- i
				}
			}
		}(start, end)
	}

	go func() {
		wg2.Wait()
		close(primesChan)
	}()

	primes := []int{}
	for p := range primesChan {
		primes = append(primes, p)
	}

	return primes
}

// func main() {
// 	var n int
// 	fmt.Print("Masukkan angka maksimal: ")
// 	fmt.Scan(&n)

// 	primes := parallelSieve(n)
// 	fmt.Println("Bilangan prima ditemukan:")
// 	fmt.Println(primes)
// }
