package main

import (
	"fmt"
	"github.com/jacktrusler/goutils"
	"math"
	"slices"
	"strings"
)

func day5part1and2(input string) {
	instr := strings.Split(input, "\n\n")
	rules := strings.Split(instr[0], "\n")
	updates := strings.Split(instr[1], "\n")

	// rule map
	rM := make(map[int][]int)

	var lhs, rhs int
	for _, rule := range rules {
		fmt.Sscanf(rule, "%d|%d", &lhs, &rhs)
		rM[lhs] = append(rM[lhs], rhs)
	}

	correctUpdates := [][]int{}
	incorrectUpdates := [][]int{}

	for _, update := range updates {
		pages := strings.Split(update, ",")
		pagesNum := goutils.StringArrAtoI(pages)
		correctOrder := true

		for i := 0; i < len(pagesNum); {
			for _, pagesAfter := range rM[pagesNum[i]] {
				afterPageIndex := slices.Index(pagesNum, pagesAfter)
				// if after pages come before current page, fail
				if afterPageIndex != -1 && afterPageIndex < i {
					correctOrder = false

					// Part 2, swap incorrect pages
					temp := pagesNum[i]
					pagesNum[i] = pagesAfter
					pagesNum[afterPageIndex] = temp
					// start the loop over if pages need swapped
					i = 0
				}
			}
			i++
		}
		if correctOrder {
			correctUpdates = append(correctUpdates, pagesNum)
		} else {
			incorrectUpdates = append(incorrectUpdates, pagesNum)
		}
	}

	ans1 := 0
	ans2 := 0
	// assume all updates have an odd number of pages
	for _, update := range correctUpdates {
		middleI := int(math.Floor(float64(len(update) / 2)))
		ans1 += update[middleI]
	}
	for _, update := range incorrectUpdates {
		// assume all updates have an odd number of pages
		middleI := int(math.Floor(float64(len(update) / 2)))
		ans2 += update[middleI]
	}

	fmt.Println(ans1)
	fmt.Println("----- Part 2 -----")
	fmt.Println(ans2)

	// if key is after any value -- incorrect

}

func Day5() {
	input := goutils.FileAsString("./input/2024-05-input.txt")
	fmt.Println("----- Part 1 -----")
	day5part1and2(input)
}
