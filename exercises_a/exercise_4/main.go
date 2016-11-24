package main

import "fmt"

func main() {
	var bignumber int32
	var smallnumber int32
	fmt.Printf("Enter a big number: ")
	fmt.Scanf("%d\n", &bignumber)
	fmt.Printf("Enter a small number: ")
	fmt.Scanf("%d", &smallnumber)
	divisionresult := bignumber / smallnumber
	remainder := bignumber % smallnumber
	fmt.Printf("%d divided by %d is %d, remainder %d\n", bignumber, smallnumber, divisionresult, remainder)
}
