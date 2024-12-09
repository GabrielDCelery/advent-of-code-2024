package day_09

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_calculateFileSystemChecksumV1(t *testing.T) {
	t.Run("Successfully calculates file system checksum from input", func(t *testing.T) {
		// Given
		input := `2333133121414131402`

		// When
		result, err := calculateFileSystemChecksumV1(input)

		// Then
		assert.NoError(t, err)
		assert.Equal(t, 1928, result)
	})
}

func Test_calculateFileSystemChecksumV2(t *testing.T) {
	t.Run("Successfully calculates file system checksum from input", func(t *testing.T) {
		// Given
		input := `2333133121414131402`

		// When
		result, err := calculateFileSystemChecksumV2(input)

		// Then
		assert.NoError(t, err)
		assert.Equal(t, 2858, result)
	})
}
