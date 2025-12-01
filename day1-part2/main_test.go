package main

import "testing"

func Test_decrementWithWrapCount(t *testing.T) {
	type args struct {
		state     int
		decrement int
		base      int
	}
	tests := []struct {
		name         string
		args         args
		wantNewState int
		wantWraps    int
	}{
		{"once", args{50, 68, 100}, 82, 1},
		{"zero", args{82, 30, 100}, 52, 0},
		{"end on zero once", args{55, 55, 100}, 0, 1},
		{"many times", args{55, 1000, 100}, 55, 10},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotNewState, gotWraps := decrementWithWrapCount(tt.args.state, tt.args.decrement, tt.args.base)
			if gotNewState != tt.wantNewState {
				t.Errorf("decrementWithWrapCount() gotNewState = %v, want %v", gotNewState, tt.wantNewState)
			}
			if gotWraps != tt.wantWraps {
				t.Errorf("decrementWithWrapCount() gotWraps = %v, want %v", gotWraps, tt.wantWraps)
			}
		})
	}
}
