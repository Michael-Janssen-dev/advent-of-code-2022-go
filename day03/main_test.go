package main

import (
	"testing"

	"github.com/michael-janssen-dev/advent-of-code-2022-go/common"
)

func TestPart1(t *testing.T) {
	input := common.ReadFile("input/test.txt")
	result := Part1(input)
	if result != 157 {
		t.Fatalf("Day3Part1: expected 157, got %d", result)
	}

	result = Part2(input)
	if result != 70 {
		t.Fatalf("Day3Part2: expected 70, got %d", result)
	}
}
