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
	result := 0

	for _, line := range parsed {
		line := strings.Split(line, ":")
		answer := cast.ToInt(line[0])
		problems := parseInt(line[1])

		if isValid(answer, problems[0], problems[1:], false) {
			result += answer
		}
	}

	return result
}

func part2(input string) int {
	parsed := parseInput(input)
	_ = parsed
	result := 0

	for _, line := range parsed {
		line := strings.Split(line, ":")
		answer := cast.ToInt(line[0])
		problems := parseInt(line[1])

		if isValid(answer, problems[0], problems[1:], true) {
			result += answer
		}
	}

	return result
}

func parseInput(input string) (ans []string) {
	for _, line := range strings.Split(input, "\n") {
		ans = append(ans, line)
	}
	return ans
}

func parseInt(input string) (ans []int) {
	for _, line := range strings.Split(strings.Trim(input, " "), " ") {
		ans = append(ans, cast.ToInt(line))
	}
	return ans
}

func concatInts(a int, b int) int {
	return cast.ToInt(cast.ToString(a) + cast.ToString(b))
}

func isValid(target, current int, numbers []int, concat bool) bool {
	if len(numbers) == 0 && target == current {
		return true
	}

	if len(numbers) == 0 {
		return false
	}

	return isValid(target, current+numbers[0], numbers[1:], concat) || isValid(target, current*numbers[0], numbers[1:], concat) ||
		(concat && isValid(target, concatInts(current, numbers[0]), numbers[1:], concat))
}
