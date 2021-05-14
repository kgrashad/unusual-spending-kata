package main

type Tracker struct {
	currentMonthPayments  map[int][]Payment
	previousMonthPayments map[int][]Payment
}

type Payment struct {
	Description string
	Category    string
	Price       float64
}

type PaymentCategorySummary struct {
	Category           string
	PreviousMonthTotal float64
	CurrentMonthTotal  float64
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

func (t *Tracker) GetUserPaymentCategorySummaries(userId int) []PaymentCategorySummary {
	curr, prev := t.GetUserPayments(userId)

	currCategories := groupPaymentsByCategories(curr)
	prevCategories := groupPaymentsByCategories(prev)

	return mergePaymentCategories(prevCategories, currCategories)
}

func groupPaymentsByCategories(payments []Payment) (results map[string]float64) {

	results = make(map[string]float64)

	for _, p := range payments {
		_, ok := results[p.Category]

		if !ok {
			results[p.Category] = 0
		}

		results[p.Category] += p.Price
	}

	return
}

func mergePaymentCategories(prev, curr map[string]float64) (results []PaymentCategorySummary) {

	results = make([]PaymentCategorySummary, 0)

	for k, v := range prev {
		results = append(results, PaymentCategorySummary{
			Category:           k,
			PreviousMonthTotal: v,
			CurrentMonthTotal:  curr[k],
		})
	}

	for k, v := range curr {
		_, ok := prev[k]

		if !ok {
			results = append(results, PaymentCategorySummary{
				Category:           k,
				PreviousMonthTotal: 0,
				CurrentMonthTotal:  v,
			})
		}
	}

	return
}
