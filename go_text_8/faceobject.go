package main

import (
	"fmt"
)

//in go language we don't need to know how the memory assign ,because the system can automatically identify how to assign
type treeNode struct {
	value       int
	left, right *treeNode
}

// pay attention to: Local variables don't be released when the Local function run ending
func factory(values int) *treeNode {
	return &treeNode{value: values} //this is Local variables but it can be used for External function
}

//identify the struct func
// func + ( structName + structType ) + functionName ( ParameterName + Parametertype ) + { functionContent  }
func (node treeNode) print() {
	fmt.Println(node.value)
}

//pay attention to : it is not to change the value because the object transfer is value not reference
// if we need only to change the value we can add a '*' in front of the structType
func (node treeNode) setValues(value int) {
	node.value = value
}

func (node *treeNode) trave() {
	if node == nil {
		return
	}
	//pay attention to : we doesn't need to judge whether the tree is nil beacuse the system can help me to judge
	node.left.trave()
	fmt.Print(" ", node.value)
	node.right.trave()
}

func (node *treeNode) setValue(value int) {
	node.value = value
}

func main() {
	branch1 := treeNode{3, nil, nil}
	branch2 := treeNode{4, nil, nil}
	root := treeNode{value: 3, left: &branch1, right: &branch2}
	branch1.left = &root
	branch2.left = &root

	nodes := []treeNode{
		{3, nil, nil},
		{},
		{value: 3},
	}

	fmt.Println(root)
	fmt.Println(branch1)
	fmt.Println(branch2)
	fmt.Println(nodes)

	//text
	root.setValues(10)
	root.print()
	root.setValue(20)
	root.print()

	fmt.Println("******************华丽的分割线********************")

	roots := treeNode{value: 3}
	roots.left = &treeNode{value: 5}
	roots.right = &treeNode{value: 4}
	roots.left.right = &treeNode{value: 6}
	roots.right.left = &treeNode{value: 7}
	roots.trave()
}
