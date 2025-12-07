package y2025

import (
	"aoc/lib"
	"math"
	"net/http"
	"path/filepath"
	"time"

	"github.com/starfederation/datastar-go/datastar"
)

func findLocalMax(bank string, startInd, remaining int) (ind, maxBattery int) {
	for i := startInd + 1; i < len(bank)-remaining; i++ {
		currentDigit := int(bank[i] - '0')
		if currentDigit > maxBattery {
			maxBattery = currentDigit
			ind = i
		}
	}
	return ind, maxBattery
}

func calculateJoltage(bank string, numBatteries int) int {
	maxBattery := 0
	nextIndex := -1
	lineValue := 0
	for i := numBatteries - 1; i >= 0; i-- {
		nextIndex, maxBattery = findLocalMax(bank, nextIndex, i)
		lineValue += maxBattery * int(math.Pow(10.0, float64(i)))
	}
	return lineValue
}

func D3SSE(w http.ResponseWriter, r *http.Request) {
	sse := datastar.NewSSE(w, r)
	inputPath, err := filepath.Abs("web/y2025/input/2025-03-input.txt")
	if err != nil {
		return
	}
	allBanks := lib.SplitFileOnNewLines(inputPath)

	totalJoltage1 := 0
	for _, bank := range allBanks {
		joltage := calculateJoltage(bank, 2)
		totalJoltage1 += joltage
		sse.PatchElementTempl(Part1Joltage(joltage, totalJoltage1))
		time.Sleep(4 * time.Millisecond)
	}

	totalJoltage2 := 0
	for _, bank := range allBanks {
		joltage := calculateJoltage(bank, 12)
		totalJoltage2 += joltage
		sse.PatchElementTempl(Part2Joltage(joltage, totalJoltage2))
		time.Sleep(4 * time.Millisecond)
	}
}
