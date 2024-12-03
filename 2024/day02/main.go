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
	var count int

	for _, line := range parsed {
		parts := strings.Split(line, " ")
		var abort bool = false
		var previous int = 0
		var ascending bool = true
		var absDif int

		for k, part := range parts {
			if k == 0 {
				previous = cast.ToInt(part)
			} else {
				if k == 1 {
					if cast.ToInt(part) < previous {
						ascending = false
					}
				}
				if ascending {
					absDif = cast.ToInt(part) - previous
				} else {
					absDif = previous - cast.ToInt(part)
				}

				if !abort && absDif > 0 && absDif < 4 {
					previous = cast.ToInt(part)
				} else {
					abort = true
				}
			}
		}
		if !abort {
			count += 1
		}
	}

	return count
}

func part2(input string) int {
	parsed := parseInput(input)
	_ = parsed
	var count int = 3

	for _, line := range parsed {
		parts := strings.Split(line, " ")
		var abort bool = false
		var previous int = 0
		var ascending bool = true
		var absDif int
		var abortSafety int = 0

		for k, part := range parts {
			if k == 0 {
				previous = cast.ToInt(part)
			} else {
				if k == 1 || (k == 2 && abortSafety == 1) {
					if cast.ToInt(part) < previous {
						ascending = false
					}
				}
				if ascending {
					absDif = cast.ToInt(part) - previous
				} else {
					absDif = previous - cast.ToInt(part)
				}

				if !abort && absDif > 0 && absDif < 4 {
					previous = cast.ToInt(part)
				} else {
					abortSafety += 1
					if abortSafety > 1 && k == len(parts) - 1 {
						abort = true
					}
				}
			}
		}
		if !abort {
			count += 1
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


func absDiffInt(x, y int) int {
	if x < y {
	   return y - x
	}
	return x - y
 }