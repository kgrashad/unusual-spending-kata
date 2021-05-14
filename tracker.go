package main

type Tracker struct{}
type Payment struct{}

func (t *Tracker) GetUserPayments(userId int) (currMonth, prevMonth []Payment) {
	return
}
