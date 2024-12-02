package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"runtime/pprof"
	"strconv"
	"strings"
)

const (
	demoInputFile = "demo.txt"
	inputFile     = "input.txt"
)

var (
	demo        = flag.Bool("demo", false, "Use demo input")
	enablePprof = flag.Bool("pprof", false, "Enable pprof")
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

	reports, err := parseInput(input)
	if err != nil {
		log.Fatal(err)
	}

	processReports(reports)
}

func readInputFile(file string) (string, error) {
	bytes, err := os.ReadFile(file)
	if err != nil {
		return "", err
	}
	return string(bytes), nil
}

func parseInput(input string) ([][]int, error) {
	var reports [][]int
	for _, line := range strings.Split(input, "\n") {
		if line == "" {
			continue
		}
		var report []int
		for _, char := range strings.Fields(line) {
			level, err := strconv.Atoi(char)
			if err != nil {
				return nil, err
			}
			report = append(report, level)
		}
		reports = append(reports, report)
	}
	return reports, nil
}

func processReports(reports [][]int) {
	rawSafeReports := 0
	dampenedSafeReports := 0
	for _, report := range reports {
		if isReportSafe(report) {
			rawSafeReports++
		}
		if isReportSafeDampened(report) {
			dampenedSafeReports++
		}
	}
	fmt.Printf("Safe reports: %d\n", rawSafeReports)
	fmt.Printf("Dampened safe reports: %d\n", dampenedSafeReports)
}

func isReportSafeDampened(report []int) bool {
	for i := 0; i < len(report); i++ {
		// Copy the report
		dampened := make([]int, len(report))
		copy(dampened, report)
		// Remove the current level
		dampened = remove(dampened, i)
		if isReportSafe(dampened) {
			return true
		}
	}
	return false
}

func isReportSafe(report []int) bool {
	prev := 0
	increasing := false
	decreasing := false
	for i, level := range report {
		if i == 0 {
			prev = level
			continue
		}
		if i == 1 {
			if level > prev {
				increasing = true
			} else if level < prev {
				decreasing = true
			}
		}
		if increasing && level < prev {
			return false
		}
		if decreasing && level > prev {
			return false
		}
		diff := diff(level, prev)
		if diff < 1 || diff > 3 {
			return false
		}
		prev = level
	}
	return true
}

func remove(slice []int, s int) []int {
	return append(slice[:s], slice[s+1:]...)
}

func diff(a, b int) int {
	if a > b {
		return a - b
	}
	return b - a
}
