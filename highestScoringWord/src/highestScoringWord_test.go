package src

import (
	"testing"
)

func TestHigh(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{name: "", args: args{"man i need a taxi up to ubud"}, want: "taxi"},
		{name: "", args: args{"what time are we climbing up the volcano"}, want: "volcano"},
		{name: "", args: args{"take me to semynak"}, want: "semynak"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := High(tt.args.s); got != tt.want {
				t.Errorf("High() = %v, want %v", got, tt.want)
			}
		})
	}
}
