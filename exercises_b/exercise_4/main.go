package main

import "fmt"

func main() {
	fmt.Println("The answer is: ", (true && false) || (false && true) || !(false && false))
}
