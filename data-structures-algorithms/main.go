package main

import (
	"fmt"

	"github.com/nixpig/go-collection/data-structures-algorithms/grids"
)

func main() {
	// data := []int{23, 13, 7, 14, 666, 42, 69}
	//
	// straightLineComparison := func(test, target int) float64 {
	// 	d := test - target
	// 	if d < 0 {
	// 		return float64(d * -1)
	// 	} else {
	// 		return float64(d)
	// 	}
	// }
	//
	// nearest, err := nearest_neighbor.LinearScan(data, 50, straightLineComparison)
	// if err != nil {
	// 	fmt.Println(err)
	// 	os.Exit(1)
	// }
	//
	// fmt.Println("nearest:", nearest)
	// fmt.Println("---")
	//
	// type point struct {
	// 	x float64
	// 	y float64
	// }
	//
	// euclidianComparison := func(test, target point) float64 {
	// 	testPoint := math.Pow((test.x - test.y), 2)
	// 	targetPoint := math.Pow((target.x - target.y), 2)
	// 	distance := math.Sqrt(testPoint + targetPoint)
	//
	// 	fmt.Println("testPoint:", testPoint, "| targetPoint:", targetPoint, "| distance:", distance)
	//
	// 	return distance
	// }
	//
	// points := []point{
	// 	{x: 21, y: 32},
	// 	{x: 11, y: 12},
	// 	{x: 20, y: 2},
	// 	{x: 4, y: 40},
	// }
	//
	// closestPoint, err := nearest_neighbor.LinearScan(points, point{x: 2, y: 2}, euclidianComparison)
	// if err != nil {
	// 	fmt.Println(err)
	// 	os.Exit(1)
	// }
	// fmt.Println("closest point:", closestPoint)

	grid := grids.Grid{
		NumXBins: 10,
		NumYBins: 10,
		XStart:   1,
		YStart:   1,
		XEnd:     10,
		YEnd:     10,
	}

	ins := grid.Insert(3, 8)
	fmt.Println("inserted?:", ins)

	got, err := grid.Get(3, 8)
	if err != nil {
		fmt.Println("Not got:", err)
	}
	fmt.Println("Got:", got)

	var del bool

	del = grid.Delete(3, 8)
	fmt.Println("deleted?:", del)

	del = grid.Delete(3, 8)
	fmt.Println("deleted?:", del)

	got, err = grid.Get(3, 8)
	if err != nil {
		fmt.Println("Not got:", err)
	}
	fmt.Println("Got:", got)

}
