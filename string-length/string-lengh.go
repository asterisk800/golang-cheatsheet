package main

import (
	"fmt"
	"unicode/utf8"
)

func main() {
	name := "ጄምስ"
	fmt.Println(name)

	// the len function print the number of byes, not the number of runes (codepint)
	fmt.Println(len(name))

	// This prints the correct number of charactors in the string literal
	fmt.Println(utf8.RuneCountInString(name))
}
