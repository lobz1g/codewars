package src

import "testing"

func TestDblLinear(t *testing.T) {
	type args struct {
		n int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"0", args{10}, 22},
		{"1", args{50}, 175},
		{"2", args{100}, 447},
		{"3", args{500}, 3355},
		{"4", args{1000}, 8488},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := DblLinear(tt.args.n); got != tt.want {
				t.Errorf("DblLinear() = %v, want %v", got, tt.want)
			}
		})
	}
}
