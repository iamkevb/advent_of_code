package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	// Part1()
	Part2()
}

func readInput(path string) [][]rune {
	f, _ := os.Open(path)
	defer f.Close()

	input := [][]rune{}
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		t := scanner.Text()
		input = append(input, []rune(t))
	}
	return input
}

func findStartingPosition(grid [][]rune) (int, int) {
	for r, row := range grid {
		for c, v := range row {
			if v == '^' {
				return r, c
			}
		}
	}
	panic("Starting point not found")
}

func step(r, c int, dir string, grid [][]rune) (int, int, string) {
	switch dir {
	case "N":
		if r < 0 {
			return -1, -1, "N"
		}
		if grid[r-1][c] == '#' {
			return step(r, c, "E", grid)
		}
		return r - 1, c, "N"
	case "E":
		if c == len(grid)-1 {
			return -1, -1, "E"
		}
		if grid[r][c+1] == '#' {
			return step(r, c, "S", grid)
		}
		return r, c + 1, "E"
	case "S":
		if r == len(grid[0])-1 {
			return -1, -1, "S"
		}
		if grid[r+1][c] == '#' {
			return step(r, c, "W", grid)
		}
		return r + 1, c, "S"
	case "W":
		if c < 0 {
			return -1, -1, "W"
		}
		if grid[r][c-1] == '#' {
			return step(r, c, "N", grid)
		}
		return r, c - 1, "W"
	}
	panic(fmt.Sprintf("Direction invalid %s", dir))
}

// PART1
func Part1() {
	// grid := readInput("./input/test.txt")
	grid := readInput("./input/input.txt")

	r, c := findStartingPosition(grid)
	dir := "N"

	visited := map[string]struct{}{}

	for r > 0 && r < len(grid[0]) && c > 0 && c < len(grid) {
		visited[fmt.Sprintf("%d,%d", r, c)] = struct{}{}
		r, c, dir = step(r, c, dir, grid)
	}
	fmt.Println("VISITED", len(visited))

}

// Part 2
func Part2() {
	// grid := readInput("./input/test.txt")
	grid := readInput("./input/input.txt")

	startingRow, startingCol := findStartingPosition(grid)
	startingDir := "N"

	visited := map[Position]struct{}{}
	r := startingRow
	c := startingCol
	dir := startingDir

	//Find path, obstacle in non visited positions won't affect the output.
	for r > 0 && r < len(grid[0]) && c > 0 && c < len(grid) {
		visited[Position{R: r, C: c, Dir: dir}] = struct{}{}
		r, c, dir = step(r, c, dir, grid)
	}

	obstacles := map[Obstacle]struct{}{}
	for p := range visited {
		// assume p is an obstacle (unless it's the starting position)
		if p.R == startingRow && p.C == startingCol {
			continue
		}

		// assume position is an obstacle and check for loop.
		testGrid := DeepCopy(grid)
		testGrid[p.R][p.C] = '#'

		r = startingRow
		c = startingCol
		dir = startingDir

		testVisits := map[Position]struct{}{}
		for r > 0 && r < len(testGrid[0]) && c > 0 && c < len(testGrid) {
			testVisits[Position{R: r, C: c, Dir: dir}] = struct{}{}
			r, c, dir = step(r, c, dir, testGrid)
			v := Position{R: r, C: c, Dir: dir}
			if _, exists := testVisits[v]; exists {
				obstacles[Obstacle{p.R, p.C}] = struct{}{}
				r = -1
			}
		}
	}
	fmt.Println(len(obstacles))
}

type Position struct {
	R, C int
	Dir  string
}
type Obstacle struct {
	R, C int
}

func DeepCopy(input [][]rune) [][]rune {
	// Create a new 2D slice with the same outer length
	dup := make([][]rune, len(input))

	for i, row := range input {
		// Create a new slice for each inner slice
		dup[i] = make([]rune, len(row))
		// Copy the contents of the inner slice
		copy(dup[i], row)
	}

	return dup
}
