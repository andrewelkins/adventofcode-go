package main

import (
	_ "embed"
	"flag"
	"fmt"
	"strings"

	"github.com/andrewelkins/adventofcode-go/cast"
	"github.com/andrewelkins/adventofcode-go/util"
)

//go:embed input.txt
var input string

func init() {
	// do this in init (not main) so test file has same input
	input = strings.TrimRight(input, "\n")
	if len(input) == 0 {
		panic("empty input.txt file")
	}
}

func main() {
	var part int
	flag.IntVar(&part, "part", 1, "part 1 or 2")
	flag.Parse()
	fmt.Println("Running part", part)

	if part == 1 {
		ans := part1(input)
		util.CopyToClipboard(fmt.Sprintf("%v", ans))
		fmt.Println("Output:", ans)
	} else {
		ans := part2(input)
		util.CopyToClipboard(fmt.Sprintf("%v", ans))
		fmt.Println("Output:", ans)
	}
}

func part1(input string) int {
	parsed := parseInput(input)
	_ = parsed
	var count int = 0
	var xmasGrid = map[int]string{}

	for i, line := range parsed {
		xmasGrid[i] = line
	}

	// on the line look for X search forward, down, and diagonally down for XMAS
	// on the line look for S search forward, down, and diagonally down for SMAX
	for y, line := range xmasGrid {
		for x, char := range line {
			if cast.ToString(char) == "X" {
				// search forward
				if gridGet(xmasGrid, x+1, y) == "M" && gridGet(xmasGrid, x+2, y) == "A" && gridGet(xmasGrid, x+3, y) == "S" {
					count++
				}
				// search down
				if gridGet(xmasGrid, x, y+1) == "M" && gridGet(xmasGrid, x, y+2) == "A" && gridGet(xmasGrid, x, y+3) == "S" {
					count++
				}
				// search diagonally down right
				if gridGet(xmasGrid, x+1, y+1) == "M" && gridGet(xmasGrid, x+2, y+2) == "A" && gridGet(xmasGrid, x+3, y+3) == "S" {
					count++
				}
				// search diagonally down left
				if gridGet(xmasGrid, x-1, y+1) == "M" && gridGet(xmasGrid, x-2, y+2) == "A" && gridGet(xmasGrid, x-3, y+3) == "S" {
					count++
				}
			}
			if cast.ToString(char) == "S" {
				// search forward
				if gridGet(xmasGrid, x+1, y) == "A" && gridGet(xmasGrid, x+2, y) == "M" && gridGet(xmasGrid, x+3, y) == "X" {
					count++
				}
				// search down
				if gridGet(xmasGrid, x, y+1) == "A" && gridGet(xmasGrid, x, y+2) == "M" && gridGet(xmasGrid, x, y+3) == "X" {
					count++
				}
				// search diagonally down right
				if gridGet(xmasGrid, x+1, y+1) == "A" && gridGet(xmasGrid, x+2, y+2) == "M" && gridGet(xmasGrid, x+3, y+3) == "X" {
					count++
				}
				// search diagonally down left
				if gridGet(xmasGrid, x-1, y+1) == "A" && gridGet(xmasGrid, x-2, y+2) == "M" && gridGet(xmasGrid, x-3, y+3) == "X" {
					count++
				}
			}
		}
	}

	return count
}

func part2(input string) int {
	parsed := parseInput(input)
	_ = parsed
	var count int = 0
	var xmasGrid = map[int]string{}

	for i, line := range parsed {
		xmasGrid[i] = line
	}

	// on the line look for X search forward, down, and diagonally down for XMAS
	// on the line look for S search forward, down, and diagonally down for SMAX
	for y, line := range xmasGrid {
		for x, char := range line {
			if cast.ToString(char) == "A" {
				// search M.S
				if gridGet(xmasGrid, x-1, y-1) == "M" && gridGet(xmasGrid, x+1, y-1) == "S" && gridGet(xmasGrid, x-1, y+1) == "M" && gridGet(xmasGrid, x+1, y+1) == "S" {
					count++
				}
				// search S.M
				if gridGet(xmasGrid, x-1, y-1) == "S" && gridGet(xmasGrid, x+1, y-1) == "M" && gridGet(xmasGrid, x-1, y+1) == "S" && gridGet(xmasGrid, x+1, y+1) == "M" {
					count++
				}
				// search M.M
				if gridGet(xmasGrid, x-1, y-1) == "M" && gridGet(xmasGrid, x+1, y-1) == "M" && gridGet(xmasGrid, x-1, y+1) == "S" && gridGet(xmasGrid, x+1, y+1) == "S" {
					count++
				}
				// search S.S
				if gridGet(xmasGrid, x-1, y-1) == "S" && gridGet(xmasGrid, x+1, y-1) == "S" && gridGet(xmasGrid, x-1, y+1) == "M" && gridGet(xmasGrid, x+1, y+1) == "M" {
					count++
				}
			}
		}
	}

	return count
}

func parseInput(input string) (ans []string) {
	for _, line := range strings.Split(input, "\n") {
		ans = append(ans, line)
	}
	return ans
}

func gridGet(grid map[int]string, x int, y int) string {
    if (x < 0 || y < 0) || (x >= len(grid[y]) || y >= len(grid)){
		return "."
	}
    return cast.ToString(grid[y][x])
}
    