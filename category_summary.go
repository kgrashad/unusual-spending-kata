package main

type PaymentCategorySummary struct {
	Category           string
	PreviousMonthTotal float64
	CurrentMonthTotal  float64
}

func GroupPaymentsByCategory(prevMonth, currMonth []Payment) []PaymentCategorySummary {

	currCategories := groupPaymentsByCategories(currMonth)
	prevCategories := groupPaymentsByCategories(prevMonth)

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
