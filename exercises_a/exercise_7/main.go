package main

import "fmt"

func main() {
	startnumber := 1
	endnumber := 1000
	total := 0
	for i := startnumber; i < endnumber; i++ {
		if i%3 == 0 || i%5 == 0 {
			total += i
		}
	}
	fmt.Printf("The total is: %d\n", total)
}
