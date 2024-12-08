package day_07

import (
	"fmt"
	"log"
	"os"
	"slices"
	"strconv"
	"strings"
)

type Operation int

const (
	Null Operation = iota
	Multiply
	Add
	Concatenate
)

type Calibration struct {
	target     int
	components []int
}

type Node struct {
	value     int
	total     int
	operation Operation
	children  []*Node
}

func buildDFSForCalibration(components []int, operation Operation, validOperations []Operation) *Node {
	rootNode := &Node{
		value:     components[0],
		total:     0,
		operation: operation,
		children:  []*Node{},
	}

	if len(components) > 1 {
		for _, operation := range validOperations {
			rootNode.children = append(rootNode.children, buildDFSForCalibration(components[1:], operation, validOperations))
		}
	}

	return rootNode
}

func computePossibleTotalsForDFS(node *Node, parentNodeTotal int, totals *[]int) {
	switch node.operation {
	case Null:
		node.total = node.value
	case Multiply:
		node.total = parentNodeTotal * node.value
	case Add:
		node.total = parentNodeTotal + node.value
	case Concatenate:
		total, err := strconv.Atoi(fmt.Sprintf("%d%d", parentNodeTotal, node.value))
		if err != nil {
			log.Fatalln(err)
		}
		node.total = total
	}
	if len(node.children) == 0 {
		*totals = append(*totals, node.total)
	}
	for _, child := range node.children {
		computePossibleTotalsForDFS(child, node.total, totals)
	}
}

func transformInputLineToCalibration(line string) (Calibration, error) {
	targetAndComponents := strings.Split(line, ":")
	targetAsString := targetAndComponents[0]
	componentsAsSingleString := strings.TrimSpace(targetAndComponents[1])

	target, err := strconv.Atoi(targetAsString)

	if err != nil {
		return Calibration{}, err
	}

	componentsAsString := strings.Split(componentsAsSingleString, " ")
	components := []int{}
	for _, componentAsString := range componentsAsString {
		component, err := strconv.Atoi(componentAsString)
		if err != nil {
			return Calibration{}, err
		}
		components = append(components, component)
	}
	calibration := Calibration{target: target, components: components}
	return calibration, nil
}

func doesCalibrationProducesTestResult(calibration Calibration, validOperations []Operation) bool {
	rootNode := buildDFSForCalibration(calibration.components, Null, validOperations)
	possibleTotals := []int{}
	computePossibleTotalsForDFS(rootNode, 0, &possibleTotals)
	result := slices.Contains(possibleTotals, calibration.target)
	return result
}

func sumCalibrationsThatPassTest(input string, validOperations []Operation) (int, error) {
	lines := strings.Split(strings.TrimSpace(input), "\n")
	total := 0
	for _, line := range lines {
		calibration, err := transformInputLineToCalibration(line)
		if err != nil {
			return 0, err
		}
		if doesCalibrationProducesTestResult(calibration, validOperations) {
			total += calibration.target
		}
	}
	return total, nil
}

func SolveDay7Part1() (int, error) {
	inputPath := os.Getenv("AOC_INPUT_PATH")

	file, err := os.ReadFile(inputPath)

	if err != nil {
		return 0, err
	}

	validOperations := []Operation{Multiply, Add}

	solution, err := sumCalibrationsThatPassTest(string(file), validOperations)

	if err != nil {
		return 0, err
	}

	return solution, nil
}

func SolveDay7Part2() (int, error) {
	inputPath := os.Getenv("AOC_INPUT_PATH")

	file, err := os.ReadFile(inputPath)

	if err != nil {
		return 0, err
	}

	validOperations := []Operation{Multiply, Add, Concatenate}

	solution, err := sumCalibrationsThatPassTest(string(file), validOperations)

	if err != nil {
		return 0, err
	}

	return solution, nil
}
