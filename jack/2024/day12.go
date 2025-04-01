package main

import (
	"fmt"
	u "github.com/jacktrusler/goutils"
	"slices"
	"strings"
)

var (
	garden   []string
	plantSet u.Set[byte]
	plots    = make(map[byte][][]u.Point)
)

func counter(y, x int, garden []string, visited map[u.Point]bool, target byte) int {
	start := u.Point{Y: y, X: x}
	rows := len(garden)    // y
	cols := len(garden[0]) // x
	queue := []u.Point{start}

	plantLoc := []u.Point{}
	plantLoc = append(plantLoc, start)

	perimeter := 0
	area := 0

	for len(queue) > 0 {
		current := queue[0]
		queue = queue[1:]
		curr := u.Point{Y: current.Y, X: current.X}
		if visited[curr] {
			continue
		} else {
			visited[curr] = true
		}
		area++

		// find perimeter of current plant
		// and move to surrounding plants
		surroundingPlants := 0
		for _, dir := range u.Dirs {
			newY, newX := current.Y+dir[0], current.X+dir[1]
			if newY >= 0 && newY < rows && newX >= 0 && newX < cols && garden[newY][newX] == target {
				surroundingPlants++
				if !visited[u.Point{Y: newY, X: newX}] {
					newPoint := u.Point{Y: newY, X: newX}
					queue = append(queue, u.Point{Y: newY, X: newX})
					if !slices.Contains(plantLoc, newPoint) {
						plantLoc = append(plantLoc, u.Point{Y: newY, X: newX})
					}
				}
			}
		}
		perimeter += 4 - surroundingPlants
	}
	plots[target] = append(plots[target], plantLoc)
	fmt.Printf("%+v\n\n", plantLoc)
	return perimeter * area
}

func makePlantSet(garden []string) u.Set[byte] {
	plantSet = u.NewSet[byte]()
	for _, line := range garden {
		for _, plant := range line {
			plantSet.Add(byte(plant))
		}
	}
	return plantSet
}

func day12part1() {
	plantSet := makePlantSet(garden)

	visited := make(map[u.Point]bool)

	total := 0
	for plant := range plantSet {
		for y, line := range garden {
			for x, p := range line {
				if byte(p) == plant && !visited[u.Point{Y: y, X: x}] {
					ans := counter(y, x, garden, visited, byte(p))
					total += ans
				}
			}
		}
	}
	fmt.Printf("%+v\n", plots)
}

func calcArea(arr []u.Point, p byte) int {
	rows := len(garden)    // y
	cols := len(garden[0]) // x
	dirMap := make(map[string]u.Set[u.Point])

	if dirMap["North"] == nil {
		dirMap["North"] = make(u.Set[u.Point])
	}
	if dirMap["East"] == nil {
		dirMap["East"] = make(u.Set[u.Point])
	}
	if dirMap["South"] == nil {
		dirMap["South"] = make(u.Set[u.Point])
	}
	if dirMap["West"] == nil {
		dirMap["West"] = make(u.Set[u.Point])
	}

	for _, Point := range arr {
		for i, dir := range u.Dirs {
			newY, newX := Point.Y+dir[0], Point.X+dir[1]
			if newY < 0 || newY >= rows || newX < 0 || newX >= cols || garden[newY][newX] != p {
				// N
				if i == 0 && (newY < 0 || garden[newY][Point.X] != p) {
					dirMap["North"][Point] = true
				}
				if i == 1 && (newX >= cols || garden[Point.Y][newX] != p) {
					dirMap["East"][Point] = true
				}
				if i == 2 && (newY >= rows || garden[newY][Point.X] != p) {
					dirMap["South"][Point] = true
				}
				if i == 3 && (newX < 0 || garden[Point.Y][newX] != p) {
					dirMap["West"][Point] = true
				}
			}
		}
	}

	sides := 0
	for dir, v := range dirMap {
		// the band
		oneDirection := make([]u.Point, 0)
		for k2 := range v {
			oneDirection = append(oneDirection, k2)
		}

		if dir == "North" || dir == "South" {
			// North / South Sort
			slices.SortFunc(oneDirection, func(a, b u.Point) int {
				if a.Y == b.Y {
					return a.X - b.X
				}
				return a.Y - b.Y
			})

			y := -1
			prevX := -1
			for _, Point := range oneDirection {
				if prevX == -1 {
					prevX = Point.X
				}
				if Point.Y != y {
					y = Point.Y
					sides++
					prevX = Point.X
					continue
				}

				if Point.X-prevX > 1 {
					sides++
				}
				prevX = Point.X

			}
		} else {
			// East / West Sort
			slices.SortFunc(oneDirection, func(a, b u.Point) int {
				if a.X == b.X {
					return a.Y - b.Y
				}
				return a.X - b.X
			})

			x := -1
			prevY := -1
			for _, Point := range oneDirection {
				if prevY == -1 {
					prevY = Point.Y
				}
				if Point.X != x {
					x = Point.X
					sides++
					prevY = Point.Y
					continue
				}

				if Point.Y-prevY > 1 {
					sides++
				}
				prevY = Point.Y

			}
		}
	}
	// len(arr) is total squares a.k.a. area
	return sides * len(arr)
}

// find edges
func day12part2() {
	plantSet := makePlantSet(garden)
	total2 := 0
	for plant := range plantSet {
		// put perimeters into map
		// if out of bounds or doesn't match target... must be a side
		for _, plot := range plots[plant] {
			total2 += calcArea(plot, plant)
		}
	}
	fmt.Println(total2)
}

func Day12() {
	input := u.FileAsString("./input/2024-12-input.txt")
	garden = strings.Split(input, "\n")
	fmt.Println("----- Part 1 -----")
	day12part1()
	fmt.Println("----- Part 2 -----")
	day12part2()
}
