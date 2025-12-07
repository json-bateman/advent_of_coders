package y2025

import (
	"aoc/lib"
	"net/http"
	"path/filepath"

	"github.com/starfederation/datastar-go/datastar"
)

func Oob(y, x, rows, cols int) bool {
	return y < 0 || y >= rows || x < 0 || x >= cols
}

func removePaper(rollRuneArray [][]rune) (toRemove [][2]int) {
	for y, roll := range rollRuneArray {
		for x, r := range roll {
			if r != '@' {
				continue
			}
			surroundingRolls := 0
			for _, dir := range lib.Surround {
				// if a surrounding point has an @, count it
				if !Oob(y+dir[1], x+dir[0], len(rollRuneArray), len(roll)) && rollRuneArray[y+dir[1]][x+dir[0]] == '@' {
					surroundingRolls++
				}
			}
			if surroundingRolls < 4 {
				toRemove = append(toRemove, [2]int{y, x})
			}
		}
	}
	return toRemove
}

func D4SSE(w http.ResponseWriter, r *http.Request) {
	sse := datastar.NewSSE(w, r)
	inputPath, err := filepath.Abs("web/y2025/input/2025-04-input.txt")
	if err != nil {
		return
	}
	strArr := lib.SplitFileOnNewLines(inputPath)
	allRolls := lib.MakeRuneArr(strArr)

	accessibleRolls := 0
	part1 := 0
	iterations := 0
	for {
		toRemove := removePaper(allRolls)
		iterations++
		if len(toRemove) == 0 {
			break
		}
		accessibleRolls += len(toRemove)
		for _, pos := range toRemove {
			allRolls[pos[0]][pos[1]] = '.'
		}

		// Answer after first iteration in Part 1
		if iterations == 1 {
			part1 = accessibleRolls
		}
	}
	sse.PatchElementTempl(BothPartsD4(part1, accessibleRolls))
}
