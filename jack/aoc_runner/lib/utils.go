package lib

import (
	"fmt"
	"log"
	"math"
	"os"
	"strconv"
	"strings"
)

func MakeRuneArr(strArr []string) [][]rune {
	runeArr := make([][]rune, len(strArr))
	for i, s := range strArr {
		runeArr[i] = []rune(s)
	}
	return runeArr
}

// Marginally more efficient grid creation because go allocates every element to 0 value when using make()
func MakeZeroGrid(colsX, rowsY int) [][]int {
	grid := make([][]int, rowsY)
	for y := range rowsY {
		grid[y] = make([]int, colsX)
	}
	return grid
}

// In case you need 0's for some reason, -1 to indicate it hasn't been traveled to
func MakeNegativeOneGrid(colsX, rowsY int) [][]int {
	grid := make([][]int, rowsY)
	for y := range rowsY {
		grid[y] = make([]int, colsX)
		for x := range colsX {
			grid[y][x] = -1
		}
	}
	return grid
}

func FileAsString(file string) string {
	content, err := os.ReadFile(file)
	if err != nil {
		log.Fatal(err)
	}

	fileAsString := string(content)
	fileAsString = strings.TrimSuffix(fileAsString, "\n")
	return fileAsString
}

func SplitFileOnNewLines(file string) []string {
	content, err := os.ReadFile(file)
	if err != nil {
		log.Fatal(err)
	}

	fileAsString := string(content)
	fileAsString = strings.TrimSuffix(fileAsString, "\n")
	strArr := strings.Split(fileAsString, "\n")
	return strArr
}

func StringArrAtoI(arr []string) []int {
	arr2 := make([]int, 0)
	for _, i := range arr {
		j, err := strconv.Atoi(i)
		if err != nil {
			panic(err)
		}
		arr2 = append(arr2, j)
	}
	return arr2
}

func StringArrToFloat(arr []string) []float64 {
	arr2 := make([]float64, 0)
	for _, i := range arr {
		j, err := strconv.Atoi(i)
		if err != nil {
			panic(err)
		}
		arr2 = append(arr2, float64(j))
	}
	return arr2
}

func PrintMap(someMap map[any]any) {
	for k, v := range someMap {
		fmt.Printf("key: %+v | value: %+v\n", k, v)
	}
}

func IntArrayToLinkedList(arr []int) *Node {
	if len(arr) == 0 {
		return nil
	}

	head := &Node{Val: arr[0]}
	current := head

	for _, val := range arr[1:] {
		newNode := &Node{Val: val}
		current.Next = newNode
		current = newNode
	}

	return head
}

// Create a Set type
type Set[T comparable] map[T]bool

func NewSetFromArr[T comparable](items []T) {
	s := NewSet[T]()
	for _, item := range items {
		s.Add(item)
	}
	fmt.Println(s)
}

func NewSet[T comparable]() Set[T] {
	return make(Set[T])
}

func (s Set[T]) Add(e T) Set[T] {
	s[e] = true
	return s
}

func (s Set[T]) Remove(e T) Set[T] {
	delete(s, e)
	return s
}

func (s Set[T]) Has(e T) bool {
	_, ok := s[e]
	return ok
}

// Depth First Search
func DFS(x, y int, grid []string, visited map[Point]bool, target byte) *Point {
	// Row: Y | Col: X
	rows := len(grid)
	cols := len(grid[0])

	if y < 0 || y >= rows || x < 0 || x >= cols || visited[Point{Y: y, X: x}] {
		// oob or already visited!
		return nil
	}

	if grid[y][x] == target {
		return &Point{y, x}
	}

	visited[Point{Y: y, X: x}] = true

	for _, dir := range Dirs {
		newY, newX := y+dir[0], x+dir[1]
		if point := DFS(newY, newX, grid, visited, target); point != nil {
			return point
		}
	}

	return nil
}

// Breadth First Search
func BFS(x, y int, grid []string, target byte) *Point {
	visited := make(map[Point]bool)

	rows := len(grid)
	cols := len(grid[0])
	start := Point{y, x}

	queue := []Point{start}
	visited[start] = true

	for len(queue) > 0 {
		current := queue[0]
		queue = queue[1:]

		if grid[current.Y][current.X] == target {
			return &Point{current.Y, current.X}
		}

		for _, dir := range Dirs {
			newY, newX := current.Y+dir[0], current.X+dir[1]
			// Check in bounds and not visited
			if newY >= 0 && newY < rows && newX >= 0 && newX < cols && !visited[Point{newY, newX}] {
				queue = append(queue, Point{newY, newX})
				visited[Point{newY, newX}] = true
			}
		}
	}
	return nil
}

func IsWholeNumber(f float64) bool {
	tolerance := 1e-6
	return math.Abs(f-math.Round(f)) < tolerance
}

func ManhattanD(x1, x2, y1, y2 int) int {
	return Abs(x1-x2) + Abs(y1-y2)
}

func EuclideanD(x1, x2, y1, y2 int) float64 {
	xSquared := (x1 - x2) * (x1 - x2)
	ySquared := (y1 - y2) * (y1 - y2)
	return math.Sqrt(float64(Abs(xSquared) + Abs(ySquared)))
}

// Helper function to calculate absolute value of an integer using ints
// the math.Abs library only uses float64
func Abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}
