package main

import (
	_ "embed"
	"flag"
	"fmt"
	"strings"
	"time"

	"github.com/andrewelkins/adventofcode-go/util"
)

//go:embed input.txt
var input string

var preventRecursion = 0

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

	start := time.Now()

	parsed := parseInput(input)
	_ = parsed
	xcount := 0
	var do bool = true
	lab := make(map[int][]string)
	cursor := [2]int{0, 0}

	for i, line := range parsed {
		lab[i] = strings.Split(line, "")
		start := indexOf(lab[i], "^") // ^ > v <  u l r d
		if start != -1 && cursor[0] == 0 {
			cursor = [2]int{i, start}
		}
	}

	for do {
		if !cursorCheck(lab, cursor) {
			do = false
			break
		}
		switch lab[cursor[0]][cursor[1]] {
		case "^":
			lab[cursor[0]][cursor[1]] = "X"
			if lab[cursor[0]-1][cursor[1]] == "#" {
				cursor[1]++
				if !cursorCheck(lab, cursor) {
					do = false
				} else {
					lab[cursor[0]][cursor[1]] = ">"
				}
			} else {
				cursor[0]--
				if !cursorCheck(lab, cursor) {
					do = false
				} else {
					lab[cursor[0]][cursor[1]] = "^"
				}
			}
		case ">":
			lab[cursor[0]][cursor[1]] = "X"
			if lab[cursor[0]][cursor[1]+1] == "#" {
				cursor[0]++
				if !cursorCheck(lab, cursor) {
					do = false
				} else {
					lab[cursor[0]][cursor[1]] = "v"
				}
			} else {
				cursor[1]++
				if !cursorCheck(lab, cursor) {
					do = false
				} else {
					lab[cursor[0]][cursor[1]] = ">"
				}
			}
		case "v":
			lab[cursor[0]][cursor[1]] = "X"
			if cursor[0]+1 >= len(lab[cursor[0]]) {
				do = false
				break
			}
			if lab[cursor[0]+1][cursor[1]] == "#" {
				cursor[1]--
				if !cursorCheck(lab, cursor) {
					do = false
				} else {
					lab[cursor[0]][cursor[1]] = "<"
				}
			} else {
				cursor[0]++
				if !cursorCheck(lab, cursor) {
					do = false
				} else {
					lab[cursor[0]][cursor[1]] = "v"
				}
			}
		case "<":
			lab[cursor[0]][cursor[1]] = "X"
			if lab[cursor[0]][cursor[1]-1] == "#" {
				cursor[0]--
				if !cursorCheck(lab, cursor) {
					do = false
				} else {
					lab[cursor[0]][cursor[1]] = "^"
				}
			} else {
				cursor[1]--
				if !cursorCheck(lab, cursor) {
					do = false
				} else {
					lab[cursor[0]][cursor[1]] = "<"
				}
			}
		default:
			do = false
		}
		// // fmt.Print("\033[2J") // clear the screen
		// fmt.Print("====================================\n")
		// for i, _ := range lab[0] {
		// 	fmt.Println(lab[i])
		// }

	}

	for i, _ := range lab[0] {
		for _, char := range lab[i] {
			if char == "X" {
				xcount++
			}
		}
	}

	elapsed := time.Since(start)
	fmt.Printf("page took %s \n", elapsed)

	return xcount
}

func part2(input string) int {

	start := time.Now()

	parsed := parseInput(input)
	_ = parsed
	var do bool = true
	lab := make(map[int][]string)
	cursor := [2]int{0, 0}
	circle := 1

	for i, line := range parsed {
		lab[i] = strings.Split(line, "")
		start := indexOf(lab[i], "^") // ^ > v <  u l r d
		if start != -1 && cursor[0] == 0 {
			cursor = [2]int{i, start}
		}
	}
	startPosition := cursor
	uniqueSpots := make(map[string]bool)

	for do {
		preventRecursion = 0
		if !cursorCheck(lab, cursor) {
			do = false
			break
		}
		circleCursor := cursor
		circleLab := lab
		switch lab[cursor[0]][cursor[1]] {
		case "^":
			lab[cursor[0]][cursor[1]] = "^"
			if lab[cursor[0]-1][cursor[1]] == "#" {
				cursor[1]++
				if !cursorCheck(lab, cursor) {
					do = false
				} else {
					lab[cursor[0]][cursor[1]] = ">"
				}
			} else {
				cursor[0]--
				if !cursorCheck(lab, cursor) {
					do = false
				} else {
					fmt.Print("cursorLAB", circleCursor, "\n")
					circleLab[cursor[0]][cursor[1]] = "O"
					if cursorMove(circleLab, circleCursor, ">") {
						uniqueSpots[fmt.Sprintf("%v", cursor)] = true
						if startPosition != cursor {
							circle++
						}
					}
					lab[cursor[0]][cursor[1]] = "^"
				}
			}
		case ">":
			lab[cursor[0]][cursor[1]] = ">"
			if lab[cursor[0]][cursor[1]+1] == "#" {
				cursor[0]++
				if !cursorCheck(lab, cursor) {
					do = false
				} else {
					lab[cursor[0]][cursor[1]] = "v"
				}
			} else {
				cursor[1]++
				if !cursorCheck(lab, cursor) {
					do = false
				} else {

					fmt.Print("cursorLAB", circleCursor, "\n")
					circleLab[cursor[0]][cursor[1]] = "O"
					if cursorMove(circleLab, circleCursor, "v") {
						uniqueSpots[fmt.Sprintf("%v", cursor)] = true
						if startPosition != cursor {
							circle++
						}
					}
					lab[cursor[0]][cursor[1]] = ">"
				}
			}
		case "v":
			lab[cursor[0]][cursor[1]] = "v"
			if cursor[0]+1 >= len(lab[cursor[0]]) {
				do = false
				break
			}
			if lab[cursor[0]+1][cursor[1]] == "#" {
				cursor[1]--
				if !cursorCheck(lab, cursor) {
					do = false
				} else {
					lab[cursor[0]][cursor[1]] = "<"
				}
			} else {
				cursor[0]++
				if !cursorCheck(lab, cursor) {
					do = false
				} else {

					fmt.Print("cursorLAB", circleCursor, "\n")
					circleLab[cursor[0]][cursor[1]] = "O"
					if cursorMove(circleLab, circleCursor, "<") {
						uniqueSpots[fmt.Sprintf("%v", cursor)] = true
						if startPosition != cursor {
							circle++
						}
					}
					lab[cursor[0]][cursor[1]] = "v"
				}
			}
		case "<":
			lab[cursor[0]][cursor[1]] = "<"
			if lab[cursor[0]][cursor[1]-1] == "#" {
				cursor[0]--
				if !cursorCheck(lab, cursor) {
					do = false
				} else {
					lab[cursor[0]][cursor[1]] = "^"
				}
			} else {
				cursor[1]--
				if !cursorCheck(lab, cursor) {
					do = false
				} else {

					fmt.Print("cursorLAB", circleCursor, "\n")
					circleLab[cursor[0]][cursor[1]] = "O"
					if cursorMove(circleLab, circleCursor, "^") {
						uniqueSpots[fmt.Sprintf("%v", cursor)] = true
						if startPosition != cursor {
							circle++
						}
					}
					lab[cursor[0]][cursor[1]] = "<"
				}
			}
		default:
			do = false
		}
	}

	cc := 0
	for i, _ := range uniqueSpots {
		fmt.Print(i, "\n")
		cc++
	}

	elapsed := time.Since(start)
	fmt.Printf("page took %s \n", elapsed)
	fmt.Print("circle", circle, "\n")
	return circle
}

func parseInput(input string) (ans []string) {
	for _, line := range strings.Split(input, "\n") {
		ans = append(ans, line)
	}
	return ans
}

func indexOf(haystack []string, needle string) int {
	for i, v := range haystack {
		if v == needle {
			return i
		}
	}
	return -1
}

func cursorCheck(lab map[int][]string, cursor [2]int) bool {
	if cursor[0] < 0 || cursor[0] > len(lab) || cursor[1] < 0 || cursor[1] > len(lab[cursor[0]]) {
		return false
	}
	return true
}

func cursorMove(lab map[int][]string, cursor [2]int, direction string) bool {
	if preventRecursion > 1000 {
		return false
	}
	preventRecursion++
	// fmt.Println("cursorMove", direction)
	// for i, _ := range lab[0] {
	// 	fmt.Println(lab[i])
	// }
	switch direction {
	case "^":
		for i := cursor[0]; i >= 0; i-- {
			y := lab[i][cursor[1]]
			// fmt.Println("y", i, cursor[1], y)
			if y == "O" {
				return true
			}
			if y == "#" {
				cursor[0] = i + 1
				if lab[cursor[0]][cursor[1]+1] == "#" {
					cursor[0]++
					return cursorMove(lab, cursor, "v")
				} else {
					return cursorMove(lab, cursor, ">")
				}
			}
		}
	case ">":
		for i := cursor[1]; i < len(lab[cursor[0]]); i++ {
			x := lab[cursor[0]][i]
			// fmt.Println("x", cursor[0], i, x)
			if x == "O" {
				return true
			}
			if x == "#" {
				cursor[1] = i - 1
				if lab[cursor[0]+1][cursor[1]] == "#" {
					cursor[1]++
					return cursorMove(lab, cursor, "<")
				} else {
					return cursorMove(lab, cursor, "v")
				}
			}
		}
	case "v":
		for i := cursor[0]; i < len(lab); i++ {
			y := lab[i][cursor[1]]
			// fmt.Println("y", i, cursor[1], y)
			if y == "O" {
				return true
			}
			if y == "#" {
				cursor[0] = i - 1
				if lab[cursor[0]][cursor[1]-1] == "#" {
					cursor[1]++
					return cursorMove(lab, cursor, "v")
				} else {
					return cursorMove(lab, cursor, "<")
				}
			}
		}
	case "<":
		for i := cursor[1]; i >= 0; i-- {
			x := lab[cursor[0]][i]
			// fmt.Println("x", cursor[0], i, x)
			if x == "O" {
				return true
			}
			if x == "#" {
				cursor[1] = i + 1
				if lab[cursor[0]-1][cursor[1]] == "#" {
					cursor[0]++
					return cursorMove(lab, cursor, "v")
				} else {
					return cursorMove(lab, cursor, "^")
				}
			}
		}
	}
	// for i, _ := range lab[0] {
	// 	fmt.Println(lab[i])
	// }
	return false
}
