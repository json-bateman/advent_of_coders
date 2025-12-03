package y2025

import (
	"aoc/lib"
	"log/slog"
	"net/http"
	"path/filepath"
	"strconv"
	"strings"
	"time"

	"github.com/dlclark/regexp2"
	"github.com/starfederation/datastar-go/datastar"
)

func D2SSE(w http.ResponseWriter, r *http.Request) {
	sse := datastar.NewSSE(w, r)

	inputPath, _ := filepath.Abs("web/y2025/input/2025-02-input.txt")
	input := lib.FileAsString(inputPath)
	inputArr := strings.SplitSeq(input, ",")
	finalSum1 := 0
	finalSum2 := 0
	for idRange := range inputArr {
		rangeArr := strings.Split(idRange, "-")
		startStr := rangeArr[0]
		endStr := rangeArr[1]

		start, _ := strconv.Atoi(startStr)
		end, _ := strconv.Atoi(endStr)

		part1 := regexp2.MustCompile(`\b(\d+)\1\b`, 0)
		part2 := regexp2.MustCompile(`\b(\d+)\1+\b`, 0)

		for start <= end {
			matches1, _ := part1.FindStringMatch(strconv.Itoa(start))
			matches2, _ := part2.FindStringMatch(strconv.Itoa(start))
			if matches1 != nil {
				patch := MatchBox(strconv.Itoa(start), finalSum1)
				if err := sse.PatchElementTempl(patch); err != nil {
					slog.Error("Error patching day2 part1 2025")
					return
				}
				finalSum1 += start
			}
			if matches2 != nil {
				patch := MatchBox2(strconv.Itoa(start), finalSum2)
				if err := sse.PatchElementTempl(patch); err != nil {
					slog.Error("Error patching day2 part2 2025")
					return
				}
				finalSum2 += start
				time.Sleep(4 * time.Millisecond)
			}
			start++
		}

	}
	patch := MatchBox("Complete!", finalSum1)
	if err := sse.PatchElementTempl(patch); err != nil {
		slog.Error("Error patching day2 part1 2025")
		return
	}
	patch2 := MatchBox2("Complete!", finalSum2)
	if err := sse.PatchElementTempl(patch2); err != nil {
		slog.Error("Error patching day2 part2 2025")
		return
	}
}
