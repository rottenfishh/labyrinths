package pathfinding_algorithm

import (
	"container/heap"
	"math"

	"gitlab.education.tbank.ru/backend-academy-go-2025/homeworks/hw2-labyrinths/internal/domain"
	"gitlab.education.tbank.ru/backend-academy-go-2025/homeworks/hw2-labyrinths/internal/domain/labyrinth"
	"gitlab.education.tbank.ru/backend-academy-go-2025/homeworks/hw2-labyrinths/internal/domain/queue_utils"
)

type VertexHeuristic struct {
	point            domain.Coords
	distanceToVertex int
	distanceToEnd    int
	dist             int
	inOpenList       bool
	parent           *VertexHeuristic
}

func newVertexHeuristic(point domain.Coords) *VertexHeuristic {
	return &VertexHeuristic{point, inf, inf, inf, false, nil}
}

func (v *VertexHeuristic) Distance() int {
	return v.dist
}

func (v *VertexHeuristic) Prev() queue_utils.Node {
	return queue_utils.Node(v.parent)
}

func (v *VertexHeuristic) Coords() domain.Coords {
	return v.point
}

func initGraphAStar(m *labyrinth.Maze, startCell domain.Coords, graph map[domain.Coords]queue_utils.Node, visited map[domain.Coords]bool) {
	graph[startCell] = newVertexHeuristic(startCell)
	neighbours := GetNeighbourNodes(m, startCell)
	for _, cell := range neighbours {
		if visited[cell.Coords] {
			continue
		}
		vertex := newVertexHeuristic(cell.Coords)
		graph[cell.Coords] = vertex
		visited[cell.Coords] = true
		initGraphAStar(m, cell.Coords, graph, visited)
	}
}

func ManhattanDistanceToVertex(from domain.Coords, to domain.Coords) int {
	return int(math.Abs(float64(to.X-from.X)) + math.Abs(float64(to.Y-from.Y)))
}

func AStar(m *labyrinth.Maze, graph map[domain.Coords]queue_utils.Node, dest domain.Coords, start domain.Coords) {
	closedList := make(map[domain.Coords]*VertexHeuristic)
	openList := &queue_utils.PriorityQueue{}

	heap.Init(openList)
	startNode := newVertexHeuristic(start)
	startNode.distanceToVertex = 0
	startNode.dist = 0
	heap.Push(openList, startNode)
	startNode.inOpenList = true
	graph[start] = startNode
	for openList.Len() > 0 {
		currCoords := heap.Pop(openList).(*VertexHeuristic)
		currentNode := graph[currCoords.point].(*VertexHeuristic)
		closedList[currCoords.point] = currentNode
		if currCoords.point == dest {
			return
		}

		neighbours := GetNeighbourNodes(m, currentNode.point)
		for _, neighNode := range neighbours {
			if closedList[neighNode.Coords] != nil {
				continue
			}
			neighGraph := graph[neighNode.Coords].(*VertexHeuristic)
			newDistToVertex := currentNode.distanceToVertex + 1
			newDistToEnd := ManhattanDistanceToVertex(neighGraph.point, dest)
			if (newDistToVertex) < neighGraph.distanceToVertex {
				neighGraph.distanceToVertex = newDistToVertex
				neighGraph.distanceToEnd = newDistToEnd
				neighGraph.dist = neighGraph.distanceToVertex + neighGraph.distanceToEnd
				neighGraph.parent = currentNode
			}

			if !neighGraph.inOpenList {
				heap.Push(openList, neighGraph)
				neighGraph.inOpenList = true
			}
		}
	}
}

func FindPathAStar(m *labyrinth.Maze, startCell domain.Coords, destCell domain.Coords) bool {
	graph := make(map[domain.Coords]queue_utils.Node)
	visited := make(map[domain.Coords]bool)
	initGraphAStar(m, startCell, graph, visited)
	AStar(m, graph, destCell, startCell)
	if graph[destCell] == nil || graph[destCell].(*VertexHeuristic).dist == inf {
		return false
	}
	RestorePath(m, graph, startCell, destCell)
	return true
}
