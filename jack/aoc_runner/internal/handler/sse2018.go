package handler

import (
	"aoc-runner/internal/lib"
	"aoc-runner/views/y2018"
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
	inputPath, _ := filepath.Abs("internal/handler/input/2018-01-input.txt")
	inputArr := lib.SplitFileOnNewLines(inputPath)

	sse := datastar.NewSSE(w, r)

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
