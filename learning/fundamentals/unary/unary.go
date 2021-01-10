package main

import "fmt"

func main() {
	x := 1
	y := 2

	// only postfix - prefix doesn't exist.
	x++ // x += 1 or x = x + 1
	fmt.Println(x)

	y-- // y -= 1 ou y - 1

	fmt.Println(y)

	// Not accepted
	// fmt.Println(++x == y--)
}
