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

func createUniqueAntennaPairs(antennaMap map[string][]Vector) [][2]Vector {
	antennaPairsCoveredMap := map[string]bool{}
	antennaPairs := [][2]Vector{}

	for key := range antennaMap {
		for a, antennaA := range antennaMap[key] {
			for b, antennaB := range antennaMap[key] {
				if a == b {
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

func getAntinoeVectorsFromAntennaPair(antennaPair [2]Vector) [2]Vector {
	vec1 := antennaPair[0]
	vec2 := antennaPair[1]

	antinode1 := addVectors(vec1, subtractVectors(vec1, vec2))
	antinode2 := addVectors(vec2, subtractVectors(vec2, vec1))

	return [2]Vector{antinode1, antinode2}
}

func calculateNumOfUniqueAntinodes(input string) (int, error) {
	antennaMap := map[string][]Vector{}
	lines := strings.Split(strings.TrimSpace(input), "\n")
	for y, line := range lines {
		chars := strings.Split(line, "")
		for x, char := range chars {
			if char == "." {
				continue
			}
			_, ok := antennaMap[char]
			if !ok {
				antennaMap[char] = []Vector{}
			}
			antennaMap[char] = append(antennaMap[char], Vector{y: y, x: x})
		}
	}

	antennaPairs := createUniqueAntennaPairs(antennaMap)

	areaHeight := len(lines)
	areaWidth := len(lines[0])

	uniqueAntinodesMap := map[string]bool{}

	for _, antennaPair := range antennaPairs {
		antinodes := getAntinoeVectorsFromAntennaPair(antennaPair)
		for _, antinode := range antinodes {
			isAntinodeWithinArea := antinode.y >= 0 &&
				antinode.y < areaHeight &&
				antinode.x >= 0 &&
				antinode.x < areaWidth
			if !isAntinodeWithinArea {
				continue
			}
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
