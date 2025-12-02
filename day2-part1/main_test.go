package main

import "testing"

func Test_checkID(t *testing.T) {
	type args struct {
		id int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"single", args{id: 1}, 0},
		{"not symmetrical", args{id: 110234}, 0},
		{"odd", args{id: 1102345}, 0},
		{"short symmetrical", args{id: 33}, 33},
		{"medium symmetrical", args{id: 330330}, 330330},
		{"long symmetrical", args{id: 4433044330}, 4433044330},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := checkID(tt.args.id); got != tt.want {
				t.Errorf("checkID() = %v, want %v", got, tt.want)
			}
		})
	}
}
