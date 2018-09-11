package main

import (
	"fmt"
	"go_text/go_text_10/Queue"
)

// pay attention to : the queue is variable means when the code running the queue possible to change the cap and len
//because the queue is a []int -----slice
func main() {
	queue := Queue.Queue{}
	fmt.Println(queue)
	queue.Push(6)
	queue.Push(7)
	fmt.Println(queue)
	fmt.Println(queue.Pop())
	fmt.Println(queue.Pop())
	fmt.Println(queue)
	fmt.Print(queue.IsEmpty())
	fmt.Println()
	//queue.Push("aaaaaa")
	fmt.Println(queue)

	//fmt.Println(queue.Pop())
	fmt.Println(queue)
}
