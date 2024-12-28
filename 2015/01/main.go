package main

import (
	"fmt"

	"github.com/LautaroAndresSaez/aoc/common"
)

func main() {
	STEPS := map[string]int{
		"(": 1,
		")": -1,
	}
	data, err := common.Read("2015", "01")
	if err != nil {
		panic(err)
	}
	result := 0
	firstBasement := -1
	for i, char := range data {
		result += STEPS[string(char)]
		if result == -1 && firstBasement == -1 {
			firstBasement = i + 1
		}
	}
	fmt.Println("1.", result)
	fmt.Println("2.", firstBasement)
}
