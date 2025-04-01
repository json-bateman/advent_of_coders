package main

import (
	"fmt"
	"github.com/jacktrusler/goutils"
	"slices"
	"strconv"
	"strings"
)

func convertToInt(eq string) (int, []int) {
	e := strings.Split(eq, ":")
	lhs, err := strconv.Atoi(e[0])
	if err != nil {
		panic(err)
	}
	rhsStr := strings.Trim(e[1], " ")
	rhs := goutils.StringArrAtoI(strings.Split(rhsStr, " "))

	return lhs, rhs
}

func day7part1(cE []string) {
	finalArr := []int{}
	for _, eq := range cE {
		lhs, rhs := convertToInt(eq)

		// Possible Values
		// 2 then 4 then 8 then 16 etc..
		pV := make([]int, 2)
		pV[0] = rhs[0] + rhs[1]
		pV[1] = rhs[0] * rhs[1]
		for i := 2; i < len(rhs); i++ {
			temp := []int{}

			for j := 0; j < len(pV); j++ {
				temp = append(temp, rhs[i]*pV[j])
				temp = append(temp, rhs[i]+pV[j])
			}
			pV = temp
		}
		if slices.Contains(pV, lhs) {
			finalArr = append(finalArr, lhs)
		}
	}
	var total int
	for _, val := range finalArr {
		total += val
	}
	fmt.Println(total)
}

func day7part2(cE []string) {
	finalArr := []int{}
	for _, eq := range cE {
		lhs, rhs := convertToInt(eq)

		pV := make([]int, 3)
		pV[0] = rhs[0] + rhs[1]
		pV[1] = rhs[0] * rhs[1]
		r := strconv.Itoa(rhs[0])
		pj := strconv.Itoa(rhs[1])
		rpj, _ := strconv.Atoi(r + pj)
		pV[2] = rpj

		for i := 2; i < len(rhs); i++ {
			temp := []int{}

			for j := 0; j < len(pV); j++ {
				temp = append(temp, rhs[i]*pV[j])
				temp = append(temp, rhs[i]+pV[j])
				r := strconv.Itoa(rhs[i])
				pj := strconv.Itoa(pV[j])
				rpj, _ := strconv.Atoi(pj + r)
				temp = append(temp, rpj)
			}
			pV = temp
		}
		if slices.Contains(pV, lhs) {
			finalArr = append(finalArr, lhs)
		}
	}
	var total int
	for _, val := range finalArr {
		total += val
	}
	fmt.Println(total)
}

func Day7() {
	input := goutils.FileAsString("./input/2024-07-input.txt")
	calibrationEquations := strings.Split(input, "\n")
	fmt.Println("----- Part 1 -----")
	day7part1(calibrationEquations)
	fmt.Println("----- Part 2 -----")
	day7part2(calibrationEquations)
}
