package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	// Part1()
	Part2()
}

type Rule struct {
	Before, After int
}

func (r Rule) Equal(other Rule) bool {
	return r.Before == other.Before && r.After == other.After
}

func readInput(path string) (map[Rule]struct{}, [][]int) {
	f, _ := os.ReadFile(path)
	s := strings.Split(string(f), "\n\n")
	rules := s[0]
	updates := s[1]
	return parseRules(rules), parseUpdates(updates)
}

func parseRules(input string) map[Rule]struct{} {
	rules := make(map[Rule]struct{})
	strs := strings.Split(input, "\n")
	for _, v := range strs {
		var i, p int
		fmt.Sscanf(v, "%d|%d", &i, &p)
		rules[Rule{i, p}] = struct{}{}
	}
	return rules
}

func parseUpdates(input string) [][]int {
	updates := [][]int{}
	strs := strings.Split(input, "\n")
	for _, v := range strs {
		pages := strings.Split(v, ",")
		update := []int{}
		for _, pgS := range pages {
			pg, _ := strconv.Atoi(pgS)
			update = append(update, pg)
		}
		updates = append(updates, update)
	}
	return updates
}

// PART1
func Part1() {
	// rules, updates := readInput("./input/test.txt")
	rules, updates := readInput("./input/input.txt")

	validUpdates := 0
	for _, update := range updates {
		validUpdates += updateIsValid(update, rules)
	}
	fmt.Println("Valid updates: ", validUpdates)
}

func updateIsValid(update []int, rules map[Rule]struct{}) int {
	for i := 0; i < len(update)-1; i++ {
		for j := i; j < len(update)-1; j++ {
			r := Rule{update[j+1], update[i]}
			_, exists := rules[r]
			if exists {
				return 0
			}
		}
	}
	//find middle number
	return update[len(update)/2]
}

// Part 2
func Part2() {
	// rules, updates := readInput("./input/test.txt")
	rules, updates := readInput("./input/input.txt")
	updated := 0
	for _, update := range updates {
		if updateIsValid(update, rules) == 0 {
			updated += repairUpdate(update, rules)
		}
	}
	fmt.Println("Updated: ", updated)
}

func repairUpdate(update []int, rules map[Rule]struct{}) int {
	for i := 0; i < len(update)-1; i++ {
		for j := i; j < len(update)-1; j++ {
			r := Rule{update[j+1], update[i]}
			_, exists := rules[r]
			if exists {
				repaired := []int{}
				repaired = append(repaired, update[:i]...)
				repaired = append(repaired, update[j+1])
				repaired = append(repaired, update[i:j+1]...)
				repaired = append(repaired, update[j+2:]...)
				return repairUpdate(repaired, rules)
			}
		}
	}

	//find middle number
	return update[len(update)/2]
}
