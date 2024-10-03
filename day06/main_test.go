package main

import (
	"testing"

	"github.com/michael-janssen-dev/advent-of-code-2022-go/common"
)

func TestPart1(t *testing.T) {
	input := common.ReadFile("input/test.txt")
	result := Part1(input)
	if result != 7 {
		t.Fatalf("Expected 7, got %d", result)
	}

	result = Part2(input)
	if result != 19 {
		t.Fatalf("Expected 19, got %d", result)
	}

}
