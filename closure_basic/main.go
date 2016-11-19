package main

import "fmt"

func main() {
	x := 0
	// increment is assigned a function, and the function returns an int
	// the function is an anonymous function (it has no name).  aka: a closure
	increment := func() int {
		x++
		return x
	}

	fmt.Printf("The type of the increment variable is: %T\n", increment)
	fmt.Println(increment())
	fmt.Println(increment())
}
