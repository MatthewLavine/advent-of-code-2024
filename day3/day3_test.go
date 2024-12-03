package main

import (
	"regexp"
	"testing"
)

func TestParseInput(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		regex    *regexp.Regexp
		expected []Instruction
		wantErr  bool
	}{
		{
			name:  "Valid mul instructions",
			input: "mul(2,3)\nmul[7,2]___mul(4,5)\n!!mul(6,7)",
			regex: partOneInstructionRegex,
			expected: []Instruction{
				{"mul", 2, 3},
				{"mul", 4, 5},
				{"mul", 6, 7},
			},
			wantErr: false,
		},
		{
			name:  "Valid add instructions",
			input: "add(2,3)\nadd[7,2]___add(4,5)\n!!add(6,7)",
			regex: partOneInstructionRegex,
			expected: []Instruction{
				{"add", 2, 3},
				{"add", 4, 5},
				{"add", 6, 7},
			},
			wantErr: false,
		},
		{
			name:  "Valid do and don't instructions",
			input: "do()\nmul(2,3)\ndon't()\nmul(4,5)\ndo()\nmul(6,7)",
			regex: partTwoInstructionRegex,
			expected: []Instruction{
				{"mul", 2, 3},
				{"mul", 6, 7},
			},
			wantErr: false,
		},
		{
			name:     "No instructions",
			input:    "",
			regex:    partOneInstructionRegex,
			expected: nil,
			wantErr:  true,
		},
		{
			name:     "Invalid instructions",
			input:    "invalid(1,2)",
			regex:    partOneInstructionRegex,
			expected: nil,
			wantErr:  true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := parseInput(tt.input, tt.regex)
			if (err != nil) != tt.wantErr {
				t.Errorf("parseInput() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !equalInstructions(got, tt.expected) {
				t.Errorf("parseInput() = %v, expected %v", got, tt.expected)
			}
		})
	}
}

func TestCompute(t *testing.T) {
	tests := []struct {
		name         string
		instructions []Instruction
		expected     int
		wantErr      bool
	}{
		{
			name: "Valid mul instructions",
			instructions: []Instruction{
				{"mul", 2, 3},
				{"mul", 4, 5},
				{"mul", 6, 7},
			},
			expected: 68,
			wantErr:  false,
		},
		{
			name: "Valid add instructions",
			instructions: []Instruction{
				{"add", 2, 3},
				{"add", 4, 5},
				{"add", 6, 7},
			},
			expected: 27,
			wantErr:  false,
		},
		{
			name:         "Empty instructions",
			instructions: []Instruction{},
			expected:     0,
			wantErr:      false,
		},
		{
			name: "Single mul instruction",
			instructions: []Instruction{
				{"mul", 3, 3},
			},
			expected: 9,
			wantErr:  false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := compute(tt.instructions)
			if (err != nil) != tt.wantErr {
				t.Errorf("compute() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.expected {
				t.Errorf("compute() = %v, expected %v", got, tt.expected)
			}
		})
	}
}

func equalInstructions(a, b []Instruction) bool {
	if len(a) != len(b) {
		return false
	}
	for i := range a {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}
