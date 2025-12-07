package main

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
)

const TEST = "./input-test.txt"
const DATA = "./input.txt"

func main() {
	grid := readInput(DATA)
	part1(grid)
	part2(grid)
}

func findStart(row []rune) int {
	for i, v := range row {
		if v == 'S' {
			return i
		}
	}
	panic("No start found")
}

func countSplits(grid [][]rune, beams []bool, count int) int {
	if len(grid) == 0 {
		return count
	}

	newBeams := beams
	for i, v := range grid[0] {
		if v == '^' && beams[i] {
			newBeams[i-1] = true
			newBeams[i] = false
			newBeams[i+1] = true
			count++
		}
	}

	return countSplits(grid[1:], newBeams, count)
}

func countTimelines(grid [][]rune, timelines []int64) []int64 {
	if len(grid) == 0 {
		return timelines
	}

	newTimelines := make([]int64, len(timelines))
	for i := range timelines {
		v := timelines[i]

		if grid[0][i] == '^' && v > 0 {
			newTimelines[i] = 0
			newTimelines[i-1] += v
			newTimelines[i+1] += v
		} else {
			newTimelines[i] += v
		}
	}
	return countTimelines(grid[1:], newTimelines)
}

func part1(grid [][]rune) {
	start := findStart(grid[0])
	beams := make([]bool, len(grid[0]))
	beams[start] = true
	splits := countSplits(grid[1:], beams, 0)
	fmt.Println("Part 1:", splits)
}

func part2(grid [][]rune) {
	start := findStart(grid[0])
	timelines := make([]int64, len(grid[0]))
	timelines[start] = 1
	timelines = countTimelines(grid, timelines)
	sum := int64(0)
	for _, v := range timelines {
		sum += v
	}

	fmt.Println("Part 2:", sum)
}

func readInput(path string) [][]rune {
	f, _ := os.ReadFile(path)
	scanner := bufio.NewScanner(bytes.NewReader(f))
	lines := [][]rune{}
	for scanner.Scan() {
		line := scanner.Text()
		lines = append(lines, []rune(line))
	}
	return lines
}
