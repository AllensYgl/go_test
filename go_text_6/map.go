package main

import "fmt"

//hash map
func mapCreate() {
	// way 1
	//		key   value
	m := map[string]string{
		"name":    "allen guo",
		"course":  "golang",
		"site":    "imooc",
		"quality": "notbad",
	}
	l()

	//输出完全随机 不分先后
	fmt.Println("m=", m)
	l()
	//way 2
	m2 := make(map[string]int)
	fmt.Println("m2=", m2) // m2 == empty map[]
	l()
	//way 3
	var m3 map[int]string
	fmt.Println("m3=", m3) // m3 == nil
	l()

	//map 遍历
	for i, v := range m {
		fmt.Println(i, v)
	}
	l()

	//获得某个 key 的 value
	courseName := m["course"]
	fmt.Println("courseName =", courseName)
	l()
	//注意  key 不存在的情况下 会输出 空
	cases := m["cases"]
	fmt.Println("cases =", cases)
	l()

	//判断 key 是否存在
	_, ok := m["course"]
	fmt.Println("the key have the course  ? ", ok)
	_, oks := m["cour"]
	fmt.Println("the key have the cour  ? ", oks)
	l()

	//配合 if 语句使用
	if name, ok1 := m["name"]; ok1 {
		fmt.Println("the name is =", name)
	} else {
		fmt.Printf("the name %s is not have\n", name)
	}
	l()

	//删除map元素
	delete(m, "name")
	fmt.Println("delete succed")
	if name, ok1 := m["name"]; ok1 {
		fmt.Println("the name is =", name)
	} else {
		fmt.Printf("the name %s is not have\n", name)
	}
	l()

}

func l() {
	fmt.Println("-----------------华丽分割线------------------")
}

func main() {
	mapCreate()
}
