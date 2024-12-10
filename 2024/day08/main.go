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

	maxY := len(parsed)
	maxX := len(parsed[0])
	charMap := make(map[string][][]int)
	antinodes := make(map[string]bool)
	count := 0

	fmt.Print(maxX, maxY, "\n")

	// result := 0
	for y, line := range parsed {
		for x, char := range strings.Split(line, "") {
			if char != "." {
				value, exists := charMap[char]
				if exists {
					// check for antinode
					for _, node := range value {
						diffY := y - node[0]
						diffX := x - node[1]
						// Check Y, then X
						if node[0]-diffY >= 0 && node[1]-diffX > 0 && node[1]-diffX < maxX {
							if !antinodes[cast.ToString(node[0]-diffY)+","+cast.ToString(node[1]-diffX)] {
								antinodes[cast.ToString(node[0]-diffY)+","+cast.ToString(node[1]-diffX)] = true
								count++
							}
						}
						if y+diffY <= maxY && x+diffX > 0 && x+diffX < maxX {
							if !antinodes[cast.ToString(y+diffY)+","+cast.ToString(x+diffX)] {
								antinodes[cast.ToString(y+diffY)+","+cast.ToString(x+diffX)] = true
								count++
							}
						}
					}
				}
				charMap[char] = append(charMap[char], []int{y, x})
			}
		}
	}

	return count
}

func part2(input string) int {
	parsed := parseInput(input)

	maxY := len(parsed) - 1
	maxX := len(parsed[0]) - 1
	charMap := make(map[string][][]int)
	antinodes := make(map[string]bool)
	count := 0

	for y, line := range parsed {
		for x, char := range strings.Split(line, "") {
			if char != "." {
				value, exists := charMap[char]
				if exists {
					// check for antinode
					for _, node := range value {
						diffY := y - node[0]
						diffX := x - node[1]
						loops := 50

						// up
						less := node[1] - diffX
						if node[0]-diffY >= 0 && less >= 0 && less <= maxX {
							for i := 1; i < loops; i++ {
								ymult := node[0] - (diffY * i)
								xmult := node[1] - (diffX * i)
								antinodeString := cast.ToString(ymult) + "," + cast.ToString(xmult)
								if !antinodes[antinodeString] && ymult <= maxY && xmult <= maxX && ymult >= 0 && xmult >= 0 {
									antinodes[antinodeString] = true
								}
							}
						}
						// down
						more := x + diffX
						if y+diffY <= maxY && more >= 0 && more <= maxX {
							for i := 1; i < loops; i++ {
								ymult := y + (diffY * i)
								xmult := x + (diffX * i)
								antinodeString := cast.ToString(ymult) + "," + cast.ToString(xmult)
								if !antinodes[antinodeString] && ymult <= maxY && xmult <= maxX && ymult >= 0 && xmult >= 0 {
									antinodes[antinodeString] = true
								}
							}
						}
					}
				}
				antinodes[cast.ToString(y)+","+cast.ToString(x)] = true
				charMap[char] = append(charMap[char], []int{y, x})
			}
		}
	}
	for k, _ := range antinodes {
		_ = k
		count++
	}

	return count
}

func parseInput(input string) (ans []string) {
	for _, line := range strings.Split(input, "\n") {
		ans = append(ans, line)
	}
	return ans
}

func leastNumberLoops(multY, multX int) int {
	loops := 0
	if multY > multX {
		loops = multY
	} else {
		loops = multX
	}
	return loops
}
