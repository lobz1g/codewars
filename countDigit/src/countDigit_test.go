package src

import "testing"

func TestNbDig(t *testing.T) {
	type args struct {
		n int
		d int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"", args{550, 5}, 213},
		{"", args{5750, 0}, 4700},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NbDig(tt.args.n, tt.args.d); got != tt.want {
				t.Errorf("NbDig() = %v, want %v", got, tt.want)
			}
		})
	}
}
