package day_13

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_calculateTheMinimumNumberOfTokensToNeededForMachine(t *testing.T) {
	t.Run("Correctly calculates that it takes 280 tokens to win prize on the first machine among the examples", func(t *testing.T) {
		// Given
		machine := Machine{
			buttons: []Button{
				{direction: Vector{x: 94, y: 34}, cost: 3},
				{direction: Vector{x: 22, y: 67}, cost: 1},
			},
			prize: Vector{x: 8400, y: 5400},
		}
		maxNumOfButtonPresses := 100

		// When
		result := calculateTheMinimumNumberOfTokensToNeededForMachine(machine, maxNumOfButtonPresses)

		// Then
		assert.Equal(t, 280, result)
	})

	t.Run("Correctly calculates that it takes 200 tokens to win prize on the third machine among the examples", func(t *testing.T) {
		// Given
		machine := Machine{
			buttons: []Button{
				{direction: Vector{x: 17, y: 86}, cost: 3},
				{direction: Vector{x: 84, y: 37}, cost: 1},
			},
			prize: Vector{x: 7870, y: 6450},
		}
		maxNumOfButtonPresses := 100

		// When
		result := calculateTheMinimumNumberOfTokensToNeededForMachine(machine, maxNumOfButtonPresses)

		// Then
		assert.Equal(t, 200, result)
	})

	t.Run("Correctly calculates that the prize can not be won on the second machine among the examples", func(t *testing.T) {
		// Given
		machine := Machine{
			buttons: []Button{
				{direction: Vector{x: 26, y: 66}, cost: 3},
				{direction: Vector{x: 67, y: 21}, cost: 1},
			},
			prize: Vector{x: 12748, y: 12176},
		}
		maxNumOfButtonPresses := 100

		// When
		result := calculateTheMinimumNumberOfTokensToNeededForMachine(machine, maxNumOfButtonPresses)

		// Then
		assert.Equal(t, -1, result)
	})

	t.Run("Correctly calculates that the prize can not be won on the fourth machine among the examples", func(t *testing.T) {
		// Given
		machine := Machine{
			buttons: []Button{
				{direction: Vector{x: 69, y: 23}, cost: 3},
				{direction: Vector{x: 27, y: 71}, cost: 1},
			},
			prize: Vector{x: 18641, y: 10279},
		}
		maxNumOfButtonPresses := 100

		// When
		result := calculateTheMinimumNumberOfTokensToNeededForMachine(machine, maxNumOfButtonPresses)

		// Then
		assert.Equal(t, -1, result)
	})
}

func Test_calculateMinNumberOfTokensNeededToWinMostNumOfPrizes(t *testing.T) {
	t.Run("Correctly sums of the total number of tokens required to win the most amount of prizes in the example", func(t *testing.T) {
		// Given
		input := `
Button A: X+94, Y+34
Button B: X+22, Y+67
Prize: X=8400, Y=5400

Button A: X+26, Y+66
Button B: X+67, Y+21
Prize: X=12748, Y=12176

Button A: X+17, Y+86
Button B: X+84, Y+37
Prize: X=7870, Y=6450

Button A: X+69, Y+23
Button B: X+27, Y+71
Prize: X=18641, Y=10279`

		maxNumOfButtonPresses := 100

		// When
		result, error := calculateMinNumberOfTokensNeededToWinMostNumOfPrizes(input, maxNumOfButtonPresses)

		// Then
		assert.NoError(t, error)
		assert.Equal(t, 480, result)
	})

}

// func Test_calculateMinNumberOfTokensNeededToWinMostNumOfAdjustedPrizes(t *testing.T) {
// 	t.Run("Correctly calculates that it takes X tokens to solve the second example when the values are adjusted", func(t *testing.T) {
// 		// Given
// 		machine := Machine{
// 			buttons: []Button{
// 				{direction: Vector{x: 26, y: 66}, cost: 3},
// 				{direction: Vector{x: 67, y: 21}, cost: 1},
// 			},
// 			prize: Vector{x: 10000000012748, y: 10000000012176},
// 		}
// 		maxNumOfButtonPresses := -1
//
// 		// When
// 		result := calculateTheMinimumNumberOfTokensToNeededForMachine(machine, maxNumOfButtonPresses)
//
// 		// Then
// 		assert.Equal(t, -1, result)
// 	})
// }
