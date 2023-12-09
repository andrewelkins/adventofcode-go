package main

import (
	_ "embed"
	"flag"
	"fmt"
	"strings"
	"sort"

	"github.com/andrewelkins/adventofcode-go/cast"
	"github.com/andrewelkins/adventofcode-go/util"
)

//go:embed input.txt
var input string

var jWild bool

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
	jWild = false
	parsed := parseInput(input)
	total := 0
	handsGrouped := []string{}

	for _, line := range parsed {
		lineSplit := strings.Split(line, " ")
		hand := strings.Map(rules, lineSplit[0])
		handValue := NewStringSet(strings.Split(hand, ""))
		cc := cast.ToString(cardType(handValue))+" "+hand + " " + lineSplit[1]
		handsGrouped = append(handsGrouped, cc)
	}

	sort.Sort(sort.StringSlice(handsGrouped))

	for x, line := range handsGrouped {
		lineSplit := strings.Split(line, " ")
		total += cast.ToInt(lineSplit[2]) * (x+1)
	}

	return total
}

func part2(input string) int {
	jWild = true
	parsed := parseInput(input)
	total := 0
	handsGrouped := []string{}

	for _, line := range parsed {
		lineSplit := strings.Split(line, " ")
		hand := strings.Map(rules, lineSplit[0])
		handValue := NewStringSet(strings.Split(hand, ""))
		cc := cast.ToString(cardType(handValue))+" "+hand + " " + lineSplit[1]
		fmt.Println(cc)
		handsGrouped = append(handsGrouped, cc)
	}

	sort.Sort(sort.StringSlice(handsGrouped))

	fmt.Println(handsGrouped)

	for x, line := range handsGrouped {
		lineSplit := strings.Split(line, " ")
		total += cast.ToInt(lineSplit[2]) * (x+1)
	}

	return total
}

func parseInput(input string) (ans []string) {
	for _, line := range strings.Split(input, "\n") {
		ans = append(ans, line)
	}
	return ans
}

type StringSet map[string]int

// NewStringSet initializes a set with the values form the input string slice
func NewStringSet(stringSlice []string) StringSet {
	set := StringSet{}
	for _, v := range stringSlice {
		if set.Has(v) {
			set[v] += 1 
		} else {
			set[v] = 1
		}
	}
	// fmt.Println(set)
	return set
}

// Has returns true if the value if found in the underlying set
func (s StringSet) Has(val string) bool {
	_, ok := s[val]
	return ok
}

// Add a value to the set
func (s StringSet) Add(val string) {
	s[val] = 1
}

// Remove a value from the set
func (s StringSet) Remove(val string) {
	delete(s, val)
}

// Keys returns a slice of all keys in the set
func (s StringSet) Keys() []string {
	var keys []string
	for k := range s {
		keys = append(keys, k)
	}
	return keys
}

// Return in of max number of keys
func (s StringSet) Max() int {
	max := 0
	for k, v := range s {
		_ = k
		if v > max {
			max = v
		}
	}
	return max
}

func cardType(hand StringSet) int {
	handLength := hand.Keys()
	max := hand.Max()

	if jWild && hand.Has("1") {
		if len(handLength) != 1 {
			if len(handLength) == 4 {
				max = 3
			} else if len(handLength) == 3{
				max += 1
			} else if len(handLength) == 2{
				max += 1
			}
			handLength = handLength[:len(handLength)-1]
		}
	}

	if len(handLength) == 5 {
		return 1 //"high-card"
	} else if len(handLength) == 4 {
		return 2 //"one-pair"
	} else if len(handLength) == 3 {
		if max == 2 {
			return 3 //"two-pair"
		}
		return 4 //"three-of-a-kind"
	} else if len(handLength) == 2 {
		if max == 3 {
			return 5 //"full-house"
		}
		return 6 //"four-of-a-kind"
	} else if len(handLength) == 1 {
		return 7 //"five-of-a-kind"
	}
	return 1
}

func rules(r rune) rune {
    switch r {
		case 'T':
			return 'A'
		case 'J':
			return '1'
		case 'Q':
			return 'C'
		case 'K':
			return 'D'
		case 'A':
			return 'E'
		default:
			return r
    }
}