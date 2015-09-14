package main

import "fmt"

// This is a really simple program that will loop from 0-100 and print
// out the different representations of the integer value.
func main() {
	fmt.Println("Here are the hex, binary, and Unicode representations of 0-100:")
	for i := 0; i <= 100; i++ {
		fmt.Println("i value: ", i)
		fmt.Printf("As hex: %X\n", i)
		fmt.Printf("As binary (base 2): %b\n", i)
		fmt.Printf("As octal: %o\n", i)
		fmt.Printf("As unicode: %U\n", i)
	}
}
