package seive

import "testing"

func TestNewPrimeSeive(t *testing.T) {
	ps := NewPrimeSeive(10)

	primes := ps.Primes
	if len(primes) != 10 {
		t.Errorf("Expexted 10 prime numbers but got %d", len(primes))
	}
	for _, v := range primes {
		if !isPrime(v) {
			t.Errorf("%d is not a prime number", v)
		}
	}
}

func TestSieveOfEratosthenes(t *testing.T) {
	primes := sieveOfEratosthenes(100)
	if len(primes) != 25 {
		t.Errorf("Expexted 3 prime numbers but got %d", len(primes))
	}
	for _, v := range primes {
		if !isPrime(v) {
			t.Errorf("%d is not a prime number", v)
		}
	}
}

func TestIsPrime(t *testing.T) {
	t.Run("PrimeNumber", func(t *testing.T) {
		t.Parallel()
		if isPrime(6) {
			t.Error("6 is not a prime number")
		}
	})
	t.Run("PrimeNumber", func(t *testing.T) {
		t.Parallel()
		if !isPrime(2) {
			t.Error("2 is a prime number")
		}
	})
}
