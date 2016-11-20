package main

import "fmt"

const (
	_  = iota             // the iota starts with 0
	KB = 1 << (iota * 10) // i.e.  1 << (1 * 10) since the iota is now 1
	// the resulting number is: 10000000000
	MB = 1 << (iota * 10) // i.e.  1 << (2 * 10) since the iota is now 2
	// the resulting number is: 100000000000000000000
)

func main() {
	fmt.Println("binary\t\tdecimal")
	fmt.Printf("%b\t", KB)
	fmt.Printf("\t%d\n", KB)
	fmt.Printf("%b\t", MB)
	fmt.Printf("%d\n", MB)
}
