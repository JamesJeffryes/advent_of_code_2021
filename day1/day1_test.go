package main

import "testing"

func TestIncrease(t *testing.T) {
	readings := getReadings("test_input.txt")
	inc := countIncreases(readings)
	if inc != 7 {
		t.Errorf("Expected 7 increases. Got %d", inc)
	}
}

func TestWindowIncrease(t *testing.T) {
	readings := getReadings("test_input.txt")
	inc := countWindowIncreases(readings, 3)
	if inc != 5 {
		t.Errorf("Expected 5 increases. Got %d", inc)
	}
}