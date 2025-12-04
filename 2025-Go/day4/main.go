package main

import (
	"bufio"
	"fmt"
	"os"
)

const TEST = "./input-test.txt"
const DATA = "./input.txt"

func main() {
	input := readInput(DATA)
	part1(input)
	part2(input)
}

func part1(input [][]rune) {
	sum := int64(0)
	for r, row := range input {
		for c := range row {
			if isAccessible(input, r, c) {
				sum++
			}
		}
	}
	fmt.Println("Part 1: ", sum)
}

func part2(input [][]rune) {
	removed := removeAccessible(input, 0)
	fmt.Println("Part 2: ", removed)
}

func removeAccessible(grid [][]rune, removed int) int {
	for r, row := range grid {
		for c := range row {
			if isAccessible(grid, r, c) {
				grid[r][c] = 'x'
				removed++
				removed = removeAccessible(grid, removed)
			}
		}
	}
	return removed
}

func isAccessible(grid [][]rune, row int, col int) bool {
	if grid[row][col] != '@' {
		return false
	}
	adjacent := 0
	adjacent += isRoll(grid, row-1, col-1)
	adjacent += isRoll(grid, row-1, col)
	adjacent += isRoll(grid, row-1, col+1)
	adjacent += isRoll(grid, row, col-1)
	adjacent += isRoll(grid, row, col+1)
	adjacent += isRoll(grid, row+1, col-1)
	adjacent += isRoll(grid, row+1, col)
	adjacent += isRoll(grid, row+1, col+1)

	return adjacent < 4
}

func isRoll(grid [][]rune, row int, col int) int {
	if row < 0 || col < 0 {
		return 0
	}
	if row >= len(grid) || col >= len(grid[row]) {
		return 0
	}
	if grid[row][col] != '@' {
		return 0
	}
	return 1
}

func readInput(path string) [][]rune {
	f, _ := os.Open(path)
	scanner := bufio.NewScanner(f)
	input := [][]rune{}
	for scanner.Scan() {
		row := []rune(scanner.Text())
		input = append(input, row)
	}
	return input
}
