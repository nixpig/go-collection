package insertion_sort

func InsertionSort(arr []int) []int {
outer:
	for i := 1; i < len(arr); i++ {
		currentItem := arr[i]

	inner:
		for j := i - 1; j >= 0; j-- {
			if currentItem >= arr[j] {
				arr[j+1] = currentItem
				continue outer
			} else {
				arr[j+1] = arr[j]
				arr[j] = currentItem
				continue inner
			}
		}
	}

	return arr
}
