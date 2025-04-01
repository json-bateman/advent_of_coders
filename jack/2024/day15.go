package main

import (
	"fmt"
	u "github.com/jacktrusler/goutils"
	"strings"
)

// maps N,E,S,W with any 4 runes
// --- i.e. ---
// { '^':{0,1}, '>':{1,0}, 'v':{0,-1}, '<':{-1,0} }
func dirMap(runes [4]rune) map[rune][2]int {
	dM := make(map[rune][2]int)
	dM[runes[0]] = [2]int{0, -1} // N
	dM[runes[1]] = [2]int{1, 0}  // E
	dM[runes[2]] = [2]int{0, 1}  // S
	dM[runes[3]] = [2]int{-1, 0} // W

	return dM
}

func printWarehouse(w [][]byte) {
	for _, line := range w {
		fmt.Println(string(line))
	}
}

func day15part1(warehouse [][]byte, instructions string) {

	dM := dirMap([4]rune{'^', '>', 'v', '<'})

	for y, line := range warehouse {
		for x, r := range line {
			if r == '@' {
				// start location
				curr := u.Point{X: x, Y: y}
				for _, ins := range instructions {
					// warehouse[curr.Y][curr.X] = '@'
					// printWarehouse(warehouse)
					warehouse[curr.Y][curr.X] = '.'
					var dx, dy int
					dx = dM[ins][0]
					dy = dM[ins][1]

					var next u.Point
					next.X, next.Y = curr.X+dx, curr.Y+dy

					if warehouse[next.Y][next.X] == '#' {
						continue
					} else if warehouse[next.Y][next.X] == 'O' {
						// collect stones
						stoneCache := []u.Point{{Y: next.Y, X: next.X}}
						var next2 u.Point
						next2.X, next2.Y = next.X+dx, next.Y+dy
						for {

							if warehouse[next2.Y][next2.X] == 'O' {
								stoneCache = append(stoneCache, u.Point{Y: next2.Y, X: next2.X})
								next2.X, next2.Y = next2.X+dx, next2.Y+dy
								continue
							}
							if warehouse[next2.Y][next2.X] == '.' {
								//move stones to empty space
								warehouse[stoneCache[0].Y][stoneCache[0].X] = '.'
								for _, stone := range stoneCache {
									warehouse[stone.Y+dy][stone.X+dx] = 'O'
									next2.X, next2.Y = next2.X+dx, next2.Y+dy
								}
								curr.X = next.X
								curr.Y = next.Y
								break
							}
							if warehouse[next2.Y][next2.X] == '#' {
								break
							}
						}
					} else {
						curr.X = next.X
						curr.Y = next.Y
						continue
					}
				}
			}
		}
	}

	total := 0
	for y, line := range warehouse {
		for x, b := range line {
			if b == 'O' {
				total += y*100 + x
			}
		}
	}
	fmt.Println(total)

}

func day15part2(warehouse [][]byte, dirs string) {

}

func Day15() {
	input := u.FileAsString("./input/2024-15-input.txt")
	both := strings.Split(input, "\n\n")
	warehouse := strings.Split(both[0], "\n")
	var w [][]byte
	for _, line := range warehouse {
		w = append(w, []byte(line))
	}
	dirs := strings.ReplaceAll(both[1], "\n", "")

	fmt.Println("----- Part 1 -----")
	day15part1(w, dirs)
	fmt.Println("----- Part 2 -----")
	day15part2(w, dirs)
}
