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

func Lcm(a, b *big.Int) *big.Int {

	gcd := Gcd(a, b)

	zero := new(big.Int)
	zero.SetInt64(0)

	if gcd.Cmp(zero) == 0 {
		return zero
	}

	lcm := new(big.Int)
	lcm.Mul(a, b)
	lcm.Div(lcm, gcd)
	return lcm
}
