package generation_algorithm

import (
	"math/rand"

	"gitlab.education.tbank.ru/backend-academy-go-2025/homeworks/hw2-labyrinths/internal/domain"
	"gitlab.education.tbank.ru/backend-academy-go-2025/homeworks/hw2-labyrinths/internal/domain/labyrinth"
	"gitlab.education.tbank.ru/backend-academy-go-2025/homeworks/hw2-labyrinths/internal/domain/queue_utils"
)

var stack = queue_utils.Stack{}

func Dfs(m *labyrinth.Maze) {
	startCell := domain.Coords{}
	m.Field[startCell.X][startCell.Y].Visited = true
	m.Field[startCell.X][startCell.Y].CellType = "start"

	stack.Push(&m.Field[startCell.X][startCell.Y])
	for len(stack.Items) != 0 {
		newCell, res := stack.Pop()
		if !res {
			return
		}
		cell, ok := newCell.(*labyrinth.Cell)
		if !ok {
			return
		}

		neighbours := GetNeighbours(m, cell.Coords)
		rand.Shuffle(len(neighbours), func(i, j int) {
			neighbours[i], neighbours[j] = neighbours[j], neighbours[i]
		})

		for _, elem := range neighbours {
			if !elem.Visited {
				stack.Push(cell)

				wall := FindWall(cell.Coords, elem.Coords)
				m.Field[wall.X][wall.Y].CellType = "cell"

				elem.Visited = true
				elem.CellType = "cell"
				stack.Push(elem)
				break
			}
		}
	}
}
