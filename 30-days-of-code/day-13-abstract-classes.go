package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

type Book struct {
	Title  string
	Author string
}

type MyBook struct {
	Book
	Price int
}

func (m *MyBook) Display() {
	display := "Title: " + m.Title + "\n" + "Author: " + m.Author + "\n" + "Price: " + strconv.Itoa(m.Price)
	fmt.Println(display)
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	scanner.Scan()
	TitleInput := scanner.Text()
	scanner.Scan()
	AuthorInput := scanner.Text()
	scanner.Scan()
	PriceInput, err := strconv.Atoi(scanner.Text())
	checkError(err)

	book := new(MyBook)
	book.Title = TitleInput
	book.Author = AuthorInput
	book.Price = PriceInput
	book.Display()
}

func CheckError(err error) {
	if err != nil {
		panic(err)
	}
}
