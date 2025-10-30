package infrastructure

import (
	"fmt"
	"os"
	"strings"

	"gitlab.education.tbank.ru/backend-academy-go-2025/homeworks/hw2-labyrinths/internal/domain/labyrinth"
)

func ReadMaze(fileName string) *labyrinth.Maze {
	data, err := os.ReadFile(fileName)
	if err != nil {
		fmt.Println(err)
		return nil
	}
	lines := strings.Split(string(data), "\n")
	width := len(lines[0])
	height := len(lines)
	field := make([][]labyrinth.Cell, width)
	for i := range field {
		field[i] = make([]labyrinth.Cell, height)
	}

	for i, str := range lines {
		for j, char := range str {
			var cellType string
			switch char {
			case '#':
				cellType = "wall"
			default:
				cellType = "cell"
			}
			field[i][j] = labyrinth.NewCell(i, j, cellType, "outer")
		}
	}
	return labyrinth.BuildMaze(field, width, height)
}
