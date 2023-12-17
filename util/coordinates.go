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

func (c Coordinate) Equals(other Coordinate) bool {
	return c.x == other.x && c.y == other.y
}

func (c Coordinate) MDistanceTo(c2 Coordinate) int {
	var x, y int
	if c.x > c2.x {
		x = c.x - c2.x
	} else {
		x = c2.x - c.x
	}
	if c.y > c2.y {
		y = c.y - c2.y
	} else {
		y = c2.y - c.y
	}
	return x + y
}
