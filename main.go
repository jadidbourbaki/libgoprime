package main

import (
	"fmt"
	"math/big"
)

const maxPrime = 65537

func main() {
	for i := 3; i <= maxPrime; i += 2 {
		num := big.NewInt(int64(i))

		if num.ProbablyPrime(0) {
			fmt.Printf("Found new prime: %s.\n", num.String())
		}
	}
}
