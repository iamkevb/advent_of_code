package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	Part1()
	Part2()
}

func readInput(path string) [][]int {
	f, _ := os.Open(path)
	input := [][]int{}
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		t := scanner.Text()
		row := []int{}
		for _, r := range t {
			row = append(row, int(r-'0'))
		}
		input = append(input, row)
	}
	return input
}

type Position struct {
	col, row int
	value    int
	next     []string
}

func mapKey(col, row int) string {
	return fmt.Sprintf("%d-%d", col, row)
}

func createGridMap(input [][]int) map[string]Position {
	grid := map[string]Position{}
	for r, row := range input {
		for c, value := range row {
			next := []string{}
			//up
			if r > 0 && input[r-1][c] == value+1 {
				next = append(next, mapKey(c, r-1))
			}
			//down
			if r < len(input)-1 && input[r+1][c] == value+1 {
				next = append(next, mapKey(c, r+1))
			}
			//left
			if c > 0 && input[r][c-1] == value+1 {
				next = append(next, mapKey(c-1, r))
			}
			//right
			if c < len(row)-1 && input[r][c+1] == value+1 {
				next = append(next, mapKey(c+1, r))
			}
			grid[mapKey(c, r)] = Position{c, r, value, next}
		}
	}
	return grid
}

func findStartingPositions(grid map[string]Position) []Position {
	starts := []Position{}
	for _, p := range grid {
		if p.value == 0 {
			starts = append(starts, p)
		}
	}
	return starts
}

// PART1
func Part1() {
	// input := readInput("./input/test.txt")
	input := readInput("./input/input.txt")
	grid := createGridMap(input)
	startPositions := findStartingPositions(grid)

	total := 0
	for _, start := range startPositions {
		peaks := countPeaks(start, grid, make(map[string]struct{}))
		total += len(peaks)
	}
	fmt.Println(total)
}

func countPeaks(start Position, grid map[string]Position, peaks map[string]struct{}) map[string]struct{} {
	if start.value == 9 {
		peaks[mapKey(start.col, start.row)] = struct{}{}
		return peaks
	}

	for _, next := range start.next {
		nextP := grid[next]
		peaks = countPeaks(nextP, grid, peaks)
	}
	return peaks
}

// PART2
func Part2() {
	// input := readInput("./input/test.txt")
	input := readInput("./input/input.txt")
	grid := createGridMap(input)
	startPositions := findStartingPositions(grid)

	total := 0
	for _, start := range startPositions {
		total += findTrails(start, grid, 0)
	}
	fmt.Println(total)
}

func findTrails(start Position, grid map[string]Position, score int) int {
	if start.value == 9 {
		return score + 1
	}

	for _, next := range start.next {
		nextP := grid[next]
		score = findTrails(nextP, grid, score)
	}
	return score
}
