package main

import "fmt"

//切片  把数组切割出一块区域用来使用  切片方法是左闭右开的
// slice 是 array 的一个视图 view
//传递 slice 是值传递的指针 可以进行修改
func slic(array []int) {
	sum := 0
	for i := range array {
		sum += array[i]
	}
	fmt.Printf("sum=%d\n", sum)
	fmt.Printf("源:array=%v\n", array)
	array[0] = 100
	fmt.Printf("改:array=%v\n", array)
}

//切片的性质及操作
func slice() {
	arr := [...]int{2, 4, 6, 8, 10}
	fmt.Println("arr[2:5]", arr[2:5])
	fmt.Println("arr[:5]", arr[:5])
	fmt.Println("arr[2:]", arr[2:])
	fmt.Println("arr[:]", arr[:])
	slic(arr[:])
	fmt.Println(arr)

	//切片操作只要数组足够长 就算切片的大小不足以继续切出更大的区域依旧可以实现切片
	// slice 可以向后扩展 不可以向前扩展
	arrs := [...]int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	fmt.Println(arrs)

	/*

					  slice
								   cap
		ptr		------len-----------|-----
		|	    |_______|________	     |
		|		|				|		 |
		**********************************
	*/

	ss := arrs[2:6]
	fmt.Printf("ss=%v,len(ss)=%d,cap(ss)=%d\n", ss, len(ss), cap(ss))
	sss := ss[2:6] //超出下标依旧可以实现切片 只要数组元素足够长
	fmt.Printf("sss=%v,len(sss)=%d,cap(sss)=%d\n", sss, len(sss), cap(sss))

	fmt.Println("*********************************************************")
	//添加元素超过cap的时候系统会自动分配一个更长的数组
	arrss := [...]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	s1 := arrss[2:4]
	fmt.Printf("s1=%v,len(s1)=%d,cap(s1)=%d\n", s1, len(s1), cap(s1))
	s2 := append(s1, 55)
	fmt.Printf("s2=%v,len(s2)=%d,cap(s2)=%d\n", s2, len(s2), cap(s2))
	//fmt.Println(arrss)		append如果再数组范围内则会修改而不是添加元素
	fmt.Printf("arrss=%v,len(arrss)=%d,cap(arrss)=%d\n", arrss, len(arrss), cap(arrss))

	fmt.Println("*********************************************************")
	//超过范围则会新建一个数组
	s3 := arrss[6:10]
	fmt.Printf("s3=%v,len(s3)=%d,cap(s3)=%d\n", s3, len(s3), cap(s3))
	s4 := append(s3, 55)
	fmt.Printf("s4=%v,len(s4)=%d,cap(s4)=%d\n", s4, len(s4), cap(s4))
	//数组最后一个超出10的的部分无法在源数组显示 而是重新创建一个新的数组
	fmt.Printf("arrss=%v,len(arrss)=%d,cap(arrss)=%d\n", arrss, len(arrss), cap(arrss))
}

//cap 的增长是按照2^？次方进行增长的
func printSlice() {
	var array []int
	for i := 0; i < 10; i++ {
		fmt.Printf("array=%v   len(array)=%d  cap(array)=%d\n", array, len(array), cap(array))
		array = append(array, i)
	}
	fmt.Println()
}

func printfslice(array []int) {
	fmt.Printf("array=%v len(array)=%d cap(array)=%d \n", array, len(array), cap(array))
}

func copyslice() {
	//copy操作	长slice copy给 短slice 会截取长slice的部分copy 而不是扩展短slice
	array1 := []int{1, 2, 3, 4}
	array2 := []int{5, 6, 7, 8, 9}
	printfslice(array1)
	printfslice(array2)
	fmt.Println("------------------copy---------------------")
	copy(array1, array2)
	printfslice(array1)
	printfslice(array2)
}

//delet操作不会修改cap
func deletslice() {
	array := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	printfslice(array)
	fmt.Println("-----------------delet '5' ----------------------")
	array = append(array[:4], array[5:]...)
	printfslice(array)
}

//pop操作会改变cap
func popfirst() {
	array := []int{1, 2, 3, 4, 5, 6, 7}
	printfslice(array)

	frist := array[0]
	fmt.Printf("frist: %d\n", frist)

	array = array[1:]
	printfslice(array)
}

//pop操作会改变cap
func poplast() {
	array := []int{1, 2, 3, 4, 5, 6, 7}
	printfslice(array)

	frist := array[len(array)-1]
	fmt.Printf("last: %d\n", frist)

	array = array[:len(array)-1]
	printfslice(array)
}

func main() {
	//slice()
	//printSlice()

	//创建一个slice
	//slice := []int{1, 2, 3, 4}
	//slic(slice)

	//创建一个固定长度的slice  type len	默认cap等于len
	//slice1 := make([]int, 10)
	//创建固定长度和cap的slice	type len cap
	//slice2 := make([]int, 10, 20)

	//printfslice(slice1)
	//printfslice(slice2)

	//copyslice()
	//deletslice()
	fmt.Println("---------------------------------------------------")
	popfirst()

	fmt.Println("---------------------------------------------------")

	poplast()
	fmt.Println("---------------------------------------------------")
}
