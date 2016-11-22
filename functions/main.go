package main

import "fmt"

func main() {
	// one way is to call average() with a list of float64's
	// n := average(43, 56, 87, 12, 45, 57)

	// another way is to call average() with a slice containing float64's, but
	// then you need to feed them in one at a time.  Like this:
	data := []float64{43, 56, 87, 12, 45, 57}
	n := average(data...) // unroll? the slice
	fmt.Println(n)
}

func average(sf ...float64) float64 {
	// variadic - i.e. takes 0 or more float64's, returns a single float64
	//
	// another way to declare this function would have been:
	// func average (sf []float64) float64
	// which would take a slice of float64's, and return a single float64
	fmt.Println(sf)
	fmt.Printf("%T \n", sf)
	total := 0.0

	for _, v := range sf {
		// i.e. loop through all the numbers in sf
		total += v
	}

	return total / float64(len(sf)) // i.e. the total divided by the number of items in sf
}
