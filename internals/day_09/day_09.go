package day_09

import (
	"math"
	"os"
	"strconv"
	"strings"
)

func transformStringToDiskSpace(input string) ([]int, error) {
	diskSpace := []int{}
	chars := strings.Split(input, "")
	for i, char := range chars {
		numOfBlocks, err := strconv.Atoi(char)
		if err != nil {
			return nil, err
		}
		isFile := math.Mod(float64(i), float64(2)) == 0
		for range numOfBlocks {
			if isFile {
				value := int(i / 2)
				diskSpace = append(diskSpace, value)
			} else {

				diskSpace = append(diskSpace, -1)
			}

		}
	}
	return diskSpace, nil
}

func deFragmentDiskSpace(diskSpace []int) {
	pointerA := 0
	pointerB := len(diskSpace) - 1
	for {
		if pointerA == pointerB {
			return
		}
		if diskSpace[pointerA] >= 0 {
			pointerA += 1
			continue
		}
		if diskSpace[pointerB] == -1 {
			pointerB -= 1
			continue
		}
		diskSpace[pointerA] = diskSpace[pointerB]
		diskSpace[pointerB] = -1
	}
}

func calculateCheckSumForDiskSpace(diskSpace []int) int {
	total := 0
	for i, value := range diskSpace {
		if value == -1 {
			break
		}
		total += i * int(value)
	}
	return total
}

func calculateFileSystemChecksum(input string) (int, error) {
	diskSpace, err := transformStringToDiskSpace(input)
	if err != nil {
		return 0, err
	}
	deFragmentDiskSpace(diskSpace)
	total := calculateCheckSumForDiskSpace(diskSpace)
	return total, nil
}

func SolveDay9Part1() (int, error) {
	inputPath := os.Getenv("AOC_INPUT_PATH")

	file, err := os.ReadFile(inputPath)

	if err != nil {
		return 0, err
	}

	solution, err := calculateFileSystemChecksum(strings.TrimSpace(string(file)))

	if err != nil {
		return 0, err
	}

	return solution, nil
}
