package lib

type Direction int

type RuneGrid [][]int

const (
	North Direction = iota
	East
	South
	West
)

var (
	// N, E, S, W  ---  { X, Y }
	Dirs = [][]int{{0, -1}, {1, 0}, {0, 1}, {-1, 0}}
	// N, NE, E, SE, S, SW, W, NW  ---  { X, Y }
	Surround = [][]int{{0, -1}, {1, -1}, {1, 0}, {1, 1}, {0, 1}, {-1, 1}, {-1, 0}, {-1, -1}}
)

type Node struct {
	Val  int
	Next *Node
}

type Point struct {
	X int
	Y int
}
