package part2

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
)

func Part2() {
	input := readInput("./input/input.txt")
	r := regexp.MustCompile(`(\d)`)
	total := int64(0)
	for _, s := range input {
		fmt.Println(s)
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
		lines = append(lines, fixNumbers(line))
	}
	return lines
}

var replacements = map[string]int{
	"zero":  0,
	"one":   1,
	"two":   2,
	"three": 3,
	"four":  4,
	"five":  5,
	"six":   6,
	"seven": 7,
	"eight": 8,
	"nine":  9,
}

var re = regexp.MustCompile(`one|two|three|four|five|six|seven|eight|nine`)

func fixNumbers(line string) string {
	lineEdit := ""
	window := ""
	for _, c := range line {
		if n, e := strconv.Atoi(string(c)); e == nil {
			if len(window) > 0 {
				lineEdit += window[:len(window)-1]
			}
			lineEdit += strconv.Itoa(n)
			window = ""
		} else {
			window = window + string(c)
			n := containedDigit(window)
			if n != -1 {
				lineEdit += window[:len(window)-1] + strconv.Itoa(n)
				window = window[len(window)-1:]
			}
		}
	}
	lineEdit += window
	return lineEdit
}

func containedDigit(s string) int {
	b := re.Find([]byte(s))
	if b != nil {
		return replacements[string(b)]
	}
	return -1
}
