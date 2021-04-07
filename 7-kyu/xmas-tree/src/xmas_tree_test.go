package src

import (
	"fmt"
	"reflect"
	"testing"
)

func TestXMasTree(t *testing.T) {
	tests := []struct {
		height int
		want   []string
	}{{
		height: 0, want: []string{},
	},
		{
			height: 2,
			want: []string{
				"_#_",
				"###",
				"_#_",
				"_#_",
			},
		},
		{
			height: 3,
			want: []string{
				"__#__",
				"_###_",
				"#####",
				"__#__",
				"__#__",
			},
		},
		{
			height: 4,
			want: []string{
				"___#___",
				"__###__",
				"_#####_",
				"#######",
				"___#___",
				"___#___",
			},
		},
		{
			height: 5,
			want: []string{
				"____#____",
				"___###___",
				"__#####__",
				"_#######_",
				"#########",
				"____#____",
				"____#____",
			},
		},
		{
			height: 6,
			want: []string{
				"_____#_____",
				"____###____",
				"___#####___",
				"__#######__",
				"_#########_",
				"###########",
				"_____#_____",
				"_____#_____",
			},
		},
		{
			height: 7,
			want: []string{
				"______#______",
				"_____###_____",
				"____#####____",
				"___#######___",
				"__#########__",
				"_###########_",
				"#############",
				"______#______",
				"______#______",
			},
		},
	}
	for _, tt := range tests {
		t.Run(fmt.Sprintf("%d", tt.height), func(t *testing.T) {
			if got := XMasTree(tt.height); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("XMasTree() = %v, want %v", got, tt.want)
			}
		})
	}
}
