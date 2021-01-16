package main

import (
	"fmt"
	"math/rand"
	"time"
)

func test(n1, n2 float64) string {
	if n1 > n2 {
		return "N1 is greater then N2"
	} else if n1 != n2 {
		return "N2 is greater then N1"
	} else {
		return "N1 is equal N2"
	}
}

func numRandom() int {
	s := rand.NewSource(time.Now().UnixNano())
	r := rand.New(s)
	return r.Intn(10)
}

func main() {
	fmt.Println(test(2, 1))
	fmt.Println(test(1, 2))
	fmt.Println(test(1, 1))

	if i := numRandom(); i > 5 {
		fmt.Println("Win!")
	} else {
		fmt.Println("Lost!")
	}

	// scope. I is defined in that if body only
	// fmt.Println(i) -> i is undefined
}
