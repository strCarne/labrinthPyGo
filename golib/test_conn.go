package main

import (
	"C"
	"fmt"
)

//export Hello
func Hello(name *C.char) {
	goName := C.GoString(name)
	fmt.Println("Hello,", goName)
}

//export Factorial
func Factorial(cN *C.char) *C.char {
	n := atoi(cN)
	res := 1
	for i := 2; i <= n; i++ {
		res *= i
	
	}
	return itoa(res)
}