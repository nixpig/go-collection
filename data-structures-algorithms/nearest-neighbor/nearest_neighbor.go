package nearest_neighbor

import (
	"errors"
	"fmt"
)

func LinearScan(data []int, target int, compare func(test, target int) int) (int, error) {
	if len(data) == 0 {
		return 0, errors.New("data collection to search is empty")
	}

	candidate := data[0]
	fmt.Println("candidate:", candidate)
	closestDistance := compare(candidate, target)
	fmt.Println("closestDistance:", closestDistance)

	fmt.Println("---")

	for i, v := range data {
		currentDistance := compare(v, target)

		fmt.Println("distance:", target, v, v-target)

		if currentDistance < closestDistance {
			closestDistance = currentDistance
			candidate = data[i]
		}
	}

	return candidate, nil
}
