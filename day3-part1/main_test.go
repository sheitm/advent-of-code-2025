package main

import "testing"

func Test_highestJoltageInBank(t *testing.T) {
	type args struct {
		bank string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"1", args{bank: "987654321111111"}, 98},
		{"2", args{bank: "811111111111119"}, 89},
		{"3", args{bank: "234234234234278"}, 78},
		{"4", args{bank: "818181911112111"}, 92},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := highestJoltageInBank(tt.args.bank); got != tt.want {
				t.Errorf("highestJoltageInBank() = %v, want %v", got, tt.want)
			}
		})
	}
}
