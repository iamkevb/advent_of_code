package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

type Point struct {
	r, c int
}

func main() {
	cave := readFile("./input.txt")
	lows := lowPoints(cave)
	basins := basinSizes(cave, lows)
	sort.Sort(Sizes(basins))
	product := 1
	for i := 0; i < 3; i++ {
		product *= basins[len(basins)-1-i]
	}
	fmt.Println(product)
}

func basinSizes(cave [][]int, lowPoints []Point) []int {
	sizes := []int{}
	for _, lp := range lowPoints {
		basin := make(map[Point]bool, 1)
		basin[lp] = true
		mapBasin(cave, lp, &basin)
		sizes = append(sizes, len(basin))
	}
	return sizes
}

func mapBasin(cave [][]int, low Point, basin *map[Point]bool) {
	row := low.r
	col := low.c
	val := cave[row][col]

	addPoint := func(row, col int) {
		if cave[row][col] == 9 {
			return
		}
		point := Point{r: row, c: col}
		if !(*basin)[point] {
			(*basin)[point] = true
			mapBasin(cave, point, basin)
		}
	}

	//north
	if row > 0 && val < cave[row-1][col] {
		addPoint(row-1, col)
	}

	//south
	if row+1 < len(cave) && val < cave[row+1][col] {
		addPoint(row+1, col)
	}

	// west
	if col > 0 && val < cave[row][col-1] {
		addPoint(row, col-1)
	}
	// east
	if col+1 < len(cave[row]) && val < cave[row][col+1] {
		addPoint(row, col+1)
	}
}

func lowPoints(cave [][]int) []Point {
	lows := []Point{}
	for row := range cave {
		for col := range cave[row] {
			val := cave[row][col]
			isLow := true
			// north
			if row > 0 && val >= cave[row-1][col] {
				isLow = false
			}
			// south
			if isLow && row+1 < len(cave) && val >= cave[row+1][col] {
				isLow = false
			}
			// west
			if isLow && col > 0 && val >= cave[row][col-1] {
				isLow = false
			}
			// east
			if isLow && col+1 < len(cave[row]) && val >= cave[row][col+1] {
				isLow = false
			}
			if isLow {
				lows = append(lows, Point{r: row, c: col})
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

type Sizes []int

func (a Sizes) Len() int           { return len(a) }
func (a Sizes) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a Sizes) Less(i, j int) bool { return a[i] < a[j] }
