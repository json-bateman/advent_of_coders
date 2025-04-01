package main

import (
	"fmt"
	"log"
	"strconv"
	"strings"
)

func Day2() {
	fileAsString := FileAsString("day2.txt")
	fileAsString = (strings.TrimSuffix)(fileAsString, "\n")
	var lines []string = strings.Split(fileAsString, ",")
	var nums []int
	for i := 0; i < len(lines); i++ {
		if lines[i] == "" {
			continue
		}
		num, err := strconv.Atoi(lines[i])
		if err != nil {
			log.Fatal(err)
		}
		nums = append(nums, num)
	}
	// Problem says replace index 1 and 2 before starting
	nums[1] = 40
	nums[2] = 19
	count := 0
	for i := 0; i < len(nums); i = i + 4 {
		count++
		switch nums[i] {
		case 1:
			nums[nums[i+3]] = nums[nums[i+1]] + nums[nums[i+2]]
		case 2:
			nums[nums[i+3]] = nums[nums[i+1]] * nums[nums[i+2]]
		case 99:
			break
		}
	}
	fmt.Println(nums)
}
