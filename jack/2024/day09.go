package main

import (
	"fmt"
	"github.com/jacktrusler/goutils"
	"strings"
)

// 2333133121414131402
// 00...111...2...333.44.5555.6666.777.888899
func createDiskMap(input string) []int {
	blocks := []int{}
	nums := strings.Split(input, "")
	diskMap := goutils.StringArrAtoI(nums)

	var counter int
	for i := 0; i < len(diskMap); i += 2 {
		for j := 0; j < diskMap[i]; j++ {
			blocks = append(blocks, counter)
		}
		if i+1 >= len(diskMap) {
			break
		}
		for k := 0; k < diskMap[i+1]; k++ {
			blocks = append(blocks, -1)
		}
		counter++
	}
	return blocks
}
func day9part1(input string) {
	blocks := createDiskMap(input)
	for i := len(blocks) - 1; i > 0; i-- {
		for j := 0; j < i; j++ {
			if blocks[j] == -1 {
				temp := blocks[j]
				blocks[j] = blocks[i]
				blocks[i] = temp
				break
			}
		}
	}
	var part1 int
	for i, val := range blocks {
		if val == -1 {
			break
		}
		part1 += val * i

	}
	fmt.Println(part1)
}

func findFreeSpaceIndex(blocks []int, spaceNeeded int, endIndex int) int {
	space := 0
	freeSpaceIndex := -1
	for j := 0; j < endIndex; j++ {
		if blocks[j] == -1 {
			freeSpaceIndex = j
			space++
		} else {
			space = 0
			freeSpaceIndex = -1
		}

		if space == spaceNeeded {
			return freeSpaceIndex - space + 1
		}
	}
	return -1
}

func day9part2(input string) {
	blocks := createDiskMap(input)

	startId := -1
	for i := len(blocks) - 1; i > 0; i-- {
		if blocks[i] != -1 {
			startId = blocks[i]
			break
		}
	}

	tempArr := []int{}
	for i := startId; i > 0; i-- {
		// check for number of times id appears from end of array
		// save the indexes where it appears
		for j := len(blocks) - 1; j > 0; j-- {
			if blocks[j] == i {
				tempArr = append(tempArr, j)
			}
		}
		// check for first index with amount of free space needed
		ind := findFreeSpaceIndex(blocks, len(tempArr), tempArr[len(tempArr)-1])
		if ind == -1 {
			tempArr = []int{}
			continue
		}
		// swap indexes
		for j := 0; j < len(tempArr); j++ {
			temp := blocks[ind+j]
			blocks[ind+j] = blocks[tempArr[j]]
			blocks[tempArr[j]] = temp
		}
		tempArr = []int{}

	}
	var part2 int
	for i, val := range blocks {
		if val == -1 {
			continue
		}
		part2 += val * i

	}
	fmt.Println(part2)
}

func Day9() {
	input := goutils.FileAsString("./input/2024-09-input.txt")
	fmt.Println("----- Part 1 -----")
	day9part1(input)
	fmt.Println("----- Part 2 -----")
	day9part2(input)
}
