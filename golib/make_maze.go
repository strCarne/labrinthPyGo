package main

import (
	"C"
	"math/rand"
)
import "strconv"


const (
	downLeft = iota
	upLeft
	upRight
	downRight
	center

	cellIsAvailable = 1
	cellIsNotAvailable = 2
)

var maze []byte
var size int
var walls []int

//export MakeMaze
func MakeMaze(cSize *C.char, cExit *C.char) *C.char {

	// Convertions
	size = atoi(cSize)
	i, j := calcExit(atoi(cExit))

	maze = make([]byte, size*size)
	for i := range maze {
		maze[i] = cellIsNotAvailable
	}

	visitedCells := make([]int, 0, size*size)
	revisitedCells := make([]int, 0, size*size)
	walls = make([]int, size*size)

	maze[i*size+j] = cellIsAvailable

	for continueBuildingMaze := true; continueBuildingMaze; {
		visitableNeighbors := checkNeighbors(i, j)
		if len(visitableNeighbors) != 0 {
			d := rand.Intn(len(visitableNeighbors))
			position := visitableNeighbors[d]
			maze[position] = cellIsAvailable
			visitedCells = append(visitedCells, position)
			i, j = position/size, position%size
		}
		if len(visitableNeighbors) == 0 {
			if len(visitedCells) > 0 {
				k := len(visitedCells) - 1
				i, j = visitedCells[k]/size, visitedCells[k]%size
				revisitedCells = append(revisitedCells, visitedCells[k])
				visitedCells = visitedCells[:k]
			} else {
				continueBuildingMaze = false
			}

		}
	}

	playerPosition := revisitedCells[rand.Intn(len(revisitedCells))]

	return C.CString(string(maze) + strconv.Itoa(playerPosition))
}

func calcExit(exit int) (int, int) {
	switch exit {
	case downLeft:
		return size - 2, 1
	case upLeft:
		return 1, 1
	case upRight:
		return 1, size - 2
	case downRight:
		return size - 2, size - 2
	case center:
		return size - 1, size / 2
	}
	panic("calcExit: wrong params")
}

func checkNeighbors(i, j int) []int {
	neighbors := [][][2]int{
		{{i, j - 1}, {i - 1, j - 2}, {i, j - 2}, {i + 1, j - 2}, {i - 1, j - 1}, {i + 1, j - 1}}, // left
		{{i, j + 1}, {i - 1, j + 2}, {i, j + 2}, {i + 1, j + 2}, {i - 1, j + 1}, {i + 1, j + 1}}, // right
		{{i - 1, j}, {i - 2, j - 1}, {i - 2, j}, {i - 2, j + 1}, {i - 1, j - 1}, {i - 1, j + 1}}, // up
		{{i + 1, j}, {i + 2, j - 1}, {i + 2, j}, {i + 2, j + 1}, {i + 1, j - 1}, {i + 1, j + 1}}, // down
	}

	visitableNeighbors := make([]int, 0, size*size)

	for _, neighbor := range neighbors {
		if neighbor[0][0] > 0 && neighbor[0][0] < (size-1) && neighbor[0][1] > 0 && neighbor[0][1] < (size-1) {
			if maze[neighbor[1][0]*size+neighbor[1][1]] == 1 ||
				maze[neighbor[2][0]*size+neighbor[2][1]] == 1 ||
				maze[neighbor[3][0]*size+neighbor[3][1]] == 1 ||
				maze[neighbor[4][0]*size+neighbor[4][1]] == 1 ||
				maze[neighbor[5][0]*size+neighbor[5][1]] == 1 {
				walls = append(walls, neighbor[0][0]*size+neighbor[0][1])
			} else {
				visitableNeighbors = append(visitableNeighbors, neighbor[0][0]*size+neighbor[0][1])
			}
		}
	}

	return visitableNeighbors
}
