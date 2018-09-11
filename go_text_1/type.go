package main

import (
	"fmt"
	"math"
	"math/cmplx"
)

func variable() {
	var b int    //初值为0
	var a string //初值为 “”
	//fmt.Print(a, b)
	fmt.Printf("%q %d", a, b)
}

func variables() {
	var a, b int = 3, 4
	fmt.Println(a, b)
}

func varisbleType() {
	var a = 3
	var b = "string"
	fmt.Println(a, b)
}

func variableShprt() {
	a, b, c, d := 1, 2, 3, 4 // := 不可以在函数外面定义
	fmt.Println(a, b, c, d)
}

//包内部变量
var a = 'a'

//简便写法
var (
	s = 1
	b = 2
	c = 3
)

//char 已经没了 用 rune 表示字符
//uintptr  指针类型
//complex 表示复数

func ss() {
	a := 3 + 4i
	fmt.Println(cmplx.Abs(a))
}

//欧拉公式
func euler() {
	fmt.Println(cmplx.Pow(math.E, 1i*math.Pi) + 1)
}

func gg() {
	a, b := 3, 4
	c := math.Sqrt(float64(a*a + b*b))
	fmt.Println(c)
}

func consts() {
	//定义常量
	//首字母大写表示public
	const (
		filename = "aaa.txt"
		pi       = 3.1415926
		a, b     = 3, 4
	)
	fmt.Println(math.Sqrt(a*a + b*b))
}

//枚举类型
func enums() {
	const (
		cpp    = 0
		java   = 1
		python = 2
		golang = 3
	)
	fmt.Println(cpp, java, python, golang)
	//简化版本
	const (
		cpps = iota
		_    //表示无内容
		pythons
		golangs
		javascript
	)
	fmt.Println(cpps, javascript, pythons, golangs)

	const (
		bs = 1 << (10 * iota)
		kb
		mb
		gb
		tb
		pb
	)
	fmt.Println(bs, kb, mb, gb, tb, pb)

	const (
		ss = iota + 30
		tt
		uu
		vv
		ww
	)
	fmt.Println(ss, tt, uu, vv, ww)
}

func main() {
	fmt.Println("hello world")
	variables()
	varisbleType()
	variableShprt()
	fmt.Println(a, b, c, s)
	ss()
	euler()
	fmt.Printf("%.3f\n", cmplx.Exp(1i*math.Pi)+1) //欧拉公式  EXP表示 math.E 的 ?? 次方
	gg()
	consts()
	enums()

}
