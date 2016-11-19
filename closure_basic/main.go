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

	fmt.Println(increment())
	fmt.Println(increment())
}
