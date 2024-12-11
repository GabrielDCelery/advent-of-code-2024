package day_11

import (
	"math"
	"os"
	"strconv"
	"strings"
	"sync"
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

func calculateNumOfStonesForNumber(wg *sync.WaitGroup, number int, targetBlinkCount int) int {
	defer wg.Done()

	stoneCount := 0

	currentStone := NewStone(number, 0, nil)

	for {
		if currentStone.blinkCount == targetBlinkCount {
			stoneCount += 1
			if currentStone.parent == nil {
				break
			} else {
				childStone := currentStone
				currentStone = currentStone.parent
				currentStone.child = nil
				childStone.parent = nil
				continue
			}
		}

		if currentStone.value == 0 {
			currentStone.value = 1
			currentStone.blinkCount += 1
			continue
		}

		numOfDecimalDigits := int(math.Log10(float64(currentStone.value))) + 1
		if numOfDecimalDigits%2 == 0 {
			divider := int(math.Pow10(numOfDecimalDigits / 2))
			left := currentStone.value % divider
			right := int(math.Floor(float64(currentStone.value / divider)))
			currentStone.value = left
			currentStone.blinkCount += 1
			currentStone = NewStone(right, currentStone.blinkCount, currentStone)
			continue
		}

		currentStone.value = currentStone.value * 2024
		currentStone.blinkCount += 1
	}

	return stoneCount
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

	var wg sync.WaitGroup

	stoneCountChan := make(chan int)

	for _, number := range numbers {
		wg.Add(1)
		go func(wgPtr *sync.WaitGroup, stoneCountChan chan int, number int) {
			stoneCountChan <- calculateNumOfStonesForNumber(wgPtr, number, targetBlinkCount)
		}(&wg, stoneCountChan, number)
	}

	go func(wgPtr *sync.WaitGroup, stoneCountChan chan int) {
		wgPtr.Wait()
		close(stoneCountChan)
	}(&wg, stoneCountChan)

	totalStoneCount := 0

	for stoneCount := range stoneCountChan {
		totalStoneCount += stoneCount
	}

	return totalStoneCount, nil
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
