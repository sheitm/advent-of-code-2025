package main

import "testing"

func Test_compute(t *testing.T) {
	accum := map[int]string{
		0: "*9*",
		1: "91*",
		2: "86*",
		3: "62*",
	}

	op := "+"

	total := compute(accum, op)

	expectedTotal := 986 + 9162

	if total != expectedTotal {
		t.Errorf("Expected %d, got %d", expectedTotal, total)
	}
}

func Test_compute2(t *testing.T) {
	accum := map[int]string{
		0: "921*",
		1: "716*",
		2: "138*",
		3: "455*",
	}

	op := "+"

	total := compute(accum, op)

	expectedTotal := 986 + 9162

	if total != expectedTotal {
		t.Errorf("Expected %d, got %d", expectedTotal, total)
	}
}
