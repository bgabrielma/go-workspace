package main

// Alias
// m "math"
import (
	"fmt"
	"math"
)

func main() {
	// By default, a float uses float64
	const PI float64 = math.Pi
	var raio = 3.2 // tipo (float64) infered by compiler

	// another way to create and assign a var
	// Initializer
	area := PI * math.Pow(raio, 2)
	fmt.Println("Area of this circle is", area)

	const (
		a = 1
		b = 2
	)

	var (
		c = 3
		d = 4
	)
	fmt.Println(a, b, c, d)

	// Inline variable with initializer
	var e, f bool = true, false
	fmt.Println(e, f)

	g, h, i := 2, false, "nice!"
	fmt.Println(g, h, i)
}
