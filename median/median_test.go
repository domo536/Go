package median

import (
	"testing"
)

func TestFindMedian(t *testing.T) {
	tests := []struct {
		name     string
		input    []int
		expected int
	}{
		{
			name:     "Test with odd number of elements",
			input:    []int{7, 3, 1, 9, 2},
			expected: 3,
		},
		{
			name:     "Test with even number of elements",
			input:    []int{7, 3, 1, 9, 2, 6},
			expected: 4,
		},
		{
			name:     "Test with repeated elements",
			input:    []int{5, 5, 5, 5},
			expected: 5,
		},
		{
			name:     "Test with a single element",
			input:    []int{42},
			expected: 42,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			actual := FindMedian(test.input)
			if actual != test.expected {
				t.Errorf("Expected median: %d, but got: %d", test.expected, actual)
			}
		})
	}
}
