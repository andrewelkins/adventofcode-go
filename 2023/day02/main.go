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

type Game struct {
	gameId int
	games []string
}

type Match struct {
	red int
	green int
	blue int
}

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
	gamesGood := 0

	for _, line := range parsed {
		var gameId, games = parseGameIdAndGames(line)

		maxCount := generateMaxCountMap(games)

		if maxCount.red <= 12 && maxCount.green <= 13 && maxCount.blue <= 14 {
			gamesGood += gameId
		}
	}

	return gamesGood
}

func part2(input string) int {
	parsed := parseInput(input)
	gamesGood := 0

	for _, line := range parsed {
		var gameId, games = parseGameIdAndGames(line)
		_ = gameId

		maxCount := generateMaxCountMap(games)

		gamesGood += maxCount.red * maxCount.green * maxCount.blue
	}

	return gamesGood
}

func parseInput(input string) (ans []string) {
	for _, line := range strings.Split(input, "\n") {
		ans = append(ans, line)
	}
	return ans
}

func parseGameIdAndGames(input string) (gameId int, games []string) {
	inputSplit := strings.Split(input, ":")

	gameId = parseGameId(inputSplit[0])
	games = parseGames(inputSplit[1])
	return gameId, games
}

func parseGameId(input string) (ans int) {
	return cast.ToInt(strings.ReplaceAll(input, "Game ", ""))
}

func parseGames(input string) (ans []string) {
	return strings.Split(input, ";")
}

func parseGame(input string) (ans []string) {
	return strings.Split(input, ",")
}

func generateMaxCountMap(games []string) (maxCount Match) {
	maxCount = Match{}
	maxCount.red = 0
	maxCount.green = 0
	maxCount.blue = 0
	for _, line := range games {
		games := parseGame(line)

		for _, game := range games {
			handSplit := strings.Split(strings.Trim(game, " "), " ")
			color := handSplit[1]
			count := cast.ToInt(handSplit[0])

			switch c := color ; c {
				case "blue":
					if count > maxCount.blue {
						maxCount.blue = count
					}
				case "red":
					if count > maxCount.red {
						maxCount.red = count
					}
				case "green":
					if count > maxCount.green {
						maxCount.green = count
					}
				default:
			}
		}
	}
	return maxCount
}