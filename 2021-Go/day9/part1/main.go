package main

import (
	"bufio"
	"fmt"
	"os"
)

type Point struct {
	x, y int
}

func main() {
	cave := readFile("./input.txt")
	lows := lowPoints(cave)
	total := 0
	for _, l := range lows {
		total += 1 + cave[l.x][l.y]
	}
	fmt.Println("total: ", total)
}

func lowPoints(cave [][]int) []Point {
	lows := []Point{}
	for row := range cave {
		for col := range cave[row] {
			val := cave[row][col]
			isLow := true
			if row > 0 && val >= cave[row-1][col] {
				isLow = false
			}
			// below
			if isLow && row+1 < len(cave) && val >= cave[row+1][col] {
				isLow = false
			}
			// left
			if isLow && col > 0 && val >= cave[row][col-1] {
				isLow = false
			}
			// right
			if isLow && col+1 < len(cave[row]) && val >= cave[row][col+1] {
				isLow = false
			}
			if isLow {
				lows = append(lows, Point{x: row, y: col})
			}
		}
	}
	return lows
}

func readFile(path string) [][]int {
	f, e := os.Open(path)
	if e != nil {
		panic("error opening file")
	}
	defer f.Close()

	cave := [][]int{}

	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		terowt := scanner.Text()
		row := []int{}
		for _, v := range terowt {
			row = append(row, int(v)-'0')
		}
		cave = append(cave, row)
	}
	return cave
}
