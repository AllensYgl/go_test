package main

import (
	"fmt"
)

func main() {
	a := "你好，世界！"
	b := "你"
	var s string
	for i, v := range a {
		for _, k := range b {
			if rune(k) == rune(v) {
				s = a[:i]
				fmt.Printf("%d  %c\n", i, v)
			}
		}
	}
	fmt.Println(s)
}
