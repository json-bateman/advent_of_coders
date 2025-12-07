package y2025

import (
	"aoc/lib"
	"fmt"
	"net/http"
	"path/filepath"
	"slices"
	"strconv"
	"strings"

	"github.com/starfederation/datastar-go/datastar"
)

func D5SSE(w http.ResponseWriter, r *http.Request) {
	sse := datastar.NewSSE(w, r)
	inputPath, err := filepath.Abs("web/y2025/input/2025-05-test-input.txt")
	if err != nil {
		fmt.Println("Error finding filepath")
		return
	}
	inputStr := lib.FileAsString(inputPath)
	parts := strings.Split(inputStr, "\n\n")
	idRanges := strings.Split(parts[0], "\n")
	ids := strings.Split(parts[1], "\n")

	var arr [][2]int
	for _, idRange := range idRanges {
		ranges := strings.Split(idRange, "-")
		start, _ := strconv.Atoi(ranges[0])
		end, _ := strconv.Atoi(ranges[1])
		arr = append(arr, [2]int{start, end})
	}

	slices.SortFunc(arr, func(a, b [2]int) int {
		return a[0] - b[0]
	})

	// Merge intervals
	merged := [][2]int{arr[0]}
	for i := 1; i < len(arr); i++ {
		current := arr[i]
		last := &merged[len(merged)-1]

		if current[0] <= last[1] {
			if current[1] > last[1] {
				last[1] = current[1]
			}
		} else {
			merged = append(merged, current)
		}
	}

	part1 := 0
	for _, idStr := range ids {
		id, _ := strconv.Atoi(idStr)
		for _, r := range merged {
			if r[0] <= id && id <= r[1] {
				part1++
				break
			}
		}
	}
	// Throw out ingredient ids, just add up ranges
	part2 := 0
	for _, r := range merged {
		part2 += r[1] - r[0] + 1
	}
	sse.PatchElementTempl(BothPartsD5(part1, part2))
}
