package main

import (
	"fmt"
	"github.com/jacktrusler/goutils"
	"math"
	"slices"
	"strconv"
	"strings"
)

type Occurances map[int]int

func makeIdListsAndMaps(allIds []string) (l1, l2 []int, m2 Occurances) {
	l1 = make([]int, len(allIds))
	l2 = make([]int, len(allIds))
	m2 = make(Occurances)

	for i, ids := range allIds {
		matches := strings.Fields(ids)
		var match1, match2 int
		var err error
		match1, err = strconv.Atoi(matches[0])
		match2, err = strconv.Atoi(matches[1])
		if err != nil {
			panic("failed string conv")
		}

		l1[i] = match1
		l2[i] = match2
		m2[match2]++
	}
	return l1, l2, m2
}

func part1and2(input string) {
	allIds := strings.Split(input, "\n")
	list1, list2, m2 := makeIdListsAndMaps(allIds)

	slices.Sort(list1)
	slices.Sort(list2)

	var total float64
	var total2 int
	for i, id := range list1 {
		distance := list1[i] - list2[i]
		total += math.Abs(float64(distance))
		if m2[id] != 0 {
			total2 += id * m2[id]
		}
	}
	fmt.Println("----- Part 1 -----")
	fmt.Println(int(total))
	fmt.Println("----- Part 2 -----")
	fmt.Println(total2)
}

func Day1() {
	input := goutils.FileAsString("./input/2024-01-input.txt")
	part1and2(input)
}
