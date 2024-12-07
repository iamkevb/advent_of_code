package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	// Part1()
	Part2()
}

type Equation struct {
	value   int64
	numbers []int64
}

func readInput(path string) []Equation {
	f, _ := os.Open(path)
	defer f.Close()

	input := []Equation{}
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		t := scanner.Text()
		split := strings.Split(t, ": ")

		v, _ := strconv.ParseInt(split[0], 10, 64)
		n := []int64{}
		for _, s := range strings.Split(split[1], " ") {
			i, _ := strconv.ParseInt(s, 10, 64)
			n = append(n, i)
		}
		input = append(input, Equation{v, n})
	}
	return input
}

func evaluate(eq Equation, ops []string) int64 {
	for _, o := range ops {
		result := eq.numbers[0]
		for i, oc := range []rune(o) {
			switch oc {
			case '*':
				result *= eq.numbers[i+1]
			case '+':
				result += eq.numbers[i+1]
			case '|':
				s := fmt.Sprintf("%d%d", result, eq.numbers[i+1])
				result, _ = strconv.ParseInt(s, 10, 64)
			default:
				panic("bad operator")
			}
		}
		if result == eq.value {
			return result
		}
	}
	return 0
}

// PART1
func Part1() {
	// equations := readInput("./input/test.txt")
	equations := readInput("./input/input.txt")

	var total int64 = 0
	for _, equation := range equations {
		ops := generateOperators1(len(equation.numbers) - 1)
		total += evaluate(equation, ops)
	}
	fmt.Println(total)
}

func generateOperators1(slots int) []string {
	ops := []string{"*", "+"}
	for i := 1; i < slots; i++ {
		next := []string{}
		for _, v := range ops {
			next = append(next,
				v+"*",
				v+"+",
			)
		}
		ops = next
	}
	return ops
}

// Part 2
func Part2() {
	// equations := readInput("./input/test.txt")
	equations := readInput("./input/input.txt")

	var total int64 = 0
	for _, equation := range equations {
		ops := generateOperators2(len(equation.numbers) - 1)
		total += evaluate(equation, ops)
	}
	fmt.Println(total)
}

func generateOperators2(slots int) []string {
	ops := []string{"*", "+", "|"}
	for i := 1; i < slots; i++ {
		next := []string{}
		for _, v := range ops {
			next = append(next,
				v+"*",
				v+"+",
				v+"|",
			)
		}
		ops = next
	}
	return ops
}
