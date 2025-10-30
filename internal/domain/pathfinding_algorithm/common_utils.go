package pathfinding_algorithm

import (
	"gitlab.education.tbank.ru/backend-academy-go-2025/homeworks/hw2-labyrinths/internal/domain"
	"gitlab.education.tbank.ru/backend-academy-go-2025/homeworks/hw2-labyrinths/internal/domain/labyrinth"
	"gitlab.education.tbank.ru/backend-academy-go-2025/homeworks/hw2-labyrinths/internal/domain/queue_utils"
)

func GetNeighbourNodes(maze *labyrinth.Maze, node domain.Coords) []*labyrinth.Cell {
	var neighbours []*labyrinth.Cell
	coordsX := []int{+1, -1, 0, 0}
	coordsY := []int{0, 0, +1, -1}
	for i := 0; i < len(coordsX); i++ {
		newX := coordsX[i] + node.X
		newY := coordsY[i] + node.Y
		if (newX < maze.Width && newX >= 0) && (newY < maze.Height && newY >= 0) {
			if maze.Field[newX][newY].CellType != "wall" {
				neighbours = append(neighbours, &maze.Field[newX][newY])
			}
		}
	}
	return neighbours
}

func RestorePath(m *labyrinth.Maze, graph map[domain.Coords]queue_utils.Node, start, dest domain.Coords) {
	for v := graph[dest]; v.Prev() != nil && v.Coords() != start; v = v.Prev() {
		c := v.Coords()
		m.Field[c.X][c.Y].CellType = "path"
	}
	m.Field[start.X][start.Y].CellType = "start"
	m.Field[dest.X][dest.Y].CellType = "end"
}
