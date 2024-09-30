package main

import (
	"testing"

	"github.com/michael-janssen-dev/advent-of-code-2022-go/common"
)

func TestDay2(t *testing.T) {
	input := common.ReadFile("input/test.txt")
	result := Part1(input)
	if result != 15 {
		t.Fatalf("Day2Part1: expected 15, got %d", result)
	}

	result = Part2(input)
	if result != 12 {
		t.Fatalf("Day2Part2: expected 12, got %d", result)
	}
}
