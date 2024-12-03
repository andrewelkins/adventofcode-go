package main

import (
	_ "embed"
	"flag"
	"fmt"
	"strings"
	"sort"

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

	var leftValues []int
	var rightValues []int
	var total int

	for _, line := range parsed {
		parts := strings.Split(line, "   ")
		leftValues = append(leftValues, cast.ToInt(parts[0]))
		rightValues = append(rightValues, cast.ToInt(parts[1]))
	}

	sort.Ints(leftValues[:])
	sort.Ints(rightValues[:])

	for i := 0; i < len(leftValues); i++ { 
		total += absDiffInt(rightValues[i], leftValues[i])
    } 

	return total
}

func part2(input string) int {
	return 0
}

func parseInput(input string) (ans []string) {
	for _, line := range strings.Split(input, "\n") {
		ans = append(ans, line)
	}
	return ans
}


func absDiffInt(x, y int) int {
	if x < y {
	   return y - x
	}
	return x - y
 }