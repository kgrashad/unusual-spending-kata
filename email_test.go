package main

import "testing"

func TestGenerateEmail(t *testing.T) {

	t.Run("When no unusual spending, return an error", func(t *testing.T) {
		_, _, err := GenerateEmail([]PaymentCategorySummary{})

		if err != ErrEmptySpendingInput {
			t.Errorf("want %v, got %v", ErrEmptySpendingInput, err)
		}
	})

	t.Run("When unusual spending, don't return an error", func(t *testing.T) {
		input := []PaymentCategorySummary{
			{
				Category:           "Reading",
				PreviousMonthTotal: 100,
				CurrentMonthTotal:  150,
			},
		}

		_, _, err := GenerateEmail(input)

		if err != nil {
			t.Errorf("want %v, got %v", nil, err)
		}
	})
}
