package main

import (
	_ "embed"
	"flag"
	"fmt"
	"strings"
	"regexp"
	"math"

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

	re := regexp.MustCompile("\\d+")
	times := re.FindAllString(parsed[0], -1)
	distances := re.FindAllString(parsed[1], -1)
	win := 1

	for x, time := range times {
		intervalWinningCount := 0
		timeInt := cast.ToInt(time)
		distance := cast.ToInt(distances[x])
		for i := 0; i < timeInt; i++ {
			if i * (timeInt-i) > distance {
				intervalWinningCount++
			}
		}
		win = win * intervalWinningCount
	}

	return win
}

func part2(input string) int {
	parsed := parseInput(input)

	time := cast.ToFloat(strings.ReplaceAll(strings.Split(parsed[0], ":")[1], " ", ""))
	distance := cast.ToFloat(strings.ReplaceAll(strings.Split(parsed[1], ":")[1], " ", ""))

	low := math.Floor((time - math.Sqrt(math.Pow(time, 2) - 4 * distance))/2)
	high := math.Ceil((time + math.Sqrt(math.Pow(time, 2) - 4 * distance))/2)

	return int(high - low - 1)
}

func parseInput(input string) (ans []string) {
	for _, line := range strings.Split(input, "\n") {
		ans = append(ans, line)
	}
	return ans
}
