package part1

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
)

func Part1() {
	input := readInput("./input/test1.txt")
	r := regexp.MustCompile(`(\d)`)
	total := int64(0)
	for _, s := range input {
		matches := r.FindAll([]byte(s), -1)
		if len(matches) == 0 {
			break
		}
		c1 := matches[0]
		c2 := matches[len(matches)-1]
		val, _ := strconv.ParseInt(string(c1)+string(c2), 10, 32)
		total += val
	}
	fmt.Println(total)
}

func readInput(path string) []string {
	f, _ := os.Open(path)
	scanner := bufio.NewScanner(f)
	lines := []string{}
	for scanner.Scan() {
		line := scanner.Text()
		lines = append(lines, line)
	}
	return lines
}
