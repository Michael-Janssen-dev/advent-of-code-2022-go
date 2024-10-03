package main

import (
	"fmt"
	"math"
	"strconv"
	"strings"

	"github.com/michael-janssen-dev/advent-of-code-2022-go/common"
)

type Dir struct {
	Files  map[string]int
	Dirs   map[string]*Dir
	Parent *Dir
}

func NewDir(parent *Dir) *Dir {
	return &Dir{Files: make(map[string]int), Dirs: make(map[string]*Dir), Parent: parent}
}

func FindDirsSmallerThan(dir *Dir, size int) (int, int) {
	count := 0
	sum := 0
	for _, file := range dir.Files {
		sum += file
	}
	for _, subdir := range dir.Dirs {
		total, tmp := FindDirsSmallerThan(subdir, size)
		sum += total
		count += tmp
	}
	if sum <= size {
		count += sum
	}
	return sum, count
}

func TotalSize(dir *Dir) int {
	sum := 0
	for _, subdir := range dir.Dirs {
		sum += TotalSize(subdir)
	}
	for _, file := range dir.Files {
		sum += file
	}
	return sum
}

func FindSmallestDirLargerThan(dir *Dir, size int) (int, int) {
	sum := 0
	smallest := math.MaxInt
	for _, subdir := range dir.Dirs {
		tmpSum, tmp := FindSmallestDirLargerThan(subdir, size)
		if tmp < smallest {
			smallest = tmp
		}
		sum += tmpSum
	}
	for _, file := range dir.Files {
		sum += file
	}
	if sum >= size && sum < smallest {
		return sum, sum
	}
	return sum, smallest
}

func BuildDirs(commands []string) *Dir {
	root := NewDir(nil)
	currentDir := root
	for _, command := range commands[1:] {
		lines := strings.Split(command, "\n")
		cmd := strings.Split(lines[0], " ")
		switch cmd[0] {
		case "cd":
			if cmd[1] == ".." {
				currentDir = currentDir.Parent
			} else {
				currentDir = currentDir.Dirs[cmd[1]]
			}
		case "ls":
			for _, line := range lines[1 : len(lines)-1] {
				splitted := strings.Split(line, " ")
				if splitted[0] == "dir" {
					currentDir.Dirs[splitted[1]] = NewDir(currentDir)
				} else {
					currentDir.Files[splitted[1]], _ = strconv.Atoi(splitted[0])
				}
			}
		}
	}
	return root
}

func Part1(input string) int {
	commands := strings.Split(input, "$ ")[1:]
	root := BuildDirs(commands)
	_, res := FindDirsSmallerThan(root, 100000)
	return res
}

func Part2(input string) int {
	commands := strings.Split(input, "$ ")[1:]
	root := BuildDirs(commands)
	used := TotalSize(root)
	_, res := FindSmallestDirLargerThan(root, 30000000-(70000000-used))
	return res
}

func main() {
	input := common.ReadFile("input/inp.txt")
	fmt.Println(Part1(input))
	fmt.Println(Part2(input))
}
