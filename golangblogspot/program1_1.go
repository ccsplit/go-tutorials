package main

import (
	"bufio"
	"fmt"
	"os"
	s "strings"
)

// A very simple program that will retrieve the user's name from stdin
// and print it back to them as a Hello, <Person>.
func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("What is your name?")
	n, _ := reader.ReadString('\n')
	n = s.Trim(n, "\n")
	fmt.Printf("Hello, %s.\n", n)
}
