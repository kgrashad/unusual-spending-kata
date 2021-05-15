package main

import (
	"errors"
	"fmt"
)

var ErrEmptySpendingInput = errors.New("can't generate email if no unusual spending passed.")

const EmailSubjectTemplate = "Unusual spending of $%0.0f detected!"
const EmailBodyTemplate string = `Hello card user!

We have detected unusually high spending on your card in these categories:

* You spent $%0.0f on %v

Love,

The Credit Card Company`

func GenerateEmail(categories []PaymentCategorySummary) (subject, body string, err error) {
	if categories == nil || len(categories) == 0 {
		return "", "", ErrEmptySpendingInput
	}

	subject = fmt.Sprintf(EmailSubjectTemplate, getTotal(categories))
	body = fmt.Sprintf(EmailBodyTemplate, categories[0].CurrentMonthTotal, categories[0].Category)

	return
}

func getTotal(categories []PaymentCategorySummary) (total float64) {
	for _, c := range categories {
		total += c.CurrentMonthTotal
	}

	return
}
