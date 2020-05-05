package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {

	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	input := scanner.Text()

	converted, err := strconv.Atoi(input)
	if err != nil {
		fmt.Println("Bad String")
	} else {
		fmt.Println(converted)
	}
}
