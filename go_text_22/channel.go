// package main

// import (
// 	"fmt"
// )

// func work(c chan int) {
// 	for {
// 		n := <-c
// 		fmt.Println(n)
// 	}
// }

// func chanDemo() {
// 	// a := make(chan int)
// 	// go work(a)
// 	// a <- 1
// 	// a <- 2
// 	// time.Sleep(time.Millisecond)

// 	a := make(<-chan int, 3)
// 	a := make(<-chan string, 4)
// 	var a [10]<-chan int
// 	for i := 0; i < 10; i++ {
// 		a[i] = make(<-chan int, 5)
// 	}

// 	for {
// 		if i, ok := <-a; !ok {
// 			break
// 		} else {
// 			fmt.Println(i)
// 		}
// 	}

// 	close(a)
// }

// func main() {
// 	chanDemo()
// }