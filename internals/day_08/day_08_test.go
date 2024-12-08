package day_08

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_calculateNumOfUniqueAntinodes(t *testing.T) {
	t.Run("Successfully calculates the number of unique antinodes", func(t *testing.T) {
		// Given
		input := `............
........0...
.....0......
.......0....
....0.......
......A.....
............
............
........A...
.........A..
............
............`

		// When
		result, err := calculateNumOfUniqueAntinodes(input)

		// Then
		assert.NoError(t, err)
		assert.Equal(t, 14, result)
	})
}
