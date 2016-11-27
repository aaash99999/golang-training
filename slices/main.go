package main

import "fmt"

func main() {
	myslice := []int{1, 2, 3, 4}
	fmt.Println("myslice: ", myslice)
	fmt.Println("length of myslice: ", len(myslice))
	fmt.Println("capacity of myslice: ", cap(myslice))

	section := myslice[0:2]
	fmt.Println("section is: ", section)
	fmt.Println("length of section: ", len(section))
	fmt.Println("capacity of section: ", cap(section))

	myslice2 := make([]int, 4, 4) // length 4, capacity 4
	myslice2[2] = 7
	fmt.Println("myslice2 is: ", myslice2)
	fmt.Println("length of myslice2: ", len(myslice2))
	fmt.Println("capacity of myslice2: ", cap(myslice2))

	fmt.Println("appending to myslice2...")
	myslice2 = append(myslice2, 9)
	fmt.Println("myslice2 is: ", myslice2)
	fmt.Println("length of myslice2: ", len(myslice2))
	fmt.Println("capacity of myslice2: ", cap(myslice2))
}
