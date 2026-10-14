package main

import (
	"fmt"
	"slices"
	"testing"
)

func TestGenerateRandomElements(t *testing.T) {
	for _, size := range []int{-10, -1, 0, 1, CHUNKS, 1_000} {
		t.Run(fmt.Sprintf("size_%d", size), func(t *testing.T) {
			data := generateRandomElements(size)
			wantLength := size
			if size < 0 {
				wantLength = 0
			}
			if len(data) != wantLength {
				t.Fatalf("len(generateRandomElements(%d)) = %d, want %d", size, len(data), wantLength)
			}
			for index, value := range data {
				if value <= 0 || value > 1_000_000 {
					t.Errorf("data[%d] = %d, want a value in [1, 1_000_000]", index, value)
				}
			}
		})
	}
}

func TestMaximum(t *testing.T) {
	testMaximum(t, maximum)
}

func TestMaxChunks(t *testing.T) {
	testMaximum(t, maxChunks)
}

func testMaximum(t *testing.T, findMaximum func([]int) int) {
	t.Helper()
	maxInt := int(^uint(0) >> 1)
	minInt := -maxInt - 1
	testCases := []struct {
		name string
		data []int
		want int
	}{
		{name: "nil", data: nil, want: 0},
		{name: "empty", data: []int{}, want: 0},
		{name: "single_positive", data: []int{42}, want: 42},
		{name: "single_zero", data: []int{0}, want: 0},
		{name: "single_negative", data: []int{-42}, want: -42},
		{name: "maximum_at_start", data: []int{9, 4, 2}, want: 9},
		{name: "maximum_in_middle", data: []int{2, 9, 4}, want: 9},
		{name: "maximum_at_end", data: []int{2, 4, 9}, want: 9},
		{name: "equal_values", data: []int{5, 5, 5}, want: 5},
		{name: "all_zero", data: []int{0, 0, 0}, want: 0},
		{name: "all_negative", data: []int{-9, -2, -5}, want: -2},
		{name: "mixed_signs", data: []int{-9, 0, 12, -5}, want: 12},
		{name: "integer_limits", data: []int{minInt, 0, maxInt}, want: maxInt},
		{name: "minimum_integer", data: []int{minInt}, want: minInt},
		{name: "eight_elements", data: []int{1, 9, 2, 3, 4, 5, 6, 7}, want: 9},
		{name: "negative_full_chunks", data: []int{-8, -7, -6, -5, -4, -3, -2, -1}, want: -1},
		{name: "maximum_in_remainder", data: []int{1, 2, 3, 4, 5, 6, 7, 8, 99}, want: 99},
	}
	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			original := slices.Clone(testCase.data)
			if got := findMaximum(testCase.data); got != testCase.want {
				t.Errorf("maximum(%v) = %d, want %d", testCase.data, got, testCase.want)
			}
			if !slices.Equal(testCase.data, original) {
				t.Errorf("input changed: got %v, want %v", testCase.data, original)
			}
		})
	}
}

func TestMaxChunksEveryPosition(t *testing.T) {
	for size := 1; size <= 3*CHUNKS; size++ {
		for maxIndex := 0; maxIndex < size; maxIndex++ {
			t.Run(fmt.Sprintf("size_%d_max_at_%d", size, maxIndex), func(t *testing.T) {
				data := make([]int, size)
				for index := range data {
					data[index] = -index - 2
				}
				data[maxIndex] = -1
				if got := maxChunks(data); got != -1 {
					t.Errorf("maxChunks(%v) = %d, want -1", data, got)
				}
			})
		}
	}
}
