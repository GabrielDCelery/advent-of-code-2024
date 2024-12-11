package day_11

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_blinkNTimesAndCountNumberOfStones(t *testing.T) {
	t.Run("Correctly counts the number of stones after blinking has been applied", func(t *testing.T) {
		// Given
		input := `125 17`
		blinkCount := 6

		// When
		result, err := blinkNTimesAndCountNumberOfStones(input, blinkCount)

		// Then
		assert.NoError(t, err)
		assert.Equal(t, 22, result)
	})

	t.Run("Correctly counts the number of stones after blinking has been applied", func(t *testing.T) {
		// Given
		input := `125 17`
		blinkCount := 25

		// When
		result, err := blinkNTimesAndCountNumberOfStones(input, blinkCount)

		// Then
		assert.NoError(t, err)
		assert.Equal(t, 55312, result)
	})
}
