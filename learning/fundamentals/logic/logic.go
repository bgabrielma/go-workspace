package main

import "fmt"

func shop(work1, work2 bool) (bool, bool, bool) {
	buyTv50 := work1 && work2     // AND
	buyTv32 := work1 != work2     // XOR
	buyIceceram := work1 || work2 // OR

	return buyTv50, buyTv32, buyIceceram
}

func main() {
	tv50, tv32, iceCream := shop(true, true)
	fmt.Printf("Tv50: %t, Tv32: %t, IceCream: %t, Healthy: %t", tv50, tv32, iceCream, !iceCream)
}
