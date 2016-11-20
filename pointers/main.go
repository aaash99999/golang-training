package main

import "fmt"

func main() {
	a := 43
	fmt.Printf("a has a value of: %d\n", a) // 43
	fmt.Printf("a is type: %T\n", a)

	var b *int = &a
	fmt.Printf("b has a value of: %x\n", b)   // 10bd81b8  (e.g. only)
	fmt.Printf("*b has a value of: %d\n", *b) // 43, since *b means "what's in the box?"
	fmt.Printf("b is type: %T\n", b)
	fmt.Printf("*b is type: %T\n", *b)

	/* sample output:
	a has a value of: 43
	a is type: int
	b has a value of: 10bd81b8
	*b has a value of: 43
	b is type: *int
	*b is type: int
	*/
}
