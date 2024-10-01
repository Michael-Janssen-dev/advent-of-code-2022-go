package main

import (
	"testing"

	"github.com/michael-janssen-dev/advent-of-code-2022-go/common"
)

func TestDay5(t *testing.T) {
	input := common.ReadFile("input/test.txt")
	result := Part1(input)
	if result != "CMZ" {
		t.Fatalf("Day5Part1: expected CMZ, got %s", result)
	}

	result = Part2(input)
	if result != "MCD" {
		t.Fatalf("Day5Part2: expected MCD, got %s", result)
	}
}
