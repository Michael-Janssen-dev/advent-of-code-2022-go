package main

import (
	"testing"

	"github.com/michael-janssen-dev/advent-of-code-2022-go/common"
)

func TestPart1(t *testing.T) {
	input := common.ReadFile("input/test.txt")
	result := Part1(input)
	if result != 2 {
		t.Fatalf("Day4Part1: expected 2, got %d", result)
	}

	result = Part2(input)
	if result != 4 {
		t.Fatalf("Day4Part2: expected 4, got %d", result)
	}
}
