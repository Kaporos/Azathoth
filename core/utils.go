package core

import "math"

type Coordinate struct {
	x int
	y int
}

func NewCoordinate(x int, y int) Coordinate {
	return Coordinate{x, y}
}

func (c *Coordinate) distanceTo(a Coordinate) float64 {
	return math.Sqrt(
		math.Pow((float64)(c.x-a.x), 2) +
			math.Pow((float64)(c.y-a.y), 2))
}

func (c *Coordinate) GetX() int {
	return c.x
}
func (c *Coordinate) GetY() int {
	return c.y
}
