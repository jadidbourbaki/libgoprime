package libgoprime

import (
	"crypto/rand"
	"math/big"
)

func FermatTest(n *big.Int, k uint) bool {
	if n.Cmp(new(big.Int).SetInt64(0)) == -1 {
		panic("cannot be negative")
	}

	if n.Cmp(new(big.Int).SetInt64(1)) == 0 ||
		n.Cmp(new(big.Int).SetInt64(4)) == 0 {
		return false
	}

	if n.Cmp(new(big.Int).SetInt64(2)) == 0 ||
		n.Cmp(new(big.Int).SetInt64(3)) == 0 {
		return true
	}

	for i := uint(0); i < k; i++ {
		nMinusFour := new(big.Int)
		nMinusFour.Sub(n, new(big.Int).SetInt64(4))

		a, _ := rand.Int(rand.Reader, nMinusFour)
		a.Add(a, new(big.Int).SetInt64(2))

		nMinusOne := new(big.Int)
		nMinusOne.Sub(n, new(big.Int).SetInt64(1))

		result := new(big.Int)
		result.Exp(a, nMinusOne, n)

		if result.Cmp(new(big.Int).SetInt64(1)) != 0 {
			return false
		}

	}

	return true
}

func MillerRabinTest(n *big.Int, k uint) bool {
	if n.Cmp(big.NewInt(2)) == 0 {
		return true
	}

	z := new(big.Int)
	z.Mod(n, big.NewInt(2))

	if z.Cmp(big.NewInt(0)) == 0 {
		return false
	}

	r := big.NewInt(0)
	s := new(big.Int)
	s.Set(n)
	s.Sub(s, big.NewInt(1))

	checkS := new(big.Int)
	checkS.Mod(s, big.NewInt(2))

	for checkS.Cmp(big.NewInt(0)) == 0 {
		r.Add(r, big.NewInt(1))
		s.Div(s, big.NewInt(2))
		checkS.Mod(s, big.NewInt(2))
	}

	for i := uint(0); i < k; i++ {
		nMinusThree := new(big.Int)
		nMinusThree.Sub(n, big.NewInt(3))

		a, _ := rand.Int(rand.Reader, nMinusThree)
		a.Add(a, big.NewInt(2))

		x := new(big.Int)
		x.Exp(a, s, n)

		nMinusOne := new(big.Int)
		nMinusOne.Sub(n, big.NewInt(1))

		if x.Cmp(big.NewInt(1)) == 0 || x.Cmp(nMinusOne) == 0 {
			continue
		}

		rMinusOne := new(big.Int)
		rMinusOne.Sub(r, big.NewInt(1))

		for j := big.NewInt(0); j.Cmp(rMinusOne) == -1; j.Add(j, big.NewInt(1)) {
			x.Exp(x, big.NewInt(2), n)
			if x.Cmp(nMinusOne) == 0 {
				break
			}
		}

		return false
	}

	return true
}
