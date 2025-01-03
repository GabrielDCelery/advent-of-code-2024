package day_07

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
	"sync"
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
	operation Operation
	children  []*Node
}

func buildDFSForCalibration(components []int, operation Operation, validOperations []Operation) *Node {
	rootNode := &Node{
		value:     components[0],
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

func doesDFSHasAPathThatAddsUpToTarget(node *Node, target int, parentNodeTotal int, hasFoundTarget *bool) {
	// if we found the target number there is no need to continue traversing the tree
	if *hasFoundTarget == true {
		return
	}

	currentNodeTotal := parentNodeTotal

	switch node.operation {
	case Null:
		currentNodeTotal = node.value
	case Multiply:
		currentNodeTotal = parentNodeTotal * node.value
	case Add:
		currentNodeTotal = parentNodeTotal + node.value
	case Concatenate:
		total, err := strconv.Atoi(fmt.Sprintf("%d%d", parentNodeTotal, node.value))
		if err != nil {
			log.Fatalln(err)
		}
		currentNodeTotal = total
	}

	// if there are move child nodes continue traversing
	if len(node.children) > 0 {
		for _, child := range node.children {
			doesDFSHasAPathThatAddsUpToTarget(child, target, currentNodeTotal, hasFoundTarget)
		}
		return
	}

	// once we reaached the leaf node check if the totals add up the the total number we are seeking
	if currentNodeTotal == target {
		*hasFoundTarget = true
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
	hasFoundTarget := false
	doesDFSHasAPathThatAddsUpToTarget(rootNode, calibration.target, 0, &hasFoundTarget)
	return hasFoundTarget
}

func sumCalibrationsThatPassTest(input string, validOperations []Operation) (int, error) {
	lines := strings.Split(strings.TrimSpace(input), "\n")
	var wg sync.WaitGroup
	resultsChan := make(chan int)
	errorChan := make(chan error)
	for _, line := range lines {
		wg.Add(1)
		go func(wg *sync.WaitGroup, resultsChan chan int, errorChan chan error) {
			defer wg.Done()
			calibration, err := transformInputLineToCalibration(line)
			if err != nil {
				errorChan <- err
				return
			}
			if doesCalibrationProducesTestResult(calibration, validOperations) {
				resultsChan <- calibration.target
				return
			}
			resultsChan <- 0
		}(&wg, resultsChan, errorChan)
	}
	go func(wg *sync.WaitGroup, resultsChan chan int, errorChan chan error) {
		wg.Wait()
		close(resultsChan)
		close(errorChan)
	}(&wg, resultsChan, errorChan)
	total := 0
	for {
		select {
		case result, ok := <-resultsChan:
			if !ok {
				return total, nil
			}
			total += result
		case err, ok := <-errorChan:
			if !ok {
				return total, nil
			}
			if err != nil {
				return 0, err
			}
		default:
		}
	}
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
