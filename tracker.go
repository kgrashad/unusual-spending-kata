package main

type Tracker struct {
	currentMonthPayments  map[int][]Payment
	previousMonthPayments map[int][]Payment
}
type Payment struct {
	Description string
}
type Month string

const (
	CurrentMonth  Month = "Current Month"
	PreviousMonth Month = "Previous Month"
)

func NewTracker() *Tracker {
	var t Tracker
	t.currentMonthPayments = make(map[int][]Payment)
	t.previousMonthPayments = make(map[int][]Payment)

	return &t
}

func (t *Tracker) TrackUserPayment(userId int, payment Payment, month Month) {
	switch month {
	case CurrentMonth:
		t.currentMonthPayments[userId] = append(t.currentMonthPayments[userId], payment)
	case PreviousMonth:
		t.previousMonthPayments[userId] = append(t.previousMonthPayments[userId], payment)
	}
}

func (t *Tracker) GetUserPayments(userId int) (currMonth, prevMonth []Payment) {
	return t.currentMonthPayments[userId], t.previousMonthPayments[userId]
}
