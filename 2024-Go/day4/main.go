package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	// Part1()
	Part2()
}

func readInput(path string) []string {
	f, _ := os.Open(path)
	defer f.Close()

	input := []string{}
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		t := scanner.Text()
		input = append(input, t)
	}
	return input
}

// PART1
func Part1() {
	// input := readInput("./input/test.txt")
	input := readInput("./input/input.txt")
	total := 0

	//Count XMAS in rows
	total += countXMAS(input)
	fmt.Println(total)

	//Count XMAS in columns
	transposed := colsToRows(input)
	total += countXMAS(transposed)
	fmt.Println(total)

	//Count XMAS in diagonals
	diags := diagonalToRows(input)
	fmt.Println(diags)
	total += countXMAS(diags)

	diags = diagonalToRows2(input)
	fmt.Println(diags)
	total += countXMAS(diags)

	fmt.Println(total)
}

func countXMAS(input []string) int {
	t := 0
	for _, r := range input {
		t += strings.Count(r, "XMAS")
		t += strings.Count(r, "SAMX")
		// fmt.Println(r, "XMAS: ", strings.Count(r, "XMAS"), "SAMX: ", strings.Count(r, "SAMX"))
	}
	return t
}

func colsToRows(input []string) []string {
	rows := len(input)
	cols := len(input[0])

	transposed := make([][]rune, cols)
	for i := range transposed {
		transposed[i] = make([]rune, rows)
	}

	for i, s := range input {
		for j, c := range []rune(s) {
			transposed[j][i] = rune(c)
		}
	}

	result := []string{}
	for _, rs := range transposed {
		result = append(result, string(rs))
	}
	return []string(result)
}

func diagonalToRows(input []string) []string {
	runes := [][]rune{}
	for _, v := range input {
		runes = append(runes, []rune(v))
	}

	rows := len(runes)
	cols := len(runes[0])
	diagonals := make([]string, rows+cols-1)

	for y := 0; y < rows; y++ {
		for x := 0; x < cols; x++ {
			if len(diagonals) <= x {
				diagonals = append(diagonals, "")
			}
			diagonals[x+y] = diagonals[x+y] + string(runes[x][y])
		}
	}
	return diagonals
}

func diagonalToRows2(input []string) []string {
	runes := [][]rune{}
	for _, v := range input {
		runes = append(runes, []rune(v))
	}

	rows := len(runes)
	cols := len(runes[0])
	diagonals := make([]string, rows+cols-1)

	for y := rows - 1; y >= 0; y-- {
		for x := 0; x < cols; x++ {
			diagonals[rows-1-y+x] = diagonals[rows-1-y+x] + string(runes[x][y])
		}
	}
	return diagonals
}

// Part 2

func Part2() {
	// input := readInput("./input/test.txt")
	input := readInput("./input/input.txt")

	runes := [][]rune{}
	for _, v := range input {
		runes = append(runes, []rune(v))
	}

	total := 0

	for x := 1; x < len(runes[0])-1; x++ {
		for y := 1; y < len(runes)-1; y++ {
			if runes[x][y] == 'A' {
				str := string([]rune{
					runes[x-1][y-1],
					'A',
					runes[x+1][y+1],
				})
				if str != "MAS" && str != "SAM" {
					continue
				}
				str = string([]rune{
					runes[x-1][y+1],
					'A',
					runes[x+1][y-1],
				})
				if str == "MAS" || str == "SAM" {
					total += 1
				}
			}
		}
	}

	fmt.Println(total)
}
