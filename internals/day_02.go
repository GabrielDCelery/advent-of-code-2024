package internals

import (
	"bufio"
	"math"
	"os"
	"strconv"
	"strings"
)

func readInputIntoLevelsList(inputPath string) ([][]int, error) {
	file, err := os.Open(inputPath)

	if err != nil {
		return nil, err
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)

	levelsList := [][]int{}

	for scanner.Scan() {
		line := scanner.Text()

		levelsAsStr := strings.Split(line, " ")

		levels := []int{}

		for _, levelAsStr := range levelsAsStr {
			level, err := strconv.Atoi(levelAsStr)

			if err != nil {
				return nil, err
			}

			levels = append(levels, level)
		}

		levelsList = append(levelsList, levels)
	}

	if err := scanner.Err(); err != nil {
		return nil, err
	}

	return levelsList, nil
}

func areLevelsSafe(levels []int) bool {
	if len(levels) == 1 {
		return true
	}

	isIncreasing := false
	isDecreasing := false

	for i := 0; i < len(levels)-1; i++ {
		currLevel := levels[i]
		nextLevel := levels[i+1]

		if nextLevel == currLevel {
			return false
		}

		if nextLevel > currLevel {
			isIncreasing = true
		}

		if nextLevel < currLevel {
			isDecreasing = true
		}

		if isIncreasing == true && isDecreasing == true {
			return false
		}

		diff := int(math.Abs(float64(nextLevel - currLevel)))

		if diff < 1 || diff > 3 {
			return false
		}
	}

	return true
}

func removeLevel(levels []int, index int) []int {
	clone := []int{}
	for i, level := range levels {
		if i == index {
			continue
		}
		clone = append(clone, level)
	}
	return clone
}

func areLevelsSafeUsingDampener(levels []int) bool {
	if areLevelsSafe(levels) {
		return true
	}

	for i := range levels {
		modifiedLevels := removeLevel(levels, i)
		if areLevelsSafe(modifiedLevels) {
			return true
		}
	}

	return false
}

func SolveDay2Part1() (int, error) {
	inputPath := os.Getenv("AOC_INPUT_PATH")

	levelsList, err := readInputIntoLevelsList(inputPath)

	if err != nil {
		return 0, err
	}

	totalCountOfSaveLevels := 0

	for _, levels := range levelsList {
		if areLevelsSafe(levels) {
			totalCountOfSaveLevels += 1
		}
	}

	return totalCountOfSaveLevels, nil
}

func SolveDay2Part2() (int, error) {
	inputPath := os.Getenv("AOC_INPUT_PATH")

	levelsList, err := readInputIntoLevelsList(inputPath)

	if err != nil {
		return 0, err
	}

	totalCountOfSaveLevels := 0

	for _, levels := range levelsList {
		if areLevelsSafeUsingDampener(levels) {
			totalCountOfSaveLevels += 1
		}
	}

	return totalCountOfSaveLevels, nil
}
