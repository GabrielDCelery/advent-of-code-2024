package internals

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_solveHowManyTimesTemplatesAppearInInput(t *testing.T) {
	t.Run("Correctly calculates how many times templates appear in an input", func(t *testing.T) {
		// Given

		input := `MMMSXXMASM
MSAMXMSMSA
AMXSXMAAMM
MSAMASMSMX
XMASAMXAMM
XXAMMXXAMA
SMSMSASXSS
SAXAMASAAA
MAMMMXMMMM
MXMXAXMASX`

		templates := `XMAS
$$$$
X...
.M..
..A.
...S
$$$$
X
M
A
S
$$$$
...X
..M.
.A..
S...
$$$$
SAMX
$$$$
S...
.A..
..M.
...X
$$$$
S
A
M
X
$$$$
...S
..A.
.M..
X...
`

		// When
		result := solveHowManyTimesTemplatesAppearInInput(input, templates)

		// Then
		assert.Equal(t, 18, result)
	})

	t.Run("Correctly calculates how many times templates appear in an input", func(t *testing.T) {
		// Given

		input := `MMMSXXMASM
MSAMXMSMSA
AMXSXMAAMM
MSAMASMSMX
XMASAMXAMM
XXAMMXXAMA
SMSMSASXSS
SAXAMASAAA
MAMMMXMMMM
MXMXAXMASX`

		templates := `M.S
.A.
M.S
$$$$
M.M
.A.
S.S
$$$$
S.M
.A.
S.M
$$$$
S.S
.A.
M.M
`

		// When
		result := solveHowManyTimesTemplatesAppearInInput(input, templates)

		// Then
		assert.Equal(t, 9, result)
	})
}
