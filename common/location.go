package common

import "fmt"

type Location struct {
	X int
	Y int
}

func NewLocation(x int, y int) *Location {
	return &Location{
		X: x,
		Y: y,
	}
}

func (l Location) String() string {
	return fmt.Sprintf("(%d, %d)", l.X, l.Y)
}

func (l Location) Add(l2 *Location) *Location {
	return &Location{
		X: l.X + l2.X,
		Y: l.Y + l2.Y,
	}
}
