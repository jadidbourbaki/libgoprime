package libgoprime

import (
	"math/big"
	"testing"
)

func TestGcdTrivial(t *testing.T) {
	a := new(big.Int)
	a.SetInt64(0)

	b := new(big.Int)
	b.SetInt64(0)

	mygcd := Gcd(a, b)

	zero := new(big.Int)
	zero.SetInt64(0)

	if mygcd.Cmp(zero) != 0 {
		t.Errorf("failed")
	}
}

func TestGcdBigNumber(t *testing.T) {
	a := new(big.Int)
	a.SetString("4242424242", 10)

	b := new(big.Int)
	b.SetString("462662115", 10)

	mygcd := Gcd(a, b)

	result := new(big.Int)
	result.SetInt64(3)

	if mygcd.Cmp(result) != 0 {
		t.Errorf("failed")
	}
}

func TestLcmTrivial(t *testing.T) {
	a := new(big.Int)
	a.SetInt64(0)

	b := new(big.Int)
	b.SetInt64(0)

	mylcm := Lcm(a, b)

	zero := new(big.Int)
	zero.SetInt64(0)

	if mylcm.Cmp(zero) != 0 {
		t.Errorf("failed")
	}

}

func TestLcmBigNumber(t *testing.T) {
	a := new(big.Int)
	a.SetString("4242424242", 10)

	b := new(big.Int)
	b.SetString("462662115", 10)

	mylcm := Lcm(a, b)

	result := new(big.Int)
	result.SetString("654269657510330610", 10)

	if mylcm.Cmp(result) != 0 {
		t.Errorf("failed")
	}

}

func TestPhiTrivial(t *testing.T) {
	n := new(big.Int)
	n.SetInt64(0)

	myphi := Phi(n)

	one := new(big.Int)
	one.SetInt64(1)

	if myphi.Cmp(one) != 0 {
		t.Errorf("failed")
	}
}

func TestPhiBigNumber(t *testing.T) {
	n := new(big.Int)
	n.SetString("155215", 10)

	myphi := Phi(n)

	result := new(big.Int)
	result.SetString("120672", 10)

	if myphi.Cmp(result) != 0 {
		t.Errorf("failed")
	}
}
