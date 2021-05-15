package main

import (
	"testing"
)

func TestDetectUnusualSpending(t *testing.T) {
	t.Run("No payment categories should return empty slice", func(t *testing.T) {
		got := DetectUnusualSpending([]PaymentCategorySummary{})
		want := []PaymentCategorySummary{}
		assertPaymentCategorySummaries(want, got, t)
	})
}
