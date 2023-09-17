package src

import "testing"

func TestThirt(t *testing.T) {
	tests := []struct {
		name string
		n    int
		want int
	}{
		{
			name: "case 0",
			n:    1234567,
			want: 87,
		},
		{
			name: "case 1",
			n:    8529,
			want: 79,
		},
		{
			name: "case 2",
			n:    85299258,
			want: 31,
		},
		{
			name: "case 3",
			n:    5634,
			want: 57,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Thirt(tt.n); got != tt.want {
				t.Errorf("Thirt() = %v, want %v", got, tt.want)
			}
		})
	}
}
