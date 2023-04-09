package fibonacci

import (
	"github.com/stretchr/testify/assert"
	"math/big"
	"testing"
)

type fibonacci func(int) *big.Int

func TestFibonacci(t *testing.T) {
	one := big.NewInt(1)
	assert.Equal(t, fibonacci_loop(0), one, "fibonacci 0, should be 1")
	assert.Equal(t, fibonacci_loop(1), one, "fibonacci 1, should be 1")
	assert.Equal(t, fibonacci_loop(2), big.NewInt(2), "fibonacci 2, should be 2")
	assert.Equal(t, fibonacci_loop(3), big.NewInt(3), "fibonacci 3, should be 3")
	assert.Equal(t, fibonacci_loop(4), big.NewInt(5), "fibonacci 4, should be 5")
	assert.Equal(t, fibonacci_loop(5), big.NewInt(8), "fibonacci 5, should be 8")
}

// https://dave.cheney.net/2013/06/30/how-to-write-benchmarks-in-go
// run with go test -bench=.

func benchmarkFib(fib fibonacci, n int, b *testing.B) {
	// run the Fib function b.N times
	for i := 0; i < b.N; i++ {
		fib(n)
	}
}

func BenchmarkLoop20(b *testing.B) { benchmarkFib(fibonacci_loop, 20, b) }

func BenchmarkRecursion20(b *testing.B) { benchmarkFib(fibonacci_recursion, 20, b) }

func BenchmarkTailRecursion20(b *testing.B) { benchmarkFib(fibonacci_tail_recursion, 20, b) }

// Recursion isn't that good
// Tail recursion and loop are roughly the same until 1_000_000

func BenchmarkLoop(b *testing.B) { benchmarkFib(fibonacci_loop, 100000, b) }

func BenchmarkTailRecursion(b *testing.B) { benchmarkFib(fibonacci_tail_recursion, 100000, b) }
