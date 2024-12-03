package internals

import (
	"regexp"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_createSeriesOfInstructionsFromInput(t *testing.T) {
	t.Run("Extracts mul commands from corrupted memory", func(t *testing.T) {
		// Given
		corruptedMemory := "xmul(1,1)xdo()xdon't()xmul(2,2)xdo()xdon't()x"
		instructionConfigs := []InstructionConfig{{Id: MulID, Re: regexp.MustCompile(MulRegexp)}}

		// When
		result := createSeriesOfInstructionsFromInput([]byte(corruptedMemory), instructionConfigs)

		// Then
		expected := []Instruction{
			{Id: MulID, Position: 1, Value: []byte("mul(1,1)")},
			{Id: MulID, Position: 23, Value: []byte("mul(2,2)")},
		}
		assert.Equal(t, expected, result)
	})

	t.Run("Extracts do commands from corrupted memory", func(t *testing.T) {
		// Given
		corruptedMemory := "xmul(1,1)xdo()xdon't()xmul(2,2)xdo()xdon't()x"
		instructionConfigs := []InstructionConfig{{Id: DoID, Re: regexp.MustCompile(DoRegexp)}}

		// When
		result := createSeriesOfInstructionsFromInput([]byte(corruptedMemory), instructionConfigs)

		// Then
		expected := []Instruction{
			{Id: DoID, Position: 10, Value: []byte("do()")},
			{Id: DoID, Position: 32, Value: []byte("do()")},
		}
		assert.Equal(t, expected, result)
	})

	t.Run("Extracts don't commands from corrupted memory", func(t *testing.T) {
		// Given
		corruptedMemory := "xmul(1,1)xdo()xdon't()xmul(2,2)xdo()xdon't()x"
		instructionConfigs := []InstructionConfig{{Id: DontID, Re: regexp.MustCompile(DontRegexp)}}

		// When
		result := createSeriesOfInstructionsFromInput([]byte(corruptedMemory), instructionConfigs)

		// Then
		expected := []Instruction{
			{Id: DontID, Position: 15, Value: []byte("don't()")},
			{Id: DontID, Position: 37, Value: []byte("don't()")},
		}
		assert.Equal(t, expected, result)
	})

	t.Run("Extracts mul commands from corrupted memory", func(t *testing.T) {
		// Given
		corruptedMemory := "xmul(2,4)%&mul[3,7]!@^do_not_mul(5,5)+mul(32,64]then(mul(11,8)mul(8,5))"
		instructionConfigs := []InstructionConfig{{Id: MulID, Re: regexp.MustCompile(MulRegexp)}}

		// When
		result := createSeriesOfInstructionsFromInput([]byte(corruptedMemory), instructionConfigs)

		// Then
		expected := []Instruction{
			{Id: MulID, Position: 1, Value: []byte("mul(2,4)")},
			{Id: MulID, Position: 29, Value: []byte("mul(5,5)")},
			{Id: MulID, Position: 53, Value: []byte("mul(11,8)")},
			{Id: MulID, Position: 62, Value: []byte("mul(8,5)")},
		}
		assert.Equal(t, expected, result)
	})
}

func Test_solveInput(t *testing.T) {
	t.Run("Correctly extracts and executes instructions from corrupted input that only contain mul", func(t *testing.T) {
		// Given
		corruptedMemory := "xmul(2,4)%&mul[3,7]!@^do_not_mul(5,5)+mul(32,64]then(mul(11,8)mul(8,5))"
		instructionConfigs := []InstructionConfig{{Id: MulID, Re: regexp.MustCompile(MulRegexp)}}

		// When
		result, err := solveInput([]byte(corruptedMemory), instructionConfigs)

		// Then
		assert.NoError(t, err)
		assert.Equal(t, 161, result)
	})

	t.Run("Correctly extracts and executes instructions from corrupted input that only contain do, don't and mul", func(t *testing.T) {
		// Given
		corruptedMemory := "xmul(2,4)&mul[3,7]!^don't()_mul(5,5)+mul(32,64](mul(11,8)undo()?mul(8,5))"
		instructionConfigs := []InstructionConfig{
			{Id: MulID, Re: regexp.MustCompile(MulRegexp)},
			{Id: DoID, Re: regexp.MustCompile(DoRegexp)},
			{Id: DontID, Re: regexp.MustCompile(DontRegexp)},
		}

		// When
		result, err := solveInput([]byte(corruptedMemory), instructionConfigs)

		// Then
		assert.NoError(t, err)
		assert.Equal(t, 48, result)
	})

}
