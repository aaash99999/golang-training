package main

import "fmt"

// visit takes two parameters, a slice of ints, and a function (which takes an int)
func visit(numbers []int, callback func(int)) {
	for _, n := range numbers {
		callback(n)
	}
}

func main() {
	// here we call visit with the two parameters it needs:
	// first one is a slice of ints
	// second one is a function (anonymous function in this case) which takes an int
	visit([]int{1, 2, 3, 4}, func(n int) {
		fmt.Println(n)
	})

	// here we call visit again, this time with a different function
	visit([]int{1, 2, 3, 4}, func(j int) {
		fmt.Printf("%d\n", j*2)
	})
}
