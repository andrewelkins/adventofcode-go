package main

import (
	"bufio"
	_ "embed"
	"flag"
	"fmt"
	"os"
	"slices"
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

	parsedInput := parseInput(input)
	if part == 1 {
		ans := part1(parsedInput)
		util.CopyToClipboard(fmt.Sprintf("%v", ans))
		fmt.Println("Output:", ans)
	} else {
		ans := part2(parsedInput)
		util.CopyToClipboard(fmt.Sprintf("%v", ans))
		fmt.Println("Output:", ans)
	}
}

func part1(input []string) int {
	// Going to try scanner. Thanks Ben!
	inputFile, _ := os.Open("input.txt")
	var rules = []string{}
	var pages = []string{}
	var sumMiddleNumbers int = 0

	defer inputFile.Close()
	scanner := bufio.NewScanner(inputFile)

	for scanner.Scan() {
		if scanner.Text() == "" {
			break
		}
		rules = append(rules, scanner.Text())
	}
	for scanner.Scan() {
		pages = append(pages, scanner.Text())
	}

	for _, page := range pages {
		pageArray := strings.Split(page, ",")
		if validateRules(pageArray, rules) {
			middleNumber := cast.ToInt(pageArray[len(pageArray)/2])
			sumMiddleNumbers += middleNumber
		}
	}

	return sumMiddleNumbers
}

func part2(input []string) int {
	inputFile, _ := os.Open("input.txt")
	var rules = []string{}
	var pages = []string{}
	var sumMiddleNumbers int = 0
	// incorrectOrderedPages := [][]string{}

	defer inputFile.Close()
	scanner := bufio.NewScanner(inputFile)

	for scanner.Scan() {
		if scanner.Text() == "" {
			break
		}
		rules = append(rules, scanner.Text())
	}
	for scanner.Scan() {
		pages = append(pages, scanner.Text())
	}

	for _, page := range pages {
		pageArray := strings.Split(page, ",")
		if !validateRules(pageArray, rules) {
			orderedPages := reorderPages(pageArray, rules)
			middleNumber := cast.ToInt(orderedPages[len(orderedPages)/2])
			sumMiddleNumbers += middleNumber
		}
	}

	return sumMiddleNumbers
}

func parseInput(input string) (ans []string) {
	ans = append(ans, strings.Split(input, "\n")...)
	return ans
}

func validateRules(pages []string, rules []string) bool {
	for _, rule := range rules {
		var rs []string = strings.Split(rule, "|")
		var indexBefore int = indexOf(pages, rs[0])
		var indexAfter int = indexOf(pages, rs[1])
		// If the rule is not in the correct order, return false
		if indexAfter != -1 && indexBefore > indexAfter {
			return false
		}
	}
	return true
}

func indexOf(haystack []string, needle string) int {
	for i, v := range haystack {
		if v == needle {
			return i
		}
	}
	return -1
}

func reorderPages(pages []string, rules []string) []string {
	// Process rule by rule until all rules are satisfied :( Brute force! Do .. while
	for !validateRules(pages, rules) {
		for _, rule := range rules {
			var rs []string = strings.Split(rule, "|")
			var indexBefore int = indexOf(pages, rs[0])
			var indexAfter int = indexOf(pages, rs[1])
			if indexAfter != -1 && indexBefore > indexAfter {
				pages = slices.Replace(pages, indexBefore, indexBefore+1, rs[1])
				pages = slices.Replace(pages, indexAfter, indexAfter+1, rs[0])
			}
		}
	}
	return pages
}
