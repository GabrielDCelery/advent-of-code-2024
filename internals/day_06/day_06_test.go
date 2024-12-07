package day_06

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_Guardian_executeAction(t *testing.T) {
	t.Run("Guardian executing TURN_RIGHT successfully turns right when facing upwards", func(t *testing.T) {
		// Given
		guardian := NewGuardian(0, Vector{y: 0, x: 0}, Vector{y: -1, x: 0})

		// When
		guardian.ExecuteAction(TURN_RIGHT)

		// Then
		assert.Equal(t, guardian.location, Vector{y: 0, x: 0})
		assert.Equal(t, guardian.facing, Vector{y: 0, x: 1})
	})

	t.Run("Guardian executing TURN_RIGHT successfully turns downwards when facing right", func(t *testing.T) {
		// Given
		guardian := NewGuardian(0, Vector{y: 0, x: 0}, Vector{y: 0, x: 1})

		// When
		guardian.ExecuteAction(TURN_RIGHT)

		// Then
		assert.Equal(t, guardian.location, Vector{y: 0, x: 0})
		assert.Equal(t, guardian.facing, Vector{y: 1, x: 0})
	})

	t.Run("Guardian executing TURN_RIGHT successfully turns left when facing downwards", func(t *testing.T) {
		// Given
		guardian := NewGuardian(0, Vector{y: 0, x: 0}, Vector{y: 1, x: 0})

		// When
		guardian.ExecuteAction(TURN_RIGHT)

		// Then
		assert.Equal(t, guardian.location, Vector{y: 0, x: 0})
		assert.Equal(t, guardian.facing, Vector{y: 0, x: -1})
	})

	t.Run("Guardian executing TURN_RIGHT successfully turns upwards when facing left", func(t *testing.T) {
		// Given
		guardian := NewGuardian(0, Vector{y: 0, x: 0}, Vector{y: 0, x: -1})

		// When
		guardian.ExecuteAction(TURN_RIGHT)

		// Then
		assert.Equal(t, guardian.location, Vector{y: 0, x: 0})
		assert.Equal(t, guardian.facing, Vector{y: -1, x: 0})
	})

	t.Run("Guardian executing MOVE_FORWARD successfully moves forwards when facing downwards", func(t *testing.T) {
		// Given
		guardian := NewGuardian(0, Vector{y: 0, x: 0}, Vector{y: 1, x: 0})

		// When
		guardian.ExecuteAction(MOVE_FORWARD)

		// Then
		assert.Equal(t, guardian.location, Vector{y: 1, x: 0})
		assert.Equal(t, guardian.facing, Vector{y: 1, x: 0})
	})
}

func Test_transformInputToGameState(t *testing.T) {
	t.Run("Transforms input to correct game state", func(t *testing.T) {
		// Given
		input := `....#.....
.........#
..........
..#.......
.......#..
..........
.#..^.....
........#.
#.........
......#...`

		// When
		gameMap, guardians, err := transformInputToGameState(input)

		// Then
		expectedGameMap := GameMap{
			matrix: [][]Cell{
				{0, 0, 0, 0, 1, 0, 0, 0, 0, 0},
				{0, 0, 0, 0, 0, 0, 0, 0, 0, 1},
				{0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
				{0, 0, 1, 0, 0, 0, 0, 0, 0, 0},
				{0, 0, 0, 0, 0, 0, 0, 1, 0, 0},
				{0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
				{0, 1, 0, 0, 0, 0, 0, 0, 0, 0},
				{0, 0, 0, 0, 0, 0, 0, 0, 1, 0},
				{1, 0, 0, 0, 0, 0, 0, 0, 0, 0},
				{0, 0, 0, 0, 0, 0, 1, 0, 0, 0},
			},
		}

		expectedGuardian := Guardian{
			id:       0,
			location: Vector{y: 6, x: 4},
			facing:   Vector{y: -1, x: 0},
		}

		assert.NoError(t, err)
		assert.Equal(t, expectedGameMap, *gameMap)
		assert.Equal(t, expectedGuardian, *guardians[0])
	})
}

func Test_countTheNumberOfDistinctPositionsBeforeGuardianLeaves(t *testing.T) {
	t.Run("Correctly calculates the number of distinct fields the guardian visits before leaving", func(t *testing.T) {
		// Given
		gameMap := `....#.....
.........#
..........
..#.......
.......#..
..........
.#..^.....
........#.
#.........
......#...`

		// When
		result, err := countTheNumberOfDistinctPositionsBeforeGuardianLeaves(gameMap)

		// Then
		assert.NoError(t, err)
		assert.Equal(t, 41, result)
	})
}
