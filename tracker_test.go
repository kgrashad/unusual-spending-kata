package main

import (
	"reflect"
	"testing"
)

func TestGetUserPayments(t *testing.T) {
	t.Run("Zero payments should return nil", func(t *testing.T) {
		tracker := NewTracker()
		userId := 1

		currMonthPayments, prevMonthPayments := tracker.GetUserPayments(userId)

		assertPayments(nil, currMonthPayments, t)
		assertPayments(nil, prevMonthPayments, t)
	})

	t.Run("When current month payment is tracked, GetUserPayments should return it.", func(t *testing.T) {
		tracker := NewTracker()
		payment := Payment{}
		userId := 1

		tracker.TrackUserPayment(userId, payment, CurrentMonth)

		currMonthPayments, _ := tracker.GetUserPayments(userId)
		want := []Payment{payment}

		assertPayments(want, currMonthPayments, t)
	})

	t.Run("When previous month payment is tracked, GetUserPayments should return it.", func(t *testing.T) {
		tracker := NewTracker()
		payment := Payment{}
		userId := 1

		tracker.TrackUserPayment(userId, payment, PreviousMonth)

		_, prevMonthPayments := tracker.GetUserPayments(userId)
		want := []Payment{payment}

		assertPayments(want, prevMonthPayments, t)
	})

	t.Run("When multiple users' payments are tracked, only asked for user is returned", func(t *testing.T) {
		tracker := NewTracker()
		firstUserId := 1
		secondUserId := 2
		firstUserPayment := Payment{
			Description: "I am the first user's payment",
		}
		secondUserPayment := Payment{
			Description: "I am the second user's payment",
		}

		tracker.TrackUserPayment(firstUserId, firstUserPayment, PreviousMonth)
		tracker.TrackUserPayment(secondUserId, secondUserPayment, PreviousMonth)

		_, prevMonthPayments := tracker.GetUserPayments(firstUserId)
		want := []Payment{firstUserPayment}

		assertPayments(want, prevMonthPayments, t)
	})
}

func TestGetUserPaymentCategorySummary(t *testing.T) {
	t.Run("No tracked payments should return empty category slice", func(t *testing.T) {
		tracker := NewTracker()
		userId := 1

		got := tracker.GetUserPaymentCategorySummary(userId)
		want := []PaymentCategorySummary{}

		assertPaymentCategorySummaries(want, got, t)
	})
}

func assertPayments(want, got []Payment, t *testing.T) {

	if !reflect.DeepEqual(want, got) {
		t.Errorf("got: %v, want: %v", want, got)
	}
}

func assertPaymentCategorySummaries(want, got []PaymentCategorySummary, t *testing.T) {

	if !reflect.DeepEqual(want, got) {
		t.Errorf("got: %v, want: %v", want, got)
	}
}
