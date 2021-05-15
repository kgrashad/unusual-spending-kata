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

	t.Run("Equal payment total should return empty slice", func(t *testing.T) {
		summary := []PaymentCategorySummary{
			{
				Category:           "Reading",
				PreviousMonthTotal: 100,
				CurrentMonthTotal:  100,
			},
		}
		got := DetectUnusualSpending(summary)
		want := []PaymentCategorySummary{}
		assertPaymentCategorySummaries(want, got, t)
	})

	t.Run("49.999999999% increase in payment total should return empty slice", func(t *testing.T) {
		summary := []PaymentCategorySummary{
			{
				Category:           "Reading",
				PreviousMonthTotal: 100,
				CurrentMonthTotal:  149.999999999,
			},
		}
		got := DetectUnusualSpending(summary)
		want := []PaymentCategorySummary{}
		assertPaymentCategorySummaries(want, got, t)
	})

	t.Run("50% decrease in payment total should return empty slice", func(t *testing.T) {
		summary := []PaymentCategorySummary{
			{
				Category:           "Reading",
				PreviousMonthTotal: 100,
				CurrentMonthTotal:  50,
			},
		}
		got := DetectUnusualSpending(summary)
		want := []PaymentCategorySummary{}
		assertPaymentCategorySummaries(want, got, t)
	})

	t.Run("50% increase in payment total should return summary", func(t *testing.T) {
		summary := []PaymentCategorySummary{
			{
				Category:           "Reading",
				PreviousMonthTotal: 100,
				CurrentMonthTotal:  150,
			},
		}
		got := DetectUnusualSpending(summary)
		want := summary
		assertPaymentCategorySummaries(want, got, t)
	})

	t.Run("Zero total in previous month should return the expected result", func(t *testing.T) {
		summary := []PaymentCategorySummary{
			{
				Category:           "Reading",
				PreviousMonthTotal: 0,
				CurrentMonthTotal:  150,
			},
		}
		got := DetectUnusualSpending(summary)
		want := summary
		assertPaymentCategorySummaries(want, got, t)
	})

	t.Run("Slice with mixed summaries should return ones with 50% spike only", func(t *testing.T) {
		summary := []PaymentCategorySummary{
			{
				Category:           "Reading",
				PreviousMonthTotal: 200,
				CurrentMonthTotal:  4000,
			}, {
				Category:           "Fun",
				PreviousMonthTotal: 100,
				CurrentMonthTotal:  50,
			}, {
				Category:           "Space Travel",
				PreviousMonthTotal: 0,
				CurrentMonthTotal:  0,
			},
		}
		got := DetectUnusualSpending(summary)
		want := []PaymentCategorySummary{
			{
				Category:           "Reading",
				PreviousMonthTotal: 200,
				CurrentMonthTotal:  4000,
			},
		}
		assertPaymentCategorySummaries(want, got, t)
	})
}
