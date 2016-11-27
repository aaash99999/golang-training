package main

import "fmt"

// create a new type called Square
type Square struct {
	side float64
}

// area is a method of type Square
func (z Square) area() float64 {
	return z.side * z.side
}

// Shape is an interface.  If you fulfil the methods it contains, you can use the interface
type Shape interface {
	area() float64
}

// info is a function that takes a Shape, which is an interface.
// so in order to use info(), we need to pass it a type which
// implements the methods defined in the Shape interface.
// since the Square type has the area() method, it meets the
// requirements of the Shape interface, so we can pass a
// Square to the info() function.
func info(z Shape) {
	fmt.Println(z)
	fmt.Println(z.area())
}

func main() {
	s := Square{10}
	info(s)
}
