package main

import (
	"fmt"
)

func longset() {
	frist := 0
	flag := 0
	len := 0
	//str := "erfrertertvsdmsadmnsdfmojsadasdsa"
	str := "这回就用中文来试试你觉得怎么样？？fwer lwedjlka？？"
	var str1 string
	//遍历字符串
	for i, v := range str {

		if i-frist+1 > len {

			len = i - frist + 1
			//fmt.Printf("%d ", len)
			str1 = str[frist:i]

		}

		//把不重复的加入新字符串
		for _, k := range str1 {
			//出现重复
			fmt.Printf("v=%c   k=%c\n", v, k)

			if rune(v) == rune(k) {
				flag = 1
				break
			}
		}
		//处理重复
		if flag == 1 {
			frist = i
			flag = 0
		}

	}

	for _, ch := range str1 {
		fmt.Printf("%c", ch)
	}
	fmt.Println()
	fmt.Printf("%d \n", len)
}

func LengOfNonRepeatingSubStr(s string) int {
	last := make(map[rune]int)
	start := 0
	maxLength := 0

	for i, ch := range []rune(s) {
		if lastI, ok := last[ch]; ok && lastI >= 0 {
			start = lastI + 1
		}
		if i-start+1 > maxLength {
			maxLength = i - start + 1
		}
		last[ch] = i
	}
	return maxLength
}

func main() {
	//fmt.Println(LengOfNonRepeatingSubStr("中文也不是不可以"))
	//fmt.Println()
	longset()
}
