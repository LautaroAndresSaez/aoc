package main

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/LautaroAndresSaez/aoc/common"
)

type Box struct {
	Height int
	Length int
	Width  int
}

func NewBox(width, height, length int) *Box {
	return &Box{
		Height: height,
		Length: length,
		Width:  width,
	}
}

func (b *Box) Area() int {
	return 2*b.Height*b.Length + 2*b.Height*b.Width + 2*b.Width*b.Length
}

func (b *Box) Volume() int {
	return b.Height * b.Length * b.Width
}

func (b *Box) SmallestArea() int {
	a1 := b.Width * b.Length
	a2 := b.Width * b.Height
	a3 := b.Length * b.Height
	return min(min(a1, a2), a3)
}

func (b *Box) SmallestPerimeter() int {
	p1 := b.Width*2 + 2*b.Length
	p2 := b.Width*2 + 2*b.Height
	p3 := b.Length*2 + 2*b.Height
	return min(min(p1, p2), p3)
}

func createBoxes(data string) []*Box {
	rawBoxes := strings.Split(data, "\n")
	boxes := make([]*Box, 0, len(rawBoxes))
	for _, rawBox := range rawBoxes {
		dimensions := strings.Split(strings.TrimSpace(rawBox), "x")
		length, _ := strconv.Atoi(dimensions[0])
		width, _ := strconv.Atoi(dimensions[1])
		height, _ := strconv.Atoi(dimensions[2])
		boxes = append(boxes, NewBox(width, height, length))
	}
	return boxes
}

func main() {
	data, err := common.Read("2015", "02")
	if err != nil {
		panic(err)
	}
	boxes := createBoxes(data)
	totalArea := 0
	totalRibbon := 0
	for _, box := range boxes {
		totalArea += box.Area() + box.SmallestArea()
		totalRibbon += box.Volume() + box.SmallestPerimeter()
	}
	fmt.Println("1.", totalArea)
	fmt.Println("2.", totalRibbon)
}
