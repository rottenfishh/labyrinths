package application

import (
	"flag"
	"fmt"

	"gitlab.education.tbank.ru/backend-academy-go-2025/homeworks/hw2-labyrinths/internal/domain/generation_algorithm"
	"gitlab.education.tbank.ru/backend-academy-go-2025/homeworks/hw2-labyrinths/internal/domain/labyrinth"
	"gitlab.education.tbank.ru/backend-academy-go-2025/homeworks/hw2-labyrinths/internal/infrastructure"
)

type GeneratorConfig struct {
	algorithm string
	width     int
	height    int
	fileName  string
}

func parseCliGeneration() GeneratorConfig {
	algorithm := flag.String("algorithm", "dfs", "Generation algorithm to use "+
		"dfs | prim")
	width := flag.Int("width", 10, "width of maze. int number")
	height := flag.Int("height", 10, "height of maze. int number")
	file := flag.String("output", "", "output file name. string")
	flag.Parse()
	config := GeneratorConfig{algorithm: *algorithm, width: *width, height: *height, fileName: *file}
	return config
}

func runGenerationAlgorithm(config GeneratorConfig) *labyrinth.Maze {
	m := labyrinth.NewMaze(config.width, config.height)
	if config.algorithm == "prim" {
		generation_algorithm.Prim(&m)
	} else {
		generation_algorithm.Dfs(&m)
	}
	if config.fileName != "" {
		err := infrastructure.WriteMaze(config.fileName, &m)
		if err != nil {
			fmt.Println(err)
			return nil
		}
	} else {
		m.PrintMaze()
	}
	return &m
}

func RunGenerator() *labyrinth.Maze {
	config := parseCliGeneration()
	return runGenerationAlgorithm(config)
}
