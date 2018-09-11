package main

import (
	"fmt"
)

func throwPanic() int {
	// the defer will be run after painc throws the error
	// and in front of return
	defer func() {
		r := recover()
		if err, ok := r.(error); ok {
			fmt.Println("ERROR :", err)
		} else {
			panic(fmt.Sprintf("I don't know what to do : %v", r))
		}
	}()

	//panic(errors.New("this is an error"))
	// b := 0
	// a := 5 / b
	// fmt.Println(a)
	panic(123)
	return 1
}

func main() {
	throwPanic()
	fmt.Print("HELLO WORLD")
}
