package main

import (
	"fmt"
	"github.com/jacktrusler/goutils"
	"strconv"
	"strings"
)

func numberedReports(stringReports []string) [][]int {
	numReports := make([][]int, 0, len(stringReports))
	for _, report := range stringReports {
		oneReport := make([]int, 0)
		levels := strings.Fields(report)
		for _, level := range levels {
			n, err := strconv.Atoi(level)
			if err != nil {
				panic("failed string conv")
			}
			oneReport = append(oneReport, n)
		}
		numReports = append(numReports, oneReport)
	}
	return numReports
}

func day2part1(input string) {
	reports := strings.Split(input, "\n")
	numReports := numberedReports(reports)

	var dir string
	safeReports := 0

reports:
	for _, report := range numReports {
		for i := 0; i < len(report)-1; i++ {
			if report[0] < report[1] {
				dir = "inc"
			} else {
				dir = "dec"
			}

			switch dir {
			case "inc":
				if report[i+1] < report[i] || report[i+1]-report[i] > 3 || report[i+1]-report[i] == 0 {
					continue reports
				}
			case "dec":
				if report[i] < report[i+1] || report[i]-report[i+1] > 3 || report[i]-report[i+1] == 0 {
					continue reports
				}
			}
		}
		safeReports++
	}

	fmt.Println(safeReports)
}

func day2part2(input string) {
	reports := strings.Split(input, "\n")
	numReports := numberedReports(reports)

	var dir string
	safeReports := 0

reports:
	for _, report := range numReports {
		problems := 0
		for i := 0; i < len(report)-1; i++ {
			if report[0] < report[1] {
				dir = "inc"
			} else {
				dir = "dec"
			}

			switch dir {
			case "inc":
				if report[i+1] < report[i] || report[i+1]-report[i] > 3 || report[i+1]-report[i] == 0 {
					problems++
				}
			case "dec":
				if report[i] < report[i+1] || report[i]-report[i+1] > 3 || report[i]-report[i+1] == 0 {
					problems++
				}
			}
			if problems > 1 {
				continue reports
			}
		}
		safeReports++
	}

	fmt.Println(safeReports)
}

func Day2() {
	input := goutils.FileAsString("./input/2024-02-input.txt")
	fmt.Println("----- Part 1 -----")
	day2part1(input)
	fmt.Println("----- Part 2 -----")
	day2part2(input)
}
