package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

type Command struct {
	inputs  []string
	outputs []string
}

func main() {
	commands := readFile("./input.txt")
	total := 0
	for _, command := range commands {
		encoding := createEncoding(command.inputs)
		outputVal := value(command.outputs, &encoding)
		total += outputVal
	}
	fmt.Println(total)
}

func value(outputs []string, encoding *map[string]int) int {
	s := ""
	for _, v := range outputs {
		s += fmt.Sprintf("%d", (*encoding)[v])
	}
	i, _ := strconv.Atoi(s)
	return i
}
func createEncoding(inputs []string) map[string]int {
	encoding := make(map[string]int)
	solve1478(&encoding, inputs)
	solve235(&encoding, inputs)
	solve069(&encoding, inputs)
	return encoding
}

func commonSegments(s1, s2 string) int {
	count := 0
	if len(s2) > len(s1) {
		s1, s2 = s2, s1
	}
	for _, c := range s1 {
		if strings.ContainsRune(s2, c) {
			count++
		}
	}
	return count
}

func solve235(encoding *map[string]int, inputs []string) {
	// use 1 and 4 to differentiate 2, 3, 5
	var one, four string
	for k, v := range *encoding {
		if v == 1 {
			one = k
		}
		if v == 4 {
			four = k
		}
	}

	for _, val := range inputs {
		if len(val) != 5 {
			continue
		}
		if commonSegments(val, one) == 2 {
			(*encoding)[val] = 3
		} else if commonSegments(val, four) == 2 {
			(*encoding)[val] = 2
		} else {
			(*encoding)[val] = 5
		}
	}
}

func solve069(encoding *map[string]int, inputs []string) {
	// use 7 and 4 to differentiate 0, 6, 9
	var seven, four string
	for k, v := range *encoding {
		if v == 7 {
			seven = k
		}
		if v == 4 {
			four = k
		}
	}

	for _, val := range inputs {
		if len(val) != 6 {
			continue
		}
		if commonSegments(val, seven) == 2 {
			(*encoding)[val] = 6
		} else if commonSegments(val, four) == 4 {
			(*encoding)[val] = 9
		} else {
			(*encoding)[val] = 0
		}
	}
}

func solve1478(encoding *map[string]int, inputs []string) {
	e := *encoding
	for _, v := range inputs {
		switch len(v) {
		case 2:
			e[v] = 1
		case 3:
			e[v] = 7
		case 4:
			e[v] = 4
		case 7:
			e[v] = 8
		}
	}
}

func readFile(path string) []Command {
	f, err := os.Open(path)
	if err != nil {
		panic("Cannot open file")
	}
	defer f.Close()

	var commands []Command
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		line := scanner.Text()
		inOut := strings.Split(line, " | ")
		inputs := strings.Split(inOut[0], " ")
		outputs := strings.Split(inOut[1], " ")
		commands = append(commands, Command{inputs: sortValues(inputs), outputs: sortValues(outputs)})
	}
	return commands
}

func sortValues(ins []string) []string {
	outs := make([]string, len(ins))
	for i, v := range ins {
		r := Sortable(v)
		sort.Sort(r)
		outs[i] = string(r)
	}
	return outs
}

type Sortable []rune

func (a Sortable) Len() int           { return len(a) }
func (a Sortable) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a Sortable) Less(i, j int) bool { return a[i] < a[j] }
