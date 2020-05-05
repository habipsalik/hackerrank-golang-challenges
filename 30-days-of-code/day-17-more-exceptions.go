package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

func Calculator(n, p int) {
	if n < 0 || p < 0 {
		fmt.Println("n and p should be non-negative")
	} else {
		fmt.Println(math.Pow(float64(n), float64(p)))
	}
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	scanner.Scan()
	_ = scanner.Text() //  This input of test cases. The input not used.

	var inputs [][]string
	for scanner.Scan() {
		if scanner.Text() == "" {
			break
		}

		x := strings.Split(scanner.Text(), " ")
		inputs = append(inputs, x)
	}

	for _, input := range inputs {
		n, err := strconv.Atoi(input[0])
		checkError(err)
		p, err := strconv.Atoi(input[1])
		checkError(err)

		Calculator(n, p)
	}
}

func checkError(err error) {
	if err != nil {
		panic(err)
	}
}
