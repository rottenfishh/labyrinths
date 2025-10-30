package queue_utils

type PriorityQueue []Node

func (pq PriorityQueue) Len() int {
	return len(pq)
}
func (pq PriorityQueue) Less(i, j int) bool {
	return pq[i].Distance() < pq[j].Distance()
}

func (pq PriorityQueue) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
}

func (pq *PriorityQueue) Push(x any) {
	*pq = append(*pq, x.(Node))
}

func (pq *PriorityQueue) Pop() any {
	old := *pq
	n := len(old)
	item := old[n-1] // Get last element
	*pq = old[:n-1]  // Remove last element
	return item
}
