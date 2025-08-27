package main

import "testing"

func TestPart1(t *testing.T) {
	if SolvePart1() != 11 {
		t.FailNow()
	}
}

func TestPart2(t *testing.T) {
	if SolvePart2() != 31 {
		t.FailNow()
	}
}
