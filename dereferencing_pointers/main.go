package main

import "fmt"

func zero(z *int) {
	fmt.Printf("You called zero(), and z is %p, and *z is %d\n", z, *z)
	*z = 0
	fmt.Printf("Just set it to zero.  Proof: *z is: %d\n", *z)

}

func main() {
	x := 5
	fmt.Printf("x is: %d\n", x)
	fmt.Printf("&x is: %p\n", &x)
	zero(&x)
	fmt.Printf("x is: %d\n", x)
	fmt.Printf("&x is: %p\n", &x)

	/* sample output:
	x is: 5
	&x is: 0x10998128
	You called zero(), and z is 0x10998128, and *z is 5
	Just set it to zero.  Proof: *z is: 0
	x is: 0
	&x is: 0x10998128
	*/
}
