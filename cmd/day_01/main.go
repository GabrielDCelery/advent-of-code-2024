package main

import (
	"fmt"
	"log"

	"github.com/GabrielDCelery/advent-of-code-2024/internals/day_01"
)

func main() {
	part1Solution, err := day_01.SolveDay1Part1()

	if err != nil {
		log.Fatalln(err)
	}

	fmt.Printf("Day 1, part 1 solution: %d\n", part1Solution)

	part2Solution, err := day_01.SolveDay1Part2()

	if err != nil {
		log.Fatalln(err)
	}

	fmt.Printf("Day 1, part 2 solution: %d\n", part2Solution)
}
