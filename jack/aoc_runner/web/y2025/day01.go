package y2025

import (
	"aoc/lib"
	"fmt"
	"log/slog"
	"net/http"
	"path/filepath"
	"strconv"
	"time"

	"github.com/starfederation/datastar-go/datastar"
)

func D1P1SSE(w http.ResponseWriter, r *http.Request) {
	sse := datastar.NewSSE(w, r)

	inputPath, _ := filepath.Abs("web/y2025/input/2025-01-input.txt")
	inputArr := lib.SplitFileOnNewLines(inputPath)

	pos := 50
	actualPassword := 0
	for _, ins := range inputArr {
		dir, turn, err := parseTurn(ins)
		if err != nil {
			slog.Error("Error Parsing Turn", "Year", "2025", "Day", "1", "Part", "1")
			return
		}
		switch dir {
		case 'R':
			pos = (pos + turn) % 100
		case 'L':
			pos = (pos - turn) % 100
			if pos < 0 {
				pos += 100
			}
		}
		if pos == 0 {
			actualPassword++
		}
		dial := Dial(pos)
		tick := DialTickMark(pos)
		password := PasswordBox(actualPassword)

		if err := sse.PatchElementTempl(dial); err != nil {
			slog.Error("Error patching day1 part1 2025")
			return
		}
		if err := sse.PatchElementTempl(tick); err != nil {
			slog.Error("Error patching day1 part1 2025")
			return
		}
		if err := sse.PatchElementTempl(password); err != nil {
			slog.Error("Error patching day1 part1 2025")
			return
		}
		if pos == 0 {
			time.Sleep(10 * time.Millisecond)
		}
		time.Sleep(1 * time.Millisecond)
	}

}

func D1P2SSE(w http.ResponseWriter, r *http.Request) {
	sse := datastar.NewSSE(w, r)

	inputPath, _ := filepath.Abs("web/y2025/input/2025-01-input.txt")
	inputArr := lib.SplitFileOnNewLines(inputPath)

	pos := 50
	actualPassword := 0
	for _, ins := range inputArr {
		dir, turn, err := parseTurn(ins)
		if err != nil {
			slog.Error("Error Parsing Turn", "Year", "2025", "Day", "1", "Part", "2")
			return
		}
		// For rotations > 100
		actualPassword += turn / 100
		turn %= 100
		switch dir {
		case 'R':
			if pos+turn >= 100 {
				actualPassword++
			}
			pos = (pos + turn) % 100
		case 'L':
			if pos != 0 && pos-turn <= 0 {
				actualPassword++
			}

			pos = (pos - turn) % 100

			if pos < 0 {
				pos += 100
			}
		}
		dial := Dial(pos)
		tick := DialTickMark(pos)
		password := PasswordBox(actualPassword)

		if err := sse.PatchElementTempl(dial); err != nil {
			slog.Error("Error patching day1 part2 2025")
			return
		}
		if err := sse.PatchElementTempl(tick); err != nil {
			slog.Error("Error patching day1 part2 2025")
			return
		}
		if err := sse.PatchElementTempl(password); err != nil {
			slog.Error("Error patching day1 part2 2025")
			return
		}
		time.Sleep(1 * time.Millisecond)
	}
}

func parseTurn(ins string) (direction byte, amount int, err error) {
	direction = ins[0]
	amount, err = strconv.Atoi(ins[1:])
	if err != nil {
		return 0, 0, fmt.Errorf("invalid number in instruction %s: %w", ins, err)
	}

	return direction, amount, nil
}
