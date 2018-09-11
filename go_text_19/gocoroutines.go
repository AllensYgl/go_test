package main

import (
	"fmt"
)

func add(x, y int) {
	fmt.Println(x + y)
}

//coroutine need setting running time

func main() {
	i := 1
	j := 2
	for {
		go add(i, j)
		i++
		j++
	}
}
