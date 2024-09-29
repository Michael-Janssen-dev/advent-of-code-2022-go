package main

import (
	"fmt"
	"sort"
	"strconv"
	"strings"

	"github.com/michael-janssen-dev/advent-of-code-2022-go/common"
)

func Part1(input string) int {
	lines := strings.Split(input, "\n\n")
	max := 0
	for _, line := range lines[:len(lines)-1] {
		values := strings.Split(line, "\n")
		sum := 0
		for _, value := range values {
			asInt, err := strconv.Atoi(value)
			if err != nil {
				panic(err)
			}
			sum += asInt
		}
		if sum > max {
			max = sum
		}
	}
	return max
}

func Part2(input string) int {
	lines := strings.Split(input, "\n\n")
	sums := make([]int, len(lines)-1)
	for i, line := range lines[:len(lines)-1] {
		values := strings.Split(line, "\n")
		sums = append(sums, 0)
		for _, value := range values {
			asInt, err := strconv.Atoi(value)
			if err != nil {
				panic(err)
			}
			sums[i] += asInt
		}
	}
	sort.Slice(sums, func(i, j int) bool {
		return sums[i] > sums[j]
	})
	return sums[0] + sums[1] + sums[2]
}

func main() {
	input := common.ReadFile("input/inp.txt")

	fmt.Println(Part1(input))
	fmt.Println(Part2(input))
}
