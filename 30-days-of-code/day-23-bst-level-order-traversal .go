package main

import (
	"fmt"
)

type Node struct {
	data  int
	left  *Node
	right *Node
}

func main() {
	var t, data int
	var node *Node

	_, err := fmt.Scan(&t)
	checkError(err)

	for i := 0; i < t; i++ {
		_, err = fmt.Scan(&data)
		checkError(err)
		node = insert(node, data)
	}

	levelOrder(node)
}

func insert(node *Node, newData int) *Node {
	if node == nil {
		node = &Node{data: newData}
	} else {
		if newData <= node.data {
			node.left = insert(node.left, newData)
		} else {
			node.right = insert(node.right, newData)
		}
	}

	return node
}

func levelOrder(node *Node) {
	var nodeList []*Node
	nodeList = append(nodeList, node)

	for nodeList != nil {
		var level []*Node
		for _, n := range nodeList {
			fmt.Printf("%d ", n.data)
			if n.left != nil {
				level = append(level, n.left)
			}
			if n.right != nil {
				level = append(level, n.right)
			}
			nodeList = level
		}
	}

}

func checkError(err error) {
	if err != nil {
		panic(err)
	}
}
