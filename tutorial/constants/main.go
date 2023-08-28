package main

import "fmt"

func main() {
	const number1 int = 1000
	fmt.Println("number1 is ", number1)

	const str string = "This is a constant string"
	fmt.Println("str is ", str)

	// number1 = 10 // this will throw error! constant values cannot be reassigned!
}

/*
$ go run tutorial/constants/main.go

number1 is  1000
str is  This is a constant string
*/
