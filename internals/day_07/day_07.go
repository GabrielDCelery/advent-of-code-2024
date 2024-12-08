package day_07

import (
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

func buildDFSForCalibration(components []int, operation Operation) *Node {
	rootNode := &Node{
		value:     components[0],
		total:     0,
		operation: operation,
		children:  []*Node{},
	}

	if len(components) > 1 {
		for _, operation := range []Operation{Multiply, Add} {
			rootNode.children = append(rootNode.children, buildDFSForCalibration(components[1:], operation))
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

func doesCalibrationProducesTestResult(calibration Calibration) bool {
	rootNode := buildDFSForCalibration(calibration.components, Null)
	possibleTotals := []int{}
	computePossibleTotalsForDFS(rootNode, 0, &possibleTotals)
	result := slices.Contains(possibleTotals, calibration.target)
	return result
}

func sumCalibrationsThatPassTest(input string) (int, error) {
	lines := strings.Split(strings.TrimSpace(input), "\n")
	total := 0
	for _, line := range lines {
		calibration, err := transformInputLineToCalibration(line)
		if err != nil {
			return 0, err
		}
		if doesCalibrationProducesTestResult(calibration) {
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

	solution, err := sumCalibrationsThatPassTest(string(file))

	if err != nil {
		return 0, err
	}

	return solution, nil
}
