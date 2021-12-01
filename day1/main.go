package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
	"strconv"
)

func main() {
	f, e := os.Open("./input.txt")
	if e != nil {
		log.Fatalf("Error reading file: %s", e)
	}
	scanner := bufio.NewScanner(f)

	increases := 0
	lastDepth := math.MaxInt
	for scanner.Scan() {
		depth, e := strconv.Atoi(scanner.Text())
		if e != nil {
			log.Fatalf("Read invalid value: %s", e)
		}
		if depth-lastDepth > 0 {
			increases += 1
		}
		lastDepth = depth
	}

	fmt.Printf("Number of increases: %d\n", increases)
}
