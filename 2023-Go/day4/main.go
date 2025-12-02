package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Card struct {
	winning []int
	card    []int
}

func (c Card) matches() int {
	matches := 0
	for _, v := range c.winning {
		if containsInt(c.card, v) {
			matches++
		}
	}
	return matches
}

func main() {
	input := readInput("./puzzle.input")
	part1(input)
	part2(input)
}

func readInput(path string) []Card {
	f, _ := os.Open(path)
	defer f.Close()
	cards := []Card{}
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		line := scanner.Text()
		cards = append(cards, parseCard(line))
	}
	return cards
}

func parseCard(line string) Card {
	cardStrings := strings.Split(line, ":")
	cardStrings = strings.Split(cardStrings[1], "|")

	//winning numbers
	ws := cardStrings[0]
	w := []int{}
	winNums := strings.Split(ws, " ")
	for _, sVal := range winNums {
		if len(sVal) == 0 {
			continue
		}
		val, _ := strconv.Atoi(sVal)
		w = append(w, val)
	}

	//card numbers
	cs := cardStrings[1]
	c := []int{}
	cNums := strings.Split(cs, " ")
	for _, sVal := range cNums {
		if len(sVal) == 0 {
			continue
		}
		val, _ := strconv.Atoi(sVal)
		c = append(c, val)
	}

	return Card{
		winning: w,
		card:    c,
	}
}

func part1(cards []Card) {
	score := 0
	for _, card := range cards {
		score += scoreCard(card)
	}
	fmt.Println("Score: ", score)
}

func scoreCard(card Card) int {
	score := 0
	for i := 0; i < card.matches(); i++ {
		if score == 0 {
			score = 1
		} else {
			score *= 2
		}
	}
	return score
}

func containsInt(ints []int, val int) bool {
	for _, v := range ints {
		if v == val {
			return true
		}
	}
	return false
}

func part2(cards []Card) {
	won := []Card{}
	for i := 0; i < len(cards); i++ {
		wonCards := wonCards(i, cards)
		won = append(won, wonCards...)
	}
	fmt.Println("Cards: ", len(won))
}

var winningCardMap = map[int][]Card{}

func wonCards(idx int, cards []Card) []Card {
	won := []Card{cards[idx]}

	m := cards[idx].matches()
	for i := 0; i < m; i++ {
		c, ok := winningCardMap[idx+i+1]
		if !ok {
			c = wonCards(idx+i+1, cards)
			winningCardMap[idx+i+1] = c
		}
		won = append(won, c...)
	}

	return won
}
