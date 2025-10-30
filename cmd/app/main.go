package main

import (
	"fmt"
	"os"

	"gitlab.education.tbank.ru/backend-academy-go-2025/homeworks/hw2-labyrinths/internal/application"
)

func main() {
	action := os.Args[1]
	os.Args = append([]string{os.Args[0]}, os.Args[2:]...)
	switch action {
	case "generate":
		m := application.RunGenerator()
		if m == nil {
			break
		}
	case "solve":
		m := application.RunSolver()
		if m == nil {
			break
		}
	default:
		fmt.Println("Usage: maze-app [-hV] [COMMAND]")
		fmt.Println("Maze generator and solver CLI application.")
		fmt.Println("  -h, --help      Show this help message and exit.")
		fmt.Println("  -V, --version   Print version information and exit.")
		fmt.Println("Commands:")
		fmt.Println("  generate  Generate a maze with specified algorithm and dimensions.")
		fmt.Println("  solve     Solve a maze with specified algorithm and points.")

	}
}
