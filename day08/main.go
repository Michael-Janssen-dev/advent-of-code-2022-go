package main

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/michael-janssen-dev/advent-of-code-2022-go/common"
)

func Part1(input string) int {
	trees := PreProcess(input)

	visible := make([][]bool, len(trees))
	for i := range trees {
		visible[i] = make([]bool, len(trees[i]))
	}

	for i := range trees {
		biggest := -1
		for j := range trees[i] {
			if trees[i][j] > biggest {
				biggest = trees[i][j]
				visible[i][j] = true
			}
		}
	}

	for i := range trees {
		biggest := -1
		for j := len(trees[i]) - 1; j >= 0; j-- {
			if trees[i][j] > biggest {
				biggest = trees[i][j]
				visible[i][j] = true
			}
		}
	}

	for j := range trees[0] {
		biggest := -1
		for i := range trees {
			if trees[i][j] > biggest {
				biggest = trees[i][j]
				visible[i][j] = true
			}
		}
	}

	for j := range trees[0] {
		biggest := -1
		for i := len(trees) - 1; i >= 0; i-- {
			if trees[i][j] > biggest {
				biggest = trees[i][j]
				visible[i][j] = true
			}
		}
	}

	count := 0
	for i := range visible {
		for j := range visible[i] {
			if visible[i][j] {
				count++
			}
		}
	}

	return count
}

func PreProcess(input string) [][]int {
	splitted := strings.Split(input, "\n")
	splitted = splitted[:len(splitted)-1]

	trees := make([][]int, len(splitted))
	for i, line := range splitted {
		trees[i] = make([]int, len(line))
		for j, char := range line {
			trees[i][j], _ = strconv.Atoi(string(char))
		}
	}
	return trees
}

func Part2(input string) int {
	trees := PreProcess(input)
	biggest := 0
	for i := range trees {
		for j := range trees[i] {
			down := 0
			for k := i + 1; k < len(trees); k++ {
				down++
				if trees[k][j] >= trees[i][j] {
					break
				}
			}
			up := 0
			for k := i - 1; k >= 0; k-- {
				up++
				if trees[k][j] >= trees[i][j] {
					break
				}
			}
			right := 0
			for k := j + 1; k < len(trees[i]); k++ {
				right++
				if trees[i][k] >= trees[i][j] {
					break
				}
			}
			left := 0
			for k := j - 1; k >= 0; k-- {
				left++
				if trees[i][k] >= trees[i][j] {
					break
				}
			}
			if left*right*up*down > biggest {
				biggest = left * right * up * down
			}
		}
	}
	return biggest
}

func main() {
	input := common.ReadFile("input/inp.txt")
	fmt.Println(Part1(input))
	fmt.Println(Part2(input))
}
