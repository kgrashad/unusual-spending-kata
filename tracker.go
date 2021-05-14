package main

type Tracker struct {
	currentMonthPayments []Payment
}
type Payment struct{}
type Month string

const (
	CurrentMonth  Month = "Current Month"
	PreviousMonth Month = "Previous Month"
)

func (t *Tracker) TrackUserPayment(userId int, payment Payment, month Month) {
	t.currentMonthPayments = append(t.currentMonthPayments, payment)
}

func (t *Tracker) GetUserPayments(userId int) (currMonth, prevMonth []Payment) {
	return t.currentMonthPayments, []Payment{}
}
