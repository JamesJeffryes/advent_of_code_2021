package main

import (
	"testing"
)

func TestPt1(t *testing.T) {
	vectors := getInput("test_input.txt")
	print(vectors)
	pos := Position{0, 0, 0}
	pos = ApplyVectors(pos, vectors)
	mult := pos.x * pos.y
	if mult != 150 {
		t.Errorf("Expected 150 increases. Got %d", mult)
	}
}

func TestPt2(t *testing.T) {
	vectors := getInput("test_input.txt")
	print(vectors)
	pos := Position{0, 0, 0}
	pos = ApplyAimVectors(pos, vectors)
	expect := 900
	result := pos.x * pos.y
	if result != expect {
		t.Errorf("Expected %d increases. Got %d", expect, result)
	}
}
