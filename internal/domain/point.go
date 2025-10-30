package domain

type Coords struct {
	X, Y int
}

func NewCoords(x, y int) Coords {
	return Coords{x, y}
}
