package main

import (
	"reflect"
	"testing"
)

func TestGetUserPaymentCategorySummary(t *testing.T) {
	t.Run("No tracked payments should return empty category slice", func(t *testing.T) {

		got := GroupPaymentsByCategory([]Payment{}, []Payment{})
		want := []PaymentCategorySummary{}

		assertPaymentCategorySummaries(want, got, t)
	})

	t.Run("Category payments across both months should return expected category slice", func(t *testing.T) {

		prev := []Payment{
			{
				Category: "Reading",
				Price:    10,
			},
		}

		curr := []Payment{
			{
				Category: "Reading",
				Price:    10,
			},
			{
				Category: "Reading",
				Price:    25,
			},
		}

		got := GroupPaymentsByCategory(prev, curr)
		want := []PaymentCategorySummary{
			{
				Category:           "Reading",
				PreviousMonthTotal: 10,
				CurrentMonthTotal:  35,
			},
		}

		assertPaymentCategorySummaries(want, got, t)
	})

	t.Run("Category payments in previous month only should return expected category slice", func(t *testing.T) {

		prev := []Payment{
			{
				Category: "Reading",
				Price:    10,
			},
		}

		curr := []Payment{}

		got := GroupPaymentsByCategory(prev, curr)

		want := []PaymentCategorySummary{
			{
				Category:           "Reading",
				PreviousMonthTotal: 10,
				CurrentMonthTotal:  0,
			},
		}

		assertPaymentCategorySummaries(want, got, t)
	})

	t.Run("Category payments in current month only should return expected category slice", func(t *testing.T) {
		prev := []Payment{}

		curr := []Payment{
			{
				Category: "Reading",
				Price:    10,
			},
			{
				Category: "Reading",
				Price:    25,
			},
		}

		got := GroupPaymentsByCategory(prev, curr)
		want := []PaymentCategorySummary{
			{
				Category:           "Reading",
				PreviousMonthTotal: 0,
				CurrentMonthTotal:  35,
			},
		}

		assertPaymentCategorySummaries(want, got, t)
	})
}

func assertPaymentCategorySummaries(want, got []PaymentCategorySummary, t *testing.T) {

	if !reflect.DeepEqual(want, got) {
		t.Errorf("want: %v, got: %v", want, got)
	}
}
