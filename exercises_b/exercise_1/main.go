package main

import "fmt"

func foo(n int) (int, bool) {
	return n / 2, n%2 == 0
}

func main() {
	fmt.Println(foo(5))
	fmt.Println(foo(6))
}
