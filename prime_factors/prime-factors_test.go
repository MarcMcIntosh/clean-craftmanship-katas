package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPrimeFactors(t *testing.T) {
	var want []int
	got := prime_factors(1)
	assert.Equal(t, want, got, "no prime factors ffor 1")

	want = []int{2}
	got = prime_factors(2)
	assert.Equal(t, want, got, "prime factor of 2 is 2")

	want = []int{3}
	got = prime_factors(3)
	assert.Equal(t, want, got, "prime factor of 3 is 3")

	want = []int{2, 2}
	got = prime_factors(4)
	assert.Equal(t, want, got, "prime factor of 4 is 2 and 2")

	want = []int{5}
	got = prime_factors(5)
	assert.Equal(t, want, got, "prime factor of 5 is 5")

	want = []int{2, 3}
	got = prime_factors(6)
	assert.Equal(t, want, got, "prime factor of 5 is 2 and 3")

	want = []int{7}
	got = prime_factors(7)
	assert.Equal(t, want, got, "prime factor of 7 is 7")

	want = []int{2, 2, 2}
	got = prime_factors(8)
	assert.Equal(t, want, got, "prime factor of 8 is 2,2,2")

	want = []int{3, 3}
	got = prime_factors(9)
	assert.Equal(t, want, got, "prime factors of 9 is 3 and 3")

	want = []int{2, 2, 3, 3, 5, 7, 7, 13}
	got = prime_factors(2 * 2 * 3 * 3 * 5 * 7 * 7 * 13)
	assert.Equal(t, want, got, "yeah that works")
}
