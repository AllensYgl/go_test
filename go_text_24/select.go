package main

import (
	"fmt"
	"math/rand"
	"time"
)

func makeChan() chan int {
	out := make(chan int)
	go func() {

		for i := 0; i < 10; i++ {
			time.Sleep(time.Duration(rand.Intn(1500)) * time.Millisecond)
			out <- i
		}
	}()
	return out
}

func main() {
	a, b := makeChan(), makeChan()
	tim := time.After(time.Second * 10)
	tick := time.Tick(time.Second)
	for i := 0; i < 20; i++ {
		select {
		case n := <-a:
			fmt.Println("received from a", n)
		case n := <-b:
			fmt.Println("received from b", n)
		case <-time.After(800 * time.Second):
			fmt.Println("800s")
		case <-tick:
			fmt.Println("once tick")
		case <-tim:
			fmt.Println("Bye")
			return
		}
	}
}
