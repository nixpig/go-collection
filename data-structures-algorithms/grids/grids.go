package grids

import (
	"fmt"
	"math"
)

type Grid struct {
	NumXBins int
	NumYBins int

	XStart float64
	XEnd   float64

	YStart float64
	YEnd   float64

	Bins [10][10]*GridPoint
}

type GridPoint struct {
	x    float64
	y    float64
	next *GridPoint
}

func (g *Grid) xBinWidth() float64 {
	return (g.XEnd - g.XStart) / float64(g.NumXBins)
}

func (g *Grid) yBinWidth() float64 {
	return (g.YEnd - g.YStart) / float64(g.NumYBins)
}

func (g *Grid) Insert(x, y float64) bool {
	xBin := int(math.Floor((x - g.XStart) / g.xBinWidth()))
	fmt.Println("ins | xBin:", xBin)
	if xBin < 0 || xBin >= (g.NumXBins) {
		fmt.Println("ins | xBin out of bounds")
		return false
	}

	yBin := int(math.Floor((y - g.YStart) / g.yBinWidth()))
	fmt.Println("ins | yBin:", yBin)
	if yBin < 0 || yBin >= (g.NumYBins) {
		fmt.Println("ins | yBin out of bounds")
		return false
	}

	nextPoint := g.Bins[xBin][yBin]

	g.Bins[xBin][yBin] = &GridPoint{
		x: x, y: y, next: nextPoint,
	}

	return true
}

func (g *Grid) Delete(x, y float64) bool {
	xBin := int(math.Floor((x - g.XStart) / g.xBinWidth()))
	fmt.Println("del | xBin:", xBin)
	if xBin < 0 || xBin >= g.NumXBins {
		fmt.Println("del | xBin out of bounds")
		return false
	}

	yBin := int(math.Floor((y - g.YStart) / g.yBinWidth()))
	fmt.Println("del | yBin:", yBin)
	if yBin < 0 || yBin >= g.NumYBins {
		fmt.Println("del | Bin out of bounds")
		return false
	}

	current := g.Bins[xBin][yBin]
	if current == nil {
		return false
	}

	previous := &GridPoint{}

	for current != nil {
		if approxEqual(x, y, current.x, current.y) {
			fmt.Println("del | is approx. equal: ", x, y, current.x, current.y)
			if previous == nil {
				fmt.Println("del | previous is nil")
				g.Bins[xBin][yBin] = current.next
			} else {
				fmt.Println("del | set prev.next to current.next")
				previous.next = current.next
			}

			return true
		}

		previous = current
		current = current.next
	}

	return false
}

func approxEqual(x1, y1, x2, y2 float64) bool {
	var threshold float64 = 1

	if math.Abs(x1-x2) > threshold {
		return false
	}

	if math.Abs(y1-y2) > 1 {
		return false
	}

	return true
}
