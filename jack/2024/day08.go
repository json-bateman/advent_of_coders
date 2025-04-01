package main

import (
	"fmt"
	"github.com/jacktrusler/goutils"
	u "github.com/jacktrusler/goutils"
	"math"
	"strings"
)

var (
	city []string
	// antennae mapping
	aM = make(AntennaMap)
	// part 2 antenna are also antinodes ig
	antiNodes2 = make(map[u.Point]bool)
)

type AntennaMap map[rune][]u.Point

func day8part1() {
	for y, line := range city {
		for x, rune := range line {
			// some antennae found
			if rune != '.' {
				aM[rune] = append(aM[rune], u.Point{Y: y, X: x})
				antiNodes2[u.Point{Y: y, X: x}] = true
			}
		}
	}

	antiNodeMap := make(map[u.Point]bool)
	for _, v := range aM {
		for i := 0; i < len(v); i++ {
			for j := i + 1; j < len(v); j++ {
				var a1, a2 u.Point
				a1.X, a1.Y = v[i].X, v[i].Y
				a2.X, a2.Y = v[j].X, v[j].Y
				dy := int(math.Abs(float64(a1.Y - a2.Y)))
				dx := int(math.Abs(float64(a1.X - a2.X)))

				if a1.X < a2.X {
					dx = dx * -1
				}
				if a1.Y < a2.Y {
					dy = dy * -1
				}
				createAntiNodes(a1, dy, dx, antiNodeMap, 1)
				createAntiNodes(a2, -dy, -dx, antiNodeMap, 1)
			}
		}
	}

	fmt.Println(len(antiNodeMap))

}

func createAntiNodes(a u.Point, dy, dx int, aNM map[u.Point]bool, part int) {
	anti := u.Point{}
	anti.X = a.X + dx
	anti.Y = a.Y + dy
	// exit condition
	if anti.X < 0 || anti.Y < 0 || anti.X >= len(city[0]) || anti.Y >= len(city) {
		return
	}
	aNM[anti] = true
	if part == 2 {
		createAntiNodes(anti, dy, dx, aNM, 2)
	}
}

func day8part2() {
	for _, v := range aM {
		for i := 0; i < len(v); i++ {
			for j := i + 1; j < len(v); j++ {
				var a1, a2 u.Point
				a1.X, a1.Y = v[i].X, v[i].Y
				a2.X, a2.Y = v[j].X, v[j].Y
				dy := int(math.Abs(float64(a1.Y - a2.Y)))
				dx := int(math.Abs(float64(a1.X - a2.X)))

				if a1.Y < a2.Y {
					dy = dy * -1
				}
				if a1.X < a2.X {
					dx = dx * -1
				}
				createAntiNodes(a1, dy, dx, antiNodes2, 2)
				createAntiNodes(a2, -dy, -dx, antiNodes2, 2)
			}
		}
	}
	fmt.Println(len(antiNodes2))
}

func Day8() {
	input := goutils.FileAsString("./input/2024-08-input.txt")
	city = strings.Split(input, "\n")
	fmt.Println("----- Part 1 -----")
	day8part1()
	fmt.Println("----- Part 2 -----")
	day8part2()
}
