package main

import "fmt"

func solveMeFirst(a uint32, b uint32) uint32 {
	// Hint: Type return (a+b) below
	return a + b
}

func main() {
	var a, b, result uint32
	_, _ = fmt.Scanf("%d\n%d", &a, &b)
	result = solveMeFirst(a, b)
	fmt.Println(result)
}
