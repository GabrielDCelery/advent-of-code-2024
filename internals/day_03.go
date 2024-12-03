package internals

import (
	"fmt"
	"os"
	"regexp"
	"sort"
	"strconv"
	"strings"
)

type InstructionID int

const (
	MulID InstructionID = iota
	DoID
	DontID
)

const (
	MulRegexp  = `mul\(\d+,\d+\)`
	DoRegexp   = `do\(\)`
	DontRegexp = `don't\(\)`
)

type Instruction struct {
	Id       InstructionID
	Position int
	Value    []byte
}

type InstructionConfig struct {
	Id InstructionID
	Re *regexp.Regexp
}

func createSeriesOfInstructionsFromInput(input []byte, instructionConfigs []InstructionConfig) []Instruction {
	instructions := []Instruction{}

	for _, instructionConfig := range instructionConfigs {
		matches := instructionConfig.Re.FindAllIndex(input, -1)

		for _, match := range matches {
			instruction := Instruction{
				Id:       instructionConfig.Id,
				Position: match[0],
				Value:    input[match[0]:match[1]],
			}

			instructions = append(instructions, instruction)
		}
	}

	sort.Slice(instructions, func(i int, j int) bool {
		return instructions[i].Position < instructions[j].Position
	})

	return instructions
}

func executeInstructions(instructions []Instruction) (int, error) {
	total := 0
	isOn := true

	for _, instruction := range instructions {
		switch instruction.Id {
		case MulID:
			if isOn {
				value, err := executeMulInstruction(instruction)
				if err != nil {
					return 0, err
				}
				total += value
			}
		case DoID:
			isOn = true
		case DontID:
			isOn = false
		default:
			return 0, fmt.Errorf("Unexpexted instruction ID")
		}
	}

	return total, nil
}

func solveInput(input []byte, instructionConfigs []InstructionConfig) (int, error) {
	instructions := createSeriesOfInstructionsFromInput(input, instructionConfigs)

	solution, err := executeInstructions(instructions)

	if err != nil {
		return 0, nil
	}

	return solution, nil
}

func executeMulInstruction(instruction Instruction) (int, error) {
	mulInstruction := string(instruction.Value)
	mulInstruction = strings.ReplaceAll(mulInstruction, "mul(", "")
	mulInstruction = strings.ReplaceAll(mulInstruction, ")", "")
	numbers := strings.Split(mulInstruction, ",")
	num1, err := strconv.Atoi(numbers[0])
	if err != nil {
		return 0, err
	}
	num2, err := strconv.Atoi(numbers[1])
	if err != nil {
		return 0, err
	}
	return num1 * num2, nil
}

func SoveDay3Part1() (int, error) {
	inputPath := os.Getenv("AOC_INPUT_PATH")

	input, err := os.ReadFile(inputPath)

	if err != nil {
		return 0, err
	}

	instructionConfigs := []InstructionConfig{
		{Id: MulID, Re: regexp.MustCompile(MulRegexp)},
	}

	solution, err := solveInput(input, instructionConfigs)

	if err != nil {
		return 0, err
	}

	return solution, nil
}

func SolveDay3Part2() (int, error) {
	inputPath := os.Getenv("AOC_INPUT_PATH")

	input, err := os.ReadFile(inputPath)

	if err != nil {
		return 0, err
	}

	instructionConfigs := []InstructionConfig{
		{Id: MulID, Re: regexp.MustCompile(MulRegexp)},
		{Id: DoID, Re: regexp.MustCompile(DoRegexp)},
		{Id: DontID, Re: regexp.MustCompile(DontRegexp)},
	}

	solution, err := solveInput(input, instructionConfigs)

	if err != nil {
		return 0, err
	}

	return solution, nil
}
