package main

import "fmt"

type rectangle struct {
	length, width int
}

type square struct {
	rectangle
}

func (r rectangle) area() int {
	return r.length * r.width
}

func (r rectangle) perimeter() int {
	return 2 * (r.length + r.width)
}

// This is to work with some inheritance and also learn a little about
// structs.
func main() {
	r1 := rectangle{4, 6}
	s1 := square{rectangle{6, 6}}
	fmt.Println("Rectangle is: ", r1)
	fmt.Printf("Rectangle with names: %+v\n", r1)
	fmt.Println("Rectangle area is: ", r1.area())
	fmt.Println("Rectangle perimeter is: ", r1.perimeter())
	fmt.Println("Square is: ", s1)
	fmt.Printf("Square with names: %+v\n", s1)
	fmt.Println("Square area is: ", s1.area())
	fmt.Println("Square perimeter is: ", s1.perimeter())
}
