package main

import "fmt"

//数组定义
func arrays() {
	//无赋值定义
	var arr [5]int //默认值都为0
	//有初值定义
	arr1 := [3]int{1, 2, 3}
	arr2 := [...]int{6, 7, 8, 9, 0, 11}
	fmt.Println(arr, arr1, arr2)

	//数组遍历
	for i := range arr1 {
		fmt.Println(arr1[i])
	}

	//遍历下标和值
	for i, v := range arr2 {
		fmt.Println(i, v)
	}

	//遍历值
	for _, v := range arr {
		fmt.Println(v)
	}
}

//数组的传递也是值传递
func arrayy(array [5]int) {
	sum := 0
	array[1] = 100
	for i := range array {
		sum += array[i]
	}
	fmt.Println(sum)
	fmt.Println(array)
}

//数组引用传递
func arr(array *[5]int) {
	sum := 0
	for i := range array {
		sum += array[i]
	}
	fmt.Println(sum)
	array[0] = 100
	fmt.Println(*array)
}

func main() {
	fmt.Println("aaa")
	arrays()
	arrayss := [...]int{2, 4, 6, 8, 10}
	//arrayy(arrayss)
	//fmt.Println(arrayss)
	arr(&arrayss)
	fmt.Println(arrayss)
}
