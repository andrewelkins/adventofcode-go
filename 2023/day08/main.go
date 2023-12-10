package main

import (
	_ "embed"
	"flag"
	"fmt"
	"strings"
	"regexp"
	"math"

	"github.com/andrewelkins/adventofcode-go/cast"
	"github.com/andrewelkins/adventofcode-go/util"

)

//go:embed input.txt
var input string
var position_end string

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
	position_end = "ZZZ"
	parsed := parseInput(input)
	instructions := strings.Repeat(parsed[0], 100) 
	node_input := parsed[2:]
	re := regexp.MustCompile("\\w+")
	node_slices := [][]string{}
	node_map := make(map[string][]string)
	position := "AAA"


	for _, line := range node_input {
		node_slices = append(node_slices, re.FindAllString(line, -1))
	}
	for _, line := range node_slices {
		fmt.Println("line",line)
		node_map[line[0]] = line[1:]
	}


	return number_steps(instructions, node_map, position)
}

func part2(input string) int {
	position_end = "Z"
	parsed := parseInput(input)
	instructions := strings.Repeat(parsed[0], 100) 
	node_input := parsed[2:]
	re := regexp.MustCompile("\\w+")
	node_slices := [][]string{}
	node_map := make(map[string][]string)

	for _, line := range node_input {
		node_slices = append(node_slices, re.FindAllString(line, -1))
	}
	for _, line := range node_slices {
		fmt.Println("line",line)
		node_map[line[0]] = line[1:]
	}

	num_as := []int{}
	for key, _ := range node_map {
		if cast.ToString(key[2]) == "A" {
			num_as = append(num_as, number_steps(instructions, node_map, key))
		}
	}
	divided := reduce(num_as, lcm, 0)

	return divided
}

func parseInput(input string) (ans []string) {
	for _, line := range strings.Split(input, "\n") {
		ans = append(ans, line)
	}
	return ans
}


func lcm(a int, b int) int {
    return cast.ToInt(math.Floor(float64((a * b) / GCDEuclidean(a, b))))
}

func number_steps(instructions string, node_map map[string][]string, position string) int {
	pod := 0
	for x, instruction := range instructions {
		lr := 0
		inst := cast.ToString(instruction) 
		if inst == "R" {
			lr = 1
		}
		position = node_map[position][lr]
		if position == position_end {
			pod = x + 1
			fmt.Println("pod",pod)
			return pod
		}
	}
	return pod
}

func GCDEuclidean(a int, b int) int {
	for a != b {
		if a > b {
			a -= b
		} else {
			b -= a
		}
	}

	return a
}

func reduce[T, M any](s []T, f func(M, T) M, initValue M) M {
    acc := initValue
    for _, v := range s {
        acc = f(acc, v)
    }
    return acc
}