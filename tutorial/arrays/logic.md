Arrays in Go are used to store a fixed number of elements of the same type. They provide a way to group related data together. Here's an explanation with examples:

**Declaring and Initializing Arrays:**

You declare an array by specifying the type of its elements and the number of elements it can hold. Here's how you declare and initialize an array:

```go
package main

import "fmt"

func main() {
    // Declare an array of integers with a length of 5
    var numbers [5]int

    // Initialize elements of the array
    numbers[0] = 10
    numbers[1] = 20
    numbers[2] = 30
    numbers[3] = 40
    numbers[4] = 50

    // Access and print individual elements
    fmt.Println("Third element:", numbers[2])

    // Declare and initialize an array in one line
    fruits := [3]string{"apple", "banana", "cherry"}
    fmt.Println("Second fruit:", fruits[1])
}
```

**Iterating Over Arrays:**

You can use loops to iterate through the elements of an array:

```go
package main

import "fmt"

func main() {
    temperatures := [7]float64{23.5, 24.0, 22.8, 25.5, 26.2, 21.7, 20.5}

    // Iterate using a for loop
    for i := 0; i < len(temperatures); i++ {
        fmt.Printf("Day %d temperature: %.1f\n", i+1, temperatures[i])
    }

    // Iterate using range (index and value)
    for index, value := range temperatures {
        fmt.Printf("Day %d temperature: %.1f\n", index+1, value)
    }
}
```

**Array Length:**

The `len()` function can be used to get the length of an array:

```go
package main

import "fmt"

func main() {
    scores := [4]int{98, 87, 92, 78}

    fmt.Println("Number of scores:", len(scores))
}
```

**Arrays are Value Types:**

Arrays in Go are value types. When you assign an array to another variable or pass it as a parameter, a copy of the array is made. Changes to the copy won't affect the original array.

```go
package main

import "fmt"

func main() {
    original := [3]int{1, 2, 3}
    copy := original

    copy[0] = 100

    fmt.Println("Original:", original) // [1 2 3]
    fmt.Println("Copy:", copy)         // [100 2 3]
}
```

Arrays are useful for storing a fixed set of elements, but they have a fixed size that cannot be changed once declared. If you need more flexibility in size, consider using slices, which are more dynamic.