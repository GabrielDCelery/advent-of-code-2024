package internals

import (
	"bufio"
	"math"
	"os"
	"sort"
	"strconv"
	"strings"
)

func isIntInSortedInts(integer int, sortedIntegers []int) bool {
	index := sort.SearchInts(sortedIntegers, integer)
	if index == len(sortedIntegers) {
		return false
	}
	return sortedIntegers[index] == integer
}

func doesUpdateSatisfyRules(update []int, rules map[int][]int) bool {
	for i, pageToCompare := range update {
		pagesToCompareWith := update[i+1:]
		for _, pageToCompareWith := range pagesToCompareWith {
			rule, ok := rules[pageToCompareWith]
			if ok && isIntInSortedInts(pageToCompare, rule) {
				return false
			}
		}
	}
	return true
}

func sumTheMiddleNumbersOfCorrectUpdates(updates [][]int, rules map[int][]int) int {
	total := 0
	for _, update := range updates {
		if doesUpdateSatisfyRules(update, rules) {
			middleNumIndex := int(math.Floor(float64(len(update) / 2)))
			total += update[middleNumIndex]
		}
	}
	return total
}

func sortUpdateToBeInCorrectOrder(update []int, rules map[int][]int) {
	for !doesUpdateSatisfyRules(update, rules) {
		for i := range len(update) - 1 {
			rule, ok := rules[update[i+1]]
			if ok && isIntInSortedInts(update[i], rule) {
				n1 := update[i]
				n2 := update[i+1]
				update[i] = n2
				update[i+1] = n1
			}
		}
	}

}

func fixTheOrderOfIncorrectUpdatesThenSumTheMiddleNumbers(updates [][]int, rules map[int][]int) int {
	total := 0
	for _, update := range updates {
		if doesUpdateSatisfyRules(update, rules) {
			continue
		}
		sortUpdateToBeInCorrectOrder(update, rules)
		middleNumIndex := int(math.Floor(float64(len(update) / 2)))
		total += update[middleNumIndex]
	}
	return total
}

func readInputIntoRulesAndUpdates(inputPath string) (map[int][]int, [][]int, error) {
	file, err := os.Open(inputPath)

	if err != nil {
		return nil, nil, err
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)

	readingRules := true

	rules := map[int][]int{}
	updates := [][]int{}

	for scanner.Scan() {
		line := scanner.Text()
		if len(line) == 0 {
			readingRules = false
			continue
		}
		if readingRules == true {
			nums := strings.Split(line, "|")
			pageToUpdate, err := strconv.Atoi(nums[0])
			pageThatMustFollow, err := strconv.Atoi(nums[1])
			if err != nil {
				return nil, nil, err
			}
			_, ok := rules[pageToUpdate]
			if !ok {
				rules[pageToUpdate] = []int{}
			}
			rules[pageToUpdate] = append(rules[pageToUpdate], pageThatMustFollow)
		} else {
			numsAsStr := strings.Split(line, ",")
			update := []int{}
			for _, numAsStr := range numsAsStr {
				num, err := strconv.Atoi(numAsStr)
				if err != nil {
					return nil, nil, err
				}
				update = append(update, num)
			}
			updates = append(updates, update)
		}
	}

	for k := range rules {
		sort.Ints(rules[k])
	}

	if err := scanner.Err(); err != nil {
		return nil, nil, err
	}

	return rules, updates, nil

}

func SolveDay5Part1() (int, error) {
	inputPath := os.Getenv("AOC_INPUT_PATH")

	rules, updates, err := readInputIntoRulesAndUpdates(inputPath)

	if err != nil {
		return 0, err
	}

	solution := sumTheMiddleNumbersOfCorrectUpdates(updates, rules)

	return solution, nil
}

func SolveDay5Part2() (int, error) {
	inputPath := os.Getenv("AOC_INPUT_PATH")

	rules, updates, err := readInputIntoRulesAndUpdates(inputPath)

	if err != nil {
		return 0, err
	}

	solution := fixTheOrderOfIncorrectUpdatesThenSumTheMiddleNumbers(updates, rules)

	return solution, nil
}
