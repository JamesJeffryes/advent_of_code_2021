package main

import (
	"log"
	"testing"
)

func TestPt1(t *testing.T) {
	input := getInput("test_input.txt")
	gamma, epsilon, err := getMostCommon(input)
	if err != nil {
		log.Fatal(err)
	}
	expect := 198
	result := gamma * epsilon
	if result != expect {
		t.Errorf("Expected %d increases. Got %d", expect, result)
	}
}
