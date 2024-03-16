package libgoprime

import (
	"math/big"
	"testing"
)

func TestFermatTestTrivialPrimeNumber(t *testing.T) {
	primeNumber := new(big.Int)
	primeNumber.SetInt64(11)

	var k uint = 3

	if !FermatTest(primeNumber, k) {
		t.Errorf("failed")
	}
}

func TestFermatTestTrivialNotPrimeNumber(t *testing.T) {
	notPrimeNumber := new(big.Int)
	notPrimeNumber.SetInt64(15)

	var k uint = 3

	if FermatTest(notPrimeNumber, k) {
		t.Errorf("failed")
	}
}

func TestFermatBigPrimeNumber(t *testing.T) {
	primeNumber := new(big.Int)
	primeNumber.SetString("1417567331", 10)

	var k uint = 10

	if !FermatTest(primeNumber, k) {
		t.Errorf("failed")
	}
}

func TestFermatBigNotPrimeNumber(t *testing.T) {
	notPrimeNumber := new(big.Int)
	notPrimeNumber.SetString("14151519151", 10)

	var k uint = 10

	if FermatTest(notPrimeNumber, k) {
		t.Errorf("failed")
	}
}
