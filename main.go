package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

const (
	SIZE   = 100_000_000
	CHUNKS = 8
)

// generateRandomElements returns size strictly positive random integers.
// A nonpositive size produces an empty slice.
func generateRandomElements(size int) []int {
	if size <= 0 {
		return []int{}
	}
	data := make([]int, size)
	maxInt := int(^uint(0) >> 1)
	for index := range data {
		data[index] = rand.Intn(maxInt) + 1
	}
	return data
}

// maximum returns the largest element, or zero for an empty slice.
func maximum(data []int) int {
	if len(data) == 0 {
		return 0
	}
	result := data[0]
	for _, value := range data[1:] {
		if value > result {
			result = value
		}
	}
	return result
}

// maxChunks finds the maximum using CHUNKS independent goroutines.
func maxChunks(data []int) int {
	if len(data) == 0 {
		return 0
	}
	maxima := make([]int, CHUNKS)
	chunkSize := len(data) / CHUNKS
	var workers sync.WaitGroup
	workers.Add(CHUNKS)
	for chunk := 0; chunk < CHUNKS; chunk++ {
		start := chunk * chunkSize
		end := start + chunkSize
		// Include the remainder, even when the input has fewer than CHUNKS elements.
		if chunk == CHUNKS-1 {
			end = len(data)
		}
		go func(index int, part []int) {
			defer workers.Done()
			if len(part) == 0 {
				// Use a real input value so empty chunks cannot mask a negative maximum.
				maxima[index] = data[0]
				return
			}
			maxima[index] = maximum(part)
		}(chunk, data[start:end])
	}
	workers.Wait()
	return maximum(maxima)
}

func main() {
	fmt.Printf("Генерируем %d целых чисел\n", SIZE)
	data := generateRandomElements(SIZE)

	fmt.Println("Ищем максимальное значение в один поток")
	started := time.Now()
	max := maximum(data)
	elapsed := time.Since(started).Microseconds()
	fmt.Printf("Максимальное значение элемента: %d\nВремя поиска: %d мкс\n", max, elapsed)

	fmt.Printf("Ищем максимальное значение в %d потоков\n", CHUNKS)
	started = time.Now()
	max = maxChunks(data)
	elapsed = time.Since(started).Microseconds()
	fmt.Printf("Максимальное значение элемента: %d\nВремя поиска: %d мкс\n", max, elapsed)
}
