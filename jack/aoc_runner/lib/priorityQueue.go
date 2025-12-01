package lib

// got this from the std lib package
// https://pkg.go.dev/container/heap
type PQItem struct {
	value    Point
	priority int
	// The index is needed by update and is maintained by the heap.Interface methods.
	index int
}

// A PQ implements heap.Interface and holds PQItems.
type PQ []*PQItem

func (pq PQ) Len() int { return len(pq) }

func (pq PQ) Less(i, j int) bool {
	// We want Pop to give us the highest, not lowest, priority so we use greater than here.
	return pq[i].priority > pq[j].priority
}

func (pq PQ) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
	pq[i].index = i
	pq[j].index = j
}

func (pq *PQ) Push(x any) {
	n := len(*pq)
	item := x.(*PQItem)
	item.index = n
	*pq = append(*pq, item)
}

func (pq *PQ) Pop() any {
	old := *pq
	n := len(old)
	item := old[n-1]
	old[n-1] = nil  // don't stop the GC from reclaiming the item eventually
	item.index = -1 // for safety
	*pq = old[0 : n-1]
	return item
}
