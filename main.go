package main

import (
	"fmt"
	"io"
	"os"
)

const UserId int = 1

func Run(out io.Writer) {
	t := NewTracker()

	t.TrackUserPayment(UserId, Payment{
		Category: "groceries",
		Price:    10,
	}, PreviousMonth)

	t.TrackUserPayment(UserId, Payment{
		Category: "groceries",
		Price:    148,
	}, CurrentMonth)

	t.TrackUserPayment(UserId, Payment{
		Category: "travel",
		Price:    500,
	}, CurrentMonth)

	t.TrackUserPayment(UserId, Payment{
		Category: "travel",
		Price:    428,
	}, CurrentMonth)

	curr, prev := t.GetUserPayments(UserId)
	categories := GroupPaymentsByCategory(prev, curr)
	unusualSpending := DetectUnusualSpending(categories)
	subject, body, err := GenerateEmail(unusualSpending)

	if err == nil {
		fmt.Fprintln(out, subject)
		fmt.Fprintln(out, body)
	}
}

func main() {
	Run(os.Stdout)
}
