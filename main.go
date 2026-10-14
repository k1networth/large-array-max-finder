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

func generateRandomElements(size int) []int {
	if size <= 0 {
		return nil
	}

	data := make([]int, size)
	for index := range data {
		data[index] = rand.Intn(1_000_000) + 1
	}
	return data
}

func maximum(data []int) int {
	if len(data) == 0 {
		return 0
	}

	maxValue := data[0]
	for _, value := range data[1:] {
		if value > maxValue {
			maxValue = value
		}
	}
	return maxValue
}

func maxChunks(data []int) int {
	if len(data) == 0 {
		return 0
	}

	chunkSize := len(data) / CHUNKS
	chunkMaxima := make([]int, CHUNKS)
	var waitGroup sync.WaitGroup
	waitGroup.Add(CHUNKS)

	for chunkIndex := 0; chunkIndex < CHUNKS; chunkIndex++ {
		start := chunkIndex * chunkSize
		end := start + chunkSize
		if chunkIndex == CHUNKS-1 {
			end = len(data)
		}

		go func(index int, chunk []int) {
			defer waitGroup.Done()
			if len(chunk) == 0 {
				chunkMaxima[index] = data[0]
				return
			}
			chunkMaxima[index] = maximum(chunk)
		}(chunkIndex, data[start:end])
	}

	waitGroup.Wait()
	return maximum(chunkMaxima)
}

func main() {
	fmt.Printf("Генерируем %d целых положительных чисел\n", SIZE)
	data := generateRandomElements(SIZE)

	fmt.Println("Ищем максимальное значение в один поток")
	start := time.Now()
	maxValue := maximum(data)
	elapsed := time.Since(start).Microseconds()

	fmt.Printf("Максимальное значение элемента: %d\nВремя поиска: %d мкс\n", maxValue, elapsed)

	fmt.Printf("Ищем максимальное значение в %d горутинах\n", CHUNKS)
	start = time.Now()
	maxValue = maxChunks(data)
	elapsed = time.Since(start).Microseconds()

	fmt.Printf("Максимальное значение элемента: %d\nВремя поиска: %d мкс\n", maxValue, elapsed)
}
