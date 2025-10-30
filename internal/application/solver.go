package application

import (
	"flag"
	"fmt"
	"strconv"
	"strings"

	"gitlab.education.tbank.ru/backend-academy-go-2025/homeworks/hw2-labyrinths/internal/domain"
	"gitlab.education.tbank.ru/backend-academy-go-2025/homeworks/hw2-labyrinths/internal/domain/labyrinth"
	"gitlab.education.tbank.ru/backend-academy-go-2025/homeworks/hw2-labyrinths/internal/domain/pathfinding_algorithm"
	"gitlab.education.tbank.ru/backend-academy-go-2025/homeworks/hw2-labyrinths/internal/infrastructure"
)

type solverConfig struct {
	algorithm  string
	start      domain.Coords
	end        domain.Coords
	fileInput  string
	fileOutput string
}

func parseCoordsArgs(line string) (coords domain.Coords, err error) {
	lines := strings.Split(line, ",")
	if len(lines) != 2 {
		return coords, fmt.Errorf("Invalid point format: 11, expected format: x,y")
	}
	arr := make([]int, len(lines))
	for i, num := range lines {
		arr[i], err = strconv.Atoi(num)
		if err != nil {
			return coords, fmt.Errorf(err.Error() + "coordinates not a number")
		}
	}
	coords.X = arr[0] * 2
	coords.Y = arr[1] * 2
	return coords, nil
}

func parseCliSolver() (solverConfig, error) {
	algorithm := flag.String("algorithm", "dijkstra", "Path finding algorithm to use "+
		"dijkstra | astar")
	start := flag.String("start", "0,0", "start position. ex : 0,0")
	end := flag.String("end", "9,9", "end position. ex: 9,9")
	fileInput := flag.String("file", "input.txt", "input maze file name. string")
	fileOutput := flag.String("output", "", "output maze file name. string")
	flag.Parse()
	startCoords, err := parseCoordsArgs(*start)
	if err != nil {
		return solverConfig{}, err
	}

	endCoords, err := parseCoordsArgs(*end)
	if err != nil {
		return solverConfig{}, err
	}
	config := solverConfig{algorithm: *algorithm, start: startCoords, end: endCoords, fileInput: *fileInput, fileOutput: *fileOutput}
	return config, nil
}

func runSolverAlgorithm(config solverConfig) *labyrinth.Maze {
	m := infrastructure.ReadMaze(config.fileInput)
	if m == nil {
		return nil
	}

	var res bool
	if config.algorithm == "astar" {
		res = pathfinding_algorithm.FindPathAStar(m, config.start, config.end)
	} else {
		res = pathfinding_algorithm.FindPathDijkstra(m, config.start, config.end)
	}
	if !res {
		fmt.Println("No path found in input maze.")
		return nil
	}

	if config.fileOutput != "" {
		err := infrastructure.WriteMaze(config.fileOutput, m)
		if err != nil {
			fmt.Println(err.Error() + "writing maze to output file error.")
			return nil
		}
	} else {
		m.PrintMaze()
	}
	return m
}

func RunSolver() *labyrinth.Maze {
	config, err := parseCliSolver()
	if err != nil {
		fmt.Println(err)
		return nil
	}
	m := runSolverAlgorithm(config)
	if m == nil {
		return nil
	}
	return m
}
