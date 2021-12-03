package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

var horiz, depth, aim int

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
		depth += distance * aim
	case "up":
		aim -= distance
		if aim < 0 {
			aim = 0
		}
	case "down":
		aim += distance
	default:
		panic(fmt.Sprintf("unexpected direction value %s", direction))
	}
}
