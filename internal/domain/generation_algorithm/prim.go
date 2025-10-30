package generation_algorithm

import (
	"math/rand"

	"gitlab.education.tbank.ru/backend-academy-go-2025/homeworks/hw2-labyrinths/internal/domain"
	"gitlab.education.tbank.ru/backend-academy-go-2025/homeworks/hw2-labyrinths/internal/domain/labyrinth"
)

var borderCells []*labyrinth.Cell

func remove(slice []*labyrinth.Cell, s int) []*labyrinth.Cell {
	return append(slice[:s], slice[s+1:]...)
}

func Prim(m *labyrinth.Maze) {
	startCell := domain.Coords{}
	endCell := domain.Coords{X: m.Width - 1, Y: m.Height - 1}
	m.Field[startCell.X][startCell.Y].CellType = "inner"

	neighbours := GetNeighbours(m, m.Field[startCell.X][startCell.Y].Coords)
	for _, elem := range neighbours {
		elem.CellStatus = "border"
		borderCells = append(borderCells, elem)
	}
	for len(borderCells) > 0 {
		n := rand.Intn(len(borderCells))
		chosenCell := borderCells[n]
		chosenCell.CellType = "cell"
		chosenCell.CellStatus = "inner"
		borderNeighbours := GetNeighbours(m, chosenCell.Coords)
		borderCells = remove(borderCells, n)
		rand.Shuffle(len(borderNeighbours), func(i, j int) {
			borderNeighbours[i], borderNeighbours[j] = borderNeighbours[j], borderNeighbours[i]
		})
		flag := false
		for _, elem := range borderNeighbours {
			if !flag && elem.CellStatus == "inner" {
				wall := FindWall(chosenCell.Coords, elem.Coords)
				m.Field[wall.X][wall.Y].CellType = "cell"
				flag = true
			}
			if elem.CellStatus == "outer" {
				elem.CellStatus = "border"
				borderCells = append(borderCells, elem)
			}
		}
		if chosenCell.Coords == endCell {
			chosenCell.CellType = "end"
			return
		}
	}
}
