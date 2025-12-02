package y2018

import (
	"aoc/lib"
	"log/slog"
	"math"
	"net/http"
	"path/filepath"
	"strconv"
	"time"

	"github.com/starfederation/datastar-go/datastar"
)

func D1P1SSE(w http.ResponseWriter, r *http.Request) {
	sse := datastar.NewSSE(w, r)

	inputPath, _ := filepath.Abs("web/y2018/input/2018-01-input.txt")
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
			slog.Error("Error str to num day1")
			return
		}
		if op == '+' {
			sum += num
		} else {
			sum -= num
		}

		patch := OutputBox(runningStr)

		if err := sse.PatchElementTempl(patch); err != nil {
			slog.Error("Error patching day1 part1")
			return
		}
		time.Sleep(4 * time.Millisecond)
	}
	patch := OutputBox(strconv.Itoa(sum))
	if err := sse.PatchElementTempl(patch); err != nil {
		slog.Error("Error patching day1 part1")
		return
	}
}

// Day 1 | Part 2
func D1P2SSE(w http.ResponseWriter, r *http.Request) {
	sse := datastar.NewSSE(w, r)

	inputPath, _ := filepath.Abs("web/y2018/input/2018-01-input.txt")
	inputArr := lib.SplitFileOnNewLines(inputPath)

	sum := 0
	// Part 2 - Create a cache
	cache := make(map[int]bool)
	finalVal := math.MaxInt

	//Output cache miss once, don't keep track of running string
	//Traversing this array a bunch is expensive
	patch := OutputBox("Cache Miss")

	if err := sse.PatchElementTempl(patch); err != nil {
		slog.Error("Error patching day1 part2")
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
				slog.Error("Error str to num day1")
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
	patch2 := OutputBox(strconv.Itoa(sum))

	if err := sse.PatchElementTempl(patch2); err != nil {
		slog.Error("Error patching day1 part2")
		return
	}
}
