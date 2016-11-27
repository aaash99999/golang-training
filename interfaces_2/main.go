package main

import (
	"fmt"
	"sort"
)

func main() {
	s := []string{"Zeno", "John", "Al", "Jenny"}
	fmt.Println(s)
	//sort.StringSlice(s).Sort()

	// package sort has a Sort funcion.
	// the Sort function takes an Interface type as a parameter
	// The StringSlice type implements the required functions to fulfil the Interface interface,
	// so if we convert our slice of strings to a StringSlice type, we can
	// then pass it to sort.Sort()
	fmt.Printf("The type of sort.StringSlice(s) is: %T\n", sort.StringSlice(s))
	// we could convert it first.  i.e.
	// converted := sort.StringSlice(s)
	// then:
	// sort.Sort(converted)
	// but we can also convert as part of the call
	sort.Sort(sort.StringSlice(s))
	fmt.Println(s)
}
