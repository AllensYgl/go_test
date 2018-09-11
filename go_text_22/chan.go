package main

import (
	"fmt"
	"time"
)

func main() {

	c := make(chan int, 20)
	go func() {
		for i := 0; i < 10; i++ {
			c <- i
		}
	}()
	for i := range c {

		fmt.Print(<-c)
	}
	time.Sleep(time.Microsecond * 10)

}
