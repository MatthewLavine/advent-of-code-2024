package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"regexp"
	"runtime/pprof"
	"sort"
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

	antennaRegex = regexp.MustCompile(`[A-Za-z\d]`)
)

type coordinates struct {
	x, y int
}

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

	data := parseMap(input)

	printMap(data)

	count, err := findAntinodes(data)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Antinode Count: %d\n", count)
}

func readInputFile(file string) (string, error) {
	bytes, err := os.ReadFile(file)
	if err != nil {
		return "", err
	}
	return string(bytes), nil
}

func parseMap(input string) [][]string {
	rows := strings.Split(input, "\n")
	m := make([][]string, len(rows))
	for i, row := range rows {
		m[i] = strings.Split(row, "")
	}
	return m
}

func printMap(m [][]string) {
	for _, row := range m {
		fmt.Println(strings.Join(row, " "))
	}
	fmt.Println()
}

func findAntinodes(m [][]string) (int, error) {
	antennaCoordinates := make(map[string][]coordinates, 0)
	antinodeCoordinates := make([]coordinates, 0)
	for i, row := range m {
		for j, col := range row {
			if antennaRegex.MatchString(col) {
				if antennaCoordinates[col] != nil {
					continue
				}
				if *verbose {
					fmt.Println("----------")
					fmt.Printf("Antenna found. Row: %d, Col: %d, Frequency: %s\n", i, j, col)
				}
				antennaCoordinates[col] = findMatchingAntennas(m, col)
				if *verbose {
					fmt.Printf("Matching antennas: %v\n", antennaCoordinates[col])
				}
			}
		}
	}
	for _, frequency := range antennaCoordinates {
		for _, pair := range generateAntennaPairs(frequency) {
			sort.Slice(pair, func(i, j int) bool {
				return pair[i].x < pair[j].x && pair[i].y < pair[j].y
			})
			if *verbose {
				fmt.Println("----------")
				fmt.Printf("Finding antinodes for antenna pair: %v\n", pair)
			}
			antinodes := findAntinodesForAntennaPair(m, pair)
			if *verbose {
				fmt.Printf("Antinodes: %v\n", antinodes)
			}
			antinodeCoordinates = append(antinodeCoordinates, antinodes...)
		}
	}
	if *verbose {
		fmt.Println("----------")
	}
	antinodeCoordinates = dedupe(antinodeCoordinates)
	return len(antinodeCoordinates), nil
}

func findMatchingAntennas(m [][]string, frequency string) []coordinates {
	antennas := make([]coordinates, 0)
	for i, row := range m {
		for j, col := range row {
			if col == frequency {
				antennas = append(antennas, coordinates{i, j})
			}
		}
	}
	return antennas
}

func generateAntennaPairs(antennas []coordinates) [][]coordinates {
	pairs := make([][]coordinates, 0)
	for i, a := range antennas {
		for j, b := range antennas {
			if i == j {
				continue
			}
			pairs = append(pairs, []coordinates{a, b})
		}
	}
	return pairs
}

func findAntinodesForAntennaPair(m [][]string, antennas []coordinates) []coordinates {
	antinodes := make([]coordinates, 0)

	diffX, diffY := coordinateDiff(antennas[0], antennas[1])

	antinodeOne := coordinates{antennas[0].x + diffX, antennas[0].y + diffY}
	antinodeTwo := coordinates{antennas[1].x - diffX, antennas[1].y - diffY}

	if *verbose {
		fmt.Printf("Raw antinodes: %v, %v\n", antinodeOne, antinodeTwo)
	}

	if antinodeOne.x >= 0 && antinodeOne.y >= 0 && antinodeOne.x < len(m) && antinodeOne.y < len(m[0]) {
		if *verbose {
			fmt.Printf("Valid antinode: %v\n", antinodeOne)
		}
		antinodes = append(antinodes, antinodeOne)
	}
	if antinodeTwo.x >= 0 && antinodeTwo.y >= 0 && antinodeTwo.x < len(m) && antinodeTwo.y < len(m[0]) {
		if *verbose {
			fmt.Printf("Valid antinode: %v\n", antinodeTwo)
		}
		antinodes = append(antinodes, antinodeTwo)
	}

	return antinodes
}

func coordinateDiff(a, b coordinates) (int, int) {
	return a.x - b.x, a.y - b.y
}

func dedupe(coords []coordinates) []coordinates {
	seen := make(map[coordinates]struct{})
	deduped := make([]coordinates, 0)
	for _, c := range coords {
		if _, ok := seen[c]; ok {
			continue
		}
		seen[c] = struct{}{}
		deduped = append(deduped, c)
	}
	return deduped
}
