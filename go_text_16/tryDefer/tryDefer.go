package main

import (
	"bufio"
	"fmt"
	"os"
)

func tryDefer() {
	defer fmt.Println("1")
	defer fmt.Println("2")
	defer fmt.Println("3")
	panic("error occurred")
	//defer fmt.Println("4")

}

func Fibonacci() func() int {
	a, b := 0, 1
	return func() int {
		a, b = b, a+b
		return a
	}
}

func writeFile(filename string) {
	file, err := os.OpenFile(filename, os.O_EXCL|os.O_CREATE, 0666)
	//err = errors.New("this is a error")
	if err != nil {
		//panic(err)
		if pathError, ok := err.(*os.PathError); !ok {
			panic(err)
		} else {
			fmt.Println("Error:", err.Error())
			fmt.Println("pathError.Op:", pathError.Op)
			fmt.Println("pathError.Err:", pathError.Err)
			fmt.Println("pathError.Path:", pathError.Path)
			return
		}
	}
	defer file.Close()
	write := bufio.NewWriter(file)
	defer write.Flush()

	f := Fibonacci()
	for i := 0; i < 20; i++ {
		fmt.Fprintln(write, f())
	}

}

func main() {
	fmt.Println("ss ")
	writeFile("a.txt")
	//tryDefer()
}
