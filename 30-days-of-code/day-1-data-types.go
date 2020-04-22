package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	var _ = strconv.Itoa // Ignore this comment. You can still use the package "strconv".

	var i uint64 = 4
	var d float64 = 4.0
	var s string = "HackerRank "

	scanner := bufio.NewScanner(os.Stdin)

	// Declare second integer, double, and String variables.
	var err error
	var inputInt uint64
	var inputDouble float64
	var inputString string
	var inputs []string

	// Read and save an integer, double, and String to your variables.
	for scanner.Scan() {
		inputText := scanner.Text()
		if scanner.Err() != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		if len(inputText) != 0 {
			inputs = append(inputs, inputText)
		} else {
			break
		}
	}

	inputInt, err = strconv.ParseUint(inputs[0], 10, 64)
	inputDouble, err = strconv.ParseFloat(inputs[1], 64)
	inputString = inputs[2]
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}

	// Print the sum of both integer variables on a new line.
	fmt.Println(i + inputInt)

	// Print the sum of the double variables on a new line.
	fmt.Printf("%.1f\n", d+inputDouble)

	// Concatenate and print the String variables on a new line
	// The 's' variable above should be printed first.
	fmt.Println(s + inputString)
}
