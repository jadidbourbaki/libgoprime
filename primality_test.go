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

func TestFermatTestBigPrimeNumber(t *testing.T) {
	primeNumber := new(big.Int)
	primeNumber.SetString("1417567331", 10)

	var k uint = 10

	if !FermatTest(primeNumber, k) {
		t.Errorf("failed")
	}
}

func TestFermatTestBigNotPrimeNumber(t *testing.T) {
	notPrimeNumber := new(big.Int)
	notPrimeNumber.SetString("14151519151", 10)

	var k uint = 10

	if FermatTest(notPrimeNumber, k) {
		t.Errorf("failed")
	}
}

func TestMillerRabinTestTrivialPrimeNumber(t *testing.T) {
	primeNumber := new(big.Int)
	primeNumber.SetInt64(11)

	var k uint = 3

	if !MillerRabinTest(primeNumber, k) {
		t.Errorf("failed")
	}
}

func TestMillerRabinTestTrivialNotPrimeNumber(t *testing.T) {
	notPrimeNumber := new(big.Int)
	notPrimeNumber.SetInt64(15)

	var k uint = 3

	if MillerRabinTest(notPrimeNumber, k) {
		t.Errorf("failed")
	}
}

func TestMillerRabinTestBigPrimeNumber(t *testing.T) {
	primeNumber := new(big.Int)
	primeNumber.SetString("1417567331", 10)

	var k uint = 10

	if !MillerRabinTest(primeNumber, k) {
		t.Errorf("failed")
	}
}

func TestMillerRabinTestBigNotPrimeNumber(t *testing.T) {
	notPrimeNumber := new(big.Int)
	notPrimeNumber.SetString("14151519151", 10)

	var k uint = 10

	if MillerRabinTest(notPrimeNumber, k) {
		t.Errorf("failed")
	}
}
