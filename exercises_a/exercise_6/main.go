package main

import "fmt"

func main() {
	startnumber := 1
	endnumber := 100
	printnumber := true
	for i := startnumber; i <= endnumber; i++ {
		printnumber = true
		if i%3 == 0 {
			fmt.Printf("Fizz")
			printnumber = false
		}
		if i%5 == 0 {
			fmt.Printf("Buzz")
			printnumber = false
		}
		if printnumber {
			fmt.Printf("%d", i)
		}

		fmt.Printf("\n")
	}
}
