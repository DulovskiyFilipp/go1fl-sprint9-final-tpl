package main

import (
	"testing"
)

func TestGenerateRandomElements(t *testing.T) {
	tests := []struct {
		name string
		size int
		want int
	}{
		{"zero size", 0, 0},
		{"negative size", -5, 0},
		{"one element", 1, 1},
		{"many elements", 1000, 1000},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := generateRandomElements(tt.size)
			if got == nil {
				t.Fatal("expected non-nil slice")
			}
			if len(got) != tt.want {
				t.Fatalf("len = %d, want %d", len(got), tt.want)
			}
			for _, v := range got {
				if v <= 0 {
					t.Fatalf("expected positive value, got %d", v)
				}
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
		{"nil slice", nil, 0},
		{"empty slice", []int{}, 0},
		{"one element", []int{7}, 7},
		{"max at start", []int{9, 1, 2}, 9},
		{"max in middle", []int{1, 9, 2}, 9},
		{"max at end", []int{1, 2, 9}, 9},
		{"equal elements", []int{3, 3, 3}, 3},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := maximum(tt.data); got != tt.want {
				t.Fatalf("maximum() = %d, want %d", got, tt.want)
			}
		})
	}
}

func TestMaxChunks(t *testing.T) {
	tests := []struct {
		name string
		data []int
		want int
	}{
		{"nil slice", nil, 0},
		{"empty slice", []int{}, 0},
		{"one element", []int{5}, 5},
		{"fewer than chunks", []int{1, 8, 3}, 8},
		{"exactly chunks", []int{1, 2, 3, 4, 5, 6, 7, 8}, 8},
		{"not divisible, max in remainder", []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 99}, 99},
		{"divisible", []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16}, 16},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := maxChunks(tt.data); got != tt.want {
				t.Fatalf("maxChunks() = %d, want %d", got, tt.want)
			}
		})
	}
}

func TestMaxChunksMatchesMaximum(t *testing.T) {
	data := generateRandomElements(100_003)
	if got, want := maxChunks(data), maximum(data); got != want {
		t.Fatalf("maxChunks() = %d, maximum() = %d", got, want)
	}
}