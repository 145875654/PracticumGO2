package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestGenerateRandomElements(t *testing.T) {
	tests := []struct {
		name    string
		size    int
		wantLen int
	}{
		{name: "нулевой размер", size: 0, wantLen: 0},
		{name: "отрицательный размер", size: -5, wantLen: 0},
		{name: "один элемент", size: 1, wantLen: 1},
		{name: "обычный размер", size: 1000, wantLen: 1000},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := generateRandomElements(tt.size)

			require.Len(t, got, tt.wantLen)
			for _, value := range got {
				assert.Positive(t, value, "все числа должны быть положительными")
			}
		})
	}
}

func TestMaximum(t *testing.T) {
	tests := []struct {
		name string
		data []int
		want int
	}{
		{name: "nil-слайс", data: nil, want: 0},
		{name: "пустой слайс", data: []int{}, want: 0},
		{name: "один элемент", data: []int{7}, want: 7},
		{name: "максимум в начале", data: []int{9, 3, 5}, want: 9},
		{name: "максимум в середине", data: []int{3, 9, 5}, want: 9},
		{name: "максимум в конце", data: []int{3, 5, 9}, want: 9},
		{name: "одинаковые элементы", data: []int{4, 4, 4}, want: 4},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.want, maximum(tt.data))
		})
	}
}

func TestMaxChunks(t *testing.T) {
	tests := []struct {
		name string
		data []int
		want int
	}{
		{name: "nil-слайс", data: nil, want: 0},
		{name: "пустой слайс", data: []int{}, want: 0},
		{name: "один элемент", data: []int{7}, want: 7},
		{name: "меньше элементов, чем частей", data: []int{1, 5, 3}, want: 5},
		{name: "ровно по числу частей", data: []int{1, 2, 3, 4, 5, 6, 7, 8}, want: 8},
		{name: "максимум в первой части", data: []int{99, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15}, want: 99},
		{name: "максимум в остатке (последняя часть)", data: []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 99}, want: 99},
		{name: "длина не делится на число частей", data: []int{3, 1, 4, 1, 5, 9, 2, 6, 5, 3, 5, 8, 9, 7, 9}, want: 9},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.want, maxChunks(tt.data))
		})
	}
}

func TestMaxChunksMatchesMaximum(t *testing.T) {
	for _, size := range []int{1, CHUNKS - 1, CHUNKS, CHUNKS + 1, 1000, 100_003} {
		data := generateRandomElements(size)
		assert.Equal(t, maximum(data), maxChunks(data), "размер %d", size)
	}
}
