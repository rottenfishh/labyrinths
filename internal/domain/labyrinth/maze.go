package labyrinth

import (
	"fmt"
	"strings"
)

type Maze struct {
	Field         [][]Cell
	Width, Height int
}

func BuildMaze(field [][]Cell, width, height int) *Maze {
	return &Maze{field, width, height}
}

func NewMaze(width, height int) Maze {
	width = width*2 + 1
	height = height*2 + 1
	field := make([][]Cell, width)
	for i := range field {
		field[i] = make([]Cell, height)
		for j := range field[i] {
			if i%2 == 0 && j%2 == 0 {
				field[i][j] = NewCell(i, j, "cell", "outer")
			} else {
				field[i][j] = NewCell(i, j, "wall", "outer")
			}
		}
	}
	return Maze{field, width, height}
}

func (m Maze) PrintMaze() {
	for i := 0; i < m.Width; i++ {
		for j := 0; j < m.Height; j++ {
			if m.Field[i][j].CellType == "wall" {
				fmt.Print("# ")
			} else if m.Field[i][j].CellType == "path" {
				fmt.Print(". ")
			} else if m.Field[i][j].CellType == "start" {
				fmt.Print("0 ")
			} else if m.Field[i][j].CellType == "end" {
				fmt.Print("X ")
			} else {
				fmt.Print("  ")
			}
		}
		fmt.Println("")
	}
}

func (m Maze) String() string {
	sb := new(strings.Builder)
	for i := 0; i < m.Width; i++ {
		for j := 0; j < m.Height; j++ {
			if m.Field[i][j].CellType == "wall" {
				sb.WriteString("#")
			} else if m.Field[i][j].CellType == "path" {
				sb.WriteString(".")
			} else if m.Field[i][j].CellType == "start" {
				sb.WriteString("0")
			} else if m.Field[i][j].CellType == "end" {
				sb.WriteString("X")
			} else {
				sb.WriteString(" ")
			}
		}
		sb.WriteString("\n")
	}
	return sb.String()
}
