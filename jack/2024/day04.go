package main

import (
	"fmt"
	"github.com/jacktrusler/goutils"
	"strings"
)

var (
	ws      []string
	total   int
	totalP2 int
)

func makeSusSearch(input string) {
	lines := strings.Split(input, "\n")
	runeArr := make([]rune, len(lines[0])+2)
	for i := range runeArr {
		runeArr[i] = '0'
	}

	wordSearch := make([]string, len(lines)+2)
	wordSearch[0] = string(runeArr)
	wordSearch[len(lines)+1] = string(runeArr)
	fmt.Println(wordSearch[0])
	for i, line := range lines {
		line = "0" + line + "0"
		wordSearch[i+1] = line
		fmt.Println(line)
	}
	fmt.Println(wordSearch[len(lines)+1])
	ws = wordSearch
}

func boggleStyle(y, x int, word string) {
	//exit conditions
	if ws[y][x] == '0' {
		return
	}
	if word == "XA" || word == "XX" || word == "XS" || word == "XMS" || word == "XMX" || word == "XMM" {
		return
	}
	if len(word) >= 4 {
		return
	}

	word = word + string(ws[y][x])
	if word == "XMAS" {
		total++
		return
	}

	// W
	boggleStyle(y, x-1, word)
	// NW
	boggleStyle(y-1, x-1, word)
	// N
	boggleStyle(y-1, x, word)
	// NE
	boggleStyle(y-1, x+1, word)
	// E
	boggleStyle(y, x+1, word)
	// SE
	boggleStyle(y+1, x+1, word)
	// S
	boggleStyle(y+1, x, word)
	// SW
	boggleStyle(y+1, x-1, word)

	return
}

func travel(y, x, dy, dx int, w string) {
	if w == "XMAS" {
		total++
	}
	if ws[y][x] == '0' {
		return
	}
	if len(w) == 4 {
		return
	}
	travel(y+dy, x+dx, dy, dx, w+string(ws[y][x]))
}

func day4part1() {
	for y := 1; y < len(ws)-1; y++ {
		line := ws[y]
		for x, rune := range line {
			if rune == 'X' {
				// NW
				travel(y-1, x-1, -1, -1, "X")
				// N
				travel(y-1, x, -1, 0, "X")
				// NE
				travel(y-1, x+1, -1, 1, "X")
				// E
				travel(y, x+1, 0, 1, "X")
				// SE
				travel(y+1, x+1, 1, 1, "X")
				// S
				travel(y+1, x, 1, 0, "X")
				// SW
				travel(y+1, x-1, 1, -1, "X")
				// W
				travel(y, x-1, 0, -1, "X")
			}
		}
	}
	fmt.Println(total)

}

func day4part2() {
	for y := 1; y < len(ws)-1; y++ {
		line := ws[y]
		for x, rune := range line {
			if rune == 'A' {

				runeMap := make(map[string]byte)
				runeMap["NW"] = ws[y-1][x-1]
				runeMap["NE"] = ws[y-1][x+1]
				runeMap["SW"] = ws[y+1][x-1]
				runeMap["SE"] = ws[y+1][x+1]

				if runeMap["NW"] == 'M' && runeMap["NE"] == 'M' && runeMap["SW"] == 'S' && runeMap["SE"] == 'S' {
					totalP2++
				}
				if runeMap["NW"] == 'M' && runeMap["NE"] == 'S' && runeMap["SW"] == 'M' && runeMap["SE"] == 'S' {
					totalP2++
				}
				if runeMap["NW"] == 'S' && runeMap["NE"] == 'M' && runeMap["SW"] == 'S' && runeMap["SE"] == 'M' {
					totalP2++
				}
				if runeMap["NW"] == 'S' && runeMap["NE"] == 'S' && runeMap["SW"] == 'M' && runeMap["SE"] == 'M' {
					totalP2++
				}
			}
		}
	}
	fmt.Println(totalP2)
}

func Day4() {
	input := goutils.FileAsString("./input/2024-04-input.txt")
	makeSusSearch(input)
	fmt.Println("----- Part 1 -----")
	day4part1()
	fmt.Println("----- Part 2 -----")
	day4part2()
}
