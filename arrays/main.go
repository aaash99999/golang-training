package main

import "fmt"

func main() {
	var x [58]string

	for i := 65; i <= 122; i++ {
		// cast i to a string, and put it into the array of strings i.e. x
		x[i-65] = string(i)
	}

	fmt.Println(x)
	fmt.Println(x[42])
}
