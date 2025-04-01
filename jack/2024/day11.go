package main

import (
	"fmt"
	u "github.com/jacktrusler/goutils"
	"strconv"
	"strings"
)

func applyRulesLinkedList(head *u.Node) *u.Node {
	curr := head
	for curr != nil {
		temp := curr.Val
		rocko := []rune(strconv.Itoa(temp))

		if curr.Val == 0 {
			curr.Val = 1
			curr = curr.Next
			continue
		}

		if len(rocko)%2 == 0 {
			mid := len(rocko) / 2

			lhs := string(rocko[:mid])
			rhs := string(rocko[mid:])

			lhsI, err := strconv.Atoi(lhs)
			rhsI, err := strconv.Atoi(rhs)
			if err != nil {
				panic(err)
			}
			curr.Val = lhsI
			temp := curr.Next
			curr.Next = &u.Node{
				Val:  rhsI,
				Next: temp,
			}

			curr = curr.Next.Next
			continue
		}
		curr.Val *= 2024
		curr = curr.Next
	}
	return head
}

func day11part1ForFun(input string) {
	rocks := strings.Split(input, " ")
	rocky := u.StringArrAtoI(rocks)
	head := u.IntArrayToLinkedList(rocky)

	blinks := 25
	for i := 0; i < blinks; i++ {
		head := applyRulesLinkedList(head)
		ts := 0
		for head != nil {
			ts++
			head = head.Next
		}
	}
	totalP1 := 0
	for head != nil {
		totalP1++
		head = head.Next
	}
	fmt.Println(totalP1)
}

func applyBlink(oldMap map[int]int) (newMap map[int]int) {
	newMap = make(map[int]int)
	for k := range oldMap {
		rocko := []rune(strconv.Itoa(k))
		if k == 0 {
			newMap[1] += oldMap[k]
			continue
		}

		if len(rocko)%2 == 0 {
			mid := len(rocko) / 2

			lhs := string(rocko[:mid])
			rhs := string(rocko[mid:])

			lhsI, err := strconv.Atoi(lhs)
			rhsI, err := strconv.Atoi(rhs)
			if err != nil {
				panic(err)
			}
			newMap[lhsI] += oldMap[k]
			newMap[rhsI] += oldMap[k]
			continue
		}

		newMap[k*2024] += oldMap[k]
	}
	return newMap
}

func day11part1and2(input string, blinks int) {
	rocks := strings.Split(input, " ")
	rocky := u.StringArrAtoI(rocks)
	rockMap := make(map[int]int)

	for i := 0; i < len(rocky); i++ {
		rockMap[rocky[i]] = 1
	}

	// make map, return new map, use new map in next loop
	oldMap := rockMap
	for i := 0; i < blinks; i++ {
		oldMap = applyBlink(oldMap)
	}

	total := 0
	for _, v := range oldMap {
		total += v
	}
	fmt.Println(total)
}

func Day11() {
	input := u.FileAsString("./input/2024-11-input.txt")
	fmt.Println("----- Part 1 -----")
	day11part1and2(input, 25)
	fmt.Println("----- Part 2 -----")
	day11part1and2(input, 75)
}
