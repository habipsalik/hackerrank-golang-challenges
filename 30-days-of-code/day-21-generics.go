package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

type Printer struct {
	GenericParameter interface{}
}

func (p *Printer) printArray() {
	switch v := p.GenericParameter.(type) {
	case []int:
		for _, val := range v {
			fmt.Println(val)
		}
	case []string:
		for _, val := range v {
			fmt.Println(val)
		}
	}
}

func main() {

	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	n, _ := strconv.Atoi(scanner.Text())
	var arrInt []int
	x := 0

	for scanner.Scan() {
		if x == n {
			break
		}
		x++
		i, _ := strconv.Atoi(scanner.Text())
		arrInt = append(arrInt, i)
	}

	m, _ := strconv.Atoi(scanner.Text())
	var arrStr []string
	x = 0

	for scanner.Scan() {
		if x == m {
			break
		}
		x++
		arrStr = append(arrStr, scanner.Text())
	}

	printer := new(Printer)

	printer.GenericParameter = arrInt
	printer.printArray()

	printer.GenericParameter = arrStr
	printer.printArray()
}
