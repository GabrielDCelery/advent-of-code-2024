package day_01

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_calculateDistancesBetween(t *testing.T) {
	t.Run("orders both lists to an ascending order and sums distances between values of the same index", func(t *testing.T) {
		// Given
		firstList := []int{3, 4, 2, 1, 3, 3}
		secondList := []int{4, 3, 5, 3, 9, 3}

		// When
		result, err := calcualteDistancesBetween(firstList, secondList)

		// Then
		assert.NoError(t, err)
		assert.Equal(t, 11, result)
	})
}

func Test_calculateSimilarityScore(t *testing.T) {
	t.Run("calculates how many times numbers from the first list appear in the second", func(t *testing.T) {
		// Given
		firstList := []int{3, 4, 2, 1, 3, 3}
		secondList := []int{4, 3, 5, 3, 9, 3}

		// When
		result, err := calculateSimilarityScore(firstList, secondList)

		// Then
		assert.NoError(t, err)
		assert.Equal(t, 31, result)
	})
}
