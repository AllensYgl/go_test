package main

import (
	"fmt"
	"go_text/go_text_12/baidu"
	"go_text/go_text_12/retriever"
)

// we can look
// definition a interface(Interf) and it have a Get() method
// and we definition a struct (Inter) it have a func is Get()
// and we can go through assign the struct to the interface and use the interface to calling func

type Interf interface {
	Get(url string) string
}

func dowload(i Interf) string {
	return i.Get("http://10.204.28.137")
}

func main() {
	var r Interf
	r = baidu.Inter{"aaaa"}
	fmt.Printf(" %T  %v \n", r, r)

	fmt.Println(dowload(baidu.Inter{}))
	r = retriever.Retriever{TimeOut: 20}
	//fmt.Println(dowload(r))

	fmt.Printf(" %T  %v \n", r, r)

	// type assertion
	//go through the .( typeName ) to get the type value
	real := r.(retriever.Retriever)
	fmt.Println(real.TimeOut)
}
