package nearest_neighbor

func LinearScan[T any](data []T, target T, compare func(test, target T) float64) (T, error) {
	// if len(data) == 0 {
	// 	return nil, errors.New("data collection to search is empty")
	// }

	candidate := data[0]
	closestDistance := compare(candidate, target)

	for i, v := range data {
		currentDistance := compare(v, target)

		if currentDistance < closestDistance {
			closestDistance = currentDistance
			candidate = data[i]
		}
	}

	return candidate, nil
}
