package main

import "fmt"

type node struct {
	data  int
	left  *node
	right *node
}

func insert(value int, root *node) *node {
	if root == nil {
		root = new(node)
		root.data = value
	} else {
		if root.data >= value {
			root.left = insert(value, root.left)
		} else {
			root.right = insert(value, root.right)
		}
	}
	return root
}

func traverse(root *node) {
	if root == nil {
		fmt.Printf("%s", " ")
		return
	} else {
		traverse(root.left)
		fmt.Printf("%d", root.data)
		traverse(root.right)
	}
}

func main() {
	var root *node
	root = insert(4, root)
	root = insert(6, root)
	root = insert(3, root)
	root = insert(0, root)
	root = insert(5, root)
	root = insert(8, root)
	traverse(root)
}
