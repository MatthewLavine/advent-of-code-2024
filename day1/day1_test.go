package main

import (
	"testing"
)

func TestCalculateListDistance(t *testing.T) {
	tests := []struct {
		left     []int
		right    []int
		expected int
		err      bool
	}{
		{[]int{1, 2, 3}, []int{4, 5, 6}, 9, false},
		{[]int{1, 2, 3}, []int{1, 2, 3}, 0, false},
		{[]int{1, 2, 3}, []int{3, 2, 1}, 0, false},
		{[]int{1, 2}, []int{1, 2, 3}, -1, true},
		{[]int{1, 2, 3}, []int{1, 2}, -1, true},
	}

	for _, test := range tests {
		distance, err := calculateListDistance(test.left, test.right)
		if (err != nil) != test.err {
			t.Errorf("calculateListDistance(%v, %v) returned error %v, expected error: %v", test.left, test.right, err, test.err)
		}
		if distance != test.expected {
			t.Errorf("calculateListDistance(%v, %v) = %d, expected %d", test.left, test.right, distance, test.expected)
		}
	}
}
func TestCalculateListSimilarity(t *testing.T) {
	tests := []struct {
		left     []int
		right    []int
		expected int
		err      bool
	}{
		{[]int{1, 2, 3}, []int{4, 5, 6}, 0, false},
		{[]int{1, 2, 3}, []int{1, 2, 3}, 6, false},
		{[]int{1, 2, 3}, []int{3, 2, 1}, 6, false},
		{[]int{1, 2}, []int{1, 2, 3}, -1, true},
		{[]int{1, 2, 3}, []int{1, 2}, -1, true},
		{[]int{1, 1, 1}, []int{1, 1, 1}, 9, false},
		{[]int{1, 2, 2}, []int{2, 2, 2}, 12, false},
	}

	for _, test := range tests {
		similarity, err := calculateListSimilarity(test.left, test.right)
		if (err != nil) != test.err {
			t.Errorf("calculateListSimilarity(%v, %v) returned error %v, expected error: %v", test.left, test.right, err, test.err)
		}
		if similarity != test.expected {
			t.Errorf("calculateListSimilarity(%v, %v) = %d, expected %d", test.left, test.right, similarity, test.expected)
		}
	}
}
