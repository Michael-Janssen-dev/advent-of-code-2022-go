package main

import (
	"fmt"
	"slices"
	"strconv"
	"strings"

	"github.com/michael-janssen-dev/advent-of-code-2022-go/common"
)

func Part1(input string) string {
	stacks, lines := createStacks(input)
	for _, line := range lines[:len(lines)-1] {
		amount, from, to := parseLine(line)
		for range amount {
			stacks[to] = append(stacks[to], stacks[from][len(stacks[from])-1])
			stacks[from] = stacks[from][:len(stacks[from])-1]
		}
	}

	var res strings.Builder
	for _, stack := range stacks {
		res.WriteRune(stack[len(stack)-1])
	}

	return res.String()
}

func Part2(input string) string {
	stacks, lines := createStacks(input)
	for _, line := range lines[:len(lines)-1] {
		amount, from, to := parseLine(line)
		stacks[to] = append(stacks[to], stacks[from][len(stacks[from])-amount:]...)
		stacks[from] = stacks[from][:len(stacks[from])-amount]
	}

	var res strings.Builder
	for _, stack := range stacks {
		res.WriteRune(stack[len(stack)-1])
	}

	return res.String()
}

func parseLine(line string) (int, int, int) {
	splitted := strings.Split(line, " ")
	amount, _ := strconv.Atoi(splitted[1])
	from, _ := strconv.Atoi(splitted[3])
	to, _ := strconv.Atoi(splitted[5])
	from--
	to--
	return amount, from, to
}

func createStacks(input string) ([][]rune, []string) {
	stacks := make([][]rune, 0)

	splitted := strings.Split(input, "\n\n")

	lines := strings.Split(splitted[0], "\n")
	slices.Reverse(lines)
	for range len(lines[0])/4 + 1 {
		stacks = append(stacks, make([]rune, 0))
	}
	for _, line := range lines[1:] {
		for j := range len(line)/4 + 1 {
			char := []rune(line)[j*4+1]
			if char == ' ' {
				continue
			}
			stacks[j] = append(stacks[j], char)
		}
	}

	lines = strings.Split(splitted[1], "\n")
	return stacks, lines
}

func main() {
	input := common.ReadFile("input/inp.txt")
	fmt.Println(Part1(input))
	fmt.Println(Part2(input))
}
