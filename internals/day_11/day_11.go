package day_11

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Stone struct {
	value      int
	blinkCount int
	child      *Stone
	parent     *Stone
}

func NewStone(value int, blinkCount int, parent *Stone) *Stone {
	return &Stone{
		value:      value,
		blinkCount: blinkCount,
		parent:     parent,
		child:      nil,
	}
}

func blinkNTimesAndCountNumberOfStones(input string, targetBlinkCount int) (int, error) {
	numbersAsStr := strings.Split(strings.TrimSpace(input), " ")

	numbers := []int{}

	for _, numberAsStr := range numbersAsStr {
		number, err := strconv.Atoi(numberAsStr)
		if err != nil {
			return 0, err
		}
		numbers = append(numbers, number)
	}

	stoneCount := 0

	for _, number := range numbers {
		currentStone := NewStone(number, 0, nil)

		for {
			if currentStone.blinkCount == targetBlinkCount {
				stoneCount += 1
				if currentStone.parent == nil {
					break
				} else {
					currentStone = currentStone.parent
					continue
				}
			}

			if currentStone.value == 0 {
				currentStone.value = 1
				currentStone.blinkCount += 1
				continue
			}

			str := fmt.Sprintf("%d", currentStone.value)
			if len(str)%2 == 0 {
				mid := len(str) / 2
				left, err := strconv.Atoi(str[mid:])
				if err != nil {
					return 0, err
				}
				right, err := strconv.Atoi(str[:mid])
				if err != nil {
					return 0, err
				}
				currentStone.value = left
				currentStone.blinkCount += 1
				currentStone = NewStone(right, currentStone.blinkCount, currentStone)
				continue
			}

			currentStone.value = currentStone.value * 2024
			currentStone.blinkCount += 1
		}
	}

	return stoneCount, nil
}

func SolveDay11Part1() (int, error) {
	inputPath := os.Getenv("AOC_INPUT_PATH")

	input, err := os.ReadFile(inputPath)

	if err != nil {
		return 0, err
	}

	soltion, err := blinkNTimesAndCountNumberOfStones(string(input), 25)

	if err != nil {
		return 0, err
	}

	return soltion, nil
}

func SolveDay11Part2() (int, error) {
	inputPath := os.Getenv("AOC_INPUT_PATH")

	input, err := os.ReadFile(inputPath)

	if err != nil {
		return 0, err
	}

	soltion, err := blinkNTimesAndCountNumberOfStones(string(input), 75)

	if err != nil {
		return 0, err
	}

	return soltion, nil
}
