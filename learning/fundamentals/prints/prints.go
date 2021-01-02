package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Print("Same")
	fmt.Print(" line.")

	fmt.Println(" New")
	fmt.Println("line.")

	x := math.Pi

	// Will not work
	// fmt.Println("The value of x is" + x)

	xs := fmt.Sprint(x)
	fmt.Println("The value of x is " + xs)

	// or
	fmt.Println("The value of x is", x)
	fmt.Printf("The alue of x is %.2f", x)

	a := 1
	b := 1.9999
	c := false
	d := "hey!"
	/*
		%d = integer
		%f = float
		%t = boolean
		%s = string
		%v = default format
	*/
	fmt.Printf("\n%d %f %.1f %t %s", a, b, b, c, d)
	fmt.Printf("\n%v %v %v %v", a, b, c, d)
}
