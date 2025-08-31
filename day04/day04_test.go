package day04

import (
	"testing"
)

func TestPart1(t *testing.T) {
	expected := 18
	actual := SolvePart1()
	if expected != actual {
		t.Errorf("Expected: %d. Actual: %d", expected, actual)
	}
}

func TestPart2(t *testing.T) {
	expected := 9
	actual := SolvePart2()
	if expected != actual {
		t.Errorf("Expected: %d. Actual: %d", expected, actual)
	}
}
