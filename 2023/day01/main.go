package main

import (
	_ "embed"
	"flag"
	"fmt"
	"regexp"
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
	_ = parsed
	var values []int
	result := 0 
	re := regexp.MustCompile("\\d")

	for _, line := range parsed {
		t := re.FindAllString(line, -1)

		stringValue := t[0]
		if len(t) > 1 {
			stringValue = stringValue + t[len(t)-1]
		} else {
			stringValue = stringValue + t[0]
		}
		values = append(values, cast.ToInt(stringValue))
	}
	
	for _, v := range values {
		result += v
	}

	return result
}

func part2(input string) int {
	parsed := parseInput(input)
	_ = parsed
	var values []int
	result := 0 
	digitString := "one|two|three|four|five|six|seven|eight|nine"
	re := regexp.MustCompile("("+digitString+")|\\d")
	rereverse := regexp.MustCompile("("+mathy.Reverse(digitString)+")|\\d")

	for _, line := range parsed {
		t := re.FindAllString(line, -1)
		fmt.Println(t) 

		stringValue := cast.ToStringNumber(t[0])
		
		// second digit
		s := rereverse.FindAllString(mathy.Reverse(line), -1)
		// fmt.Println(mathy.Reverse(s[0]))

		stringValue = stringValue + cast.ToStringNumber(mathy.Reverse(s[0]))

		// fmt.Println(stringValue) 
		values = append(values, cast.ToInt(stringValue))
	}
	
	for _, v := range values {
		result += v
	}

	return result
}

func parseInput(input string) (ans []string) {
	for _, line := range strings.Split(input, "\n") {
		ans = append(ans, line)
	}
	return ans
}

