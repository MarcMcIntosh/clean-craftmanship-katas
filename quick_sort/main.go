package main

func sort(list []int) []int {
	result := []int{}
	if len(list) == 0 {
		return result
	}
	middle := list[0]

	lessers := filter(list, func(d int) bool {
		return d < middle
	})

	middles := filter(list, func(d int) bool {
		return d == middle
	})
	greaters := filter(list, func(d int) bool {
		return d > middle
	})

	result = append(result, sort(lessers)...)
	result = append(result, middles...)
	result = append(result, sort(greaters)...)

	return result
}

func filter(list []int, condition func(int) bool) []int {
	result := []int{}
	for i := range list {
		if condition(list[i]) == true {
			result = append(result, list[i])
		}
	}
	return result
}
