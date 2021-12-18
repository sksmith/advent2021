package day01

func CompareWindows(values []int) int {
	inc := 0

	for i := 0; i < len(values)-3; i++ {
		a := values[i] + values[i+1] + values[i+2]
		b := values[i+1] + values[i+2] + values[i+3]

		if b > a {
			inc++
		}
	}

	return inc
}
