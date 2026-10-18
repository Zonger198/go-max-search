package main

import (
	"reflect"
	"testing"
)

func TestGenerateRandomElements(t *testing.T) {
	for _, size := range []int{-10, 0, 1, 7, 1000} {
		data := generateRandomElements(size)
		expectedSize := size
		if expectedSize < 0 {
			expectedSize = 0
		}
		if len(data) != expectedSize {
			t.Fatalf("size %d: got length %d, want %d", size, len(data), expectedSize)
		}
		for index, value := range data {
			if value <= 0 {
				t.Fatalf("size %d: element %d is not positive: %d", size, index, value)
			}
		}
	}
}

func TestMaximumFunctions(t *testing.T) {
	maxInt := int(^uint(0) >> 1)
	minInt := -maxInt - 1
	tests := []struct {
		name string
		data []int
		want int
	}{
		{"nil", nil, 0},
		{"empty", []int{}, 0},
		{"single", []int{42}, 42},
		{"single negative", []int{-42}, -42},
		{"zeros", []int{0, 0}, 0},
		{"equal", []int{5, 5, 5}, 5},
		{"maximum first", []int{90, 3, 2}, 90},
		{"maximum middle", []int{1, 90, 2}, 90},
		{"maximum last", []int{1, 2, 90}, 90},
		{"negative values", []int{-9, -2, -7}, -2},
		{"integer limits", []int{minInt, 0, maxInt}, maxInt},
		{"minimum integer", []int{minInt}, minInt},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			for name, find := range map[string]func([]int) int{"maximum": maximum, "maxChunks": maxChunks} {
				before := append([]int{}, test.data...)
				if got := find(test.data); got != test.want {
					t.Errorf("%s: got %d, want %d", name, got, test.want)
				}
				if !reflect.DeepEqual(before, append([]int{}, test.data...)) {
					t.Errorf("%s modified the input", name)
				}
			}
		})
	}
}

// Put the maximum at every position across small, exact, and uneven partitions.
func TestMaxChunksPartitions(t *testing.T) {
	for _, size := range []int{1, 7, 8, 9, 15, 16, 17, 65} {
		for position := 0; position < size; position++ {
			data := make([]int, size)
			for index := range data {
				data[index] = -100
			}
			data[position] = -1
			if got := maxChunks(data); got != -1 {
				t.Fatalf("size %d, maximum at %d: got %d, want -1", size, position, got)
			}
		}
	}
}
