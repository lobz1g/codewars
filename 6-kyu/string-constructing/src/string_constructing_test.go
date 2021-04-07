package src

import "testing"

func TestStringConstructing(t *testing.T) {
	type args struct {
		a string
		s string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"", args{"aba", "abbabba"}, 9},
		{"", args{"aba", "abaa"}, 4},
		{"", args{"aba", "a"}, 3},
		{"", args{"a", "a"}, 1},
		{"", args{"a", "aaa"}, 3},
		{"", args{"abcdefgh", "hgfedcba"}, 64},
		{"",
			args{
				"sxdfcgdgxdfgdxxf",
				"xgdfsxgdxgfsgdfxgfdgfgdfgdgsdfxgfdxgdfxgdgdfgdfxgsdxgdfxgfgsxfgdfgsdfxgdfxgsgsfgxsgsdxgsdfxgsgsdfxgsdfx",
			},
			288,
		},
		{"", args{"bbaabcbcbc", "bbcccbabcc"}, 12},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := StringConstructing(tt.args.a, tt.args.s); got != tt.want {
				t.Errorf("StringConstructing() = %v, want %v", got, tt.want)
			}
		})
	}
}
