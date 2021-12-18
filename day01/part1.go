package day01

func CountIncreases(depths []int) int {
	increaseCount := 0

	for i := 1; i < len(depths); i++ {
		if depths[i] > depths[i-1] {
			increaseCount++
		}
	}

	return increaseCount
}
