package main

import (
	"fmt"
	"strconv"
)

func main() {
	x := 2.4
	y := 2
	fmt.Println(x / float64(y))
	fmt.Println(int(x) / y)

	assign := 6.9
	notaFinal := int(assign)
	fmt.Println(notaFinal)

	// Be careful about this. This will not convert 123 into a string. 123 -> return the ASCII value
	fmt.Println("Teste " + string(123))

	// int to string
	fmt.Println("Teste " + strconv.Itoa(123))
	fmt.Println("Teste " + fmt.Sprint(123))

	// string to int
	num, _ := strconv.Atoi("123")
	fmt.Println(num - 122)

	// string to boolean
	b, _ := strconv.ParseBool("true")

	if b {
		fmt.Println("True")
	}
}
