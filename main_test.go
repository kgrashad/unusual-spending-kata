package main

import (
	"bytes"
	"testing"
)

func TestRun(t *testing.T) {

	buffer := &bytes.Buffer{}
	Run(buffer)

	got := buffer.String()
	want := "Unusual spending of $1076 detected!\nHello card user!\n\nWe have detected unusually high spending on your card in these categories:\n\n* You spent $148 on groceries\n* You spent $928 on travel\n\nLove,\n\nThe Credit Card Company\n"

	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}
