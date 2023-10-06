package binary_search

import (
	"math"
)

func BinarySearch(arr []int, target int) int {
	indexHigh := len(arr) - 1
	indexLow := 0

	for indexLow < indexHigh {
		var indexMid = int(math.Floor(float64(indexHigh+indexLow) / 2))
		var currentValue = arr[indexMid]

		if currentValue == target {
			return indexMid
		}

		if currentValue > target {
			indexHigh = indexMid - 1
		} else {
			indexLow = indexMid + 1
		}
	}

	return -1
}
