package main

import (
	"fmt"
	"strings"
)

func sum(a ...int) int {
	return sums(a...)
}

func sums(b ...int) int {
	sum := 1
	for _, v := range b {
		sum += v
	}
	return sum
}

func add(i int) int {
	var sum int
	sum += i
	return sum
}

func adder() func(int) int {
	sum := 0
	return func(v int) int {
		sum += v
		return sum
	}
}

// this is example for ecmascript(闭包函数)  a,b is the local varisable but it can use for main
func fibonacci() func() int {
	a, b := 0, 1
	return func() int {
		a, b = b, a+b
		return a
	}
}

//define b 100

func main() {
	//a := adder()
	var a []byte
	a = []byte{1, 2, 3, 4, 5, 67}
	//b := fibonacci()
	for i := 0; i < 10; i++ {
		//b := add(i)
		//s := fmt.Sprintf("%d  ", a[i])
		_, w := strings.NewReader("sss").Read(a)
		fmt.Print(w)
		//fmt.Println(b())
		//fmt.Printf("%d ", a(i))
		//fmt.Printf("%d ", b)

	}
	//fmt.Println(b)
}
