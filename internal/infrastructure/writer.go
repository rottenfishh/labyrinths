package infrastructure

import (
	"fmt"
	"os"

	"gitlab.education.tbank.ru/backend-academy-go-2025/homeworks/hw2-labyrinths/internal/domain/labyrinth"
)

func WriteMaze(fileName string, m *labyrinth.Maze) error {
	maze := m.String()
	var err error
	if fileName != "" {
		err = os.WriteFile(fileName, []byte(maze), 0644)
	} else {
		fmt.Println(maze)
	}
	return err
}
