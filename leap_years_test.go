package main

import "testing"

func TestLeapYears01(t *testing.T) {
	actual := leapYears(1)
	expected := false
	if actual != expected {
		t.Error("Error")
	}
}
