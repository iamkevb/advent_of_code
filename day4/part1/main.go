package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type Board struct {
	squares [][]int
}

func (b Board) mark(number int) bool {
	for _, xv := range b.squares {
		for y, yv := range xv {
			if yv == number {
				xv[y] = -1
			}
		}
	}
	return b.solved()
}

func (b Board) solved() bool {
	return b.rowSolved() || b.columnSolved()
}

func (b Board) rowSolved() bool {
	for _, row := range b.squares {
		sum := 0
		for _, val := range row {
			sum += val
		}
		if sum == -5 {
			return true
		}
	}
	return false
}

func (b Board) columnSolved() bool {
	for col := 0; col < 5; col++ {
		sum := 0
		for row := 0; row < 5; row++ {
			sum += b.squares[row][col]
		}
		if sum == -5 {
			return true
		}
	}
	return false
}

func (b Board) sumRemaining() int {
	sum := 0
	for _, row := range b.squares {
		for _, square := range row {
			if square != -1 {
				sum += square
			}
		}
	}
	return sum
}

func main() {
	calledNumbers, boards := parseInput()
	for _, v := range calledNumbers {
		b := markBoards(boards, v)
		if b != nil {
			sum := b.sumRemaining()
			fmt.Println("called:", v)
			fmt.Println("sum: ", sum)
			fmt.Println("product: ", sum*v)
			return
		}
	}
}

func markBoards(boards []Board, number int) *Board {
	for _, b := range boards {
		if b.mark(number) {
			return &b
		}
	}
	return nil
}

func parseInput() ([]int, []Board) {
	f, err := os.Open("./input.txt")
	if err != nil {
		panic("Failed to open file")
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)
	called := readCalledNumbers(scanner)
	boards := readBoards(scanner)

	return called, boards
}

func readCalledNumbers(scanner *bufio.Scanner) []int {
	scanner.Scan()
	calledString := scanner.Text()
	return toInts(calledString, ",")

}

func readBoards(scanner *bufio.Scanner) []Board {
	var boards []Board

	for scanner.Scan() {
		text := scanner.Text()
		if len(text) == 0 {
			continue
		}
		rows := [][]int{
			{}, {}, {}, {}, {},
		}

		rows[0] = toInts(text, " ")
		scanner.Scan()
		rows[1] = toInts(scanner.Text(), " ")
		scanner.Scan()
		rows[2] = toInts(scanner.Text(), " ")
		scanner.Scan()
		rows[3] = toInts(scanner.Text(), " ")
		scanner.Scan()
		rows[4] = toInts(scanner.Text(), " ")

		boards = append(boards, Board{squares: rows})
	}
	return boards
}

func toInts(s string, separator string) []int {
	var vals []int
	split := strings.Split(s, separator)
	for _, s := range split {
		var v int
		n, e := fmt.Sscan(s, &v)
		if n == 1 && e == nil {
			vals = append(vals, v)
		}
	}
	return vals
}
