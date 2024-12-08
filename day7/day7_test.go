package main

import (
	"testing"
)

func TestConcatNumbers(t *testing.T) {
	tests := []struct {
		a, b     int
		expected int
	}{
		{12, 34, 1234},
		{56, 78, 5678},
		{0, 123, 123},
		{123, 0, 1230},
		{1, 1, 11},
	}

	for _, test := range tests {
		result, err := concatNumbers(test.a, test.b)
		if err != nil {
			t.Errorf("concatNumbers(%d, %d) returned error: %v", test.a, test.b, err)
		}
		if result != test.expected {
			t.Errorf("concatNumbers(%d, %d) = %d; want %d", test.a, test.b, result, test.expected)
		}
	}
}

func TestGenerateOperatorPermutations(t *testing.T) {
	tests := []struct {
		operators []string
		length    int
		expected  []string
	}{
		{[]string{"+", "*"}, 1, []string{"+", "*"}},
		{[]string{"+", "*"}, 2, []string{"++", "+*", "*+", "**"}},
		{[]string{"+", "*", "|"}, 1, []string{"+", "*", "|"}},
		{[]string{"+", "*", "|"}, 2, []string{"++", "+*", "+|", "*+", "**", "*|", "|+", "|*", "||"}},
	}

	for _, test := range tests {
		result := generateOperatorPermutations(test.operators, test.length)
		if len(result) != len(test.expected) {
			t.Errorf("generateOperatorPermutations(%v, %d) = %v; want %v", test.operators, test.length, result, test.expected)
			continue
		}
		for i, v := range result {
			if v != test.expected[i] {
				t.Errorf("generateOperatorPermutations(%v, %d) = %v; want %v", test.operators, test.length, result, test.expected)
				break
			}
		}
	}
}

func TestComputeWithPerms(t *testing.T) {
	tests := []struct {
		total     int
		numbers   []int
		operators []string
		expected  int
	}{
		{10, []int{1, 2, 3, 4}, []string{"+", "*"}, 10},
		{24, []int{2, 3, 4}, []string{"+", "*"}, 24},
		{1234, []int{12, 34}, []string{"|"}, 1234},
		{2, []int{1, 1}, []string{"+"}, 2},
		{11, []int{1, 1}, []string{"|"}, 11},
	}

	for _, test := range tests {
		result, err := computeWithPerms(test.total, test.numbers, test.operators)
		if err != nil {
			t.Errorf("computeWithPerms(%d, %v, %v) returned error: %v", test.total, test.numbers, test.operators, err)
		}
		if result != test.expected {
			t.Errorf("computeWithPerms(%d, %v, %v) = %d; want %d", test.total, test.numbers, test.operators, result, test.expected)
		}
	}
}

func TestParse(t *testing.T) {
	tests := []struct {
		input    string
		expected map[int][]int
		err      bool
	}{
		{
			input: "10: 1 2 3 4\n22: 4 7 6",
			expected: map[int][]int{
				10: {1, 2, 3, 4},
				22: {4, 7, 6},
			},
			err: false,
		},
		{
			input: "1234: 12 34",
			expected: map[int][]int{
				1234: {12, 34},
			},
			err: false,
		},
		{
			input:    "invalid",
			expected: nil,
			err:      true,
		},
		{
			input:    "10: 1 2 3 4\ninvalid\n",
			expected: nil,
			err:      true,
		},
	}

	for _, test := range tests {
		result, err := parse(test.input)
		if test.err {
			if err == nil {
				t.Errorf("parse(%q) expected error but got none", test.input)
			}
		} else {
			if err != nil {
				t.Errorf("parse(%q) returned error: %v", test.input, err)
			}
			if len(result) != len(test.expected) {
				t.Errorf("parse(%q) = %v; want %v", test.input, result, test.expected)
				continue
			}
			for k, v := range result {
				if len(v) != len(test.expected[k]) {
					t.Errorf("parse(%q) = %v; want %v", test.input, result, test.expected)
					break
				}
				for i, num := range v {
					if num != test.expected[k][i] {
						t.Errorf("parse(%q) = %v; want %v", test.input, result, test.expected)
						break
					}
				}
			}
		}
	}
}
