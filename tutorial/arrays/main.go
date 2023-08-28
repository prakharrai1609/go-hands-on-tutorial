package main

import "fmt"

func main() {
	// One-dimensional array
	var numbers [5]int
	numbers[0] = 10
	numbers[1] = 20
	numbers[2] = 30
	numbers[3] = 40
	numbers[4] = 50

	fmt.Println("One-dimensional array:")
	fmt.Println("Third element:", numbers[2])

	// Declare and initialize a two-dimensional array
	var matrix [3][3]int
	matrix[0] = [3]int{1, 2, 3}
	matrix[1] = [3]int{4, 5, 6}
	matrix[2] = [3]int{7, 8, 9}

	fmt.Println("\nTwo-dimensional array:")
	fmt.Println("Second row, third element:", matrix[1][2])

	// Iterate through the two-dimensional array
	fmt.Println("\nMatrix:")
	for row := 0; row < len(matrix); row++ {
		for col := 0; col < len(matrix[row]); col++ {
			fmt.Printf("%d ", matrix[row][col])
		}
		fmt.Println()
	}
}

/*
$ go run tutorial/arrays/main.go

One-dimensional array:
Third element: 30

Two-dimensional array:
Second row, third element: 6

Matrix:
1 2 3
4 5 6
7 8 9
*/
