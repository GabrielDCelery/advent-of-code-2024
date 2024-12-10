package day_10

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_calculateAndSumTrailHeadScores(t *testing.T) {
	t.Run("Successfully calculates and sum the trailhead scores", func(t *testing.T) {
		// Given
		input := `
89010123
78121874
87430965
96549874
45678903
32019012
01329801
10456732`

		// When
		result, err := calculateAndSumTrailHeadScores(input)

		// Then
		assert.NoError(t, err)
		assert.Equal(t, 36, result)
	})
}

func Test_calculateAndSumUniqueTrails(t *testing.T) {
	t.Run("Successfully calculates and sum the trailhead scores", func(t *testing.T) {
		// Given
		input := `
89010123
78121874
87430965
96549874
45678903
32019012
01329801
10456732`

		// When
		result, err := calculateAndSumUniqueTrails(input)

		// Then
		assert.NoError(t, err)
		assert.Equal(t, 81, result)
	})
}
