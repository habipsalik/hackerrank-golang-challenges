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
	scanner.Scan()
	arr := strings.Split(scanner.Text(), " ")

	var arrInt []int
	for _, i := range arr {
		num, _ := strconv.Atoi(i)
		arrInt = append(arrInt, num)
	}

	sumSwaps := 0
	for i := 0; i < n; i++ {
		// Track number of elements swapped during a single array traversal
		numberOfSwaps := 0

		for j := 0; j < n-1; j++ {
			// Swap adjacent elements if they are in decreasing order
			if arrInt[j] > arrInt[j+1] {
				arrInt[j], arrInt[j+1] = arrInt[j+1], arrInt[j]
				numberOfSwaps++
			}
		}

		sumSwaps += numberOfSwaps

		// If no elements were swapped during a traversal, array is sorted
		if numberOfSwaps == 0 {
			break
		}
	}

	fmt.Printf("Array is sorted in %v swaps.\n", sumSwaps)
	fmt.Printf("First Element: %v\n", arrInt[0])
	fmt.Printf("Last Element: %v\n", arrInt[len(arrInt)-1])
}
