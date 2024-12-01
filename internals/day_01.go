package internals

import (
	"bufio"
	"math"
	"os"
	"sort"
	"strconv"
	"strings"
)

func readInputIntoSlices(inputPath string) ([]int, []int, error) {
	file, err := os.Open(inputPath)

	if err != nil {
		return nil, nil, err
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)

	firstSetOfLocationIDs := []int{}
	secondSetOfLocationIDs := []int{}

	for scanner.Scan() {
		line := scanner.Text()
		locationIDs := strings.Split(line, "   ")
		locationID1, err := strconv.Atoi(locationIDs[0])
		if err != nil {
			return nil, nil, err
		}
		locationID2, err := strconv.Atoi(locationIDs[1])
		if err != nil {
			return nil, nil, err
		}
		firstSetOfLocationIDs = append(firstSetOfLocationIDs, locationID1)
		secondSetOfLocationIDs = append(secondSetOfLocationIDs, locationID2)
	}

	if err := scanner.Err(); err != nil {
		return nil, nil, err
	}

	return firstSetOfLocationIDs, secondSetOfLocationIDs, nil
}

func calcualteDistancesBetween(firstSetOfLocationIDs []int, secondSetOfLocationIDs []int) (int, error) {
	fLIDs := make([]int, len(firstSetOfLocationIDs))
	copy(fLIDs, firstSetOfLocationIDs)

	sLIDs := make([]int, len(secondSetOfLocationIDs))
	copy(sLIDs, secondSetOfLocationIDs)

	sort.Ints(fLIDs)
	sort.Ints(sLIDs)

	totalDiff := 0

	for i, sLID := range sLIDs {
		fLID := fLIDs[i]
		diff := int(math.Abs(float64(sLID - fLID)))
		totalDiff += diff
	}

	return totalDiff, nil
}

func calculateSimilarityScore(firstSetOfLocationIDs []int, secondSetOfLocationIDs []int) (int, error) {
	occurenceMap := make(map[int]int)
	for _, locationID := range secondSetOfLocationIDs {
		_, exists := occurenceMap[locationID]
		if !exists {
			occurenceMap[locationID] = 0
		}
		occurenceMap[locationID] += 1
	}
	totalSimilarity := 0
	for _, locationID := range firstSetOfLocationIDs {
		numOfOccurences, exists := occurenceMap[locationID]
		if exists {
			totalSimilarity += numOfOccurences * locationID
		}
	}
	return totalSimilarity, nil
}

func SolveDay1Part1() (int, error) {
	inputPath := os.Getenv("AOC_INPUT_PATH")

	firstSetOfLocationIDs, secondSetOfLocationIDs, err := readInputIntoSlices(inputPath)

	if err != nil {
		return 0, err
	}

	solution, err := calcualteDistancesBetween(firstSetOfLocationIDs, secondSetOfLocationIDs)

	if err != nil {
		return 0, err
	}

	return solution, nil
}

func SolveDay1Part2() (int, error) {
	inputPath := os.Getenv("AOC_INPUT_PATH")

	firstSetOfLocationIDs, secondSetOfLocationIDs, err := readInputIntoSlices(inputPath)

	if err != nil {
		return 0, err
	}

	solution, err := calculateSimilarityScore(firstSetOfLocationIDs, secondSetOfLocationIDs)

	if err != nil {
		return 0, err
	}

	return solution, nil

}
