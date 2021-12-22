package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type Cube struct {
	x, y, z int
}
type Instruction struct {
	on     bool
	x1, x2 int
	y1, y2 int
	z1, z2 int
}

var grid map[Cube]bool = make(map[Cube]bool)
var instructions []Instruction = make([]Instruction, 0)

func main() {
	path := "./test.txt"
	path = "./input.txt"
	readInput(path)

	for _, i := range instructions {
		fmt.Println("DO: ", i)
		exec(i)
	}
	fmt.Println(len(grid))
}

func exec(i Instruction) {
	for x := i.x1; x <= i.x2; x++ {
		if x > 50 || x < -50 {
			continue
		}
		for y := i.y1; y <= i.y2; y++ {
			if y > 50 || y < -50 {
				continue
			}
			for z := i.z1; z <= i.z2; z++ {
				if z > 50 || z < -50 {
					continue
				}
				c := Cube{x, y, z}

				if i.on {
					// fmt.Println("ON: ", c)
					grid[c] = i.on
				} else {
					// fmt.Println("OFF: ", c)
					delete(grid, c)
				}
			}
		}
	}
}

func readInput(path string) {
	f, _ := os.Open(path)
	defer f.Close()
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		line := scanner.Text()
		var onS string
		var x1, x2, y1, y2, z1, z2 int
		fmt.Sscanf(line, "%3s x=%d..%d,y=%d..%d,z=%d..%d", &onS, &x1, &x2, &y1, &y2, &z1, &z2)
		x1, x2 = minMax(x1, x2)
		y1, y2 = minMax(y1, y2)
		z1, z2 = minMax(z1, z2)
		i := Instruction{
			on: strings.TrimSpace(onS) == "on",
			x1: x1,
			x2: x2,
			y1: y1,
			y2: y2,
			z1: z1,
			z2: z2,
		}
		instructions = append(instructions, i)
	}
}

func minMax(i, j int) (min, max int) {
	if i < j {
		return i, j
	}
	return j, i
}
