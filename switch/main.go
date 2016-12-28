package main

import (
	"fmt"
	"unicode"
)

func main() {
	fmt.Println("Hello, playground")
	fmt.Println(unicode.IsLetter('x'))

	for i, rune := range "Hello, 世界 -_- 45 ." {
		fmt.Printf("%d: %c\n", i, rune)
		switch rune {
		case 'a':
			fmt.Printf("%c is an a\n", rune)

		case 'e':
			fmt.Printf("%c is an e\n", rune)

		case 'l':
			fmt.Printf("%c is an l\n", rune)

		case ' ':
			fmt.Printf("%c is a space\n", rune)

		default:
			fmt.Printf("%c is something else!\n", rune)
		}
	}

}
