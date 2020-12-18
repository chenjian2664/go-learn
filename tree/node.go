package main

import "fmt"

type treeNode struct {
	value int
	left, right *treeNode
}
// treeNode 的成员函数
func (node treeNode) print() {
	fmt.Println(node.value)
}
// 改变值的时候必须要传指针，go中所有的参数都是值传递，即会copy一个
func /*(node treeNode)*/ (node *treeNode) setValue(val int) {
	node.value = val
}

func createNode(value int) *treeNode {
	return &treeNode{value: value}
}


func main() {
	var root treeNode
	root = treeNode{value: 3}
	root.left = &treeNode{}
	root.right = &treeNode{5, nil, nil}

	nodes := []treeNode {
		{value:3},
		{},
		{5, nil, &treeNode{}},
	}

	fmt.Println(root)
	fmt.Println(nodes)

	root.print()
}
