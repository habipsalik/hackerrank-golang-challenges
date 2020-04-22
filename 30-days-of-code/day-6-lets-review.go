package main

import (
	"bytes"
	"fmt"
	"strings"
)

func main() {
	var t int
	var s string
	_, _ = fmt.Scanf("%d", &t)

	for i := 0; i < t; i++ {
		_, _ = fmt.Scanf("%v", &s)
		even, odd := splitToString(s)
		fmt.Println(even, odd)
	}
}

func splitToString(a string) (string, string) {
	var evenBuffer bytes.Buffer
	var oddBuffer bytes.Buffer

	str := strings.Split(a, "")
	for i := 0; i < len(str); i++ {
		if i%2 == 0 {
			evenBuffer.WriteString(str[i])
		} else {
			oddBuffer.WriteString(str[i])
		}
	}

	return evenBuffer.String(), oddBuffer.String()
}
