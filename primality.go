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
