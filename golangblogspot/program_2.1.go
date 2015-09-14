package main

import "fmt"

// This is very similar to the defaults within the tutorial. However,
// I also wanted to add what the values were within the different
// formats so I could get used to the way it is printed.
func main() {
	var i int
	fmt.Println("default int is: ", i)
	fmt.Printf("As hex: %x\n", i)
	fmt.Printf("As binary: %b\n", i)
	fmt.Printf("As octal: %o\n", i)
	fmt.Printf("As Unicode: %U\n", i)
	fmt.Printf("Address of i: 0x%X\n", &i)
	var s string
	fmt.Println("default string is: ", s)
	fmt.Printf("Address of s: 0x%X\n", &s)
	var f float64
	fmt.Println("default float64 is: ", f)
	fmt.Printf("As scientific notation: %e\n", f)
	fmt.Printf("Using %%b: %b\n", f)
	fmt.Printf("Address of f: 0x%X\n", &f)
	var arInt [3]int
	fmt.Println("default int array is: ", arInt)
	fmt.Printf("As Unicode: %U\n", arInt)
	var c complex64
	fmt.Println("default complex64 is: ", c)
	fmt.Printf("Address of c: 0x%X\n", &c)
}
