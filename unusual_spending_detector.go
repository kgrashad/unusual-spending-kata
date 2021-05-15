package main

func DetectUnusualSpending(input []PaymentCategorySummary) (result []PaymentCategorySummary) {

	result = make([]PaymentCategorySummary, 0)

	for _, s := range input {

		if (s.CurrentMonthTotal-s.PreviousMonthTotal)/s.PreviousMonthTotal >= 0.5 {
			result = append(result, s)
		}
	}

	return
}
