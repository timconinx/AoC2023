package util

import "fmt"

type Coordinate struct {
	x int
	y int
}

func (c Coordinate) String() string {
	return fmt.Sprintf("(%v,%v)", c.x, c.y)
}

func NewCoordinate(x int, y int) Coordinate {
	return Coordinate{
		x: x,
		y: y,
	}
}

func (c Coordinate) X() int {
	return c.x
}

func (c Coordinate) Y() int {
	return c.y
}
