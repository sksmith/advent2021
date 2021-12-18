package day01

import (
	"fmt"
	"testing"
)

func TestCountIncreases(t *testing.T) {
	// t.Fatal("not implemented")
	tests := []struct {
		input    []int
		expected int
	}{
		{
			input: []int{
				199, 200, 208, 210, 200, 207, 240, 269, 260, 263,
			},
			expected: 7,
		},
	}

	for i, test := range tests {
		testname := fmt.Sprintf("Test %d", i)

		t.Run(testname, func(t *testing.T) {

			result := CountIncreases(test.input)
			if result != test.expected {
				t.Errorf("got %d, want %d", result, test.expected)
			}
		})
	}
}
