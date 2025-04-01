package main

import (
	"fmt"
	"log"
	"math"
	"os"
	"strings"
)

type coords struct {
	x int
	y int
	steps int
}

func (c coords) Equals(c2 coords) bool {
	return c.x == c2.x && c.y == c2.y
}

func (c coords) Dist() float64 {
	return math.Abs(float64(c.x)) + math.Abs(float64(c.y))
}

func Day3() {
	content, err := os.ReadFile("day3.txt")
	if err != nil {
		log.Fatal(err)
	}

	fileAsString := string(content)
	var lines []string = strings.Split(fileAsString, "\n")

	var firstWire []string = strings.Split(lines[0], ",")
	var secondWire []string = strings.Split(lines[1], ",")

	wire1 := make([]coords, 0)
	x, y, dx, dy, steps := 0, 0, 0, 0, 0
	for _, move := range firstWire {
		dir, dist := byte(0), 0
		fmt.Sscanf(move, "%c%d", &dir, &dist)
		switch dir {
		case 'U':
			dx, dy = 0, -1	
		case 'R':
			dx, dy = 1, 0	
		case 'D':
			dx, dy = 0, 1	
		case 'L':
			dx, dy = -1, 0	
		}
		
		for i := dist; i > 0; i-- {
			x, y = x + dx, y + dy
			steps++
			wire1 = append(wire1, coords{x, y, steps})
		}
	}

	wire2 := make([]coords, 0)
	x, y, dx, dy, steps = 0, 0, 0, 0, 0
	for _, move := range secondWire {
		dir, dist := byte(0), 0
		fmt.Sscanf(move, "%c%d", &dir, &dist)
		switch dir {
		case 'U':
			dx, dy = 0, -1	
		case 'R':
			dx, dy = 1, 0	
		case 'D':
			dx, dy = 0, 1	
		case 'L':
			dx, dy = -1, 0	
		}
		
		for i := dist; i > 0; i-- {
			x, y = x + dx, y + dy
			steps++
			wire2 = append(wire2, coords{x, y, steps})
		}
	}
	
	min := math.MaxFloat64
	for _, w1 := range wire1 {
		for _, w2 := range wire2 {
			if w1.Equals(w2) {
				min = math.Min(min, w1.Dist())
			}
		}
	}
	minSteps := math.MaxFloat64
	for _, w1 := range wire1 {
		for _, w2 := range wire2 {
			if w1.Equals(w2) {
				minSteps = math.Min(minSteps, float64(w1.steps + w2.steps))
			}
		}
	}

	fmt.Println("part 1: ", int(min))
	fmt.Println("part 2: ", int(minSteps))
}
