package main

import "fmt"

// wrapper is a function that returns a function, and that function returns an int
func wrapper() func() int {
	x := 0
	// here we create the anonymous function (aka: a closure)
	return func() int {
		x++
		return x
	}
}

func main() {
	// x := 0   // no longer need to do this, x is inside wrapper()

	increment := wrapper()

	fmt.Printf("The type of the increment variable is: %T\n", increment)
	fmt.Println(increment())
	fmt.Println(increment())
}
