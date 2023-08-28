package main

import "fmt"

func main() {
	// Concatenating two strings
	greeting := "Hello, "
	language := "Go"
	fmt.Println(greeting + language)

	// Performing addition of two integers and printing the result
	num1 := 1
	num2 := 1
	sum := num1 + num2
	fmt.Println(num1, "+", num2, "=", sum)

	// Performing division of two floating-point numbers and printing the result
	dividend := 7.0
	divisor := 3.0
	quotient := dividend / divisor
	fmt.Println(dividend, "/", divisor, "=", quotient)

	// Using logical AND to check if both conditions are true
	condition1 := true
	condition2 := false
	andResult := condition1 && condition2
	fmt.Println(condition1, "AND", condition2, "=", andResult)

	// Using logical OR to check if at least one condition is true
	orResult := condition1 || condition2
	fmt.Println(condition1, "OR", condition2, "=", orResult)

	// Using logical NOT to negate a condition
	notResult := !condition1
	fmt.Println("NOT", condition1, "=", notResult)
}

/*
$ go run tutorial/data\ types/main.go

Hello, Go
1 + 1 = 2
7 / 3 = 2.3333333333333335
true AND false = false
true OR false = true
NOT true = false
*/
