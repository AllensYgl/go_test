package main

import (
	"fmt"
	"runtime"
	"time"
)

func main() {
	for i := 0; i < 100000; i++ {
		go func(i int) {
			for {
				fmt.Printf("Hello from go routine %d\n", i)
				runtime.Gosched() //give up control to other routine
			}
		}(i)
	}
	time.Sleep(time.Millisecond * 100)
}
