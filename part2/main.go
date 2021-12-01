package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

func main() {
	f, e := os.Open("./input.txt")
	if e != nil {
		log.Fatalf("Error reading file: %s", e)
	}
	scanner := bufio.NewScanner(f)

	depths := []int{}
	for scanner.Scan() {
		depth, e := strconv.Atoi(scanner.Text())
		if e != nil {
			log.Fatalf("Read invalid value: %s", e)
		}
		depths = append(depths, depth)
	}

	increases := 0
	for i := 0; i < len(depths)-3; i++ {
		w1 := depths[i] + depths[i+1] + depths[i+2]
		w2 := depths[i+1] + depths[i+2] + depths[i+3]
		if w2 > w1 {
			increases += 1
		}
	}
	fmt.Printf("Number of increases: %d\n", increases)
}
