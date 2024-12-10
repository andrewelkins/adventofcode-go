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
	parsed := parseInput(input)[0] // one long string
	var disk []int
	var diskValue int = 0
	var start int = 0
	var end int = 0
	var result int = 0

	for key, char := range strings.Split(parsed, "") {
		count := cast.ToInt(char)

		for i := 0; i < count; i++ {
			if key%2 == 0 {
				disk = append(disk, diskValue)
			} else {
				disk = append(disk, -1)
			}
		}
		if key%2 == 0 {
			diskValue++
		}
	}

	end = len(disk) - 1
	for start < end {
		if disk[start] == -1 {
			for disk[end] == -1 {
				end--
			}
			if end < start {
				break
			}
			disk[start] = disk[end]
			disk[end] = -1
		}
		start++
	}

	for i, item := range disk {
		if item != -1 {
			result += item * i
		}
	}

	return result
}

func part2(input string) int {
	parsed := parseInput(input)[0] // one long string
	var disk []string
	var end int = 0
	var result int = 0

	for key, char := range strings.Split(parsed, "") {
		count := cast.ToInt(char)

		for i := 0; i < count; i++ {
			if key%2 == 0 {
				disk = append(disk, cast.ToString(count)+",0")
			} else {
				disk = append(disk, cast.ToString(count)+",1")
			}
		}
	}

	end = len(disk) - 1
	for i := end; i >= 0; i-- {
		t := strings.Split(disk[i], ",")
		count, isMemory := t[0], t[1]
		if isMemory == "1" {
			continue
		} else {
			for j := 0; j <= end; j++ {
				s := strings.Split(disk[j], ",")
				fmt.Println(s)
				searchCount, searchIsMemory := s[0], s[1]
				if searchIsMemory == "0" && count <= searchCount {
					fmt.Println(s)
					starting := disk[0:j]
					ending := disk[j+1 : end]
					dataCount := cast.ToInt(searchCount) - cast.ToInt(count)
					starting = append(disk, count+",0")
					if dataCount > 0 {
						starting = append(disk, cast.ToString(dataCount)+",1")
					}

					disk = append(starting, ending...)

					break
				}
			}
		}
	}

	newDisk := []int{}
	for _, item := range disk {
		count, isMemory := cast.ToInt(strings.Split(item, ",")[0]), strings.Split(item, ",")[1]
		for i := 0; i < count; i++ {
			if isMemory == "1" {
				newDisk = append(newDisk, -1)
			} else {
				newDisk = append(newDisk, count)
			}
		}
	}

	for i, item := range newDisk {
		if item != -1 {
			result += item * i
		}
	}

	return result
}

func parseInput(input string) (ans []string) {
	for _, line := range strings.Split(input, "\n") {
		ans = append(ans, line)
	}
	return ans
}
