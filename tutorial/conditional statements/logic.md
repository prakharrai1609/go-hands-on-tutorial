There are two main types of conditional statements in Go: `if` statements and `switch` statements.

1. **`if` Statements:**

The `if` statement is used to execute a block of code only if a specified condition is true.

```go
package main

import "fmt"

func main() {
    age := 18

    if age >= 18 {
        fmt.Println("You are an adult.")
    } else {
        fmt.Println("You are not an adult.")
    }
}
```

In this example, if the `age` is greater than or equal to 18, it prints "You are an adult." Otherwise, it prints "You are not an adult."

2. **`if` Statement with Short Initialization:**

You can also include a short initialization statement before the condition in an `if` statement.

```go
package main

import "fmt"

func main() {
    if num := 10; num%2 == 0 {
        fmt.Println("The number is even.")
    } else {
        fmt.Println("The number is odd.")
    }
}
```

Here, the `num` variable is initialized within the `if` statement, and it checks if the number is even or odd.

3. **`switch` Statements:**

The `switch` statement is used to evaluate different possible values of an expression and execute corresponding code blocks.

```go
package main

import "fmt"

func main() {
    day := "Wednesday"

    switch day {
    case "Monday":
        fmt.Println("It's the start of the week.")
    case "Wednesday":
        fmt.Println("It's the middle of the week.")
    default:
        fmt.Println("It's another day of the week.")
    }
}
```

Here, based on the value of `day`, the `switch` statement executes the block of code that matches the corresponding case.

These conditional statements are essential for making your programs dynamic and responsive to different situations. They allow your code to take different paths depending on the conditions you set.

**_Let's move onto [Arrays in GO](../arrays/logic.md)_**