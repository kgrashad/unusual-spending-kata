package main

import (
	"reflect"
	"testing"
)

func TestTracker(t *testing.T) {
	t.Run("Zero payments should return empty list", func(t *testing.T) {
		tracker := Tracker{}
		userId := 1

		currMonth, prevMonth := tracker.GetUserPayments(userId)

		if len(currMonth) != 0 {
			t.Errorf("Current month payments list expected to be empty, got %v", currMonth)
		}

		if len(prevMonth) != 0 {
			t.Errorf("Previous month payments list expected to be empty, got %v", prevMonth)
		}
	})

	t.Run("When one current month payment is tracker, GetUserPayments should return it.", func(t *testing.T) {
		tracker := Tracker{}
		payment := Payment{}
		userId := 1

		tracker.TrackUserPayment(userId, payment, CurrentMonth)

		currMonthPayments, _ := tracker.GetUserPayments(userId)
		want := []Payment{payment}

		println(len(currMonthPayments))

		if !reflect.DeepEqual(currMonthPayments, want) {
			t.Errorf("got: %v, want: %v", currMonthPayments, want)
		}
	})
}
