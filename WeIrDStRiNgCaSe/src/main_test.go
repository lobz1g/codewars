package src

import (
	"testing"
)

func Test_toWeirdCase(t *testing.T) {
	type args struct {
		str string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{"", args{"abc def"}, "AbC DeF"},
		{"", args{"ABC"}, "AbC"},
		{"", args{"This is a test Looks like you passed"}, "ThIs Is A TeSt LoOkS LiKe YoU PaSsEd"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := toWeirdCase(tt.args.str); got != tt.want {
				t.Errorf("toWeirdCase() = %v, want %v", got, tt.want)
			}
		})
	}
}
