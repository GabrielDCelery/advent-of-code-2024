package day_13

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Vector struct {
	y int
	x int
}

type Button struct {
	direction Vector
	cost      int
}

type Machine struct {
	buttons []Button
	prize   Vector
}

func traverseOptionsForMachine(
	machine Machine,
	visitedPositions *map[string]bool,
	buttonPressCounts []int,
	maxNumOfButtonPresses int,
	cheapestCost *int,
) {
	visited := strings.Trim(strings.Replace(fmt.Sprint(buttonPressCounts), " ", "_", -1), "[]")
	if (*visitedPositions)[visited] {
		return
	}
	(*visitedPositions)[visited] = true

	currentPosition := Vector{x: 0, y: 0}
	totalCost := 0

	for i, buttonPressCount := range buttonPressCounts {
		currentPosition.x += machine.buttons[i].direction.x * buttonPressCount
		currentPosition.y += machine.buttons[i].direction.y * buttonPressCount
		totalCost += machine.buttons[i].cost * buttonPressCount
	}

	if currentPosition.x == machine.prize.x && currentPosition.y == machine.prize.y {
		if *cheapestCost == -1 || totalCost < *cheapestCost {
			*cheapestCost = totalCost
		}
		return

	}

	if currentPosition.x > machine.prize.x || currentPosition.y > machine.prize.y {
		return
	}

	for i, buttonPressCount := range buttonPressCounts {
		if maxNumOfButtonPresses != -1 && buttonPressCount >= maxNumOfButtonPresses {
			continue
		}
		newButtonPressCounts := append([]int{}, buttonPressCounts...)
		newButtonPressCounts[i] += 1
		traverseOptionsForMachine(machine, visitedPositions, newButtonPressCounts, maxNumOfButtonPresses, cheapestCost)
	}
}

func calculateTheMinimumNumberOfTokensToNeededForMachine(machine Machine, maxNumOfButtonPresses int) int {
	cheapestCost := -1
	visitedPositions := map[string]bool{}
	buttonPressCounts := []int{}
	for range machine.buttons {
		buttonPressCounts = append(buttonPressCounts, 0)
	}
	traverseOptionsForMachine(
		machine,
		&visitedPositions,
		buttonPressCounts,
		maxNumOfButtonPresses,
		&cheapestCost,
	)
	return cheapestCost
}

func getVectorForButtonFromInput(input string) (Vector, error) {
	components := strings.Split(input, " ")
	xAsStr := strings.ReplaceAll(strings.ReplaceAll(components[2], "X+", ""), ",", "")
	x, err := strconv.Atoi(xAsStr)
	if err != nil {
		return Vector{}, err
	}
	yAsStr := strings.ReplaceAll(components[3], "Y+", "")
	y, err := strconv.Atoi(yAsStr)
	if err != nil {
		return Vector{}, err
	}
	return Vector{y: y, x: x}, nil
}

func getVectorFromPrizeInput(input string) (Vector, error) {
	components := strings.Split(input, " ")
	xAsStr := strings.ReplaceAll(strings.ReplaceAll(components[1], "X=", ""), ",", "")
	x, err := strconv.Atoi(xAsStr)
	if err != nil {
		return Vector{}, err
	}
	yAsStr := strings.ReplaceAll(components[2], "Y=", "")
	y, err := strconv.Atoi(yAsStr)
	if err != nil {
		return Vector{}, err
	}
	return Vector{y: y, x: x}, nil
}

func splitAtEmptyLines(input string) []string {
	lines := strings.Split(input, "\n")
	blocks := []string{}
	block := []string{}

	for _, line := range lines {
		if line == "" {
			continue
		}
		block = append(block, line)
		if len(block) == 3 {
			blocks = append(blocks, strings.Join(block, "\n"))
			block = []string{}
		}
	}
	return blocks
}

func transformInputStringToMachines(input string) ([]Machine, error) {
	machines := []Machine{}
	machineInputs := splitAtEmptyLines(strings.TrimSpace(input))
	for _, machineInput := range machineInputs {
		m := strings.Split(strings.TrimSpace(machineInput), "\n")
		buttonAStr, buttonBStr, prizeStr := m[0], m[1], m[2]
		buttonA, err := getVectorForButtonFromInput(buttonAStr)
		if err != nil {
			return []Machine{}, err
		}
		buttonB, err := getVectorForButtonFromInput(buttonBStr)
		if err != nil {
			return []Machine{}, err
		}
		prize, err := getVectorFromPrizeInput(prizeStr)
		if err != nil {
			return []Machine{}, err
		}
		machine := Machine{
			buttons: []Button{{direction: buttonA, cost: 3}, {direction: buttonB, cost: 1}},
			prize:   prize,
		}
		machines = append(machines, machine)
	}
	return machines, nil
}

func calculateMinNumberOfTokensNeededToWinMostNumOfPrizes(input string, maxNumOfButtonPresses int) (int, error) {
	machines, err := transformInputStringToMachines(input)
	if err != nil {
		return 0, nil
	}
	totalTokens := 0
	for _, machine := range machines {
		tokens := calculateTheMinimumNumberOfTokensToNeededForMachine(machine, maxNumOfButtonPresses)
		if tokens != -1 {
			totalTokens += tokens
		}
	}
	return totalTokens, nil
}

func calculateMinNumberOfTokensNeededToWinMostNumOfAdjustedPrizes(input string, maxNumOfButtonPresses int, adjustor int) (int, error) {
	machines, err := transformInputStringToMachines(input)
	if err != nil {
		return 0, nil
	}
	totalTokens := 0
	for _, machine := range machines {
		machine.prize.x += adjustor
		machine.prize.y += adjustor
		tokens := calculateTheMinimumNumberOfTokensToNeededForMachine(machine, maxNumOfButtonPresses)
		if tokens != -1 {
			totalTokens += tokens
		}
	}
	return totalTokens, nil
}

func SolveDay13Part1() (int, error) {
	inputPath := os.Getenv("AOC_INPUT_PATH")

	input, err := os.ReadFile(inputPath)

	if err != nil {
		return 0, err
	}

	soltion, err := calculateMinNumberOfTokensNeededToWinMostNumOfPrizes(string(input), 100)

	if err != nil {
		return 0, err
	}

	return soltion, nil
}
