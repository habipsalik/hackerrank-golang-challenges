package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

type Node struct {
	num  int
	next *Node
}

type List struct {
	length int
	start  *Node
}

func main() {
	items := new(List)
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	_ = scanner.Text() //  This input of test cases. The input not used.

	for scanner.Scan() {
		if scanner.Text() == "" {
			break
		}
		val, err := strconv.Atoi(scanner.Text())
		checkError(err)
		node := Node{num: val}
		items.insertNode(&node)
	}
	items.display()
}

func (l *List) display() {
	list := l.start
	for list != nil {
		fmt.Printf("%v ", list.num)
		list = list.next
	}
}

func (l *List) insertNode(newNode *Node) {
	if l.length == 0 {
		l.start = newNode
	} else {
		currentNode := l.start
		for currentNode.next != nil {
			currentNode = currentNode.next
		}
		currentNode.next = newNode
	}

	l.length++
}

func checkError(err error) {
	if err != nil {
		panic(err)
	}
}
