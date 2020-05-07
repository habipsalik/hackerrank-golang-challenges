package main

import (
	"bufio"
	"fmt"
	"os"
	"reflect"

	"strconv"
)

type advancedArithmetic interface {
	divisorSum() int
}

type calculator struct {
	num int
}

func (c *calculator) divisorSum() int {
	var sum int
	for i := 1; i <= c.num; i++ {
		if c.num%i == 0 {
			sum = i + sum
		}
	}

	return sum
}

func solution(arithmetic advancedArithmetic) {
	str := reflect.TypeOf(&arithmetic)
	fmt.Printf("I implemented: %s\n", str)
	fmt.Println(arithmetic.divisorSum())
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()

	n, _ := strconv.Atoi(scanner.Text())
	calculator := &calculator{num: n}
	solution(calculator)
}
