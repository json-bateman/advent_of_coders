package handler

import (
	"aoc-runner/internal/lib"
	"aoc-runner/views/y2018"
	"fmt"
	"math"
	"net/http"
	"path/filepath"
	"strconv"
	"time"

	"github.com/starfederation/datastar-go/datastar"
)

// Day 1 | Part 1
func (h *Handler) D1P1SSE(w http.ResponseWriter, r *http.Request) {
	sse := datastar.NewSSE(w, r)

	inputPath, _ := filepath.Abs("internal/handler/input/2018-01-input.txt")
	inputArr := lib.SplitFileOnNewLines(inputPath)

	sum := 0
	runningStr := ""
	for _, ins := range inputArr {
		// first char is a operation
		op := ins[0]
		// rest is a number
		numstr := ins[1:]
		runningStr += ins + " "

		num, err := strconv.Atoi(numstr)
		if err != nil {
			h.l.Error("Error str to num day1")
			return
		}
		if op == '+' {
			sum += num
		} else {
			sum -= num
		}

		patch := y2018.RowAndOutput("y2018d1p1", sum, runningStr)

		if err := sse.PatchElementTempl(patch); err != nil {
			h.l.Error("Error patching day1 part1")
			return
		}
		time.Sleep(4 * time.Millisecond)
	}
}

// Day 1 | Part 2
func (h *Handler) D1P2SSE(w http.ResponseWriter, r *http.Request) {
	sse := datastar.NewSSE(w, r)

	inputPath, _ := filepath.Abs("internal/handler/input/2018-01-input.txt")
	inputArr := lib.SplitFileOnNewLines(inputPath)

	sum := 0
	// Part 2 - Create a cache
	cache := make(map[int]bool)
	finalVal := math.MaxInt

	//Output cache miss once, don't keep track of running string
	//Traversing this array a bunch is expensive
	patch := y2018.RowAndOutput("y2018d1p2", 0, "Cache Miss")

	if err := sse.PatchElementTempl(patch); err != nil {
		h.l.Error("Error patching day1 part2")
		return
	}

	// While running, there is no answer, because the answer is a cache hit
	for finalVal == math.MaxInt {
		for i := range inputArr {
			ins := inputArr[i]
			// first char is a operation
			op := ins[0]
			// rest is a number
			numstr := ins[1:]

			num, err := strconv.Atoi(numstr)
			if err != nil {
				h.l.Error("Error str to num day1")
				return
			}
			if op == '+' {
				sum += num
			} else {
				sum -= num
			}
			_, exists := cache[sum]
			if exists {
				finalVal = sum
				break
			} else {
				cache[sum] = true
			}

			// restart from the beginning of the array
			if i == len(inputArr)-1 {
				i = 0
			}
		}
	}
	// Final Answer, cache hit
	patch2 := y2018.RowAndOutput("y2018d1p2", sum, "Cache Hit")

	if err := sse.PatchElementTempl(patch2); err != nil {
		h.l.Error("Error patching day1 part2")
		return
	}
}

// Day 2 | Part 1
func (h *Handler) D2P1SSE(w http.ResponseWriter, r *http.Request) {
	sse := datastar.NewSSE(w, r)

	inputPath, _ := filepath.Abs("internal/handler/input/2018-02-input.txt")
	inputArr := lib.SplitFileOnNewLines(inputPath)
	twos, threes := 0, 0
	for _, line := range inputArr {
		letterCount := make(map[rune]int)
		for _, r := range line {
			letterCount[r]++
		}

		hasTwo, hasThree := false, false
		twor, threer := ' ', ' '
		for r, count := range letterCount {
			if count == 2 {
				twor = r
				hasTwo = true
			}
			if count == 3 {
				threer = r
				hasThree = true
			}
		}

		build2str := ""
		build3str := ""
		if hasTwo {
			for _, r := range line {
				if r == twor {
					build2str += fmt.Sprintf(`<span style='color: green'>%s</span>`, string(r))
				} else {
					build2str += string(r)
				}
			}
			twos++

			//Patch SSE for the twos
			sse.PatchElementTempl(y2018.Day02Part1Output(twos, threes))
			if build2str != "" {
				sse.PatchElementf(`<div id="currentLine">%s</div>`, build2str)
			}
			time.Sleep(50 * time.Millisecond)
		}
		if hasThree {
			for _, r := range line {
				if r == threer {
					build3str += fmt.Sprintf(`<span style='color: orange'>%s</span>`, string(threer))
				} else {
					build3str += string(r)
				}
			}
			threes++

			//Patch SSE for the threes
			sse.PatchElementTempl(y2018.Day02Part1Output(twos, threes))
			if build3str != "" {
				sse.PatchElementf(`<div id="currentLine">%s</div>`, build3str)
			}

			time.Sleep(50 * time.Millisecond)
		}
	}

	sse.PatchElementf(`<td id="y2018d2p1">%d</td>`, twos*threes)
}

// Day 2 | Part 2
func (h *Handler) D2P2SSE(w http.ResponseWriter, r *http.Request) {
	sse := datastar.NewSSE(w, r)

	inputPath, _ := filepath.Abs("internal/handler/input/2018-02-input.txt")
	inputArr := lib.SplitFileOnNewLines(inputPath)

	// Stop before the last item, to compare to inputArr[j]
	missedIndex := 0
	line1 := ""
	line2 := ""
	for i := 0; i < len(inputArr)-1; i++ {
		misses := 0
		for j := i + 1; j < len(inputArr); j++ {
			// Compare First and Second Line
			line1 = inputArr[i]
			line2 = inputArr[j]
			sse.PatchElementTempl(y2018.Day02Part2Output(line1, line2, 100, ""))
			for k, r := range line1 {
				if byte(r) != line2[k] {
					misses++
					missedIndex = k
				}
			}
			if misses == 1 {
				answerLine := line1[:missedIndex] + line1[missedIndex+1:]
				sse.PatchElementTempl(y2018.Day02Part2Output(line1, line2, missedIndex, answerLine))
				sse.PatchElementf(`<td id="y2018d2p2">%s</td>`, "Found!")
				return
			}
			misses = 0
			time.Sleep(5 * time.Millisecond)
		}
	}
}
