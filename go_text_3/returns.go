package main

import (
	"fmt"
	"reflect"
	"runtime"
)

//多返回值
func divs(a, b int) (q, r int, c error) {
	if b == 0 {
		return 0, 0, fmt.Errorf("unsupported operation: %d", b)
	} else {
		return a / b, a % b, nil
	}
}

//多返回值
func divss(a, b int) (q int) {
	if b == 0 {
		return 0
	} else {
		return a / b
	}
}

//方法参数为方法
func apply(op func(int, int) int, a, b int) int {
	p := reflect.ValueOf(op).Pointer()
	opName := runtime.FuncForPC(p).Name()
	fmt.Printf("Calling function %s with args (%d ,%d)\n", opName, a, b)
	return op(a, b)
}

//可变参数的方法
func sum(number ...int) int {
	s := 0
	for i := range number {
		s += number[i]
	}
	return s
}

//指针传递不存在对象复刻
func point(a *int) {
	fmt.Println(&a)
}

func pointw(a, b *int) {
	//*a, *b = *b, *a
	fmt.Println(*a, *b)
	a, b = b, a
	fmt.Println(*a, *b)
}

func main() {
	q, r, c := divs(20, 0)
	fmt.Println(q, r, c)
	fmt.Println(divs(29, 3))

	apply(divss, 10, 2)

	fmt.Println(sum(1, 2, 3, 4, 5, 6, 7, 8))

	a := 3
	m := &a
	fmt.Println(&m)

	point(m)

	q, w := 1, 2
	//fmt.Println(q, w)
	pointw(&q, &w)
	//fmt.Println(q, w)
}
