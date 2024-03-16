package libgoprime

import "math/big"

func Gcd(a, b *big.Int) *big.Int {
	zero := new(big.Int)
	zero.SetInt64(0)

	for b.Cmp(zero) != 0 {
		t := b
		b = new(big.Int).Mod(a, b)
		a = t
	}

	return a
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
