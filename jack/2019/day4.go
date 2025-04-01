package main

import (
	"fmt"
)

// range 347312 - 805915

func Day4() {
	high := 805915
	totalPwords := make([]int, 0)

	for low := 347312; low < high; low++ {
		temp := low
		matches := 0
		for temp > 0 {
			num2 := temp % 10
			temp /= 10
			num1 := temp % 10
			if (num2 < num1) {
				break;
			}
			if (num2 == num1) {
			  matches++	
			}
		}
		if temp == 0 && matches > 0 {
			totalPwords = append(totalPwords, low)
		}
	}
	fmt.Println("Part 1: ", len(totalPwords))
	
	newPwords := 0
	for _, pword := range totalPwords {
		numberMap := make(map[int]int)
		temp := pword
		for temp > 0 {
			num := temp % 10
			numberMap[num] += 1
			temp /= 10
		}
		for _, value := range numberMap {
			if (value == 2) {
				newPwords++
				break
			}
		}
	}
	fmt.Println("Part 2: ", newPwords)
}