package main

import "fmt"

func main() {
	var t, n int
	var arr []int

	_, err := fmt.Scan(&t)
	checkError(err)

	for i := 0; i < t; i++ {
		_, err = fmt.Scan(&n)
		checkError(err)
		arr = append(arr, n)
	}

	for _, value := range arr {
		for i := 2; i <= value/i; i++ {
			if value%i == 0 {
				value = 1
			}
		}

		if value == 1 {
			fmt.Println("Not prime")
		} else {
			fmt.Println("Prime")
		}
	}
}

func checkError(err error) {
	if err != nil {
		panic(err)
	}
}
