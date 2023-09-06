package median

import "sort"

// FindMedian calculates the median of a slice of integers.
func FindMedian(data []int) int {
	// Sort the data in ascending order
	sort.Ints(data)

	// Calculate the middle index
	middleIndex := len(data) / 2

	// Check if the number of elements is even
	if len(data)%2 == 0 {
		// If it's even, return the average of the two middle elements
		median := (data[middleIndex-1] + data[middleIndex]) / 2
		return median
	} else {
		// If it's odd, return the middle element
		median := data[middleIndex]
		return median
	}
}
