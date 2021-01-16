package main

import (
	"fmt"
	"time"
)

func assignToConcept(assign float64) string {
	var assignInt = int(assign)

	switch assignInt { // or we can use switch only. logic condition required case
	case 20:
		fallthrough
	case 18:
		return "A"
	case 17, 16:
		return "B"
	case 15, 13:
		return "C"
	case 10, 12:
		return "D"
	default:
		return "E"
	}
}

func typeSwitch(i interface{}) string {
	switch i.(type) {
	case int:
		return "integer"
	case float32, float64:
		return "real"
	case string:
		return "string"
	case func():
		return "function"
	default:
		return "undefined"
	}
}

func main() {
	fmt.Println(assignToConcept(20))
	fmt.Println(assignToConcept(18))
	fmt.Println(assignToConcept(17))
	fmt.Println(assignToConcept(16))

	fmt.Println(typeSwitch(2.3))
	fmt.Println(typeSwitch(1))
	fmt.Println(typeSwitch("Hi"))
	fmt.Println(typeSwitch(func() {}))
	fmt.Println(typeSwitch(time.Now()))
}
