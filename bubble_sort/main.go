package main

func sort(list []int) []int {
	if len(list) > 1 {
		for limit := len(list) - 1; limit > 0; limit-- {
			for firstIndex := 0; firstIndex < limit; firstIndex++ {
				secondIndex := firstIndex + 1
				first := list[firstIndex]
				second := list[secondIndex]
				if first > second {
					list[firstIndex] = second
					list[secondIndex] = first
				}
			}
		}
	}
	return list
}
