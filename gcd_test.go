package libgoprime

import (
	"math/big"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGcdTrivial(t *testing.T) {
	a := new(big.Int)
	a.SetInt64(0)

	b := new(big.Int)
	b.SetInt64(0)

	mygcd := Gcd(a, b)

	zero := new(big.Int)
	zero.SetInt64(0)

	assert.Equal(t, mygcd, zero)
}

func TestGcdBigNumber(t *testing.T) {
	a := new(big.Int)
	a.SetString("4242424242", 10)

	b := new(big.Int)
	b.SetString("462662115", 10)

	mygcd := Gcd(a, b)

	result := new(big.Int)
	result.SetInt64(3)

	assert.Equal(t, mygcd, result)
}

func TestLcmTrivial(t *testing.T) {
	a := new(big.Int)
	a.SetInt64(0)

	b := new(big.Int)
	b.SetInt64(0)

	mylcm := Lcm(a, b)

	zero := new(big.Int)
	zero.SetInt64(0)

	assert.Equal(t, mylcm, zero)
}

func TestLcmBigNumber(t *testing.T) {
	a := new(big.Int)
	a.SetString("4242424242", 10)

	b := new(big.Int)
	b.SetString("462662115", 10)

	mylcm := Lcm(a, b)

	result := new(big.Int)
	result.SetString("654269657510330610", 10)

	assert.Equal(t, mylcm, result)

}
