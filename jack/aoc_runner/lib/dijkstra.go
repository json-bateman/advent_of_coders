package lib

import (
	"container/heap"
	"math"
)

// Algorithm to find the shortest path
// Slower than one that uses a heuristic like A* for example
func Dijkstra(start, end Point, grid []string) []Point {
	// initialize a priority queue (heap)
	// the heap will automatically rearrange based on new priorities entering
	distances := make(map[Point]int)
	parent := make(map[Point]Point)
	// set all points in the grid to a distance of infinity
	for y, line := range grid {
		for x := range line {
			p := Point{x, y}
			distances[p] = math.MaxInt32
		}
	}
	distances[start] = 0

	pq := make(PQ, 0)
	heap.Init(&pq)
	heap.Push(&pq, &PQItem{value: start, priority: 0})

	// Walk the queue until empty or the end is found
	for pq.Len() > 0 {
		current := heap.Pop(&pq).(*PQItem)
		currentP := current.value

		if currentP == end {
			// Walk the path backward
			var path []Point
			for at := end; at != start; at = parent[at] {
				path = append([]Point{at}, path...)
			}
			return path
		}

		for _, dir := range Dirs {
			// put neighbors in queue
			nx := currentP.X + dir[0]
			ny := currentP.Y + dir[1]
			neighbor := Point{X: nx, Y: ny}
			// check if in bounds and didn't hit wall/edge
			if isValid(neighbor, grid) {
				newDist := distances[currentP] + 1
				// update the priority queue
				if newDist <= distances[neighbor] {
					parent[neighbor] = currentP
					distances[neighbor] = newDist
					heap.Push(&pq, &PQItem{value: neighbor, priority: newDist})
				}
			}

		}
	}
	return nil
}

func isValid(p Point, grid []string) bool {
	// Change '#' for whatever is OOB
	return p.X >= 0 && p.X < len(grid[0]) && p.Y >= 0 && p.Y < len(grid) && grid[p.Y][p.X] != '#'
}
