package day_09

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_calculateFileSystemChecksum(t *testing.T) {
	t.Run("Successfully calculates file system checksum from input", func(t *testing.T) {
		// Given
		input := `2333133121414131402`

		// When
		result, err := calculateFileSystemChecksum(input)

		// Then
		assert.NoError(t, err)
		assert.Equal(t, 1928, result)
	})
}
