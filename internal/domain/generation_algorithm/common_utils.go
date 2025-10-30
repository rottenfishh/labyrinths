package generation_algorithm

import (
    "gitlab.education.tbank.ru/backend-academy-go-2025/homeworks/hw2-labyrinths/internal/domain"
    "gitlab.education.tbank.ru/backend-academy-go-2025/homeworks/hw2-labyrinths/internal/domain/labyrinth"
)

func GetNeighbours(maze *labyrinth.Maze, node domain.Coords) []*labyrinth.Cell {
    var neighbours []*labyrinth.Cell
    coordsX := []int{+2, -2, 0, 0}
    coordsY := []int{0, 0, +2, -2}
    for i := 0; i < len(coordsX); i++ {
        newX := coordsX[i] + node.X
        newY := coordsY[i] + node.Y
        if (newX < maze.Width && newX >= 0) && (newY < maze.Height && newY >= 0) {
            neighbours = append(neighbours, &maze.Field[newX][newY])
        }
    }
    return neighbours
}

func FindWall(cell domain.Coords, elem domain.Coords) domain.Coords {
    x := cell.X
    y := cell.Y
    if cell.X == elem.X {
        if cell.Y > elem.Y {
            y = cell.Y - 1
        } else {
            y = cell.Y + 1
        }
    }
    if cell.Y == elem.Y {
        if cell.X > elem.X {
            x = cell.X - 1
        } else {
            x = cell.X + 1
        }
    }
    return domain.NewCoords(x, y)
}