package main

import (
	"fmt"
	"strings"

	"github.com/michael-janssen-dev/advent-of-code-2022-go/common"
)

func points(c rune) int {
	if c >= 'a' {
		return int(c - 'a' + 1)
	}
	return int(c - 'A' + 27)
}

func Part1(input string) int {
	lines := strings.Split(input, "\n")

	sum := 0

	for _, line := range lines[:len(lines)-1] {
		left, right := line[:len(line)/2], line[len(line)/2:]
		for _, c := range left {
			if strings.ContainsRune(right, c) {
				sum += points(c)
				break
			}
		}

	}
	return sum
}

func Part2(input string) int {
	lines := strings.Split(input, "\n")

	sum := 0

	for i := 0; i < len(lines)-1; i += 3 {
		first, second, third := lines[i], lines[i+1], lines[i+2]
		for _, c := range first {
			if strings.ContainsRune(second, c) && strings.ContainsRune(third, c) {
				sum += points(c)
				break
			}
		}
	}
	return sum
}

func main() {
	input := common.ReadFile("input/inp.txt")

	fmt.Println(Part1(input))
	fmt.Println(Part2(input))
}
