package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"part2/util"
	"sort"
)

var closersToOpeners map[rune]rune = map[rune]rune{
	')': '(',
	']': '[',
	'}': '{',
	'>': '<',
}

var scores map[rune]int = map[rune]int{
	'(': 1,
	'[': 2,
	'{': 3,
	'<': 4,
}

func main() {
	path := "./input.txt"
	f, _ := os.Open(path)
	defer f.Close()

	programs := []string{}
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		p := scanner.Text()
		compiled, e := compile(p)
		if e != nil {
			continue
		}
		programs = append(programs, compiled)
	}

	pScores := SortableInt{}
	for i, p := range programs {
		score := 0
		fmt.Println(i, p)
		pRunes := []rune(p)
		for j := len(p) - 1; j >= 0; j-- {
			r := pRunes[j]
			score = score*5 + scores[r]
			fmt.Println(i, j, score)
		}
		pScores = append(pScores, score)
	}

	sort.Sort(pScores)
	fmt.Println(pScores[len(pScores)/2])
}

func compile(text string) (string, error) {
	stack := util.NewStack()
	for _, r := range text {
		if isOpener(r) {
			stack.Push(r)
		} else {
			if stack.Empty() {
				return "corrupted", errors.New("corrupted program")
			}

			top := stack.Top()
			if top == closersToOpeners[r] {
				stack.Pop()
			} else {
				return "corrupted", errors.New("corrupted program")
			}
		}
	}
	return stack.String(), nil
}

func isOpener(r rune) bool {
	return r == '(' || r == '[' || r == '<' || r == '{'
}

type SortableInt []int

func (a SortableInt) Len() int           { return len(a) }
func (a SortableInt) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a SortableInt) Less(i, j int) bool { return a[i] < a[j] }
