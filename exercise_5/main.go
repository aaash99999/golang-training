package main

import "fmt"

func main() {
	startnumber := 0
	endnumber := 100
	fmt.Printf("The even numbers between %d and %d are:\n\n", startnumber, endnumber)
	for i := startnumber; i <= endnumber; i++ {
		if i%2 == 0 {
			fmt.Printf("%d is even\n", i)
		}
	}
}
