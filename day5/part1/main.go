package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

type Point struct {
	x, y int
}

type Line struct {
	start, end Point
}

func (l Line) Equals(line Line) bool {
	return l.start == line.start && l.end == line.end
}

func main() {
	lines := readInputFile("./input.txt")
	grid := mapLines(lines)

	count := 0
	for _, v := range grid {
		if v >= 2 {
			count += 1
		}
	}
	fmt.Println(count)
}

func mapLines(lines []Line) map[Point]int {
	grid := make(map[Point]int, len(lines))
	for _, l := range lines {
		var points []Point
		if l.start.x == l.end.x {
			points = verticalPoints(l)
		} else if l.start.y == l.end.y {
			points = horizontalPoints(l)
		}
		for _, p := range points {
			grid[p] += 1
		}
	}
	return grid
}

func verticalPoints(l Line) []Point {
	start, end := l.start, l.end
	if end.y < start.y {
		start, end = end, start
	}
	var points []Point
	for y := start.y; y <= end.y; y++ {
		points = append(points, Point{x: start.x, y: y})
	}
	return points
}

func horizontalPoints(l Line) []Point {
	start, end := l.start, l.end
	if end.x < start.x {
		start, end = end, start
	}
	var points []Point
	for x := start.x; x <= end.x; x++ {
		points = append(points, Point{x: x, y: start.y})
	}
	return points
}

func readInputFile(path string) []Line {
	f, e := os.Open(path)
	if e != nil {
		panic("unable to open file")
	}
	defer f.Close()
	return readInput(f)
}

func readInput(r io.Reader) []Line {
	scanner := bufio.NewScanner(r)
	lines := []Line{}
	for scanner.Scan() {
		text := scanner.Text()
		var line Line
		fmt.Sscanf(text, "%d,%d -> %d,%d", &line.start.x, &line.start.y, &line.end.x, &line.end.y)
		lines = append(lines, line)
	}
	return lines
}
