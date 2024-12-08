package day_07

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_countNumberOfWaysGuardianCanBeLockedIntoInfiniteLoop(t *testing.T) {
	t.Run("Successfully transform the string '190: 10 19' to a Calibration", func(t *testing.T) {
		// Given
		line := "190: 10 19"

		// When
		result, err := transformInputLineToCalibration(line)

		// Then
		expected := Calibration{
			target:     190,
			components: []int{10, 19},
		}
		assert.NoError(t, err)
		assert.Equal(t, expected, result)
	})

	t.Run("Successfully transform the string '3267: 81 40 27' to a Calibration", func(t *testing.T) {
		// Given
		line := "3267: 81 40 27"

		// When
		result, err := transformInputLineToCalibration(line)

		// Then
		expected := Calibration{
			target:     3267,
			components: []int{81, 40, 27},
		}
		assert.NoError(t, err)
		assert.Equal(t, expected, result)
	})
}

func Test_doesCalibrationProducesTestResult(t *testing.T) {
	t.Run("Successfully evaluates calibration '190: 10 19' to a true test result", func(t *testing.T) {
		// Given
		calibration := Calibration{
			target:     190,
			components: []int{10, 19},
		}

		// When
		result := doesCalibrationProducesTestResult(calibration)

		// Then
		assert.Equal(t, true, result)
	})

	t.Run("Successfully evaluates calibration '3267: 81 40 27' to a true test result", func(t *testing.T) {
		// Given
		calibration := Calibration{
			target:     3267,
			components: []int{81, 40, 27},
		}

		// When
		result := doesCalibrationProducesTestResult(calibration)

		// Then
		assert.Equal(t, true, result)
	})
}

func Test_sumCalibrationsThatPassTest(t *testing.T) {
	t.Run("Correctly sums calibrations that pass the test", func(t *testing.T) {
		// Given
		input := `190: 10 19
3267: 81 40 27
83: 17 5
156: 15 6
7290: 6 8 6 15
161011: 16 10 13
192: 17 8 14
21037: 9 7 18 13
292: 11 6 16 20`

		// When
		result, err := sumCalibrationsThatPassTest(input)

		// Then
		assert.NoError(t, err)
		assert.Equal(t, 3749, result)
	})
}
