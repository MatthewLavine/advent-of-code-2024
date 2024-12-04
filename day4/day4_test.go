package main

import (
	"testing"
)

func TestStartsXmas(t *testing.T) {
	tests := []struct {
		matrix   [][]string
		row, col int
		expected int
	}{
		{
			matrix: [][]string{
				{"X", "M", "A", "S"},
				{"M", "M", "C", "D"},
				{"A", "F", "A", "H"},
				{"S", "J", "K", "S"},
			},
			row:      0,
			col:      0,
			expected: 3,
		},
		{
			matrix: [][]string{
				{"S", "A", "M", "X"},
				{"A", "A", "C", "D"},
				{"M", "F", "M", "H"},
				{"X", "J", "K", "X"},
			},
			row:      3,
			col:      3,
			expected: 1,
		},
		{
			matrix: [][]string{
				{"M", "A", "M", "X"},
				{"A", "A", "C", "D"},
				{"M", "F", "M", "H"},
				{"X", "J", "K", "B"},
			},
			row:      1,
			col:      1,
			expected: 0,
		},
	}

	for _, test := range tests {
		result := startsXmas(test.matrix, test.row, test.col)
		if result != test.expected {
			t.Errorf("startsXmas(%v, %d, %d) = %d; expected %d", test.matrix, test.row, test.col, result, test.expected)
		}
	}
}

func TestStartsXMas(t *testing.T) {
	tests := []struct {
		matrix   [][]string
		row, col int
		expected int
	}{
		{
			matrix: [][]string{
				{"M", "X", "M", "S"},
				{"Y", "A", "C", "D"},
				{"S", "F", "S", "H"},
				{"S", "J", "K", "S"},
			},
			row:      0,
			col:      0,
			expected: 1,
		},
		{
			matrix: [][]string{
				{"S", "A", "S", "X"},
				{"A", "A", "C", "D"},
				{"M", "F", "M", "H"},
				{"X", "J", "K", "X"},
			},
			row:      2,
			col:      2,
			expected: 1,
		},
		{
			matrix: [][]string{
				{"M", "A", "M", "X"},
				{"A", "A", "C", "D"},
				{"M", "F", "M", "H"},
				{"X", "J", "K", "B"},
			},
			row:      1,
			col:      1,
			expected: 0,
		},
	}

	for _, test := range tests {
		result := startsMas(test.matrix, test.row, test.col)
		if result != test.expected {
			t.Errorf("startsMas(%v, %d, %d) = %d; expected %d", test.matrix, test.row, test.col, result, test.expected)
		}
	}
}
