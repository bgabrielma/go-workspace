package main

import "fmt"

func main() {
	i := 1 // 64 bits - 8 bytes

	// Go doesn't have arithmetic pointers
	var p *int = nil

	// take the corresponded address
	p = &i
	// unreferenced (take the value)
	*p++
	i++

	fmt.Println(p, *p, i, &i)
}
