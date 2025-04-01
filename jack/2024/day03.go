package main

import (
	"fmt"
	"github.com/jacktrusler/goutils"
)

func processMulString(input string) int {
	var total int
	for i := 0; i < len(input)-4; i++ {
		mul := input[i : i+4]

		st := string(mul)

		if st == "mul(" {
			// at '(' -- end of mul(
			i += 3
			var lhs, rhs int
			left := true
			for {
				i++
				if input[i] >= '0' && input[i] <= '9' {

					if left {
						lhs = (lhs * 10) + int(input[i]-'0')
						continue
					} else {
						rhs = (rhs * 10) + int(input[i]-'0')
						continue
					}

				}
				if input[i] == ',' {
					left = false
					continue
				}
				if input[i] == ')' {
					total += lhs * rhs
					break
				} else {
					//if not a closing parens or number or , the whole thing is invalid
					break
				}
			}
		}

	}
	return total
}

func day3part1(input string) {
	ans := processMulString(input)

	fmt.Println(ans)
}

func day3part2(input string) {
	// Regex can ligma and suggma and zuggma
	bigFatMulSlice := make([]byte, 0)

	toggle := true
	for i := 0; i < len(input); {

		if i+4 <= len(input) {
			do := input[i : i+4]
			if string(do) == "do()" {
				toggle = true
				i += 4
				continue
			}
		}
		if i+7 <= len(input) {
			dont := input[i : i+7]
			if string(dont) == "don't()" {
				toggle = false
				i += 7
				continue
			}
		}

		if toggle {
			bigFatMulSlice = append(bigFatMulSlice, input[i])
		}
		i++
	}

	ans := processMulString(string(bigFatMulSlice))
	fmt.Println(ans)

}

func Day3() {
	input := goutils.FileAsString("./input/2024-03-input.txt")
	fmt.Println("----- Part 1 -----")
	day3part1(input)
	fmt.Println("----- Part 2 -----")
	day3part2(input)
}
