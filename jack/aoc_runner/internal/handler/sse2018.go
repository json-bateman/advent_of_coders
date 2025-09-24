package handler

import (
	"aoc-runner/internal/lib"
	"aoc-runner/views/y2018"
	"net/http"
	"path/filepath"
	"strconv"
	"time"

	"github.com/starfederation/datastar-go/datastar"
)

func (h *Handler) D01SSE(w http.ResponseWriter, r *http.Request) {
	inputPath, _ := filepath.Abs("internal/handler/input/2018-01-input.txt")
	inputArr := lib.SplitFileOnNewLines(inputPath)

	sse := datastar.NewSSE(w, r)

	// Solve the Problem
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

		day1 := y2018.Day01(sum, runningStr)
		if err := sse.PatchElementTempl(day1); err != nil {
			h.l.Error("Error patching day1")
			return
		}
		time.Sleep(10 * time.Millisecond)
	}

}
