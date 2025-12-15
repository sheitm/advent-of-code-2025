package main

import (
	"testing"
)

func Test_newMachine(t *testing.T) {
	line := "[.##.] (3) (1,3) (2) (2,3) (0,2) (0,1) {3,5,4,7}"
	m := newMachine(line)
	_ = m
}
