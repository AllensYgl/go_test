package main

import (
	"fmt"
	"go_text/go_text_13/inter"
)

func main() {
	//var s inter.Interf

	fmt.Println(inter.AA{Value: "aaaa"}.Post())
	m := inter.AA{Value: "bbb"}
	fmt.Println(m.Post())
	//inter.Sseesio(s)
}
