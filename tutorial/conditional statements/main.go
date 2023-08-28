package main

import "fmt"

func main() {
	// Example using if-else if-else
	score := 85

	if score >= 90 {
		fmt.Println("You got an A!")
	} else if score >= 80 {
		fmt.Println("You got a B.")
	} else if score >= 70 {
		fmt.Println("You got a C.")
	} else {
		fmt.Println("You need to improve.")
	}

	// Example using switch
	day := "Tuesday"

	switch day {
	case "Monday":
		fmt.Println("It's the start of the week.")
	case "Tuesday":
		fmt.Println("It's not Monday anymore!")
	case "Friday", "Saturday", "Sunday":
		fmt.Println("It's the weekend!")
	default:
		fmt.Println("It's a regular weekday.")
	}

	// Example using switch with no expression
	age := 25

	switch {
	case age < 18:
		fmt.Println("You're a minor.")
	case age >= 18 && age < 60:
		fmt.Println("You're an adult.")
	default:
		fmt.Println("You're a senior citizen.")
	}
}

/*
$ go run tutorial/conditional\ statements/main.go

You got a B.
It's not Monday anymore!
You're an adult.
*/
