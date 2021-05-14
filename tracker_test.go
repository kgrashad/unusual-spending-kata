package main

import "testing"

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
}
