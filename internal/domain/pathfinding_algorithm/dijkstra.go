package pathfinding_algorithm

import (
	"container/heap"
	"math"

	"gitlab.education.tbank.ru/backend-academy-go-2025/homeworks/hw2-labyrinths/internal/domain"
	"gitlab.education.tbank.ru/backend-academy-go-2025/homeworks/hw2-labyrinths/internal/domain/labyrinth"
	"gitlab.education.tbank.ru/backend-academy-go-2025/homeworks/hw2-labyrinths/internal/domain/queue_utils"
)

type Vertex struct {
	visited bool
	dist    int
	point   domain.Coords
	parent  *Vertex
}

func newVertex(coords domain.Coords) *Vertex {
	return &Vertex{false, inf, coords, nil}
}

func (v *Vertex) Distance() int {
	return v.dist
}

func (v *Vertex) Prev() queue_utils.Node {
	return queue_utils.Node(v.parent)
}

func (v *Vertex) Coords() domain.Coords {
	return v.point
}

const (
	inf = math.MaxInt
)

func initGraph(m *labyrinth.Maze, startCell domain.Coords, graph map[domain.Coords]queue_utils.Node, visited map[domain.Coords]bool) {
	graph[startCell] = newVertex(startCell)
	neighbours := GetNeighbourNodes(m, startCell)
	for _, cell := range neighbours {
		if visited[cell.Coords] {
			continue
		}
		vertex := newVertex(cell.Coords)
		graph[cell.Coords] = vertex
		visited[cell.Coords] = true
		initGraph(m, cell.Coords, graph, visited)
	}
}

func Dijkstra(m *labyrinth.Maze, startCell domain.Coords, destCell domain.Coords) map[domain.Coords]queue_utils.Node {
	graph := make(map[domain.Coords]queue_utils.Node)
	visited := make(map[domain.Coords]bool)
	initGraph(m, startCell, graph, visited)
	queue := &queue_utils.PriorityQueue{}
	solution := make([]domain.Coords, 0, len(graph))
	heap.Init(queue)
	start := graph[startCell].(*Vertex)
	start.dist = 0
	heap.Push(queue, start)
	for queue.Len() > 0 {
		vertex := heap.Pop(queue).(*Vertex)
		if vertex.point == destCell {
			return graph
		}
		if vertex.visited {
			continue
		}
		vertex.visited = true
		neighbours := GetNeighbourNodes(m, vertex.point)
		for _, cell := range neighbours {
			toVertex := graph[cell.Coords].(*Vertex)
			if toVertex == nil {
				continue
			}
			newDist := vertex.dist + 1
			if toVertex.dist > newDist {
				toVertex.dist = newDist
				toVertex.parent = vertex
				heap.Push(queue, toVertex)
				solution = append(solution, toVertex.point)
			}
		}
	}
	return graph
}

func FindPathDijkstra(m *labyrinth.Maze, startCell domain.Coords, destCell domain.Coords) bool {
	graph := Dijkstra(m, startCell, destCell)
	if graph[destCell] == nil || graph[destCell].(*Vertex).dist == inf {
		return false
	}
	RestorePath(m, graph, startCell, destCell)
	return true
}
