package main

import "testing"

func Test_area(t *testing.T) {
	type args struct {
		p point
		q point
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"1", args{point{2, 5}, point{9, 7}}, 24},
		{"2", args{point{7, 1}, point{11, 7}}, 35},
		{"3", args{point{7, 3}, point{2, 3}}, 6},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := area(tt.args.p, tt.args.q); got != tt.want {
				t.Errorf("area() = %v, want %v", got, tt.want)
			}
		})
	}
}
