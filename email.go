package main

import "errors"

var ErrEmptySpendingInput = errors.New("can't generate email if no unusual spending passed.")

func GenerateEmail(categories []PaymentCategorySummary) (subject, body string, err error) {
	if categories == nil || len(categories) == 0 {
		return "", "", ErrEmptySpendingInput
	}

	return "", "", nil
}
