package main

import (
	"fmt"

	"github.com/michael-janssen-dev/advent-of-code-2022-go/common"
)

func findMarker(line string, length int) int {
	i := 0
	last := make([]rune, length)
outer:
	for j, letter := range line {
		for k := 0; k < i; k++ {
			if last[k] == letter {
				last[0] = letter
				i = 1
				continue outer
			}
		}
		last[i] = letter
		i++
		if i == length {
			return j + 1
		}
	}
	return 0
}

func Part1(input string) int {
	return findMarker(input, 4)
}

func Part2(input string) int {
	return findMarker(input, 14)
}

func main() {
	input := common.ReadFile("input/inp.txt")
	fmt.Println(Part1(input))
	fmt.Println(Part2(input))
}
