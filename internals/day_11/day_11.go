package day_11

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Rule interface {
	doesRuleApply(int) (bool, error)
	applyRule(int) ([]int, error)
}

type IfZeroReplaceWithOneRule struct{}

func (r *IfZeroReplaceWithOneRule) doesRuleApply(number int) (bool, error) {
	return number == 0, nil
}

func (r *IfZeroReplaceWithOneRule) applyRule(number int) ([]int, error) {
	return []int{1}, nil
}

type IfEvenSplitRule struct{}

func (r *IfEvenSplitRule) doesRuleApply(number int) (bool, error) {
	str := fmt.Sprintf("%d", number)
	return len(str)%2 == 0, nil

}

func (r *IfEvenSplitRule) applyRule(number int) ([]int, error) {
	str := fmt.Sprintf("%d", number)
	mid := len(str) / 2
	left, err := strconv.Atoi(str[mid:])
	if err != nil {
		return nil, err
	}
	right, err := strconv.Atoi(str[:mid])
	if err != nil {
		return nil, err
	}
	return []int{left, right}, nil
}

type MultiplyBy2024Rule struct{}

func (r *MultiplyBy2024Rule) doesRuleApply(number int) (bool, error) {
	return true, nil
}

func (r *MultiplyBy2024Rule) applyRule(number int) ([]int, error) {
	return []int{2024 * number}, nil
}

func applyRulesToNumber(number int, rules []Rule) ([]int, error) {
	for _, rule := range rules {
		doesRuleApply, err := rule.doesRuleApply(number)
		if err != nil {
			return nil, err
		}
		if doesRuleApply {
			result, err := rule.applyRule(number)
			if err != nil {
				return nil, err
			}
			return result, nil
		}
	}
	return nil, fmt.Errorf("Could not find a rule that would apply to %d", number)
}

func blinkNTimesAndCountNumberOfStones(input string, blinkCount int) (int, error) {
	numbersAsStr := strings.Split(strings.TrimSpace(input), " ")
	numbers := []int{}
	for _, numberAsStr := range numbersAsStr {
		number, err := strconv.Atoi(numberAsStr)
		if err != nil {
			return 0, err
		}
		numbers = append(numbers, number)
	}
	rules := []Rule{&IfZeroReplaceWithOneRule{}, &IfEvenSplitRule{}, &MultiplyBy2024Rule{}}
	iterationCount := 0
	for iterationCount < blinkCount {
		updatedNumbers := []int{}
		for _, number := range numbers {
			numbersAfter, err := applyRulesToNumber(number, rules)
			if err != nil {
				return 0, err
			}
			for _, numberAfter := range numbersAfter {
				updatedNumbers = append(updatedNumbers, numberAfter)
			}
		}
		numbers = updatedNumbers
		iterationCount += 1
	}
	stoneCount := len(numbers)
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
