package main

import (
	"fmt"
	"strconv"
)

var INPUT = []int64{0, 37551, 469, 63, 1, 791606, 2065, 9983586}
var TEST_INPUT = []int64{125, 17}

func main() {
	// Part1()
	Part2()
}

func rule1(_ int64) int64 {
	return 1
}

func rule2(stone int64) []int64 {
	str := strconv.FormatInt(stone, 10)
	l := len(str) / 2
	s1 := str[:l]
	s2 := str[l:]
	i1, _ := strconv.Atoi(s1)
	i2, _ := strconv.Atoi(s2)
	return []int64{int64(i1), int64(i2)}
}

func rule3(stone int64) int64 {
	return stone * 2024
}

// PART1
func Part1() {
	const BLINKS = 25
	stones := INPUT
	for i := 0; i < BLINKS; i++ {
		stones = blink(stones)
	}
	fmt.Println(len(stones))
}

func blink(stones []int64) []int64 {
	result := []int64{}
	for _, stone := range stones {
		if stone == 0 {
			result = append(result, rule1(stone))
		} else if len(strconv.FormatInt(stone, 10))%2 == 0 {
			result = append(result, rule2(stone)...)
		} else {
			result = append(result, rule3(stone))
		}
	}
	return result
}

// PART2
func Part2() {
	const BLINKS = 75
	stones := INPUT
	var total int64 = 0
	for _, stone := range stones {
		total += handleBlink(stone, BLINKS)
	}
	fmt.Println(total)
}

var cache = map[Key]int64{}

type Key struct {
	value, blinks int64
}

func handleBlink(stone int64, blinks int64) int64 {
	if blinks == 0 {
		return 1
	}

	key := Key{stone, blinks}
	if v, exists := cache[key]; exists {
		return v
	}
	var r int64 = 0
	if stone == 0 {
		r = handleBlink(1, blinks-1)
	} else if len(strconv.FormatInt(stone, 10))%2 == 0 {
		str := strconv.FormatInt(stone, 10)
		l := len(str) / 2
		s1 := str[:l]
		s2 := str[l:]
		i1, _ := strconv.ParseInt(s1, 10, 64)
		i2, _ := strconv.ParseInt(s2, 10, 64)
		r = handleBlink(i1, blinks-1) + handleBlink(i2, blinks-1)
	} else {
		r = handleBlink(stone*2024, blinks-1)
	}
	cache[key] = r
	return r
}
