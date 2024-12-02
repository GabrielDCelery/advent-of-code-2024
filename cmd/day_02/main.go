package main

import (
	"fmt"
	"log"

	"github.com/GabrielDCelery/advent-of-code-2024/internals"
)

func main() {
	part1Solution, err := internals.SolveDay2Part1()

	if err != nil {
		log.Fatalln(err)
	}

	fmt.Printf("Day 2, part 1 solution: %d\n", part1Solution)

	part2Solution, err := internals.SolveDay2Part2()

	if err != nil {
		log.Fatalln(err)
	}

	fmt.Printf("Day 2, part 2 solution: %d\n", part2Solution)
}
