package util

import (
	"bufio"
	"io"
	"strconv"
)

// ReadIntArr converts an input stream to an array of ints
// each value should be separated by a newline. Anything
// that isn't a valid integer is ignored.
func ReadIntArr(r io.Reader) []int {
	result := make([]int, 0)
	scanner := bufio.NewScanner(r)

	for scanner.Scan() {
		v, err := strconv.Atoi(scanner.Text())
		if err != nil {
			continue
		}
		result = append(result, v)
	}

	return result
}

func ReadStringArr(r io.Reader) []string {
	result := make([]string, 0)
	scanner := bufio.NewScanner(r)

	for scanner.Scan() {
		v := scanner.Text()
		result = append(result, v)
	}

	return result
}
