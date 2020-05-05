package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

// Stack
type Stack struct {
	data []string
}

// Queue
type Queue struct {
	data []string
}

// Stack Constructor
func NewStack() *Stack {
	return new(Stack)
}

// Queue Constructor
func NewQueue() *Queue {
	return new(Queue)
}

// Push Stack Method
func (s *Stack) pushCharacter(item string) {
	s.data = append(s.data, item)
}

// Pop Stack Method
func (s *Stack) popCharacter() string {
	if len(s.data) == 0 {
		return ""
	}
	element := s.data[len(s.data)-1]
	s.data = s.data[:len(s.data)-1]
	return element
}

// Enqueue Method
func (q *Queue) enqueueCharacter(item string) {
	q.data = append(q.data, item)
}

// Dequeue Method
func (q *Queue) dequeueCharacter() string {
	if len(q.data) == 0 {
		return ""
	}
	element := q.data[0]
	q.data = q.data[1:]
	return element
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	str := scanner.Text()

	items := strings.Split(str, "")
	isPalindrome := true
	queue := NewQueue()
	stack := NewStack()

	for _, item := range items {
		stack.pushCharacter(item)
		queue.enqueueCharacter(item)
	}

	for i := 0; i <= len(items)/2; i++ {
		if stack.popCharacter() != queue.dequeueCharacter() {
			isPalindrome = false
			break
		}
	}

	if isPalindrome == true {
		fmt.Println("The word, " + str + ", is a palindrome.")
	} else {
		fmt.Println("The word, " + str + ", is not a palindrome.")
	}
}

// Error Handle
func checkError(err error) {
	if err != nil {
		panic(err)
	}
}
