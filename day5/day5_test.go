package main

import (
	"reflect"
	"testing"
)

func TestParseInput(t *testing.T) {
	tests := []struct {
		input           string
		expectedRules   []rule
		expectedUpdates [][]int
		expectError     bool
	}{
		{
			input: "1|2\n3,4,5\n6|7\n8,9,10",
			expectedRules: []rule{
				{l: 1, r: 2},
				{l: 6, r: 7},
			},
			expectedUpdates: [][]int{
				{3, 4, 5},
				{8, 9, 10},
			},
			expectError: false,
		},
		{
			input:           "1|a\n3,4,5",
			expectedRules:   nil,
			expectedUpdates: nil,
			expectError:     true,
		},
		{
			input: "1|2\n3,4,5\n6|7\n8,9,10\n",
			expectedRules: []rule{
				{l: 1, r: 2},
				{l: 6, r: 7},
			},
			expectedUpdates: [][]int{
				{3, 4, 5},
				{8, 9, 10},
			},
			expectError: false,
		},
	}

	for _, test := range tests {
		rules, updates, err := parseInput(test.input)
		if test.expectError {
			if err == nil {
				t.Errorf("Expected error but got none")
			}
		} else {
			if err != nil {
				t.Errorf("Unexpected error: %v", err)
			}
			if !reflect.DeepEqual(rules, test.expectedRules) {
				t.Errorf("Expected rules %v, but got %v", test.expectedRules, rules)
			}
			if !reflect.DeepEqual(updates, test.expectedUpdates) {
				t.Errorf("Expected updates %v, but got %v", test.expectedUpdates, updates)
			}
		}
	}
}
func TestComputePart1(t *testing.T) {
	tests := []struct {
		rules       []rule
		updates     [][]int
		expectedSum int
		expectError bool
	}{
		{
			rules: []rule{
				{l: 1, r: 2},
				{l: 6, r: 7},
			},
			updates: [][]int{
				{1, 2, 3},
				{6, 7, 8},
			},
			expectedSum: 9,
			expectError: false,
		},
		{
			rules: []rule{
				{l: 1, r: 2},
			},
			updates: [][]int{
				{1, 3, 2},
			},
			expectedSum: 3,
			expectError: false,
		},
		{
			rules: []rule{
				{l: 1, r: 2},
			},
			updates: [][]int{
				{2, 1, 3},
				{4, 5, 6},
			},
			expectedSum: 5,
			expectError: false,
		},
		{
			rules: []rule{
				{l: 1, r: 2},
				{l: 3, r: 4},
			},
			updates: [][]int{
				{1, 2, 3, 4, 7},
				{4, 3, 2, 1, 1},
			},
			expectedSum: 3,
			expectError: false,
		},
	}

	for _, test := range tests {
		sum, err := computePart1(test.rules, test.updates)
		if test.expectError {
			if err == nil {
				t.Errorf("Expected error but got none")
			}
		} else {
			if err != nil {
				t.Errorf("Unexpected error: %v", err)
			}
			if sum != test.expectedSum {
				t.Errorf("Expected sum %d, but got %d", test.expectedSum, sum)
			}
		}
	}
}
func TestComputePart2(t *testing.T) {
	tests := []struct {
		rules       []rule
		updates     [][]int
		expectedSum int
		expectError bool
	}{
		{
			rules: []rule{
				{l: 1, r: 2},
				{l: 6, r: 7},
			},
			updates: [][]int{
				{2, 1, 3},
				{8, 7, 6},
			},
			expectedSum: 8,
			expectError: false,
		},
		{
			rules: []rule{
				{l: 1, r: 2},
			},
			updates: [][]int{
				{2, 3, 1},
			},
			expectedSum: 3,
			expectError: false,
		},
		{
			rules: []rule{
				{l: 1, r: 2},
			},
			updates: [][]int{
				{2, 1, 3},
				{4, 5, 6},
			},
			expectedSum: 2,
			expectError: false,
		},
		{
			rules: []rule{
				{l: 1, r: 2},
				{l: 3, r: 4},
			},
			updates: [][]int{
				{1, 2, 4, 3, 7},
				{4, 3, 2, 1, 1},
			},
			expectedSum: 4,
			expectError: false,
		},
		{
			rules: []rule{
				{l: 1, r: 2},
				{l: 3, r: 4},
			},
			updates: [][]int{
				{2, 4, 3, 1, 7},
				{4, 3, 2, 1, 1},
			},
			expectedSum: 5,
			expectError: false,
		},
	}

	for _, test := range tests {
		sum, err := computePart2(test.rules, test.updates)
		if test.expectError {
			if err == nil {
				t.Errorf("Expected error but got none")
			}
		} else {
			if err != nil {
				t.Errorf("Unexpected error: %v", err)
			}
			if sum != test.expectedSum {
				t.Errorf("Expected sum %d, but got %d", test.expectedSum, sum)
			}
		}
	}
}
func TestMakeUpdateSatisfyRule(t *testing.T) {
	tests := []struct {
		update         []int
		rule           rule
		expectedUpdate []int
	}{
		{
			update:         []int{1, 2, 3},
			rule:           rule{l: 1, r: 2},
			expectedUpdate: []int{1, 2, 3},
		},
		{
			update:         []int{2, 1, 3},
			rule:           rule{l: 1, r: 2},
			expectedUpdate: []int{1, 2, 3},
		},
		{
			update:         []int{4, 3, 2, 1},
			rule:           rule{l: 1, r: 2},
			expectedUpdate: []int{4, 3, 1, 2},
		},
		{
			update:         []int{4, 3, 2, 1},
			rule:           rule{l: 3, r: 4},
			expectedUpdate: []int{3, 4, 2, 1},
		},
		{
			update:         []int{4, 1, 3, 2},
			rule:           rule{l: 1, r: 2},
			expectedUpdate: []int{4, 1, 3, 2},
		},
	}

	for _, test := range tests {
		result := makeUpdateSatisfyRule(test.update, test.rule)
		if !reflect.DeepEqual(result, test.expectedUpdate) {
			t.Errorf("Expected update %v, but got %v", test.expectedUpdate, result)
		}
	}
}
