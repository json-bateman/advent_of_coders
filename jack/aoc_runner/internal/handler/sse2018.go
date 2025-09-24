package handler

import (
	"aoc-runner/internal/lib"
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
	inputPath, _ := filepath.Abs("internal/handler/input/2018-01-input.txt")
	inputArr := lib.SplitFileOnNewLines(inputPath)

	sse := datastar.NewSSE(w, r)

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

		if err := sse.PatchElements(
			fmt.Sprintf(`<th id="y2018d1p1">%d</th>`, sum) +
				fmt.Sprintf(`<div class="output" id="y2018running">%s</div>`, runningStr),
		); err != nil {
			h.l.Error("Error patching day1 part1")
			return
		}
		time.Sleep(4 * time.Millisecond)
	}
}

// Day 1 | Part 2
func (h *Handler) D1P2SSE(w http.ResponseWriter, r *http.Request) {
	inputPath, _ := filepath.Abs("internal/handler/input/2018-01-input.txt")
	inputArr := lib.SplitFileOnNewLines(inputPath)

	sse := datastar.NewSSE(w, r)

	sum := 0
	// Part 2 - Create a cache
	cache := make(map[int]bool)
	finalVal := math.MaxInt

	//Output cache miss once, don't keep track of running string
	//Traversing this array a bunch is expensive
	if err := sse.PatchElements(
		`<div class="output" id="y2018running">Cache Miss</div>`,
	); err != nil {
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
	if err := sse.PatchElements(
		fmt.Sprintf(`<th id="y2018d1p2">%d</th>`, finalVal) +
			`<div class="output" id="y2018running">Cache Hit!</div>`,
	); err != nil {
		h.l.Error("Error patching day1 part2")
		return
	}
}
