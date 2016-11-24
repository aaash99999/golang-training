package main

import "fmt"

// makeGreeter() is a function which returns a function, and that function
// returns a string.
func makeGreeter() func() string {
	return func() string {
		return "Hello world!"
	}
}

func main() {
	greet := makeGreeter()
	fmt.Println(greet())
	fmt.Printf("%T\n", greet)
	// sample output
	/*
		Hello world!
		func() string
	*/
}
