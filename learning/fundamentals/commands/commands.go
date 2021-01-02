package main

import "fmt"

func main() {
	fmt.Printf("Other program using %s!\n", "Go")
}

// Open doc offline mode
// godoc -http:<port>

// Check code status report
// go vet <file>.go

// Go build
// go build <file>.go

// Go build and compile file
// go run <file>.go
