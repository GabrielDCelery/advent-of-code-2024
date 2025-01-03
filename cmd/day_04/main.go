package main

import (
	"fmt"
	"log"

	"github.com/GabrielDCelery/advent-of-code-2024/internals/day_04"
)

func main() {
	part1Solution, err := day_04.SolveDay4Part1()

	if err != nil {
		log.Fatalln(err)
	}

	fmt.Printf("Day 4, part 1 solution: %d\n", part1Solution)

	part2Solution, err := day_04.SolveDay4Part2()

	if err != nil {
		log.Fatalln(err)
	}

	fmt.Printf("Day 4, part 2 solution: %d\n", part2Solution)
}
