package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	// Part1()
	Part2()
}

func readInput(path string) [][]int {
	f, _ := os.Open(path)
	scanner := bufio.NewScanner(f)
	reports := [][]int{}
	for scanner.Scan() {
		line := scanner.Text()
		levelStr := strings.Fields(line)
		levels := []int{}
		for _, str := range levelStr {
			l, _ := strconv.Atoi(str)
			levels = append(levels, l)
		}
		reports = append(reports, levels)
	}
	return reports
}

// PART1
func Part1() {
	reports := readInput("./input/test.txt")
	// reports := readInput("./input/input.txt")
	validReports := 0
	for _, report := range reports {
		if isValid(report) {
			validReports += 1
		}
	}
	fmt.Println("valid reports: ", validReports)
}

func isValid(report []int) bool {
	// input has no single value reports. Assume this will never happen
	direction := 1
	if report[0] > report[1] {
		direction = -1
	}
	for i := 1; i < len(report); i++ {
		diff := report[i] - report[i-1]
		if direction == -1 {
			diff = -diff
		}
		if diff < 1 || diff > 3 {
			return false
		}
	}
	return true
}

// Part 2
func Part2() {
	// reports := readInput("./input/test.txt")
	reports := readInput("./input/input.txt")
	validReports := 0
	for _, report := range reports {
		if isValidWithTolerance(report) {
			validReports += 1
		}
	}
	fmt.Println("valid reports: ", validReports)
}

func isValidWithTolerance(report []int) bool {
	// input has no single value reports. Assume this will never happen
	if isValid(report) {
		return true
	}

	for i := 0; i < len(report); i++ {

		var noI []int
		noI = append(noI, report[:i]...)
		noI = append(noI, report[i+1:]...)

		if isValid(noI) {
			return true
		}
	}

	return false
}
