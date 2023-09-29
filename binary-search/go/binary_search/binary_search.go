package binary_search

import (
	"math"
)

func main() {

}

func BinarySearch(arr []int, target int) int {
	indexHigh := len(arr) - 1
	indexLow := 0
	indexMid := indexLow

	for indexLow < indexHigh {
		indexMid = int(math.Floor(float64(indexHigh+indexMid) / 2))

		if arr[indexMid] == target {
			return indexMid
		}

		if arr[indexMid] > target {
			indexHigh = indexMid - 1
		} else {

			indexLow = indexMid + 1
		}
	}

	return -1
}
