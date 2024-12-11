package main

import (
	"testing"
)

func TestNextPos(t *testing.T) {
	tests := []struct {
		direction   string
		currRow     int
		currCol     int
		expectedRow int
		expectedCol int
	}{
		{"up", 1, 1, 0, 1},
		{"down", 1, 1, 2, 1},
		{"left", 1, 1, 1, 0},
		{"right", 1, 1, 1, 2},
	}

	for _, test := range tests {
		row, col := nextPos(test.direction, test.currRow, test.currCol)
		if row != test.expectedRow || col != test.expectedCol {
			t.Errorf("nextPos(%s, %d, %d) = (%d, %d); want (%d, %d)", test.direction, test.currRow, test.currCol, row, col, test.expectedRow, test.expectedCol)
		}
	}
}

func TestTurn(t *testing.T) {
	tests := []struct {
		direction string
		expected  string
	}{
		{"up", "right"},
		{"right", "down"},
		{"down", "left"},
		{"left", "up"},
	}

	for _, test := range tests {
		result := turn(test.direction)
		if result != test.expected {
			t.Errorf("turn(%s) = %s; want %s", test.direction, result, test.expected)
		}
	}
}
func TestParseMap(t *testing.T) {
	tests := []struct {
		input       string
		expectedMap [][]string
		expectedCol int
		expectedRow int
		expectedErr bool
	}{
		{
			input: "...\n.^.\n...",
			expectedMap: [][]string{
				{".", ".", "."},
				{".", "^", "."},
				{".", ".", "."},
			},
			expectedCol: 1,
			expectedRow: 1,
			expectedErr: false,
		},
		{
			input: "###\n#^#\n###",
			expectedMap: [][]string{
				{"#", "#", "#"},
				{"#", "^", "#"},
				{"#", "#", "#"},
			},
			expectedCol: 1,
			expectedRow: 1,
			expectedErr: false,
		},
		{
			input: "^..\n...\n...",
			expectedMap: [][]string{
				{"^", ".", "."},
				{".", ".", "."},
				{".", ".", "."},
			},
			expectedCol: 0,
			expectedRow: 0,
			expectedErr: false,
		},
	}

	for _, test := range tests {
		m, col, row, err := parseMap(test.input)
		if (err != nil) != test.expectedErr {
			t.Errorf("parseMap(%q) error = %v, expectedErr %v", test.input, err, test.expectedErr)
			continue
		}
		if !compareMaps(m, test.expectedMap) || col != test.expectedCol || row != test.expectedRow {
			t.Errorf("parseMap(%q) = (%v, %d, %d), want (%v, %d, %d)", test.input, m, col, row, test.expectedMap, test.expectedCol, test.expectedRow)
		}
	}
}

func compareMaps(a, b [][]string) bool {
	if len(a) != len(b) {
		return false
	}
	for i := range a {
		if len(a[i]) != len(b[i]) {
			return false
		}
		for j := range a[i] {
			if a[i][j] != b[i][j] {
				return false
			}
		}
	}
	return true
}
func TestTraverse(t *testing.T) {
	tests := []struct {
		input       string
		startingCol int
		startingRow int
		expected    int
	}{
		{
			input:       "...\n.^.\n...",
			startingCol: 1,
			startingRow: 1,
			expected:    2,
		},
		{
			input:       "^..\n...\n...",
			startingCol: 0,
			startingRow: 0,
			expected:    1,
		},
		{
			input:       ".#.\n.^#\n...",
			startingCol: 1,
			startingRow: 1,
			expected:    2,
		},
	}

	for _, test := range tests {
		m, _, _, err := parseMap(test.input)
		if err != nil {
			t.Fatalf("parseMap(%q) error = %v", test.input, err)
		}
		result := traverse(m, test.startingCol, test.startingRow)
		if result != test.expected {
			t.Errorf("traverse(%q, %d, %d) = %d; want %d", test.input, test.startingCol, test.startingRow, result, test.expected)
		}
	}
}

func BenchmarkTraverse(b *testing.B) {
	m := [][]string{
		{".", ".", "#", ".", "."},
		{".", ".", ".", ".", "#"},
		{".", "#", ".", ".", "."},
		{".", ".", ".", ".", "."},
		{".", ".", "^", ".", "."},
		{".", ".", ".", ".", "."},
		{".", ".", ".", ".", "."},
		{"#", ".", ".", ".", "."},
		{".", ".", ".", "#", "."},
	}
	for i := 0; i < b.N; i++ {
		traverse(m, 2, 4)
	}
}
