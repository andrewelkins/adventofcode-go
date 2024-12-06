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

	return 0
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
