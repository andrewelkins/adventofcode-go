package main

import "github.com/andrewelkins/adventofcode-go/scripts/aoc"

func main() {
	day, year, cookie := aoc.ParseFlags()
	aoc.GetPrompt(day, year, cookie)
}
