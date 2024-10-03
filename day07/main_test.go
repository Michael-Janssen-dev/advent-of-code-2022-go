package main

import (
	"testing"

	"github.com/michael-janssen-dev/advent-of-code-2022-go/common"
)

func TestPart1(t *testing.T) {
	input := common.ReadFile("input/test.txt")
	result := Part1(input)
	if result != 95437 {
		t.Fatalf("Expected 95437, got %d", result)
	}
	result = Part2(input)
	if result != 24933642 {
		t.Fatalf("Expected 24933642, got %d", result)
	}
}
