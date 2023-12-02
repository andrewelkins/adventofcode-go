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
	gamesGood := 0

	for _, line := range parsed {
		// Find game id
		inputSplit := strings.Split(line, ":")
		gameId := cast.ToInt(strings.ReplaceAll(inputSplit[0], "Game ", ""))
		games := strings.Split(inputSplit[1], ";")

		maxCount := make(map[string]int)
		maxCount["red"] = 0
		maxCount["green"] = 0
		maxCount["blue"] = 0
		for _, line := range games {
			games := strings.Split(line, ",")

			for _, game := range games {
				handSplit := strings.Split(strings.Trim(game, " "), " ")
				color := handSplit[1]
				count := cast.ToInt(handSplit[0])

				if count > maxCount[color] {
					maxCount[color] = count
				}
			}
		}

		if maxCount["red"] <= 12 && maxCount["green"] <= 13 && maxCount["blue"] <= 14 {
			gamesGood = gamesGood + gameId
		}
	}

	return gamesGood
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
