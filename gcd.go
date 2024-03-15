package main

import "math/big"

func Gcd(a, b *big.Int) *big.Int {
	zero := new(big.Int)
	zero.SetInt64(0)

	if b.Cmp(zero) == 0 {
		return a
	}

	return Gcd(b, new(big.Int).Mod(a, b))
}
