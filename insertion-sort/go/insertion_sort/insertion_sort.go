package insertion_sort

func InsertionSort(arr []byte) []byte {
	var current byte
	var j int

	i := 1

	for i < len(arr) {
		current = arr[i]

		j = i - 1

		for j >= 0 && arr[j] > current {
			arr[j+1] = arr[j]

			j -= 1
		}

		arr[j+1] = current

		i += 1
	}
	return arr
}
