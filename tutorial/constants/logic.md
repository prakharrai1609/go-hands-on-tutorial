**Go supports constants of character, string, boolean, and numeric values.**

```go
package main

import "fmt"

func main() {
    // This is how you declare a constant in Go.
    const pi = 3.14159

    // Constants can't be changed once set.
    // Uncommenting the line below will result in an error.
    // pi = 3.14

    fmt.Println("The value of pi is:", pi)

    // You can also declare multiple constants together.
    const (
        daysInWeek = 7
        hoursInDay = 24
    )

    fmt.Println("There are", daysInWeek, "days in a week and", hoursInDay, "hours in a day.")
}
```

In this example, we declare a constant named `pi` with the value `3.14159`. Once a constant is defined, you can't change its value, as shown by the commented line that would produce an error.

Then we declare two more constants, `daysInWeek` and `hoursInDay`, using a technique called "constant grouping." These constants hold the values `7` and `24` respectively.

By using constants, you make your code more readable and maintainable because you give meaningful names to values that won't change.
