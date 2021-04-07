package src

import "testing"

func TestMultiple3And5(t *testing.T) {
	type args struct {
		number int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"", args{10}, 23},
		{"", args{20}, 78},
		{"", args{0}, 0},
		{"", args{1}, 0},
		{"", args{200}, 9168},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Multiple3And5(tt.args.number); got != tt.want {
				t.Errorf("Multiple3And5() = %v, want %v", got, tt.want)
			}
		})
	}
}
