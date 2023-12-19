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

func AreaOf(figure []Coordinate) int {
	var sum int
	for i := 0; i < len(figure)-1; i++ {
		sum = sum + (figure[i].X() * figure[i+1].Y()) - (figure[i].Y() * figure[i+1].X())
	}
	return Abs(sum / 2)
}

func CircumferenceOf(figure []Coordinate) int {
	var sum int
	for i := 0; i < len(figure)-1; i++ {
		sum += figure[i].MDistanceTo(figure[i+1])
	}
	return sum
}
