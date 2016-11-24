package main

import "fmt"

func biggest(numbers ...int) int {
	winner := 0
	for _, num := range numbers {
		if num > winner {
			winner = num
		}
	}
	return winner
}

func main() {
	result := biggest(1, 7, 3, 4)
	fmt.Println("biggest number is: ", result)
}
