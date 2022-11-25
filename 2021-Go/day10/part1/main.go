package main

import (
	"bufio"
	"fmt"
	"os"
	"part1/util"
)

var scores map[rune]int = map[rune]int{
	')': 3,
	']': 57,
	'}': 1197,
	'>': 25137,
}

var closersAndOpeners map[rune]rune = map[rune]rune{
	')': '(',
	']': '[',
	'}': '{',
	'>': '<',
}

func main() {
	path := "./input.txt"
	f, _ := os.Open(path)
	defer f.Close()

	score := 0
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		score += findScore(scanner.Text())
	}
	fmt.Println(score)
}

func findScore(text string) int {
	stack := util.NewStack()
	for _, r := range text {
		if isOpener(r) {
			stack.Push(r)
		} else {
			if stack.Empty() {
				return scores[r]
			}

			top := stack.Top()
			if top == closersAndOpeners[r] {
				stack.Pop()
			} else {
				return scores[r]
			}
		}
	}
	return 0
}

func isOpener(r rune) bool {
	return r == '(' || r == '[' || r == '<' || r == '{'
}
