package main

func prime_factors(num int) []int {
	var factors []int

	for divisor := 2; num > 1; divisor++ {
		for ; num%divisor == 0; num /= divisor {
			factors = append(factors, divisor)
		}
	}

	return factors
}
