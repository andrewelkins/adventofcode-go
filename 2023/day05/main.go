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
	seeds := getSeeds(parsed[0])
	board := buildMap(parsed[1:])
	location := 0

	for _, seed := range seeds {
		seedToSoil := pullValueFromMap(board, "seed-to-soil", seed)
		soilToFertilizer := pullValueFromMap(board, "soil-to-fertilizer", seedToSoil)
		fertilizerToWater := pullValueFromMap(board, "fertilizer-to-water", soilToFertilizer)
		waterToLight := pullValueFromMap(board, "water-to-light", fertilizerToWater)
		lightToTemperature := pullValueFromMap(board, "light-to-temperature", waterToLight)
		temperatureToHumidity := pullValueFromMap(board, "temperature-to-humidity", lightToTemperature)
		humidityToLocation := pullValueFromMap(board, "humidity-to-location", temperatureToHumidity)

		if location == 0 || humidityToLocation < location {
			location = humidityToLocation
		}
	}
	
	return location
}

func part2(input string) int {
	parsed := parseInput(input)
	seeds := getSeedMultipliers(getSeeds(parsed[0]))
	board := buildMap(parsed[1:])
	location := 0

	for _, seed := range seeds {
		seedToSoil := pullValueFromMap(board, "seed-to-soil", seed)
		soilToFertilizer := pullValueFromMap(board, "soil-to-fertilizer", seedToSoil)
		fertilizerToWater := pullValueFromMap(board, "fertilizer-to-water", soilToFertilizer)
		waterToLight := pullValueFromMap(board, "water-to-light", fertilizerToWater)
		lightToTemperature := pullValueFromMap(board, "light-to-temperature", waterToLight)
		temperatureToHumidity := pullValueFromMap(board, "temperature-to-humidity", lightToTemperature)
		humidityToLocation := pullValueFromMap(board, "humidity-to-location", temperatureToHumidity)

		if location == 0 || humidityToLocation < location {
			location = humidityToLocation
		}
	}
	
	return location
}

func parseInput(input string) (ans []string) {
	for _, line := range strings.Split(input, "\n") {
		ans = append(ans, line)
	}
	return ans
}

func getSeeds(input string) (board []int) {
	seeds := []int{}
	for _, seed := range strings.Split(input, " ")[1:] {
		seeds = append(seeds, cast.ToInt(seed))
	}
	return seeds
}

func getSeedMultipliers(input []int) (seeds []int) {

	for x, seed := range input {
		if x %2 == 0 {
			for _, y := range mathy.MakeRange(x, x+seed[x+1]) {
				seeds = append(seeds, seed*y)
			}
			seeds = append(seeds, x)
		} 
	return seeds
}

func buildMap(input []string) (board map[string][][]int) {
	
	board = make(map[string][][]int)
	mapType := ""
	parseMapType := false
	for _, line := range input {
		// Blank line. Set next line to be map type
		if line == "" {
			parseMapType = true
			continue
		}
		// Get map type
		if parseMapType {
			mapType = strings.Split(line, " ")[0]
			parseMapType = false
			continue
		}

		// Get map
		parsedLine := strings.Split(line, " ")
		destination := cast.ToInt(parsedLine[0])
		source := cast.ToInt(parsedLine[1])
		count := cast.ToInt(parsedLine[2])
		comboDestSourceCount := []int{destination, source, count}
		board[mapType] = append(board[mapType], comboDestSourceCount)
	}
	return board
}

func pullValueFromMap(board map[string][][]int, mapType string, entry  int) (ans int) {
	currentMap := board[mapType]
	for _, combo := range currentMap {
		destination, source, count := combo[0], combo[1], combo[2]
		if source <= entry && entry < (source + count) {
			// Return mapped value
			return destination + (entry - source)
		}
	}
	// If not, return entry
	return entry
}