package main

import (
	_ "embed"
	"flag"
	"fmt"
	"math"
	"strings"

	"github.com/andrewelkins/adventofcode-go/cast"
	"github.com/andrewelkins/adventofcode-go/util"
)

//go:embed input.txt
var input string

type Corner struct {
	x, y int
}

type Side struct {
	x, y, k     int
	orientation string
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
	grid := parseGrid(parsed)
	fullGrid := printGrid(len(grid[0]), len(grid))
	farm := parseFarm(grid)
	regions := getRegions(farm, fullGrid)

	return countRegions(farm, regions)
}

func part2(input string) int {
	parsed := parseInput(input)
	grid := parseGrid(parsed)
	fullGrid := printGrid(len(grid[0]), len(grid))
	farm := parseFarm(grid)
	sides := parseSides(grid)
	regions := getRegions(farm, fullGrid)

	return countRegions(sides, regions)
}

func parseInput(input string) (ans []string) {
	for _, line := range strings.Split(input, "\n") {
		ans = append(ans, line)
	}
	return ans
}

func parseGrid(parsed []string) (ans [][]rune) {
	for _, line := range parsed {
		ans = append(ans, []rune(line))
	}
	return ans
}

func printGrid(x int, y int) (result map[Corner]struct{}) {
	result = make(map[Corner]struct{})
	for i := 0; i < x; i++ {
		for j := 0; j < y; j++ {
			result[Corner{i, j}] = struct{}{}
		}
	}
	return result
}

func parseFarm(grid [][]rune) (farm map[Corner]map[Corner]struct{}) {
	farm = make(map[Corner]map[Corner]struct{})
	maxX, maxY := len(grid[0]), len(grid)
	corners := []Corner{{-1, 0}, {0, -1}, {0, 1}, {1, 0}}
	for i := 0; i < maxX; i++ {
		for x := 0; x < maxY; x++ {
			current := Corner{i, x}
			if _, exists := farm[current]; !exists {
				farm[current] = make(map[Corner]struct{})
			}
			for _, corner := range corners {
				cornerI, cornerX := i+corner.x, x+corner.y
				if cornerI >= 0 && cornerI < maxX && cornerX >= 0 && cornerX < maxY {
					if grid[i][x] == grid[cornerI][cornerX] {
						farm[current][Corner{cornerI, cornerX}] = struct{}{}
					}
				}
			}
		}
	}
	return farm
}

func parseSides(grid [][]rune) (farm map[Corner]map[Side]struct{}) {
	farm = make(map[Corner]map[Side]struct{})
	maxX, maxY := len(grid[0]), len(grid)
	corners := []Corner{{-1, 0}, {0, -1}, {0, 1}, {1, 0}}
	for i := 0; i < maxX; i++ {
		for x := 0; x < maxY; x++ {
			current := Corner{i, x}
			if _, exists := farm[current]; !exists {
				farm[current] = make(map[Side]struct{})
			}
			for key, corner := range corners {
				cornerI, cornerX := i+corner.x, x+corner.y
				if !(cornerI >= 0 && cornerI < maxX && cornerX >= 0 && cornerX < maxY) {
					orientation := "vh"[int(math.Abs(float64(corner.x))):][:1]
					farm[current][Side{i + cast.ToInt(math.Max(cast.ToFloat(corner.x), 0)), x + cast.ToInt(math.Max(cast.ToFloat(corner.y), 0)), key, orientation}] = struct{}{}
				}
			}
		}
	}
	return farm
}

func getRegion(farm map[Corner]map[Corner]struct{}, corner Corner) (region map[Corner]struct{}) {
	region = make(map[Corner]struct{})
	remaining := map[Corner]struct{}{corner: {}}

	for len(remaining) > 0 {
		var current Corner
		for p := range remaining {
			current = p
			break
		}
		delete(remaining, current)
		region[current] = struct{}{}

		for neighbor := range farm[current] {
			if _, inRegion := region[neighbor]; !inRegion {
				remaining[neighbor] = struct{}{}
			}
		}
	}

	return region
}

func getRegions(farm map[Corner]map[Corner]struct{}, full map[Corner]struct{}) (regions []map[Corner]struct{}) {
	for len(full) > 0 {
		var current Corner
		for c := range full {
			current = c
			break
		}

		region := getRegion(farm, current)
		regions = append(regions, region)

		for c := range region {
			delete(full, c)
		}
	}

	return regions
}

func calculatePerimeter(region map[Corner]struct{}, farm map[Corner]map[Corner]struct{}) int {
	perimeter := 0
	for c := range region {
		perimeter += 4 - len(farm[c])
	}
	return perimeter
}

func calculateArea(region map[Corner]struct{}) int {
	return len(region)
}

func calculateSides(region map[Corner]struct{}, sides map[Corner]map[Side]struct{}) int {
	sideCount := 0
	for c := range region {
		sideCount += len(sides[c])
	}
	return sideCount
}

func countRegions(farm map[Corner]map[Corner]struct{}, regions []map[Corner]struct{}) (answer int) {
	for _, region := range regions {
		answer += calculatePerimeter(region, farm) * calculateArea(region)
	}
	return answer
}

func countSides(sides map[Corner]map[Side]struct{}, regions []map[Corner]struct{}) (answer int) {
	for _, region := range regions {
		answer += calculateSides(region, sides) * calculateArea(region)
	}
	return answer
}
