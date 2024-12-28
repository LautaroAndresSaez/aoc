package main

import (
	"fmt"

	"github.com/LautaroAndresSaez/aoc/common"
)

type Move string

const (
	UP    Move = "^"
	DOWN  Move = "v"
	LEFT  Move = "<"
	RIGHT Move = ">"
)

type Trip struct {
	visited map[string]bool
	santa   *common.Location
}

func NewTrip() *Trip {
	trip := Trip{
		visited: make(map[string]bool),
		santa:   common.NewLocation(0, 0),
	}
	trip.visited[trip.santa.String()] = true
	return &trip
}

func (t *Trip) AddStep(move Move) {
	switch move {
	case UP:
		t.santa = t.santa.Add(common.NewLocation(0, 1))
	case DOWN:
		t.santa = t.santa.Add(common.NewLocation(0, -1))
	case LEFT:
		t.santa = t.santa.Add(common.NewLocation(-1, 0))
	case RIGHT:
		t.santa = t.santa.Add(common.NewLocation(1, 0))
	}
	t.visited[t.santa.String()] = true
}

func main() {
	data, err := common.Read("2015", "03")
	trip := NewTrip()
	if err != nil {
		panic(err)
	}
	for _, charCode := range data {
		move := Move(string(charCode))
		trip.AddStep(move)
	}
	fmt.Println("1.", len(trip.visited))
}
