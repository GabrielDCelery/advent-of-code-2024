package internals

import (
	"testing"
)

func Test_calculateDistancesBetween(t *testing.T) {
	t.Run("orders both lists to an ascending order and sums distances between values of the same index", func(t *testing.T) {
		// Given
		firstList := []int{3, 4, 2, 1, 3, 3}
		secondList := []int{4, 3, 5, 3, 9, 3}

		// When
		dist, err := calcualteDistancesBetween(firstList, secondList)

		// Then
		expected := 11

		if err != nil {
			t.Errorf("Unexpexted error has happened: %s", err.Error())
		}

		if dist != expected {
			t.Errorf("got %d, expected %d", dist, expected)
		}
	})
}

func Test_calculateSimilarityScore(t *testing.T) {
	t.Run("calculates how many times numbers from the first list appear in the second", func(t *testing.T) {
		// Given
		firstList := []int{3, 4, 2, 1, 3, 3}
		secondList := []int{4, 3, 5, 3, 9, 3}

		// When
		dist, err := calculateSimilarityScore(firstList, secondList)

		// Then
		expected := 31

		if err != nil {
			t.Errorf("Unexpexted error has happened: %s", err.Error())
		}

		if dist != expected {
			t.Errorf("got %d, expected %d", dist, expected)
		}
	})
}
