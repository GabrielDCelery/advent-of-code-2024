package main

import (
	"fmt"
	"log"

	"github.com/GabrielDCelery/advent-of-code-2024/internals/day_11"
)

func main() {
	part1Solution, err := day_11.SolveDay11Part1()

	if err != nil {
		log.Fatalln(err)
	}

	fmt.Printf("Day 11, part 1 solution: %d\n", part1Solution)

	// part2Solution, err := day_11.SolveDay11Part2()
	//
	// if err != nil {
	// 	log.Fatalln(err)
	// }
	//
	// fmt.Printf("Day 11, part 2 solution: %d\n", part2Solution)
}
