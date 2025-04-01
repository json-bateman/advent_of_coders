package main

import (
	"fmt"
	u "github.com/jacktrusler/goutils"
	"regexp"
)

func day13part1(theClaaawww [][]string) {
	tokens := 0.0
	for _, match := range theClaaawww {
		intArr := u.StringArrToFloat(match)
		aX, aY, bX, bY, prizeX, prizeY := intArr[0], intArr[1], intArr[2], intArr[3], intArr[4], intArr[5]

		// Pre-Math Era coding
		// for a := 0; a <= 100; a++ {
		// 	for b := 0; b <= 100; b++ {
		// 		if prizeX-(a*aX)-(b*bX) == 0 && prizeY-(a*aY)-(b*bY) == 0 {
		// 			tokens += (a * 3) + b
		// 		}
		// 	}
		// }
		// Post-Math Era algebra
		bPresses := (aX*prizeY - aY*prizeX) / (aX*bY - aY*bX)
		aPresses := (prizeX - bPresses*bX) / (aX)

		// Kinda a hack, if bPresses or aPresses is not a whole number, it fails
		if !u.IsWholeNumber(aPresses) || !u.IsWholeNumber(bPresses) {
			continue
		}

		tokens += (aPresses * 3) + bPresses

	}
	fmt.Println(tokens)
}

func day13part2(theClaaawww [][]string) {
	tokens := 0.0
	for _, match := range theClaaawww {
		intArr := u.StringArrToFloat(match)
		aX, aY, bX, bY, prizeX, prizeY := intArr[0], intArr[1], intArr[2], intArr[3], intArr[4], intArr[5]
		nPrizeX := prizeX + 10000000000000
		nPrizeY := prizeY + 10000000000000

		// lmao algebra time, hold onto your hats
		// prizeX = aPresses * aX + bPresses * bX
		// prizeY = aPresses * aY + bPresses * bY
		// ??? (had to sharpen a pencil for this one)
		// bPresses = (aX * prizeY - aY * prizeX) / (aX * bY - aY * bX)

		bPresses := (aX*nPrizeY - aY*nPrizeX) / (aX*bY - aY*bX)
		aPresses := (nPrizeX - bPresses*bX) / aX

		if !u.IsWholeNumber(aPresses) || !u.IsWholeNumber(bPresses) {
			continue
		}

		tokens += (aPresses * 3) + bPresses
	}
	fmt.Println(int(tokens))
}

func Day13() {
	input := u.FileAsString("./input/2024-13-input.txt")
	re := regexp.MustCompile(`A: X\+(\d+), Y\+(\d+)\n.+?X\+(\d+), Y\+(\d+)\n.+?X=(\d+), Y=(\d+)`)
	matches := re.FindAllStringSubmatch(input, -1)
	newM := [][]string{}
	for _, match := range matches {
		newM = append(newM, match[1:])
	}
	fmt.Println("----- Part 1 -----")
	day13part1(newM)
	fmt.Println("----- Part 2 -----")
	day13part2(newM)
}
