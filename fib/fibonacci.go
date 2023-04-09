package fibonacci

import (
	"math/big"
)

func fibonacci_recursion(n int) *big.Int {
	if n > 1 {
		n1 := fibonacci_recursion(n - 1)
		n2 := fibonacci_recursion(n - 2)
		return n1.Add(n1, n2)
	}
	return big.NewInt(1)
}

//Fibonacci in a tail-recursive version
func fiboT(n int, first, second *big.Int) *big.Int {
	if n == 0 {
		return second
	}
	return fiboT(n-1, second, first.Add(first, second))
}

func fibonacci_tail_recursion(n int) *big.Int {
	return fiboT(n, big.NewInt(0), big.NewInt(1))
}

func fibonacci_loop(n int) *big.Int {

	a, b := big.NewInt(0), big.NewInt(1)

	for i := 0; i < n; i++ {
		a.Add(a, b)
		a, b = b, a
	}

	return b
}
