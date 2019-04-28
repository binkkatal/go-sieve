package seive

import (
	"fmt"
	"strings"
)

var (
	primeNumberStore = sieveOfEratosthenes(10000000)
)

// PrimeSeive is the struct which stores the the prime numbers starting from 2 ..  a given count
type PrimeSeive struct {
	Primes []int
}

// NewPrimeSeive initializes a new PrimeSeive struct by taking a count and calculating prime numbers upto the count
func NewPrimeSeive(count int) PrimeSeive {
	primeNumbers := primeNumberStore[0:count]
	return PrimeSeive{
		Primes: primeNumbers,
	}
}

// printHeader prints table header for multiplication table
func (ps PrimeSeive) printHeader() {
	s := strings.Trim(strings.Join(strings.Fields(fmt.Sprint(ps.Primes)), " | "), "[]")
	fmt.Printf("| x | %+v |\n", s)
	fmt.Printf("|%+v\n", strings.Repeat("--|", len(ps.Primes)+1))
}

// printRow prints row for multiplication table
func (ps PrimeSeive) printRow(num int) {
	row := []int{}

	for _, v := range ps.Primes {
		row = append(row, num*v)
	}
	s := strings.Trim(strings.Join(strings.Fields(fmt.Sprint(row)), " | "), "[]")
	fmt.Printf("| %d | %+v |\n", num, s)
}

// PrintTable prints the multiplication table of prime numbers
func (ps PrimeSeive) PrintTable() {
	ps.printHeader()
	for _, v := range ps.Primes {
		ps.printRow(v)
	}
}

// sieveOfEratosthenes calculates prime numbers upto n.
// ref: https://en.wikipedia.org/wiki/Sieve_of_Eratosthenes
func sieveOfEratosthenes(n int) []int {
	// Create a boolean array "prime[0..n]" and initialize
	// all entries it as true. A value in prime[i] will
	// finally be false if i is Not a prime, else true.
	integers := make([]bool, n+1)
	for i := 2; i < n+1; i++ {
		integers[i] = true
	}

	for p := 2; p*p <= n; p++ {
		// If integers[p] is not changed, then it is a prime
		if integers[p] == true {
			// Update all multiples of p
			for i := p * 2; i <= n; i += p {
				integers[i] = false
			}
		}
	}

	// return all prime numbers <= n
	var primes []int
	for p := 2; p <= n; p++ {
		if integers[p] == true {
			primes = append(primes, p)
		}
	}

	return primes
}
