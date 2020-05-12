package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
)

type Node struct {
	data  int
	left  *Node
	right *Node
}

func insert(node *Node, newData int) *Node {
	if node == nil {
		node = &Node{data: newData}
	} else {
		if node.data <= newData {
			node.right = insert(node.right, newData)
		} else {
			node.left = insert(node.left, newData)
		}
	}

	return node
}

func getHeight(node *Node) float64 {
	if node == nil {
		return 0
	} else if node.right == nil && node.left == nil {
		return math.Max(getHeight(node.left), getHeight(node.right)) + 0
	} else {
		return math.Max(getHeight(node.left), getHeight(node.right)) + 1
	}
}

func main() {

	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	_ = scanner.Text() //  This input of test cases. The input not used.

	tree := new(Node)
	for scanner.Scan() {
		if scanner.Text() == "" {
			break
		}
		num, err := strconv.Atoi(scanner.Text())
		checkError(err)
		insert(tree, num)
	}

	fmt.Println(int(getHeight(tree)) - 1)
}

func checkError(err error) {
	if err != nil {
		panic(err)
	}
}
