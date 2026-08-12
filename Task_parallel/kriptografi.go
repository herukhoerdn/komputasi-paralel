package main

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"runtime"
	"sync"
)

func ParallelHashCipher(plaintext string) string {
	numThreads := runtime.NumCPU()
	chunks := splitString(plaintext, numThreads)

	hashes := make([]string, len(chunks))
	var wg sync.WaitGroup

	for i := 0; i < len(chunks); i++ {
		wg.Add(1)
		go func(idx int) {
			defer wg.Done()
			hashes[idx] = hashMD5(chunks[idx])
		}(i)
	}

	wg.Wait()
	return mergeHashes(hashes)
}

func hashMD5(input string) string {

	hash := md5.Sum([]byte(input))
	return hex.EncodeToString(hash[:])
}

func splitString(text string, parts int) []string {
	length := len(text)
	chunkSize := length / parts
	var chunks []string

	for i := 0; i < length; i += chunkSize {
		end := i + chunkSize
		if end > length {
			end = length
		}
		chunks = append(chunks, text[i:end])
	}

	return chunks
}

func mergeHashes(hashes []string) string {

	combined := ""
	for _, h := range hashes {
		combined += h
	}
	sum := md5.Sum([]byte(combined))
	return hex.EncodeToString(sum[:])
}
func main() {
	plaintext := "Heru Khoerudin Kelas Pagi A Nim 231351059"

	finalHash := ParallelHashCipher(plaintext)
	fmt.Println("Final Hash:", finalHash)
}
