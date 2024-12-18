package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"runtime/pprof"
	"strings"
	"time"
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

	m, startingCol, startingRow, err := parseMap(input)
	if err != nil {
		log.Fatal(err)
	}

	positions := traverse(m, startingCol, startingRow, 0, 0, 0)

	fmt.Printf("Unique Positions: %d\n", positions)

	blocks := part2(m, startingCol, startingRow)

	fmt.Printf("New Blocks: %d\n", blocks)
}

func readInputFile(file string) (string, error) {
	bytes, err := os.ReadFile(file)
	if err != nil {
		return "", err
	}
	return string(bytes), nil
}

func parseMap(input string) ([][]string, int, int, error) {
	rows := strings.Split(input, "\n")
	m := make([][]string, len(rows))
	startingCol := 0
	startingRow := 0
	for i, row := range rows {
		cols := strings.Split(row, "")
		m[i] = cols
		for j, col := range cols {
			if col == "^" {
				startingCol = j
				startingRow = i
			}
		}
	}
	return m, startingCol, startingRow, nil
}

func printMap(m [][]string) {
	for _, row := range m {
		fmt.Println(strings.Join(row, " "))
	}
	fmt.Println()
}

func clear() {
	fmt.Print("\033[H\033[2J")
}

func traverse(m [][]string, startingCol, startingRow, traversals, newCol, newRol int) int {
	m = deepCopyMap(m)
	uniquePositions := 0
	sequentialRevisits := 0

	currCol := startingCol
	currRow := startingRow
	direction := "up"

	if *verbose {
		printMap(m)
	}

	for {
		nextRow, nextCol := nextPos(direction, currRow, currCol)
		if *verbose {
			clear()
			fmt.Printf("Traversals: %d\n", traversals)
			fmt.Printf("New Blocker: %d, %d\n", newCol, newRol)
			fmt.Printf("Curr: %d, %d, Next: %d, %d, Uniq: %d, Sequential: %d\n", currRow, currCol, nextRow, nextCol, uniquePositions, sequentialRevisits)
			printMap(m)
		}
		if nextRow < 0 || nextRow >= len(m) || nextCol < 0 || nextCol >= len(m[0]) {
			uniquePositions++
			break
		}
		if m[nextRow][nextCol] == "#" {
			direction = turn(direction)
			continue
		}
		if m[nextRow][nextCol] == "X" {
			sequentialRevisits++
			if sequentialRevisits == 1000 {
				// Loop detected.
				return -1
			}
		} else {
			sequentialRevisits = 0
			uniquePositions++
		}
		m[currRow][currCol] = "X"
		currCol = nextCol
		currRow = nextRow
		m[nextRow][nextCol] = charForDirection(direction)
		if *verbose {
			time.Sleep(50 * time.Millisecond)
		}
	}

	return uniquePositions
}

func nextPos(direction string, currRow, currCol int) (int, int) {
	switch direction {
	case "up":
		return currRow - 1, currCol
	case "down":
		return currRow + 1, currCol
	case "left":
		return currRow, currCol - 1
	case "right":
		return currRow, currCol + 1
	}
	return 0, 0
}

func turn(direction string) string {
	switch direction {
	case "up":
		return "right"
	case "right":
		return "down"
	case "down":
		return "left"
	case "left":
		return "up"
	}
	return ""
}

func charForDirection(direction string) string {
	switch direction {
	case "up":
		return "^"
	case "down":
		return "v"
	case "left":
		return "<"
	case "right":
		return ">"
	}
	return ""
}

func part2(m [][]string, startingCol, startingRow int) int {
	m = deepCopyMap(m)
	blocks := 0
	traversals := 0

	for i, row := range m {
		for j, col := range row {
			set := false
			if col != "." {
				continue
			}
			if col == "." {
				set = true
				m[i][j] = "#"
			}
			tmpM := deepCopyMap(m)
			result := traverse(tmpM, startingCol, startingRow, traversals, i, j)
			traversals++
			if result == -1 {
				blocks++
			}
			if set {
				m[i][j] = "."
			}
		}
	}

	if *verbose {
		fmt.Printf("Traversals: %d\n", traversals)
	}

	return blocks
}

func deepCopyMap(m [][]string) [][]string {
	newMap := make([][]string, len(m))
	for i, row := range m {
		newMap[i] = make([]string, len(row))
		copy(newMap[i], row)
	}
	return newMap
}
