package src

import "testing"

func TestMxDifLg(t *testing.T) {
	type args struct {
		a1 []string
		a2 []string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			"",
			args{
				[]string{"hoqq", "bbllkw", "oox", "ejjuyyy", "plmiis", "xxxzgpsssa", "xxwwkktt", "znnnnfqknaz", "qqquuhii", "dvvvwz"},
				[]string{"cccooommaaqqoxii", "gggqaffhhh", "tttoowwwmmww"},
			},
			13,
		},
		{
			"",
			args{
				[]string{"ejjjjmmtthh", "zxxuueeg", "aanlljrrrxx", "dqqqaaabbb", "oocccffuucccjjjkkkjyyyeehh"},
				[]string{"bbbaaayddqbbrrrv"},
			},
			10,
		},
		{
			"",
			args{
				[]string{"ccct", "tkkeeeyy", "ggiikffsszzoo", "nnngssddu", "rrllccqqqqwuuurdd", "kkbbddaakkk"},
				[]string{"tttxxxxxxgiiyyy", "ooorcvvj", "yzzzhhhfffaaavvvpp", "jjvvvqqllgaaannn", "tttooo", "qmmzzbhhbb"},
			},
			14,
		},
		{
			"",
			args{
				[]string{},
				[]string{"cccooommaaqqoxii", "gggqaffhhh", "tttoowwwmmww"},
			},
			-1,
		},
		{
			"",
			args{
				[]string{},
				[]string{},
			},
			-1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := MxDifLg(tt.args.a1, tt.args.a2); got != tt.want {
				t.Errorf("MxDifLg() = %v, want %v", got, tt.want)
			}
		})
	}
}
