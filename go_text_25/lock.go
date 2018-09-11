package main

import (
	"fmt"
	"sync"
)

type hei struct {
	value int
	locl  sync.Mutex //add lock
}

func (a *hei) add() {
	a.locl.Lock()
	defer a.locl.Unlock()
	a.value++
}

func (a *hei) get() int {
	a.locl.Lock()
	defer a.locl.Unlock()
	return int(a.value)
}

func main() {
	var a hei
	go func() {
		a.add()
	}()
	fmt.Println(a.get())
}
