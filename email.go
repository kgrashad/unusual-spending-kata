package main

import (
	"errors"
	"fmt"
)

var ErrEmptySpendingInput = errors.New("can't generate email if no unusual spending passed.")

const EmailSubjectTemplate = "Unusual spending of $%0.0f detected!"
const EmailBodyTemplate string = `Hello card user!

We have detected unusually high spending on your card in these categories:

%s
Love,

The Credit Card Company`

const EmailBodyLineTemplate string = "* You spent $%0.0f on %v"

func GenerateEmail(categories []PaymentCategorySummary) (subject, body string, err error) {
	if categories == nil || len(categories) == 0 {
		return "", "", ErrEmptySpendingInput
	}

	return getSubject(categories), getBody(categories), nil
}

func getSubject(categories []PaymentCategorySummary) string {
	total := 0.0

	for _, c := range categories {
		total += c.CurrentMonthTotal
	}

	return fmt.Sprintf(EmailSubjectTemplate, total)
}

func getBody(categories []PaymentCategorySummary) string {
	bodyLines := ""

	for _, c := range categories {
		bodyLines += fmt.Sprintf(EmailBodyLineTemplate, c.CurrentMonthTotal, c.Category)
		bodyLines += fmt.Sprintln()
	}

	return fmt.Sprintf(EmailBodyTemplate, bodyLines)
}
