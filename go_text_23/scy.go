package main

import (
	"sync"
)

func main() {
	var wg sync.WaitGroup
	// add 20 task
	wg.Add(20)

	//do a task
	// means add -1
	wg.Done()

	//wait 20 task finish
	// means add=0 finish
	wg.Wait()
}
