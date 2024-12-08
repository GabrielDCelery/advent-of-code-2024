package day_08

import (
	"fmt"
	"os"
	"strings"
)

type Vector struct {
	y int
	x int
}

func addVectors(v1 Vector, v2 Vector) Vector {
	return Vector{
		y: v1.y + v2.y,
		x: v1.x + v2.x,
	}
}

func subtractVectors(v1 Vector, v2 Vector) Vector {
	return Vector{
		y: v1.y - v2.y,
		x: v1.x - v2.x,
	}
}

func multiplyVector(v Vector, m int) Vector {
	return Vector{
		y: v.y * m,
		x: v.x * m,
	}
}

type AntennaLayout struct {
	areaHeight      int
	areaWidth       int
	positionsByType map[string][]Vector
}

func NewAntennaLayout(input string) *AntennaLayout {
	positionsByType := map[string][]Vector{}
	lines := strings.Split(strings.TrimSpace(input), "\n")
	for y, line := range lines {
		chars := strings.Split(line, "")
		for x, char := range chars {
			if char == "." {
				continue
			}
			_, ok := positionsByType[char]
			if !ok {
				positionsByType[char] = []Vector{}
			}
			positionsByType[char] = append(positionsByType[char], Vector{y: y, x: x})
		}
	}
	areaHeight := len(lines)
	areaWidth := len(lines[0])
	return &AntennaLayout{
		areaHeight:      areaHeight,
		areaWidth:       areaWidth,
		positionsByType: positionsByType,
	}
}

func (a *AntennaLayout) isVectorWithinArea(v Vector) bool {
	return v.y >= 0 && v.y < a.areaHeight && v.x >= 0 && v.x < a.areaWidth
}

func (a *AntennaLayout) createUniqueAntennaPairs() [][2]Vector {
	antennaPairsCoveredMap := map[string]bool{}
	antennaPairs := [][2]Vector{}

	for antennaType := range a.positionsByType {
		if len(a.positionsByType[antennaType]) < 2 {
			continue
		}
		for i, antennaA := range a.positionsByType[antennaType] {
			for j, antennaB := range a.positionsByType[antennaType] {
				if i == j {
					continue
				}
				key1 := fmt.Sprintf("%d_%d_%d_%d", antennaA.y, antennaA.x, antennaB.y, antennaB.x)
				_, ok := antennaPairsCoveredMap[key1]
				if ok {
					continue
				}
				antennaPair := [2]Vector{antennaA, antennaB}
				antennaPairs = append(antennaPairs, antennaPair)
				key2 := fmt.Sprintf("%d_%d_%d_%d", antennaB.y, antennaB.x, antennaA.y, antennaA.x)
				antennaPairsCoveredMap[key1] = true
				antennaPairsCoveredMap[key2] = true
			}
		}
	}

	return antennaPairs
}

func getValidAntinodesForAntennaPair(antennaPair [2]Vector, layout AntennaLayout) []Vector {
	antinodes := []Vector{}
	antinode1 := addVectors(antennaPair[0], subtractVectors(antennaPair[0], antennaPair[1]))
	if layout.isVectorWithinArea(antinode1) {
		antinodes = append(antinodes, antinode1)
	}
	antinode2 := addVectors(antennaPair[1], subtractVectors(antennaPair[1], antennaPair[0]))
	if layout.isVectorWithinArea(antinode2) {
		antinodes = append(antinodes, antinode2)
	}
	return antinodes
}

func getValidAnitnodesForAntennaPairFactoringInHarmonics(antennaPair [2]Vector, layout AntennaLayout) []Vector {
	antinodes := []Vector{}
	distance1 := 0
	for {
		antinode := addVectors(antennaPair[0], multiplyVector(subtractVectors(antennaPair[0], antennaPair[1]), distance1))
		if !layout.isVectorWithinArea(antinode) {
			break
		}
		antinodes = append(antinodes, antinode)
		distance1 += 1
	}
	distance2 := 0
	for {
		antinode := addVectors(antennaPair[1], multiplyVector(subtractVectors(antennaPair[1], antennaPair[0]), distance2))
		if !layout.isVectorWithinArea(antinode) {
			break
		}
		antinodes = append(antinodes, antinode)
		distance2 += 1
	}
	return antinodes
}

func calculateNumOfUniqueAntinodes(input string) (int, error) {
	antennaLayout := NewAntennaLayout(input)
	antennaPairs := antennaLayout.createUniqueAntennaPairs()

	uniqueAntinodesMap := map[string]bool{}

	for _, antennaPair := range antennaPairs {
		antinodes := getValidAntinodesForAntennaPair(antennaPair, *antennaLayout)
		for _, antinode := range antinodes {
			key := fmt.Sprintf("%d_%d", antinode.y, antinode.x)
			uniqueAntinodesMap[key] = true
		}
	}

	return len(uniqueAntinodesMap), nil
}

func calculateNumOfUniqueAntinodesFactoringInHarmonics(input string) (int, error) {
	antennaLayout := NewAntennaLayout(input)
	antennaPairs := antennaLayout.createUniqueAntennaPairs()

	uniqueAntinodesMap := map[string]bool{}

	for _, antennaPair := range antennaPairs {
		antinodes := getValidAnitnodesForAntennaPairFactoringInHarmonics(antennaPair, *antennaLayout)
		for _, antinode := range antinodes {
			key := fmt.Sprintf("%d_%d", antinode.y, antinode.x)
			uniqueAntinodesMap[key] = true
		}
	}

	return len(uniqueAntinodesMap), nil
}

func SolveDay8Part1() (int, error) {
	inputPath := os.Getenv("AOC_INPUT_PATH")

	file, err := os.ReadFile(inputPath)

	if err != nil {
		return 0, err
	}

	solution, err := calculateNumOfUniqueAntinodes(string(file))

	if err != nil {
		return 0, err
	}

	return solution, nil
}

func SolveDay8Part2() (int, error) {
	inputPath := os.Getenv("AOC_INPUT_PATH")

	file, err := os.ReadFile(inputPath)

	if err != nil {
		return 0, err
	}

	solution, err := calculateNumOfUniqueAntinodesFactoringInHarmonics(string(file))

	if err != nil {
		return 0, err
	}

	return solution, nil
}
