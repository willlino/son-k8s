package main

import "testing"

func TestGreeting(t *testing.T ) {
	message := "teste"
	result := greeting(message)

	if result != "<b>" + message + "</b>" {
		t.Errorf("Value <b>%v</b> expected, but the value is %v", message, result)
	}
}