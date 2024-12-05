package main

import (
	"fmt"
	"log"

	"github.com/GabrielDCelery/advent-of-code-2024/internals/day_03"
)

func main() {
	part1Solution, err := day_03.SoveDay3Part1()

	if err != nil {
		log.Fatalln(err)
	}

	fmt.Printf("Day 3, part 1 solution: %d\n", part1Solution)

	part2Solution, err := day_03.SolveDay3Part2()

	if err != nil {
		log.Fatalln(err)
	}

	fmt.Printf("Day 3, part 2 solution: %d\n", part2Solution)
}
