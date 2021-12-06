package main

import "fmt"

func main() {
	// input := []int{3, 4, 3, 1, 2}
	input := []int{1, 4, 3, 3, 1, 3, 1, 1, 1, 2, 1, 1, 1, 4, 4, 1, 5, 5, 3, 1, 3, 5, 2, 1, 5, 2, 4, 1, 4, 5, 4, 1, 5, 1, 5, 5, 1, 1, 1, 4, 1, 5, 1, 1, 1, 1, 1, 4, 1, 2, 5, 1, 4, 1, 2, 1, 1, 5, 1, 1, 1, 1, 4, 1, 5, 1, 1, 2, 1, 4, 5, 1, 2, 1, 2, 2, 1, 1, 1, 1, 1, 5, 5, 3, 1, 1, 1, 1, 1, 4, 2, 4, 1, 2, 1, 4, 2, 3, 1, 4, 5, 3, 3, 2, 1, 1, 5, 4, 1, 1, 1, 2, 1, 1, 5, 4, 5, 1, 3, 1, 1, 1, 1, 1, 1, 2, 1, 3, 1, 2, 1, 1, 1, 1, 1, 1, 1, 2, 1, 1, 1, 1, 2, 1, 1, 1, 1, 1, 1, 4, 5, 1, 3, 1, 4, 4, 2, 3, 4, 1, 1, 1, 5, 1, 1, 1, 4, 1, 5, 4, 3, 1, 5, 1, 1, 1, 1, 1, 5, 4, 1, 1, 1, 4, 3, 1, 3, 3, 1, 3, 2, 1, 1, 3, 1, 1, 4, 5, 1, 1, 1, 1, 1, 3, 1, 4, 1, 3, 1, 5, 4, 5, 1, 1, 5, 1, 1, 4, 1, 1, 1, 3, 1, 1, 4, 2, 3, 1, 1, 1, 1, 2, 4, 1, 1, 1, 1, 1, 2, 3, 1, 5, 5, 1, 4, 1, 1, 1, 1, 3, 3, 1, 4, 1, 2, 1, 3, 1, 1, 1, 3, 2, 2, 1, 5, 1, 1, 3, 2, 1, 1, 5, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 2, 5, 1, 1, 1, 1, 3, 1, 1, 1, 1, 1, 1, 1, 1, 5, 5, 1}

	// convert the input array into a consistent size. Index is days until reproduce, value is number of fish at that time.
	fish := []int{0, 0, 0, 0, 0, 0, 0, 0, 0}
	for _, v := range input {
		fish[v] += 1
	}

	days := 256
	for i := 0; i < days; i++ {
		fish = ageFish(fish)
	}
	fmt.Println(countFish(fish))
}

func ageFish(fish []int) []int {
	src := fish
	dst := []int{}

	dst = append(dst, src[1:]...)
	dst = append(dst, src[0])
	dst[6] += src[0]
	return dst
}

func countFish(fish []int) int {
	count := 0
	for _, f := range fish {
		count += f
	}
	return count
}
