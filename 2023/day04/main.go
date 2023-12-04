package main

import (
	_ "embed"
	"flag"
	"fmt"
	"strings"

	"github.com/andrewelkins/adventofcode-go/cast"
	"github.com/andrewelkins/adventofcode-go/mathy"
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
	var winnerValue map[string]int
	winnerValue = make(map[string]int)
	winnerValue["0"] = 0
	winnerCount := mathy.MakeRange(1, 25)

	prizeMoney := 0

	for _, x := range winnerCount {
		if x == 1 {
			winnerValue[cast.ToString(1)] = 1
			continue
		}
		winnerValue[cast.ToString(x)] = winnerValue[cast.ToString(x-1)] * 2
	}
	fmt.Println("winnerValue",winnerValue)

	for _, line := range parsed {
		cardAndGames := strings.Split(line, ":")
		games := cardAndGames[1]
		winnersAndGame := strings.Split(games, "|")
		winnersString := winnersAndGame[0]
		gameString := winnersAndGame[1]
		winners := strings.Split(winnersString, " ")
		game := strings.Split(gameString, " ")
		winningNumbers := IntersectionStrings(winners, game)
		fmt.Println("winningNumbers", winningNumbers)
		prizeMoney += winnerValue[cast.ToString(len(winningNumbers)-1)]
		fmt.Println("prizeMoney", prizeMoney)
	}

	return prizeMoney
}

func part2(input string) int {
	parsed := parseInput(input)
	var winnerCalcValue map[string]int
	winnerCalcValue = make(map[string]int)
	winnerCalcValue["1"] = 1
	_ = winnerCalcValue
	var winnerValue map[string]int
	winnerValue = make(map[string]int)
	winnerValue["0"] = 0
	winnerValue["1"] = 1
	winnerCount := mathy.MakeRange(1, 225)
	winnerValueRange := mathy.MakeRange(1, len(parsed))
	for _, x := range winnerValueRange {
		winnerCalcValue[cast.ToString(x)] = 1
	}

	prizeMoney := 0

	for _, x := range winnerCount {
		if x == 1 {
			winnerValue[cast.ToString(1)] = 1
			continue
		}
		winnerValue[cast.ToString(x)] = winnerValue[cast.ToString(x-1)] * 2
	}
	fmt.Println("winnerValue",winnerValue)

	for x, line := range parsed {
		cardAndGames := strings.Split(line, ":")
		games := cardAndGames[1]
		winnersAndGame := strings.Split(games, "|")
		winnersString := winnersAndGame[0]
		gameString := winnersAndGame[1]
		winners := strings.Split(winnersString, " ")
		game := strings.Split(gameString, " ")
		winningNumbers := IntersectionStrings(winners, game)
		winningNumberRange := mathy.MakeRange(1, len(winningNumbers)-1)
		for _, winningNumber := range winningNumberRange {	
			winnerCalcValue[cast.ToString(x+1+winningNumber)] += 1 * winnerCalcValue[cast.ToString(x+1)]
		}
		prizeMoney += winnerCalcValue[cast.ToString(x+1)]
	}

	return prizeMoney
}

func parseInput(input string) (ans []string) {
	for _, line := range strings.Split(input, "\n") {
		ans = append(ans, line)
	}
	return ans
}

func IntersectionStrings(sli1, sli2 []string) []string {
	var result []string
	seen := map[string]bool{}
	for _, v := range sli1 {
		seen[v] = true
	}
	for _, v := range sli2 {
		if seen[v] {
			result = append(result, v)
			delete(seen, v)
		}
	}
	return result
}