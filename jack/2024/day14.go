package main

import (
	"fmt"
	u "github.com/jacktrusler/goutils"
	"regexp"
	"slices"
)

type Roboto struct {
	startX int
	startY int
	vX     int
	vY     int
}

func day14part1(robots [][]int) {
	// 101 wide, 103 tall
	w := 101
	h := 103
	time := 100

	robotGrid := MovedRobotoGrid(robots, time, w, h)

	robotTotal := 0
	q1 := 0
	q2 := 0
	q3 := 0
	q4 := 0
	for y, line := range robotGrid {
		for x, robot := range line {
			robotTotal += robot
			if y == h/2 || x == w/2 {
				continue
			}
			if y < h/2 && x < w/2 {
				q1 += robot
			}
			if y < h/2 && x > w/2 {
				q2 += robot
			}
			if y > h/2 && x < w/2 {
				q3 += robot
			}
			if y > h/2 && x > w/2 {
				q4 += robot
			}
		}
	}
	fmt.Println(q1 * q2 * q3 * q4)
	fmt.Println(robotTotal)
}

func MovedRobotoGrid(robots [][]int, seconds int, w int, h int) [][]int {
	robotGrid := u.MakeZeroGrid(w, h)

	// exampleGrid := u.MakeZeroGrid(11, 7)
	robotoStart := []Roboto{}
	for _, robot := range robots {
		x, y, vX, vY := robot[0], robot[1], robot[2], robot[3]
		robotoStart = append(robotoStart, Roboto{
			startX: x, startY: y, vX: vX, vY: vY,
		})
	}

	for _, rS := range robotoStart {
		endY := (rS.startY + (rS.vY * seconds)) % h
		endX := (rS.startX + (rS.vX * seconds)) % w

		eY := (h + endY) % h
		eX := (w + endX) % w
		robotGrid[eY][eX]++
	}

	return robotGrid

}

func splitTouchers(robots [][]int, start, end int, c chan int) {
	w := 101
	h := 103
	for i := start; i < end; i++ {
		time := i
		robotGrid := MovedRobotoGrid(robots, time, w, h)
		visited := make(map[u.Point]bool)
		for y, line := range robotGrid {
			for x, r := range line {
				if r > 0 {
					area := RTS(y, x, robotGrid, visited)
					if area > 100 {
						c <- i
					}
				}
			}
		}
	}
}

func day14part2(robots [][]int) {
	// Find a chrimbussa tree somehow
	c := make(chan int)
	w := 101
	h := 103
	gridArea := w * h
	for i := 0; i < gridArea; i += gridArea / 10 {
		go splitTouchers(robots, i, i+gridArea/10, c)
	}
	treeIndex := <-c

	// time := treeIndex
	// robotGrid := MovedRobotoGrid(robots, time, w, h)
	// for _, line := range robotGrid {
	// 	fmt.Println(line)
	// }

	fmt.Println(treeIndex)
}

// Robot Touchy Search
func RTS(y, x int, robots [][]int, visited map[u.Point]bool) int {
	start := u.Point{Y: y, X: x}
	rows := len(robots)    // y
	cols := len(robots[0]) // x
	queue := []u.Point{start}

	robotLoc := []u.Point{}
	robotLoc = append(robotLoc, start)

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

		for _, dir := range u.Dirs {
			newY, newX := current.Y+dir[0], current.X+dir[1]
			if newY >= 0 && newY < rows && newX >= 0 && newX < cols && robots[newY][newX] > 0 {
				if !visited[u.Point{Y: newY, X: newX}] {
					newPoint := u.Point{Y: newY, X: newX}
					queue = append(queue, u.Point{Y: newY, X: newX})
					if !slices.Contains(robotLoc, newPoint) {
						robotLoc = append(robotLoc, u.Point{Y: newY, X: newX})
					}
				}
			}
		}
	}
	return area
}

func Day14() {
	input := u.FileAsString("./input/2024-14-input.txt")
	re := regexp.MustCompile(`.*?(-?\d+),(-?\d+).*?(-?\d+),(-?\d+)`)
	matches := re.FindAllStringSubmatch(input, -1)

	robots := [][]int{}
	for _, match := range matches {
		arr := u.StringArrAtoI(match[1:])
		robots = append(robots, arr)
	}

	fmt.Println("----- Part 1 -----")
	// day14part1(robots)
	fmt.Println("----- Part 2 -----")
	day14part2(robots)
}
