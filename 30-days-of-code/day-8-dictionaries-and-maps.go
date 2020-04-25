package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	//Enter your code here. Read input from STDIN. Print output to STDOUT
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	n, _ := strconv.Atoi(scanner.Text())

	phoneBook := make(map[string]int)

	for i := 0; i < n; i++ {
		scanner.Scan()
		inputArray := strings.Split(scanner.Text(), " ")
		key := inputArray[0]
		value, _ := strconv.Atoi(inputArray[1])
		phoneBook[key] = value
	}

	for scanner.Scan() {
		input := scanner.Text()
		if value, ok := phoneBook[input]; ok {
			result := input + "=" + strconv.Itoa(value)
			fmt.Println(result)
		} else {
			fmt.Println("Not found")
		}
	}
}
