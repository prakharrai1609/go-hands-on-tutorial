package main

import "fmt"

func main() {

	for j := 7; j <= 9; j++ {
		fmt.Println(j)
	}

	var count int = 0
	for {
		fmt.Printf("This is the infinite loop counter, id %d\n", count)
		count += 1
		if count > 5 {
			break
		}
	}

	m := map[string]int{"foo": 1, "bar": 2}

	for key, value := range m {
		fmt.Println("key: ", key, ", value: ", value)
	}
}

/*
$ go run tutorial/for/main.go

7
8
9
This is the infinite loop counter, id 0
This is the infinite loop counter, id 1
This is the infinite loop counter, id 2
This is the infinite loop counter, id 3
This is the infinite loop counter, id 4
This is the infinite loop counter, id 5
key:  foo , value:  1
key:  bar , value:  2
*/
