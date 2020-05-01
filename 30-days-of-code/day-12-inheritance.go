package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Person struct {
	ID        int
	FirstName string
	LastName  string
}

type Student struct {
	Person
	Scores []int
}

func NewStudent() *Student {
	student := new(Student)
	return student
}

func (s *Student) Info() string {
	info := "Name: " + s.LastName + "," + s.FirstName + "\n" + "ID: " + strconv.Itoa(s.ID)
	return info
}

func (s *Student) Calculate() string {
	var sum int
	for _, score := range s.Scores {
		sum += score
	}

	avg := sum / len(s.Scores)

	var grade string
	switch {
	case avg >= 90 && avg <= 100:
		grade = "O"
	case avg >= 80 && avg < 90:
		grade = "E"
	case avg >= 70 && avg < 80:
		grade = "A"
	case avg >= 55 && avg < 70:
		grade = "P"
	case avg >= 40 && avg < 55:
		grade = "D"
	default:
		grade = "T"
	}

	return grade
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	scanner.Scan()
	InfoInput := strings.Split(scanner.Text(), " ")

	FirstName := InfoInput[0]
	LastName := InfoInput[1]
	Id := InfoInput[2]

	scanner.Scan()
	ScoreCount, err := strconv.Atoi(scanner.Text())
	CheckError(err)

	scanner.Scan()
	ScoresInput := strings.Split(scanner.Text(), " ")

	var ScoreArr []int
	for i := 0; i < ScoreCount; i++ {
		score, err := strconv.Atoi(ScoresInput[i])
		CheckError(err)
		ScoreArr = append(ScoreArr, score)
	}

	student := NewStudent()
	student.FirstName = FirstName
	student.LastName = LastName
	student.ID, err = strconv.Atoi(Id)
	CheckError(err)
	student.Scores = ScoreArr

	info := student.Info()
	grade := student.Calculate()

	fmt.Println(info)
	fmt.Println("Grade :", grade)
}

func CheckError(err error) {
	if err != nil {
		panic(err)
	}
}
