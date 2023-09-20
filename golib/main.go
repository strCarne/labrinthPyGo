package main

import (
	"C"
	"strconv"
)

func itoa(goInt int) *C.char {
	return C.CString((strconv.Itoa(goInt)))
}

func atoi(cStr *C.char) int {
	goInt, _ := strconv.Atoi(C.GoString(cStr))
	return goInt
}

func main() {
	//cp, ep := C.CString("18"), C.CString("16")
	//size = 5
	//maze = []byte{
	//	2, 2, 2, 2, 2,
	//	2, 1, 1, 1, 2,
	//	2, 1, 2, 1, 2,
	//	2, 1, 2, 1, 2,
	//	2, 2, 2, 2, 2,
	//}
	//
	//fmt.Println(C.GoString(FindWayOut(cp, ep)))
}
