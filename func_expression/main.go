package main

import "fmt"

func main() {
	greeting := func() {
		// this is an anonymous function, and we are assigning it to greeting
		fmt.Println("Hi!")
	}

	greeting()
}
