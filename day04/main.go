package main

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/michael-janssen-dev/advent-of-code-2022-go/common"
)

func splitSection(section string) (int, int) {
	splitted := strings.Split(section, "-")
	left, _ := strconv.Atoi(splitted[0])
	right, _ := strconv.Atoi(splitted[1])
	return left, right
}

func Part1(input string) int {
	lines := strings.Split(input, "\n")
	sum := 0
	for _, line := range lines[:len(lines)-1] {
		splitted := strings.Split(line, ",")
		ll, lr := splitSection(splitted[0])
		rl, rr := splitSection(splitted[1])
		if ll <= rl && lr >= rr {
			sum++
		} else if ll >= rl && lr <= rr {
			sum++
		}
	}
	return sum
}

func Part2(input string) int {
	lines := strings.Split(input, "\n")
	sum := 0
	for _, line := range lines[:len(lines)-1] {
		splitted := strings.Split(line, ",")
		ll, lr := splitSection(splitted[0])
		rl, rr := splitSection(splitted[1])
		if ll <= rl && lr >= rr {
			sum++
		} else if ll >= rl && lr <= rr {
			sum++
		} else if ll <= rl && lr >= rl {
			sum++
		} else if ll <= rr && lr >= rr {
			sum++
		}
	}
	return sum
}

func main() {
	input := common.ReadFile("input/inp.txt")

	fmt.Println(Part1(input))
	fmt.Println(Part2(input))
}
