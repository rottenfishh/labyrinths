package labyrinth

import "gitlab.education.tbank.ru/backend-academy-go-2025/homeworks/hw2-labyrinths/internal/domain"

type Cell struct {
	Coords   domain.Coords
	Visited  bool
	CellType string // wall or not wall
	CellStatus string
}

func NewCell(x, y int, cellType string, cellStatus string) Cell {
	return Cell{Coords: domain.Coords{x, y}, CellType: cellType, CellStatus: cellStatus}
}
