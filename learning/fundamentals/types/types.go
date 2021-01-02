package main

import (
	"fmt"
	"math"
	"reflect"
)

func main() {
	// integer numbers
	fmt.Println(1, 2, 1000)
	fmt.Println("Literal integer is", reflect.TypeOf(32000))

	// unsigned uint8(btye) uint16(short - 2 btyes) uint32(int - 4 btyes) uint64(long - 8 btyes)

	var b byte = 255

	fmt.Println("The byte is", reflect.TypeOf(b))

	// With signed int8 int16 int32 int64

	// This will depend based on our architecture. for 64 bits > int64, for 32 bits > int32
	i1 := math.MaxFloat64
	fmt.Println("The maximum value of an int is", i1)
	fmt.Println("The type of i1 is", reflect.TypeOf(i1))

	var i2 rune = 'a' // represents a mapping of Unicode table (int32)
	fmt.Println("The rune is", reflect.TypeOf(i2))
	fmt.Println(i2)

	// Real numbers (float32, float64)
	var x float32 = 49.99 // we can also use var x = float32(49.99)
	fmt.Println("The type of x is", reflect.TypeOf(x))
	fmt.Println("The literal value of x is", reflect.TypeOf(49.9))

	// boolean
	bo := true
	fmt.Println("The type of bo is", reflect.TypeOf(bo))
	fmt.Println(!bo)

	// string
	s1 := "Hi, my name is Bruno"
	fmt.Println(s1 + "!")
	fmt.Println("The lenght of s1 is", len(s1))

	// string with multiple lines
	s2 := `Hi
	my
	name
	is
	Bruno`
	fmt.Println("The lenght of s2 is", len(s2))

	// char???
	// var x char = 'b' - no, char does not exist - it will assume as int32
	char := 'a'
	fmt.Println("The of char is", reflect.TypeOf(char))
	fmt.Println(char) // 97 (unicode)
}
