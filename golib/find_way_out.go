package main

import (
	"C"
)
import "errors"

const (
	left = iota
	up
	down
	right
)

var dirs = []int{left, up, down, right}
var errOutOfMaze = errors.New("out of maze")

func calcNextPos(pos int, dir int) (int, error) {
	i, j := calcIndexes(pos)
	switch dir {
	case left:
		if j-1 < 0 {
			return 0, errOutOfMaze
		}
		return pos - 1, nil
	case right:
		if j+1 >= size {
			return 0, errOutOfMaze
		}
		return pos + 1, nil
	case up:
		if i-1 < 0 {
			return 0, errOutOfMaze
		}
		return pos - size, nil
	case down:
		if i+1 >= size {
			return 0, errOutOfMaze
		}
		return pos + size, nil
	}
	panic("wrong direction")
}

func calcDirByte(dir int) byte {
	switch dir {
	case left:
		return 'l'
	case right:
		return 'r'
	case up:
		return 'u'
	case down:
		return 'd'
	}
	panic("wrong direction")
}

func calcIndexes(position int) (int, int) {
	return position / size, position % size
}

//export FindWayOut
func FindWayOut(cPos *C.char, cExitPos *C.char) *C.char {
	pos := atoi(cPos)
	exitPos := atoi(cExitPos)

	res := new([]byte)

	for _, dir := range dirs {
		nextPos, err := calcNextPos(pos, dir)
		if err == nil {
			stack := make([]byte, 0, 64)
			char := calcDirByte(dir)
			findWayOut(nextPos, exitPos, append(stack, char), res, len(dirs)-dir-1)
		}
	}

	return C.CString(string(*res))
}

func findWayOut(pos int, exitPos int, stack []byte, result *[]byte, cameFrom int) {

	if maze[pos] == cellIsNotAvailable {
		return
	}

	if pos == exitPos {
		*result = make([]byte, len(stack))
		copy(*result, stack)
		return
	}

	for _, dir := range dirs {
		if dir == cameFrom {
			continue
		}
		if len(*result) > 0 {
			return
		}
		nextPos, err := calcNextPos(pos, dir)
		if err != nil {
			continue
		}
		char := calcDirByte(dir)
		findWayOut(nextPos, exitPos, append(stack, char), result, len(dirs)-dir-1)
	}
}
