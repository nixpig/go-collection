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

func (g *Grid) Print() {
	for _, bin := range g.Bins {
		fmt.Println(bin)
	}
}

func (g *Grid) xBinWidth() float64 {
	return (g.XEnd - g.XStart) / float64(g.NumXBins)
}

func (g *Grid) yBinWidth() float64 {
	return (g.YEnd - g.YStart) / float64(g.NumYBins)
}

func (g *Grid) Get(x, y float64) (*GridPoint, error) {
	xBin := int(math.Floor((x - g.XStart) / g.xBinWidth()))
	if xBin < 0 || xBin >= g.NumXBins {
		return nil, fmt.Errorf("'x' position falls outside the grid.")
	}

	yBin := int(math.Floor((y - g.YStart) / g.yBinWidth()))
	if yBin < 0 || yBin >= g.NumYBins {
		return nil, fmt.Errorf("'y' position falls outside the grid.")
	}

	current := g.Bins[xBin][yBin]
	if current == nil {
		return nil, fmt.Errorf("Nothing found at this position.")
	}

	for current != nil {
		if approxEqual(x, y, current.x, current.y) {
			return current, nil
		}

		current = current.next
	}

	return nil, fmt.Errorf("No item found.")

}

func (g *Grid) Insert(x, y float64) bool {
	xBin := int(math.Floor((x - g.XStart) / g.xBinWidth()))
	if xBin < 0 || xBin >= (g.NumXBins) {
		return false
	}

	yBin := int(math.Floor((y - g.YStart) / g.yBinWidth()))
	if yBin < 0 || yBin >= (g.NumYBins) {
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
	if xBin < 0 || xBin >= g.NumXBins {
		return false
	}

	yBin := int(math.Floor((y - g.YStart) / g.yBinWidth()))
	if yBin < 0 || yBin >= g.NumYBins {
		return false
	}

	current := g.Bins[xBin][yBin]
	if current == nil {
		return false
	}

	var previous *GridPoint

	for current != nil {
		if approxEqual(x, y, current.x, current.y) {
			if previous == nil {
				g.Bins[xBin][yBin] = current.next
			} else {
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

func euclidianDistance(x1, y1, x2, y2 float64) float64 {
	return math.Sqrt((x1-x2)*(x1-x2) + (y1-y2)*(y1-y2))
}
