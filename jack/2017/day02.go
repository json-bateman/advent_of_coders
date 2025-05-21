package main

import (
	"fmt"
	"math"
	"strings"

	u "github.com/jacktrusler/goutils"
)

func day02() {
	input := u.FileAsString("data/2017-02-input.txt")
	lines := strings.Split(strings.TrimSpace(input), "\n")

	part1 := 0
	part2 := 0
	for _, line := range lines {
		splitLine := strings.Fields(line)
		numLine := u.StringArrAtoI(splitLine)
		max := 0
		min := math.MaxInt

		for i, num := range numLine {
			if num > max {
				max = num
			}
			if num < min {
				min = num
			}
			// part2 another for loop :>
			for j := range numLine {
				x := numLine[i]
				y := numLine[j]
				if x%y == 0 && j != i {
					part2 += x / y
					// This only happens once per line as per puzzle instructions
					break
				}
			}

		}
		part1 += max - min
	}

	fmt.Println(part1)
	fmt.Println(part2)
}
