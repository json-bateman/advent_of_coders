package main

import (
	"fmt"
	u "github.com/jacktrusler/goutils"
	"strings"
)

var (
	// This is like pokemon ice cave low key
	ice      []string
	startPos u.Point
	rockSet  u.Set[u.Point]
)

// Imaginary Rock Cache
type IRC struct {
	Y   int
	X   int
	Dir int
}

// A new genre of music in the AoC-verse
func guardStep(y, x int, dirI int) {
	// out of bounds, therefore finished
	dy := u.Dirs[dirI][0]
	dx := u.Dirs[dirI][1]
	if x == -1 || y == -1 || y >= len(ice) || x >= len(ice[y]) {
		return
	}
	if []rune(ice[y])[x] == '#' {
		turn := (dirI + 1) % 4
		switch dirI {
		// back up a step and change dir
		case 0:
			guardStep(y+1, x, turn)
		case 1:
			guardStep(y, x-1, turn)
		case 2:
			guardStep(y-1, x, turn)
		case 3:
			guardStep(y, x+1, turn)
		}
		return
	} else {
		// Spot has been moved to
		rockSet.Add(u.Point{Y: y, X: x})
	}
	guardStep(y+dy, x+dx, dirI)
	return
}

func printIce() {
	// Print it out for fun :>
	for _, line := range ice {
		fmt.Println(line)
	}
}

func day6part1() {
	rockSet = u.NewSet[u.Point]()
	for y, line := range ice {
		for x, rune := range line {
			if rune == '^' {
				startPos = u.Point{Y: y, X: x}
				guardStep(y, x, 0)
			}
		}
	}

	fmt.Println(len(rockSet))
}

func isTrapped(y, x, dirI int, iRC u.Point) bool {
	cache := make(map[string]bool)
	for {
		dy := u.Dirs[dirI][0]
		dx := u.Dirs[dirI][1]
		if x+dx < 0 || y+dy < 0 || y+dy >= len(ice) || x+dx >= len(ice[y]) {
			// Freedom
			return false
		}

		hitRock := y+dy == iRC.Y && x+dx == iRC.X
		if ice[y+dy][x+dx] == '#' || hitRock {
			str := fmt.Sprintf("y%dx%ddir%d", y, x, dirI)
			if cache[str] {
				// clapped in the baguss
				return true
			}
			cache[str] = true
			dirI = (dirI + 1) % 4
			continue
		}
		y += dy
		x += dx
	}
}

func day6part2() {
	loops := 0

	startY := startPos.Y
	startX := startPos.X
	// Remove starting position from set... Guard will notice!
	rockSet.Remove(u.Point{Y: startY, X: startX})
	for k := range rockSet {
		imagRock := u.Point{Y: k.Y, X: k.X}

		inTheClapHouse := isTrapped(startY, startX, 0, imagRock)
		if inTheClapHouse {
			loops++
		}
	}
	fmt.Println(loops)
}

func Day6() {
	input := u.FileAsString("./input/2024-06-input.txt")
	ice = strings.Split(input, "\n")
	fmt.Println("----- Part 1 -----")
	day6part1()
	fmt.Println("----- Part 2 -----")
	day6part2()
}
