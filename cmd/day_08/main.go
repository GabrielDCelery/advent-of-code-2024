package main

import (
	"fmt"
	"log"

	"github.com/GabrielDCelery/advent-of-code-2024/internals/day_08"
)

func main() {
	part1Solution, err := day_08.SolveDay8Part1()

	if err != nil {
		log.Fatalln(err)
	}

	fmt.Printf("Day 8, part 1 solution: %d\n", part1Solution)

	// part2Solution, err := day_07.SolveDay7Part2()
	//
	// if err != nil {
	// 	log.Fatalln(err)
	// }
	//
	// fmt.Printf("Day 7, part 2 solution: %d\n", part2Solution)
}
