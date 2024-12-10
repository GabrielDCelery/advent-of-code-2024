package day_10

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

func findValidPaths(area [][]int, path []Vector, validPaths *[][]Vector) {
	curr := path[len(path)-1]

	if area[curr.y][curr.x] == 9 {
		*validPaths = append(*validPaths, path)
		return
	}

	directions := []Vector{{y: -1, x: 0}, {y: 0, x: 1}, {y: 1, x: 0}, {y: 0, x: -1}}

	for _, dir := range directions {
		next := Vector{
			y: curr.y + dir.y,
			x: curr.x + dir.x,
		}
		isNextWithinArea := next.y >= 0 && next.x >= 0 && next.y < len(area) && next.x < len(area[0])
		if !isNextWithinArea {
			continue
		}
		isNextPositionHigher := (area[next.y][next.x] - area[curr.y][curr.x]) == 1
		if !isNextPositionHigher {
			continue
		}
		clonedPath := []Vector{}
		for _, cell := range path {
			clonedPath = append(clonedPath, Vector{y: cell.y, x: cell.x})
		}
		clonedPath = append(clonedPath, Vector{y: next.y, x: next.x})
		findValidPaths(area, clonedPath, validPaths)
	}
}

func getTrailHeads(input [][]int) []Vector {
	trailHeads := []Vector{}

	for y := range input {
		for x := range input[y] {
			if input[y][x] == 0 {
				trailHeadVector := Vector{y: y, x: x}
				trailHeads = append(trailHeads, trailHeadVector)
			}
		}
	}

	return trailHeads
}

func createAreaFromInput(input string) ([][]int, error) {
	area := [][]int{}

	lines := strings.Split(strings.TrimSpace(input), "\n")
	for _, line := range lines {
		areaLine := []int{}
		for _, char := range line {
			elevation, err := strconv.Atoi(string(char))

			if err != nil {
				return nil, err
			}

			areaLine = append(areaLine, elevation)
		}
		area = append(area, areaLine)
	}
	return area, nil
}

func findValidUniquePathsForTrailHead(area [][]int, trailHead Vector) [][]Vector {
	validPaths := [][]Vector{}
	path := []Vector{{y: trailHead.y, x: trailHead.x}}
	findValidPaths(area, path, &validPaths)
	return validPaths
}

func calculateAndSumTrailHeadScores(input string) (int, error) {
	area, err := createAreaFromInput(input)

	if err != nil {
		return 0, err
	}

	trailHeads := getTrailHeads(area)

	total := 0

	for _, trailHead := range trailHeads {
		validPaths := findValidUniquePathsForTrailHead(area, trailHead)
		uniqeEndsMap := map[string]bool{}
		for _, validPath := range validPaths {
			end := validPath[len(validPath)-1]
			key := fmt.Sprintf("%d_%d", end.y, end.x)
			uniqeEndsMap[key] = true
		}
		total += len(uniqeEndsMap)
	}

	return total, nil
}

func calculateAndSumUniqueTrails(input string) (int, error) {
	area, err := createAreaFromInput(input)

	if err != nil {
		return 0, err
	}

	trailHeads := getTrailHeads(area)

	total := 0

	for _, trailHead := range trailHeads {
		validPaths := findValidUniquePathsForTrailHead(area, trailHead)
		total += len(validPaths)
	}

	return total, nil
}

func SolveDay10Part1() (int, error) {
	inputPath := os.Getenv("AOC_INPUT_PATH")

	input, err := os.ReadFile(inputPath)

	if err != nil {
		return 0, err
	}

	soltion, err := calculateAndSumTrailHeadScores(string(input))

	if err != nil {
		return 0, err
	}

	return soltion, nil
}

func SolveDay10Part2() (int, error) {
	inputPath := os.Getenv("AOC_INPUT_PATH")

	input, err := os.ReadFile(inputPath)

	if err != nil {
		return 0, err
	}

	soltion, err := calculateAndSumUniqueTrails(string(input))

	if err != nil {
		return 0, err
	}

	return soltion, nil
}
