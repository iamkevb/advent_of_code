package main

import (
	"bufio"
	"fmt"
	"os"
)

var octopi [][]int

func main() {
	path := "./input.txt"
	octopi = readInput(path)
	steps := 100
	flashes := 0
	for i := 0; i < steps; i++ {
		flashes += step()
	}
	fmt.Println(flashes)
}

func step() int {
	flashes := 0
	for row, columns := range octopi {
		for col, _ := range columns {
			flashes += incrementOctopi(row, col)
		}
	}
	for row, columns := range octopi {
		for col, _ := range columns {
			if octopi[row][col] >= 10 {
				octopi[row][col] = 0
			}
		}
	}
	return flashes
}

func incrementOctopi(row, col int) int {
	flashes := 0
	octopi[row][col]++
	if octopi[row][col] == 10 {
		flashes++
		flashes += energizeNeighbours(row, col)
	}
	return flashes
}

func energizeNeighbours(row, col int) int {
	flashes := 0

	if validIndex(row-1, col-1) {
		flashes += incrementOctopi(row-1, col-1)
	}
	if validIndex(row-1, col) {
		flashes += incrementOctopi(row-1, col)
	}
	if validIndex(row-1, col+1) {
		flashes += incrementOctopi(row-1, col+1)
	}
	if validIndex(row, col-1) {
		flashes += incrementOctopi(row, col-1)
	}
	if validIndex(row, col+1) {
		flashes += incrementOctopi(row, col+1)
	}
	if validIndex(row+1, col-1) {
		flashes += incrementOctopi(row+1, col-1)
	}
	if validIndex(row+1, col) {
		flashes += incrementOctopi(row+1, col)
	}
	if validIndex(row+1, col+1) {
		flashes += incrementOctopi(row+1, col+1)
	}
	return flashes
}

func validIndex(row, col int) bool {
	return row >= 0 && col >= 0 && row < len(octopi) && col < len(octopi[row])
}

func readInput(path string) [][]int {
	f, _ := os.Open(path)
	octopi := [][]int{}

	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		line := scanner.Text()
		row := make([]int, len(line))
		for i, v := range line {
			iv := v - '0'
			row[i] = int(iv)
		}
		octopi = append(octopi, row)
	}
	return octopi
}
