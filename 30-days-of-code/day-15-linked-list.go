package main

import (
	"fmt"
)

type Node struct {
	data int
	next *Node
}

func main() {

	var t, data int
	var node *Node

	_, err := fmt.Scan(&t)
	checkError(err)

	for i := 0; i < t; i++ {
		_, err = fmt.Scan(&data)
		checkError(err)
		node = insertNode(node, data)
	}

	display(node)
}

func display(node *Node) {
	for node != nil {
		fmt.Printf("%v ", node.data)
		node = node.next
	}
}

func insertNode(node *Node, newData int) *Node {
	if node == nil {
		node = &Node{data: newData}
	} else {
		node.next = insertNode(node.next, newData)
	}

	return node
}

func checkError(err error) {
	if err != nil {
		panic(err)
	}
}
