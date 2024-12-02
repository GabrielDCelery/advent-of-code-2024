package internals

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_areLevelsSafe(t *testing.T) {
	t.Run("The levels are either all increasing or all decreasing and any two adjacent levels differ by at least one and at most three", func(t *testing.T) {
		t.Run("Safe because the levels are all decreasing by 1 or 2", func(t *testing.T) {
			// Given
			levels := []int{7, 6, 4, 2, 1}

			// When
			result := areLevelsSafe(levels)

			// Then
			assert.Equal(t, true, result)
		})

		t.Run("Safe because the levels are all increasing by 1, 2, or 3", func(t *testing.T) {
			// Given
			levels := []int{1, 3, 6, 7, 9}

			// When
			result := areLevelsSafe(levels)

			// Then
			assert.Equal(t, true, result)
		})

		t.Run("Unsafe because 2 7 is an increase of 5", func(t *testing.T) {
			// Given
			levels := []int{1, 2, 7, 8, 9}

			// When
			result := areLevelsSafe(levels)

			// Then
			assert.Equal(t, false, result)
		})

		t.Run("Unsafe because 6 2 is a decrease of 4", func(t *testing.T) {
			// Given
			levels := []int{9, 7, 6, 2, 1}

			// When
			result := areLevelsSafe(levels)

			// Then
			assert.Equal(t, false, result)
		})

		t.Run("Unsafe because 1 3 is increasing but 3 2 is decreasing", func(t *testing.T) {
			// Given
			levels := []int{1, 3, 2, 4, 5}

			// When
			result := areLevelsSafe(levels)

			// Then
			assert.Equal(t, false, result)
		})

		t.Run("Unsafe because 4 4 is neither an increase or a decrease", func(t *testing.T) {
			// Given
			levels := []int{8, 6, 4, 4, 1}

			// When
			result := areLevelsSafe(levels)

			// Then
			assert.Equal(t, false, result)
		})
	})
}

func Test_areLevelsSafeWithDampener(t *testing.T) {
	t.Run("If removing a single level from an unsafe report would make it safe, the report instead counts as safe", func(t *testing.T) {
		t.Run("Safe without removing any level", func(t *testing.T) {
			// Given
			levels := []int{7, 6, 4, 2, 1}

			// When
			result := areLevelsSafeUsingDampener(levels)

			// Then
			assert.Equal(t, true, result)
		})

		t.Run("Safe without removing any level", func(t *testing.T) {
			// Given
			levels := []int{1, 3, 6, 7, 9}

			// When
			result := areLevelsSafeUsingDampener(levels)

			// Then
			assert.Equal(t, true, result)
		})

		t.Run("Unsafe regardless of which level is removed", func(t *testing.T) {
			// Given
			levels := []int{1, 2, 7, 8, 9}

			// When
			result := areLevelsSafeUsingDampener(levels)

			// Then
			assert.Equal(t, false, result)
		})

		t.Run("Unsafe regardless of which level is removed", func(t *testing.T) {
			// Given
			levels := []int{9, 7, 6, 2, 1}

			// When
			result := areLevelsSafeUsingDampener(levels)

			// Then
			assert.Equal(t, false, result)
		})

		t.Run("Safe by removing the second level 3", func(t *testing.T) {
			// Given
			levels := []int{1, 3, 2, 4, 5}

			// When
			result := areLevelsSafeUsingDampener(levels)

			// Then
			assert.Equal(t, true, result)
		})

		t.Run("Safe by removing the third level 4", func(t *testing.T) {
			// Given
			levels := []int{8, 6, 4, 4, 1}

			// When
			result := areLevelsSafeUsingDampener(levels)

			// Then
			assert.Equal(t, true, result)
		})
	})
}
