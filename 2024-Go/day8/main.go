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

type Point struct {
	x, y int
}

func (p Point) Equal(p2 Point) bool {
	return p.x == p2.x && p.y == p2.y
}

func (p Point) IsValid() bool {
	return p.x >= 0 && p.x < MAX_X && p.y >= 0 && p.y < MAX_Y
}

var MAX_X = 0
var MAX_Y = 0

func readInput(path string) map[string][]Point {
	f, _ := os.Open(path)
	defer f.Close()

	input := make(map[string][]Point)
	scanner := bufio.NewScanner(f)
	y := 0
	for scanner.Scan() {
		t := scanner.Text()
		MAX_X = len(t)
		for x, c := range t {
			loc := string(c)
			if loc != "." {
				arr, exists := input[loc]
				antenna := Point{x, y}
				if !exists {
					input[loc] = []Point{antenna}
				} else {
					input[loc] = append(arr, antenna)
				}
			}
		}
		y++
	}
	MAX_Y = y
	return input
}

// PART1
func Part1() {
	// antennas := readInput("./input/test.txt")
	antennas := readInput("./input/input.txt")

	antinodes := map[Point]struct{}{}
	for _, freq := range antennas {
		if len(freq) < 2 {
			continue
		}

		nodes := findAntinodes(freq)
		for p := range nodes {
			antinodes[p] = struct{}{}
		}
	}
	fmt.Println(len(antinodes))
}

func findAntinodes(antennas []Point) map[Point]struct{} {
	antinodes := make(map[Point]struct{})

	for _, p1 := range antennas {
		for _, p2 := range antennas {
			if p1.Equal(p2) {
				continue
			}
			dx := p2.x - p1.x
			dy := p2.y - p1.y
			nx := p1.x - dx
			ny := p1.y - dy
			np := Point{nx, ny}
			if np.IsValid() {
				antinodes[np] = struct{}{}
			}
		}
	}

	return antinodes
}

// Part 2
func Part2() {
	// antennas := readInput("./input/test.txt")
	antennas := readInput("./input/input.txt")

	antinodes := map[Point]struct{}{}
	for _, freq := range antennas {
		if len(freq) < 2 {
			continue
		}

		nodes := findAntinodesWithResonance(freq)
		for p := range nodes {
			antinodes[p] = struct{}{}
		}
	}
	fmt.Println(len(antinodes))
}

func findAntinodesWithResonance(antennas []Point) map[Point]struct{} {
	antinodes := make(map[Point]struct{})

	for _, p1 := range antennas {
		for _, p2 := range antennas {
			if p1.Equal(p2) {
				antinodes[p2] = struct{}{}
			} else {
				tp1 := p1
				tp2 := p2
				for tp1.IsValid() && tp2.IsValid() {
					dx := tp2.x - tp1.x
					dy := tp2.y - tp1.y
					nx := tp1.x - dx
					ny := tp1.y - dy
					np := Point{nx, ny}
					if np.IsValid() {
						antinodes[np] = struct{}{}
					}
					tp2 = tp1
					tp1 = np
				}
			}

		}
	}

	return antinodes
}
