package main

import (
	"bufio"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"strconv"
	"strings"
)

// if 写法
func openfile() {
	const (
		filename = "a.txt"
	)
	contents, err := ioutil.ReadFile(filename)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Printf("%s\n", contents)
	}

	//简化版
	if content, errs := ioutil.ReadFile(filename); errs != nil {
		fmt.Println(err)
	} else {
		fmt.Printf("%s\n", content)
	}

}

//switch 写法  默认行末都有break
func grade(gd int) string {
	switch {
	case gd < 60:
		fmt.Print("F")
	case gd < 70:
		fmt.Print("E")
	case gd < 80:
		fmt.Print("D")
	case gd < 90:
		fmt.Print("C")
	case gd < 100:
		fmt.Print("B")
	case gd == 100:
		fmt.Print("A")
	case gd < 0 || gd > 100:
		panic(fmt.Sprintf("Wrong score :%d", gd))
	}
	return ""
}

//循环语句
func convertoBin(n int) string {
	result := ""
	for ; n > 0; n /= 2 {
		lsb := n % 2
		result = strconv.Itoa(lsb) + result
	}
	return result
}

//循环语句-2
func printfile(filename string) {
	file, err := os.Open(filename)
	if err != nil {
		panic(err)
	}
	Readers(file)
}

func Readers(reader io.Reader) {
	scanner := bufio.NewScanner(reader)
	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}
}

func main() {
	fmt.Println("Hello world")
	openfile()

	// grade(1)
	// grade(80)
	// grade(70)
	// grade(10)
	// grade(100)
	// grade(101)
	// grade(-1)

	fmt.Println(convertoBin(12),
		convertoBin(2)) //1100

	printfile("a.txt")

	a := `dsfds
	sdfsd
	sdfsd
	fsd
	f
	sd
	f
	sd

	fds`
	fmt.Println()
	Readers(strings.NewReader(a))
	//死循环
	// for {
	// 	fmt.Println("a")
	// }
}
