package main

import (
	"testing"

	"github.com/michael-janssen-dev/advent-of-code-2022-go/common"
)

func TestPart1(t *testing.T) {
	input := common.ReadFile("input/test.txt")
	result := Part1(input)
	if result != 21 {
		t.Fatalf("Day8Part1: expected 21, got %d", result)
	}

	result = Part2(input)
	if result != 8 {
		t.Fatalf("Day8Part2: expected 8, got %d", result)
	}
}
