package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

var horiz int
var depth int

func main() {
	f, e := os.Open("./input.txt")
	if e != nil {
		log.Fatalf("Error reading file: %s", e)
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		line := scanner.Text()
		instrs := strings.Split(line, " ")
		if len(instrs) != 2 {
			panic("Bad data. Expected every line to include direction and distance only")
		}
		direction := instrs[0]
		distance, err := strconv.Atoi(instrs[1])
		if err != nil {
			panic("Expected distance to be an integer value")
		}
		move(direction, distance)
	}

	fmt.Printf("Horizontal: %d, Depth: %d, Product: %d\n", horiz, depth, horiz*depth)
}

func move(direction string, distance int) {
	switch direction {
	case "forward":
		horiz += distance
	case "up":
		depth -= distance
		if depth < 0 {
			depth = 0
		}
	case "down":
		depth += distance
	default:
		panic(fmt.Sprintf("unexpected direction value %s", direction))
	}
}
