package main

import "fmt"

// Ternary operation doesn't exist.
func getResult(nota float64) string {
	if nota >= 6 {
		return "Approved"
	}

	return "Reprovado"
}

func main() {
	fmt.Println(getResult(20))
}
