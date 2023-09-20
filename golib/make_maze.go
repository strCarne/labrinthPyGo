package main

import (
	"C"
	"math/rand"
	"strconv"
)

func itoa(goInt int) *C.char {
	return C.CString((strconv.Itoa(goInt)))
}

func atoi(cStr *C.char) int {
	goInt, _ := strconv.Atoi(C.GoString(cStr))
	return goInt
}

const (
	downLeft = iota
	upLeft
	upRight
	downRight
)

var maze []byte
var visitedCells []int
var revisitedCells []int
var size int
var walls []int

//export MakeMaze
func MakeMaze(cSize *C.char, cI, cJ *C.char, cExit *C.char) *C.char {

	// Convertions
	size = atoi(cSize)
	i, j := atoi(cI), atoi(cJ)
	exitI, exitJ := calcExit(atoi(cExit))

	maze = make([]byte, size*size)
	for i := range maze {
		maze[i] = 2
	}

	visitedCells = make([]int, 0, size*size)
	revisitedCells = make([]int, 0, size*size)
	walls = make([]int, 0, size*size)

	maze[i*size+j] = 1
	maze[exitI*size+exitJ] = 1

	for continueBuildingMaze := true; continueBuildingMaze; {
		visitableNeighbors := checkNeighbors(i, j)
		if len(visitableNeighbors) != 0 {
			d := rand.Intn(len(visitableNeighbors))
			position := visitableNeighbors[d]
			maze[position] = 1
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

	if exitI < size/2 {
		maze[(exitI+1)*size+exitJ] = 1
	} else {
		maze[(exitI-1)*size+exitJ] = 1
	}

	return C.CString(string(maze))
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
