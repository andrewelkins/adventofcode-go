package main

import (
	_ "embed"
	"flag"
	"fmt"
	"strings"
    "regexp"

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
	var fullline = strings.Join(parsed, "")
	var count int = 0
	re := regexp.MustCompile(`mul\(\d+,\d+\)`) // matches mul(1,2)
	reSingle := regexp.MustCompile(`\d+`) // matches 1,2

	matches := re.FindAllString(fullline, -1)
	
	for _, match := range matches {
		digits := reSingle.FindAllString(match, -1)
		count += cast.ToInt(digits[0]) * cast.ToInt(digits[1])
	}

	return count
}

func part2(input string) int {
	parsed := parseInput(input)
	_ = parsed
	var fullline = strings.Join(parsed, "")
	var count int = 0
	reStrip := regexp.MustCompile(`don't\(\)(.*?)do\(\)`) // Capture dont()...do()
	reStripEnd := regexp.MustCompile(`don't\(\)(.*)(?:do\(\))*`) // Remove ending don't to end of line
	re := regexp.MustCompile(`mul\(\d+,\d+\)`) // matches mul(1,2)
	reSingle := regexp.MustCompile(`\d+`) // matches 1,2

	fulllineCleaned := reStripEnd.ReplaceAllString(reStrip.ReplaceAllString(fullline, ""),"")
	matches := re.FindAllString(fulllineCleaned, -1)
	
	for _, match := range matches {
		digits := reSingle.FindAllString(match, -1)
		count += cast.ToInt(digits[0]) * cast.ToInt(digits[1])
	}

	return count
}

func parseInput(input string) (ans []string) {
	for _, line := range strings.Split(input, "\n") {
		ans = append(ans, line)
	}
	return ans
}
