package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

type Difference struct {
	Elements          []int
	MaximumDifference int
}

func NewDifference() *Difference {
	return new(Difference)
}

func (d *Difference) ComputeDifference() {
	for _, firstNum := range d.Elements {
		for _, secondNum := range d.Elements {
			diff := firstNum - secondNum
			diffAbs := int(math.Abs(float64(diff)))

			if diffAbs > d.MaximumDifference {
				d.MaximumDifference = diffAbs
			}
		}
	}
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	scanner.Scan()
	_ = scanner.Text()

	scanner.Scan()
	inputs := strings.Split(scanner.Text(), " ")

	var numbers []int
	for _, input := range inputs {
		num, err := strconv.Atoi(input)
		checkError(err)
		numbers = append(numbers, num)
	}

	difference := NewDifference()
	difference.Elements = numbers
	difference.ComputeDifference()

	fmt.Println(difference.MaximumDifference)
}

func CheckError(err error) {
	if err != nil {
		panic(err)
	}
}
