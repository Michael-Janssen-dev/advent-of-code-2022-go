package main

import (
	"fmt"
	"strings"

	"github.com/michael-janssen-dev/advent-of-code-2022-go/common"
)

var pointsP1 = map[string]int{
	"A X": 3 + 1,
	"A Y": 6 + 2,
	"A Z": 0 + 3,
	"B X": 0 + 1,
	"B Y": 3 + 2,
	"B Z": 6 + 3,
	"C X": 6 + 1,
	"C Y": 0 + 2,
	"C Z": 3 + 3,
}

var pointsP2 = map[string]int{
	"A X": 0 + 3,
	"A Y": 3 + 1,
	"A Z": 6 + 2,
	"B X": 0 + 1,
	"B Y": 3 + 2,
	"B Z": 6 + 3,
	"C X": 0 + 2,
	"C Y": 3 + 3,
	"C Z": 6 + 1,
}

func Part1(input string) int {
	lines := strings.Split(input, "\n")
	sum := 0
	for _, line := range lines[:len(lines)-1] {
		sum += pointsP1[line]
	}

	return sum
}

func Part2(input string) int {
	lines := strings.Split(input, "\n")
	sum := 0
	for _, line := range lines[:len(lines)-1] {
		sum += pointsP2[line]
	}

	return sum
}

func main() {
	input := common.ReadFile("input/inp.txt")

	fmt.Println(Part1(input))
	fmt.Println(Part2(input))
}
