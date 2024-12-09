package main

import (
	"fmt"
	"os"
)

func main() {
	// Part1()
	Part2()
}

type Block struct {
	ID   int
	Size int
}

func (b Block) IsFreeSpace() bool {
	return b.ID == -1
}

func readInput(path string) []Block {
	f, _ := os.ReadFile(path)
	content := string(f)
	id := 0
	blocks := []Block{}
	for i, c := range content {
		size := int(c - '0')
		if i%2 == 0 {
			blocks = append(blocks, Block{id, size})
			id++
		} else {
			blocks = append(blocks, Block{-1, size})
		}
	}
	return blocks
}

// PART1
func Part1() {
	// blocks := readInput("./input/test.txt")
	blocks := readInput("./input/input.txt")

	defrag := []int{}
	last := len(blocks) - 1

	for blockIndex, b := range blocks {
		if b.IsFreeSpace() {
			filled := 0
			for filled < b.Size && blockIndex < last {
				lastBlock := blocks[last]

				if lastBlock.IsFreeSpace() {
					last--
					continue
				}
				if lastBlock.Size <= b.Size-filled {
					// move all
					for i := 0; i < lastBlock.Size; i++ {
						defrag = append(defrag, lastBlock.ID)
					}
					filled += lastBlock.Size
					lastBlock.Size = 0
					lastBlock.ID = -1
					blocks[last] = lastBlock
					last--
				} else {
					//move part
					toFill := b.Size - filled
					for i := 0; i < toFill; i++ {
						defrag = append(defrag, lastBlock.ID)
					}
					filled += toFill
					lastBlock.Size -= toFill
					blocks[last] = lastBlock
				}
			}
		} else {
			for i := 0; i < b.Size; i++ {
				defrag = append(defrag, b.ID)
			}
		}
	}

	var checksum int64 = 0
	for i, v := range defrag {
		checksum += int64(i) * int64(v)
	}
	fmt.Println(checksum)
}
func Part2() {
	// blocks := readInput("./input/test.txt")
	blocks := readInput("./input/input.txt")

	defrag := moveLastBlock(blocks)

	arr := []int{}
	for _, v := range defrag {
		for i := 0; i < v.Size; i++ {
			arr = append(arr, v.ID)
		}
	}
	var checksum int64 = 0
	for i, v := range arr {
		if v != -1 {
			checksum += int64(i) * int64(v)
		}
	}
	fmt.Println(checksum)
}

func moveLastBlock(blocks []Block) []Block {
	for i := len(blocks) - 1; i > 0; i-- {
		b := blocks[i]
		if b.IsFreeSpace() {
			continue
		}
		for j := 0; j < i; j++ {
			v := blocks[j]
			if v.IsFreeSpace() && v.Size >= b.Size {
				newBlocks := []Block{}
				newBlocks = append(newBlocks, blocks[:j]...)
				newBlocks = append(newBlocks, b)
				if b.Size < v.Size {
					v.Size = v.Size - b.Size
					newBlocks = append(newBlocks, v)
				}
				newBlocks = append(newBlocks, blocks[j+1:i]...)
				newBlocks = append(newBlocks, Block{-1, b.Size})
				if i < len(blocks)-1 {
					newBlocks = append(newBlocks, blocks[i+1:]...)
				}
				return moveLastBlock(newBlocks)
			}
		}
	}
	return blocks
}
