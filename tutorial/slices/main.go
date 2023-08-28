package main

import "fmt"

func main() {
	// Creating a slice using a composite literal.
	numbers := []int{1, 2, 3, 4, 5}
	fmt.Println("Before appending:", numbers)

	// Appending elements to the slice.
	numbers = append(numbers, 6)
	fmt.Println("After appending:", numbers)

	// Slicing the slice to create a new slice.
	slice := numbers[1:4]
	fmt.Println("Sliced slice:", slice)

	// Copying a slice into a new slice.
	copied := make([]int, len(slice))
	copy(copied, slice)
	fmt.Println("Copied slice:", copied)

	// Length and capacity of a slice.
	fmt.Println("Length:", len(numbers), "Capacity:", cap(numbers))

	// Iterating through a slice using a range loop.
	for index, value := range numbers {
		fmt.Printf("Index: %d, Value: %d\n", index, value)
	}

	// Creating a multi-dimensional slice.
	matrix := [][]int{
		{1, 2, 3},
		{4, 5, 6},
		{7, 8, 9},
	}
	fmt.Println("Multi-dimensional slice:", matrix)
}

/*
$ go run tutorial/slices/main.go

Before appending: [1 2 3 4 5]
After appending: [1 2 3 4 5 6]
Sliced slice: [2 3 4]
Copied slice: [2 3 4]
Length: 6 Capacity: 10
Index: 0, Value: 1
Index: 1, Value: 2
Index: 2, Value: 3
Index: 3, Value: 4
Index: 4, Value: 5
Index: 5, Value: 6
Multi-dimensional slice: [[1 2 3] [4 5 6] [7 8 9]]
*/
