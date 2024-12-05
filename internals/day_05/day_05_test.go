package day_05

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_doesUpdateSatisfyRules(t *testing.T) {
	t.Run("Correctly determines that an update with 75, 47, 61, 53, 29 satisfies the rules", func(t *testing.T) {
		// Given
		rules := map[int][]int{
			29: {13},
			47: {13, 29, 53, 61},
			53: {13, 29},
			61: {13, 29, 53},
			75: {13, 29, 47, 53, 61},
			97: {13, 29, 47, 53, 61, 75},
		}
		updates := []int{75, 47, 61, 53, 29}

		// When
		result := doesUpdateSatisfyRules(updates, rules)

		// Then
		assert.Equal(t, true, result)
	})

	t.Run("Correctly determines that an update with 97, 61, 53, 29, 13 satisfies the rules", func(t *testing.T) {
		// Given
		rules := map[int][]int{
			29: {13},
			47: {13, 29, 53, 61},
			53: {13, 29},
			61: {13, 29, 53},
			75: {13, 29, 47, 53, 61},
			97: {13, 29, 47, 53, 61, 75},
		}
		updates := []int{97, 61, 53, 29, 13}

		// When
		result := doesUpdateSatisfyRules(updates, rules)

		// Then
		assert.Equal(t, true, result)
	})

	t.Run("Correctly determines that an update with 75, 29, 13 satisfies the rules", func(t *testing.T) {
		// Given
		rules := map[int][]int{
			29: {13},
			47: {13, 29, 53, 61},
			53: {13, 29},
			61: {13, 29, 53},
			75: {13, 29, 47, 53, 61},
			97: {13, 29, 47, 53, 61, 75},
		}
		updates := []int{75, 29, 13}

		// When
		result := doesUpdateSatisfyRules(updates, rules)

		// Then
		assert.Equal(t, true, result)
	})

	t.Run("Correctly determines that the update 75, 97, 47, 61, 53 would break the rule 97|75", func(t *testing.T) {
		// Given
		rules := map[int][]int{
			29: {13},
			47: {13, 29, 53, 61},
			53: {13, 29},
			61: {13, 29, 53},
			75: {13, 29, 47, 53, 61},
			97: {13, 29, 47, 53, 61, 75},
		}
		updates := []int{75, 97, 47, 61, 53}

		// When
		result := doesUpdateSatisfyRules(updates, rules)

		// Then
		assert.Equal(t, false, result)
	})

	t.Run("Correctly determines that the update with 61, 13, 29 would break the rule 29|13", func(t *testing.T) {
		// Given
		rules := map[int][]int{
			29: {13},
			47: {13, 29, 53, 61},
			53: {13, 29},
			61: {13, 29, 53},
			75: {13, 29, 47, 53, 61},
			97: {13, 29, 47, 53, 61, 75},
		}
		updates := []int{61, 13, 29}

		// When
		result := doesUpdateSatisfyRules(updates, rules)

		// Then
		assert.Equal(t, false, result)
	})

	t.Run("Correctly determines that the update with 97, 13, 75, 29, 47 breaks several rules", func(t *testing.T) {
		// Given
		rules := map[int][]int{
			29: {13},
			47: {13, 29, 53, 61},
			53: {13, 29},
			61: {13, 29, 53},
			75: {13, 29, 47, 53, 61},
			97: {13, 29, 47, 53, 61, 75},
		}
		updates := []int{97, 13, 75, 29, 47}

		// When
		result := doesUpdateSatisfyRules(updates, rules)

		// Then
		assert.Equal(t, false, result)
	})
}

func Test_sumTheMiddleNumbersOfCorrectUpdates(t *testing.T) {
	t.Run("Correctly sums up the middle numbers of the correct updates", func(t *testing.T) {
		// Given
		rules := map[int][]int{
			29: {13},
			47: {13, 29, 53, 61},
			53: {13, 29},
			61: {13, 29, 53},
			75: {13, 29, 47, 53, 61},
			97: {13, 29, 47, 53, 61, 75},
		}
		updates := [][]int{
			{75, 47, 61, 53, 29},
			{97, 61, 53, 29, 13},
			{75, 29, 13},
			{75, 97, 47, 61, 53},
			{61, 13, 29},
			{97, 13, 75, 29, 47},
		}

		// When
		result := sumTheMiddleNumbersOfCorrectUpdates(updates, rules)

		// Then
		assert.Equal(t, 143, result)
	})
}

func Test_sortUpdateToBeInCorrectOrder(t *testing.T) {
	t.Run("Sort update of 75,97,47,61,53 to be in correct order of 97,75,47,61,53", func(t *testing.T) {
		// Given
		rules := map[int][]int{
			29: {13},
			47: {13, 29, 53, 61},
			53: {13, 29},
			61: {13, 29, 53},
			75: {13, 29, 47, 53, 61},
			97: {13, 29, 47, 53, 61, 75},
		}

		update := []int{75, 97, 47, 61, 53}

		// When
		sortUpdateToBeInCorrectOrder(update, rules)

		// Then
		assert.Equal(t, []int{97, 75, 47, 61, 53}, update)
	})

	t.Run("Sort update of 61,13,29 to be in correct order of 61,29,13", func(t *testing.T) {
		// Given
		rules := map[int][]int{
			29: {13},
			47: {13, 29, 53, 61},
			53: {13, 29},
			61: {13, 29, 53},
			75: {13, 29, 47, 53, 61},
			97: {13, 29, 47, 53, 61, 75},
		}

		update := []int{61, 13, 29}

		// When
		sortUpdateToBeInCorrectOrder(update, rules)

		// Then
		assert.Equal(t, []int{61, 29, 13}, update)
	})

	t.Run("Sort update of 97,13,75,29,47 to be in correct order of 97,75,47,29,13", func(t *testing.T) {
		// Given
		rules := map[int][]int{
			29: {13},
			47: {13, 29, 53, 61},
			53: {13, 29},
			61: {13, 29, 53},
			75: {13, 29, 47, 53, 61},
			97: {13, 29, 47, 53, 61, 75},
		}

		update := []int{97, 13, 75, 29, 47}

		// When
		sortUpdateToBeInCorrectOrder(update, rules)

		// Then
		assert.Equal(t, []int{97, 75, 47, 29, 13}, update)
	})
}

func Test_fixTheOrderOfIncorrectUpdatesThenSumTheMiddleNumbers(t *testing.T) {
	t.Run("Correctly sums up the middle numbers of the correct updates", func(t *testing.T) {
		// Given
		rules := map[int][]int{
			29: {13},
			47: {13, 29, 53, 61},
			53: {13, 29},
			61: {13, 29, 53},
			75: {13, 29, 47, 53, 61},
			97: {13, 29, 47, 53, 61, 75},
		}
		updates := [][]int{
			{75, 47, 61, 53, 29},
			{97, 61, 53, 29, 13},
			{75, 29, 13},
			{75, 97, 47, 61, 53},
			{61, 13, 29},
			{97, 13, 75, 29, 47},
		}

		// When
		result := fixTheOrderOfIncorrectUpdatesThenSumTheMiddleNumbers(updates, rules)

		// Then
		assert.Equal(t, 123, result)
	})
}
