package main

import (
	"testing"

	"github.com/michael-janssen-dev/advent-of-code-2022-go/common"
)

func TestDay1(t *testing.T) {
	input := common.ReadFile("input/test.txt")

	result := Part1(input)
	if result != 24000 {
		t.Fatalf("Day1Part1: expected 24000, got %d", result)
	}

	result = Part2(input)
	if result != 45000 {
		t.Fatalf("Day1Part2: expected 0, got %d", result)
	}
}
