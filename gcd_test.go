package main

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
