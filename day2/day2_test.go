package main

import (
	"testing"
)

func TestIsReportSafe(t *testing.T) {
	tests := []struct {
		report   []int
		dampener bool
		expected bool
	}{
		{[]int{1, 2, 3, 4}, false, true},
		{[]int{4, 3, 2, 1}, false, true},
		{[]int{1, 3, 2, 4}, false, false},
		{[]int{1, 2, 8}, false, false},
		{[]int{1, 6, 3, 6}, false, false},
		{[]int{1, 6, 3, 6}, true, true},
		{[]int{5, 6, 4, 3}, true, true},
		{[]int{5, 4, 6, 2}, true, true},
		{[]int{56, 55, 57, 58, 59, 60}, true, true},
		{[]int{10, 1, 2, 3, 4, 5}, true, true},
	}

	for _, test := range tests {
		result := isReportSafe(test.report, test.dampener)
		if result != test.expected {
			t.Errorf("isReportSafe(%v, %v) = %v; expected %v", test.report, test.dampener, result, test.expected)
		}
	}
}
