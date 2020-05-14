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
	actualDate := strings.Split(scanner.Text(), " ")
	actualDay, err := strconv.Atoi(actualDate[0])
	checkError(err)
	actualMonth, err := strconv.Atoi(actualDate[1])
	checkError(err)
	actualYear, err := strconv.Atoi(actualDate[2])
	checkError(err)

	scanner.Scan()
	expectedDate := strings.Split(scanner.Text(), " ")
	expectedDay, err := strconv.Atoi(expectedDate[0])
	checkError(err)
	expectedMonth, err := strconv.Atoi(expectedDate[1])
	checkError(err)
	expectedYear, err := strconv.Atoi(expectedDate[2])
	checkError(err)

	if expectedYear < actualYear {
		fmt.Println(10000)
	} else if expectedMonth < actualMonth && expectedYear == actualYear {
		diffMonth := actualMonth - expectedMonth
		fmt.Println(int(diffMonth) * 500)
	} else if expectedDay < actualDay && expectedMonth == actualMonth {
		diffDay := actualDay - expectedDay
		fmt.Println(diffDay * 15)
	} else {
		fmt.Println(0)
	}
}

func checkError(err error) {
	if err != nil {
		panic(err)
	}
}
