package internals

import (
	"os"
	"strings"
)

func doesTemplateMatchInput(inputMatrix [][]string, templateMatrix [][]string, shiftY int, shiftX int) bool {
	for tY := range len(templateMatrix) {
		for tX := range len(templateMatrix[0]) {
			templateChar := templateMatrix[tY][tX]

			if templateChar == "." {
				continue
			}

			inputChar := inputMatrix[tY+shiftY][tX+shiftX]

			if inputChar != templateChar {
				return false
			}
		}
	}

	return true
}

func calculateHowManyTimesTemplatesAppearInInput(inpuMatrix [][]string, templateMatrices [][][]string) int {
	count := 0

	for _, templateMatrix := range templateMatrices {
		shiftYCount := len(inpuMatrix) - len(templateMatrix)
		shiftXCount := len(inpuMatrix[0]) - len(templateMatrix[0])

		for shiftY := range shiftYCount + 1 {
			for shiftX := range shiftXCount + 1 {
				if doesTemplateMatchInput(inpuMatrix, templateMatrix, shiftY, shiftX) {
					count += 1
				}
			}
		}
	}

	return count
}

func transformInputToMatrix(input string) [][]string {
	charMatrix := [][]string{}

	rows := strings.Split(strings.TrimSpace(input), "\n")

	for _, row := range rows {
		chars := strings.Split(row, "")
		if len(chars) == 0 {
			continue
		}
		charMatrix = append(charMatrix, chars)

	}

	return charMatrix
}

func transformTemplatesToMatrices(templatesFileContent string) [][][]string {
	matrices := [][][]string{}

	templates := strings.Split(templatesFileContent, "$$$$")

	for _, template := range templates {
		rows := strings.Split(strings.TrimSpace(template), "\n")
		colMatrix := [][]string{}
		for _, row := range rows {
			chars := strings.Split(row, "")
			rowMatrix := []string{}
			for _, char := range chars {
				rowMatrix = append(rowMatrix, char)
			}
			colMatrix = append(colMatrix, rowMatrix)
		}
		matrices = append(matrices, colMatrix)
	}

	return matrices
}

func solveHowManyTimesTemplatesAppearInInput(input string, templates string) int {
	inputMatrix := transformInputToMatrix(input)
	templateMatrices := transformTemplatesToMatrices(templates)
	solution := calculateHowManyTimesTemplatesAppearInInput(inputMatrix, templateMatrices)
	return solution
}

func SolveDay4Part1() (int, error) {
	inputPath := os.Getenv("AOC_INPUT_PATH")
	templatesPath := os.Getenv("AOC_PART1_TEMPLATES_PATH")

	input, err := os.ReadFile(inputPath)

	if err != nil {
		return 0, err
	}

	templates, err := os.ReadFile(templatesPath)

	if err != nil {
		return 0, err
	}

	solution := solveHowManyTimesTemplatesAppearInInput(string(input), string(templates))

	return solution, nil
}

func SolveDay4Part2() (int, error) {
	inputPath := os.Getenv("AOC_INPUT_PATH")
	templatesPath := os.Getenv("AOC_PART2_TEMPLATES_PATH")

	input, err := os.ReadFile(inputPath)

	if err != nil {
		return 0, err
	}

	templates, err := os.ReadFile(templatesPath)

	if err != nil {
		return 0, err
	}

	solution := solveHowManyTimesTemplatesAppearInInput(string(input), string(templates))

	return solution, nil
}
