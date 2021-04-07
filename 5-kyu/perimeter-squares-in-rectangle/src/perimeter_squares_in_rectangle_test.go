package src

import "testing"

func TestPerimeter(t *testing.T) {
	type args struct {
		n int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"", args{5}, 80},
		{"", args{7}, 216},
		{"", args{20}, 114624},
		{"", args{30}, 14098308},
		{"", args{70}, 3226062132197568},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Perimeter(tt.args.n); got != tt.want {
				t.Errorf("Perimeter() = %v, want %v", got, tt.want)
			}
		})
	}
}
