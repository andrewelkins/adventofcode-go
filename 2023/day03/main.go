package main

import (
	_ "embed"
	"flag"
	"fmt"
	// "maps"
	"strings"
    "strconv"
	"regexp"
	// "slices"

	"github.com/andrewelkins/adventofcode-go/cast"
	"github.com/andrewelkins/adventofcode-go/mathy"
	"github.com/andrewelkins/adventofcode-go/util"
)

//go:embed input.txt
var input string

var numbersAndGears map[string][]int

var board []string


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
	numSum := 0
	re := regexp.MustCompile("\\d+")
	numbersAndGears = make(map[string][]int)

	// TIL substring in golang returns the ascii code of the character

	for y, line := range parsed {
		// Find matches for numbers
		t := re.FindAllString(line, -1)
		numberIndexes := re.FindAllStringSubmatchIndex(line, -1)
		fmt.Println(numberIndexes)
		for i, match := range t {
			// match neightbors y above, x left, y below, x right
			xLeft := numberIndexes[i][0]-1
			xRight := numberIndexes[i][1]
			matchNum := cast.ToInt(match)
			if matchNeightbors(y-1, xLeft, y+1, xRight, matchNum, parsed) {
    			numSum += cast.ToInt(match)
			} 
		}
	}
	return numSum
}

func part2(input string) int {
	parsed := parseInput(input)
	numSum := 0
	gearSum := 0
	re := regexp.MustCompile("\\d+")
	numbersAndGears = make(map[string][]int)

	for y, line := range parsed {
		// Find matches for numbers
		t := re.FindAllString(line, -1)
		numberIndexes := re.FindAllStringSubmatchIndex(line, -1)
		fmt.Println(numberIndexes)
		for i, match := range t {
			// match neightbors y above, x left, y below, x right
			xLeft := numberIndexes[i][0]-1
			xRight := numberIndexes[i][1]
			matchNum := cast.ToInt(match)
			if matchNeightbors(y-1, xLeft, y+1, xRight, matchNum, parsed) {
    			numSum += cast.ToInt(match)
			}
		}
	}
	for _, value := range numbersAndGears {
		if len(value) == 2 {
			gearSum += value[0] * value[1]
		}
	}
	return gearSum
}

func parseInput(input string) (ans []string) {
	for _, line := range strings.Split(input, "\n") {
		ans = append(ans, line)
	}
	return ans
}

func isNumeric(s string) bool {
    _, err := strconv.ParseFloat(s, 64)
    return err == nil
}

func findSpecialCharacters(board []string) []string {
	var specialCharacters []string
	for _, line := range board {
		characters := strings.Split(line, "")
		for _, character := range characters {
			if character != "." && !isNumeric(character) {
				specialCharacters = append(specialCharacters, character)
			}
		}
	}
	return removeDuplicate(specialCharacters)
}

func removeDuplicate[T string | int](sliceList []T) []T {
    allKeys := make(map[T]bool)
    list := []T{}
    for _, item := range sliceList {
        if _, value := allKeys[item]; !value {
            allKeys[item] = true
            list = append(list, item)
        }
    }
    return list
}

func matchNeightbors(yAbove int, xLeft int, yBelow int, xRight int, number int, board []string) bool {

	yRange := mathy.MakeRange(yAbove, yBelow)
	for _, y := range yRange {
		xRange := mathy.MakeRange(xLeft, xRight)
		for _, x := range xRange {
			if y >= 0 && y < len(board) && x >= 0 && x < len(board[y]) {
				if !strings.Contains("0123456789.", string(board[y][x])) {
					if strings.Contains("*", string(board[y][x])) {
						numKey := "y" + strconv.Itoa(y) + "x" + strconv.Itoa(x)
						numbersAndGears[numKey] = append(numbersAndGears[numKey], number)
					}
					return true
				}
			}
		}
	}
	return false
}