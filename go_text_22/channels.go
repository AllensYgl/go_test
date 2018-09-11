package main

import "fmt"
import "time"

func Count(ch chan int) {
	ch <- 1
	fmt.Println("Counting")
}

func main() {
	chs := make([]chan int, 10)
	for i := 0; i < 10; i++ {
		chs[i] = make(chan int)
		go Count(chs[i])
	}

	for {
		for i := 0; i < 10; i++ {
			if _, ok := <-chs[i]; !ok {
				break
			} else {
				fmt.Println("hello", ok)
			}

		}
	}

	time.Sleep(3 * time.Second)
}
