package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("Strings:", "Test" == "Test")
	fmt.Println("!=", 3 != 2)
	fmt.Println("<", 3 < 2)
	fmt.Println(">", 3 > 2)
	fmt.Println("<=", 3 <= 2)
	fmt.Println(">=", 3 >= 2)

	d1 := time.Unix(0, 0)
	d2 := time.Unix(0, 0)

	fmt.Println("Dates:", d1 == d2)
	fmt.Println("Dates", d1.Equal(d2))

	type Pessoa struct {
		Nome string
	}

	p1 := Pessoa{"Jhon"}
	p2 := Pessoa{"Jhon"}
	fmt.Println("Persons:", p1 == p2)
}
