package main

import (
	"testing"
)

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

	t.Run("When single category is passed, expected email is generate", func(t *testing.T) {
		input := []PaymentCategorySummary{
			{
				Category:           "groceries",
				PreviousMonthTotal: 50,
				CurrentMonthTotal:  148,
			},
		}

		subject, body, _ := GenerateEmail(input)

		wantSubject := "Unusual spending of $148 detected!"
		wantBody := "Hello card user!\n\nWe have detected unusually high spending on your card in these categories:\n\n* You spent $148 on groceries\n\nLove,\n\nThe Credit Card Company"

		assertString(wantSubject, subject, t)
		assertString(wantBody, body, t)
	})

	t.Run("When multiple categories are passed, expected email is generate", func(t *testing.T) {
		input := []PaymentCategorySummary{
			{
				Category:           "groceries",
				PreviousMonthTotal: 50,
				CurrentMonthTotal:  148,
			}, {
				Category:           "travel",
				PreviousMonthTotal: 400,
				CurrentMonthTotal:  928,
			},
		}

		subject, body, _ := GenerateEmail(input)

		wantSubject := "Unusual spending of $1076 detected!"
		wantBody := "Hello card user!\n\nWe have detected unusually high spending on your card in these categories:\n\n* You spent $148 on groceries\n* You spent $928 on travel\n\nLove,\n\nThe Credit Card Company"

		assertString(wantSubject, subject, t)
		assertString(wantBody, body, t)
	})
}

func assertString(want, got string, t *testing.T) {
	if want != got {
		t.Errorf("want: %q, got: %q", want, got)
	}
}
