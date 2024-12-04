package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"runtime/pprof"
	"strings"
)

const (
	demoInputFile = "demo.txt"
	inputFile     = "input.txt"
)

var (
	demo        = flag.Bool("demo", false, "Use demo input")
	enablePprof = flag.Bool("pprof", false, "Enable pprof")
	verbose     = flag.Bool("v", false, "Enable verbose output")
)

func main() {
	flag.Parse()

	if *enablePprof {
		f, err := os.Create("profile.prof")
		if err != nil {
			panic(err)
		}
		defer f.Close()
		if err := pprof.StartCPUProfile(f); err != nil {
			panic(err)
		}
		defer pprof.StopCPUProfile()
	}

	path := inputFile
	if *demo {
		path = demoInputFile
	}

	input, err := readInputFile(path)
	if err != nil {
		log.Fatal(err)
	}

	matrix, err := parseInput(input)
	if err != nil {
		log.Fatal(err)
	}

	xmasCount, masCount := processMatrix(matrix)

	fmt.Printf("Xmas count: %d\n", xmasCount)
	fmt.Printf("Mas count: %d\n", masCount)
}

func readInputFile(file string) (string, error) {
	bytes, err := os.ReadFile(file)
	if err != nil {
		return "", err
	}
	return string(bytes), nil
}

func parseInput(input string) ([][]string, error) {
	var matrix [][]string
	for _, line := range strings.Split(input, "\n") {
		matrix = append(matrix, strings.Split(line, ""))
	}
	return matrix, nil
}

func processMatrix(matrix [][]string) (int, int) {
	xmasCount := 0
	masCount := 0
	if *verbose {
		printMatrix(matrix)
	}
	for i, row := range matrix {
		for j := range row {
			xmasCount += startsXmas(matrix, i, j)
			masCount += startsMas(matrix, i, j)
		}
	}
	return xmasCount, masCount
}

func printMatrix(matrix [][]string) {
	for _, row := range matrix {
		fmt.Println(strings.Join(row, " "))
	}
	fmt.Println()
}

// ew
func startsXmas(matrix [][]string, row, col int) int {
	found := 0
	// Left to right
	if col+3 < len(matrix[row]) {
		if matrix[row][col] == "X" && matrix[row][col+1] == "M" && matrix[row][col+2] == "A" && matrix[row][col+3] == "S" {
			if *verbose {
				fmt.Printf("Found left to right xmas at %d, %d\n", row+1, col+1)
			}
			found++
		}
	}
	// Right to left
	if col-3 >= 0 {
		if matrix[row][col] == "X" && matrix[row][col-1] == "M" && matrix[row][col-2] == "A" && matrix[row][col-3] == "S" {
			if *verbose {
				fmt.Printf("Found right to left xmas at %d, %d\n", row+1, col+1)
			}
			found++
		}
	}
	// Top to bottom
	if row+3 < len(matrix) {
		if matrix[row][col] == "X" && matrix[row+1][col] == "M" && matrix[row+2][col] == "A" && matrix[row+3][col] == "S" {
			if *verbose {
				fmt.Printf("Found top to bottom xmas at %d, %d\n", row+1, col+1)
			}
			found++
		}
	}
	// Bottom to top
	if row-3 >= 0 {
		if matrix[row][col] == "X" && matrix[row-1][col] == "M" && matrix[row-2][col] == "A" && matrix[row-3][col] == "S" {
			if *verbose {
				fmt.Printf("Found bottom to top xmas at %d, %d\n", row+1, col+1)
			}
			found++
		}
	}
	// Diagonal top left to bottom right
	if col+3 < len(matrix[row]) && row+3 < len(matrix) {
		if matrix[row][col] == "X" && matrix[row+1][col+1] == "M" && matrix[row+2][col+2] == "A" && matrix[row+3][col+3] == "S" {
			if *verbose {
				fmt.Printf("Found diagonal top left to bottom right xmas at %d, %d\n", row+1, col+1)
			}
			found++
		}
	}
	// Diagonal bottom right to top left
	if col-3 >= 0 && row-3 >= 0 {
		if matrix[row][col] == "X" && matrix[row-1][col-1] == "M" && matrix[row-2][col-2] == "A" && matrix[row-3][col-3] == "S" {
			if *verbose {
				fmt.Printf("Found diagonal bottom right to top left xmas at %d, %d\n", row+1, col+1)
			}
			found++
		}
	}
	// Diagonal top right to bottom left
	if col-3 >= 0 && row+3 < len(matrix) {
		if matrix[row][col] == "X" && matrix[row+1][col-1] == "M" && matrix[row+2][col-2] == "A" && matrix[row+3][col-3] == "S" {
			if *verbose {
				fmt.Printf("Found diagonal top right to bottom left xmas at %d, %d\n", row+1, col+1)
			}
			found++
		}
	}
	// Diagonal bottom left to top right
	if col+3 < len(matrix[row]) && row-3 >= 0 {
		if matrix[row][col] == "X" && matrix[row-1][col+1] == "M" && matrix[row-2][col+2] == "A" && matrix[row-3][col+3] == "S" {
			if *verbose {
				fmt.Printf("Found diagonal bottom left to top right xmas at %d, %d\n", row+1, col+1)
			}
			found++
		}
	}
	return found
}

// ew
func startsMas(matrix [][]string, row, col int) int {
	found := 0
	// Top to bottom
	if row+2 < len(matrix) && col+2 < len(matrix[row]) {
		if matrix[row][col] == "M" && matrix[row+1][col+1] == "A" && matrix[row+2][col+2] == "S" && matrix[row][col+2] == "M" && matrix[row+2][col] == "S" {
			if *verbose {
				fmt.Printf("Found top to bottom x-mas at %d, %d\n", row+1, col+1)
			}
			found++
		}
	}
	// Left to right
	if row+2 < len(matrix) && col+2 < len(matrix[row]) {
		if matrix[row][col] == "M" && matrix[row+1][col+1] == "A" && matrix[row+2][col+2] == "S" && matrix[row+2][col] == "M" && matrix[row][col+2] == "S" {
			if *verbose {
				fmt.Printf("Found left to right x-mas at %d, %d\n", row+1, col+1)
			}
			found++
		}
	}
	// Bottom to top
	if row-2 >= 0 && col-2 >= 0 {
		if matrix[row][col] == "M" && matrix[row-1][col-1] == "A" && matrix[row-2][col-2] == "S" && matrix[row][col-2] == "M" && matrix[row-2][col] == "S" {
			if *verbose {
				fmt.Printf("Found bottom to top x-mas at %d, %d\n", row+1, col+1)
			}
			found++
		}
	}
	// Right to left
	if row-2 >= 0 && col-2 >= 0 {
		if matrix[row][col] == "M" && matrix[row-1][col-1] == "A" && matrix[row-2][col-2] == "S" && matrix[row-2][col] == "M" && matrix[row][col-2] == "S" {
			if *verbose {
				fmt.Printf("Found right to left x-mas at %d, %d\n", row+1, col+1)
			}
			found++
		}
	}
	return found
}
