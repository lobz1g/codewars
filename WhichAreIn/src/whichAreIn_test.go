package src

import (
	"reflect"
	"testing"
)

func TestInArray(t *testing.T) {
	type args struct {
		array1 []string
		array2 []string
	}
	tests := []struct {
		name string
		args args
		want []string
	}{
		{
			"",
			args{[]string{"live", "arp", "strong"}, []string{"lively", "alive", "harp", "sharp", "armstrong"}},
			[]string{"arp", "live", "strong"},
		},
		{
			"",
			args{[]string{"arp", "mice", "bull"}, []string{"lively", "alive", "harp", "sharp", "armstrong"}},
			[]string{"arp"},
		},
		{
			"",
			args{[]string{"cod", "code", "wars", "ewar"}, []string{"lively", "alive", "harp", "sharp", "armstrong", "codewars"}},
			[]string{"cod", "code", "ewar", "wars"},
		},
		{
			"",
			args{[]string{"cod", "code", "wars", "ewar", "ar"}, []string{"lively", "alive", "harp", "sharp", "armstrong", "codewars"}},
			[]string{"ar", "cod", "code", "ewar", "wars"},
		},
		{
			"",
			args{[]string{"cod", "code", "wars", "ewar", "ar"}, []string{}},
			[]string{},
		},
		{
			"",
			args{[]string{"1295", "code", "1346", "1028", "ar"}, []string{"12951295", "ode", "46", "10281066", "par"}},
			[]string{"1028", "1295", "ar"},
		},
		{
			"",
			args{[]string{"&()", "code", "1346", "1028", "ar"}, []string{"12&()95", "coderange", "46", "1066", "par"}},
			[]string{"&()", "ar", "code"},
		},
		{
			"",
			args{[]string{"ohio", "code", "1346", "1028", "art"}, []string{"Carolina", "Ohio", "4600", "NY", "California"}},
			[]string{},
		},
		{
			"",
			args{[]string{"duplicates", "duplicates"}, []string{"duplicates", "duplicates"}},
			[]string{"duplicates"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := InArray(tt.args.array1, tt.args.array2); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("InArray() = %v, want %v", got, tt.want)
			}
		})
	}
}
