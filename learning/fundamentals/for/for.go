package main

import (
	"fmt"
	"time"
)

func main() {
	i := 1
	for i <= 10 { // while doesn't exist in Go
		fmt.Println(1)
		i++

	}

	for i := 0; i <= 2; i = +2 {
		fmt.Printf("%d", i)
	}

	fmt.Println("\nMixing...")
	for i := 1; i <= 10; i++ {
		if i%2 == 0 {
			fmt.Print("Par ")
		} else {
			fmt.Print("Impar ")
		}
	}

	for {
		// infinite loop
		fmt.Println("Forever...")
		time.Sleep(time.Second)
	}
}
