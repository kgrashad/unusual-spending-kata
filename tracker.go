package main

type Tracker struct {
	currentMonthPayments  []Payment
	previousMonthPayments []Payment
}
type Payment struct{}
type Month string

const (
	CurrentMonth  Month = "Current Month"
	PreviousMonth Month = "Previous Month"
)

func (t *Tracker) TrackUserPayment(userId int, payment Payment, month Month) {
	switch month {
	case CurrentMonth:
		t.currentMonthPayments = append(t.currentMonthPayments, payment)
	case PreviousMonth:
		t.previousMonthPayments = append(t.previousMonthPayments, payment)
	}
}

func (t *Tracker) GetUserPayments(userId int) (currMonth, prevMonth []Payment) {
	return t.currentMonthPayments, t.previousMonthPayments
}
