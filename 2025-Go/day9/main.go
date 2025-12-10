package main

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Tile struct {
	c, r int
}

const TEST = "./input-test.txt"
const DATA = "./input.txt"

func main() {
	input := TEST
	boxes := readInput(input)
	part1(boxes)
	part2(boxes)
}

func Abs(i int) int {
	if i < 0 {
		return -i
	}
	return i
}

func maxRectangleArea(tiles []Tile) int {
	maxArea := 0

	for i := 0; i < len(tiles); i++ {
		for j := i + 1; j < len(tiles); j++ {
			p1 := tiles[i]
			p2 := tiles[j]

			width := Abs(p1.c-p2.c) + 1
			height := Abs(p1.r-p2.r) + 1
			area := width * height

			if area > maxArea {
				maxArea = area
			}
		}
	}

	return maxArea
}

func part1(rt []Tile) {
	fmt.Println("Part 1:", maxRectangleArea(rt))
}
func part2(rt []Tile) {
	fmt.Println("Part 2:", 0)
}

func readInput(path string) []Tile {
	f, _ := os.ReadFile(path)
	rt := []Tile{}
	scanner := bufio.NewScanner(bytes.NewReader(f))
	for scanner.Scan() {
		line := scanner.Text()
		coord := strings.Split(line, ",")
		c, _ := strconv.Atoi(coord[0])
		r, _ := strconv.Atoi(coord[1])
		rt = append(rt, Tile{c, r})
	}
	return rt
}
