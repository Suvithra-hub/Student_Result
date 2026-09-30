package main

import (
	"fmt"
	"student-result-cli/calculator"
	"student-result-cli/grade"
)

func main() {
	var name string
	var regno string
	var m1, m2, m3 int

	fmt.Println("Student Result Processing System")

	fmt.Print("Enter Name: ")
	fmt.Scan(&name)

	fmt.Print("Enter Register No: ")
	fmt.Scan(&regno)

	fmt.Print("Enter Mark 1: ")
	fmt.Scan(&m1)

	fmt.Print("Enter Mark 2: ")
	fmt.Scan(&m2)

	fmt.Print("Enter Mark 3: ")
	fmt.Scan(&m3)

	total := calculator.Total(m1, m2, m3)
	average := total / 3

	studentGrade := grade.GetGrade(average)

	fmt.Println("\n----- RESULT -----")
	fmt.Println("Name:", name)
	fmt.Println("Register No:", regno)
	fmt.Println("Total:", total)
	fmt.Println("Average:", average)
	fmt.Println("Grade:", studentGrade)
}