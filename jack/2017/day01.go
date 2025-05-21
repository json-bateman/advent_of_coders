package main

import (
	"fmt"
	u "github.com/jacktrusler/goutils"
)

func day01() {
	input := u.FileAsString("data/2017-01-input.txt")
	length := len(input)
	sum := 0

	// Part 1
	for i := range input {
		next := (i + 1) % length
		if input[i] == input[next] {
			sum += int(input[i] - '0')
		}
	}

	// Part 2
	sum2 := 0
	for i := range input {
		next := (i + (len(input) / 2)) % length
		if input[i] == input[next] {
			sum2 += int(input[i] - '0')
		}
	}

	fmt.Println(sum)
	fmt.Println(sum2)
}
