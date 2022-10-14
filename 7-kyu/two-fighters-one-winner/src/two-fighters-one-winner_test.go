package src

import (
	"strconv"
	"testing"
)

func TestDeclareWinner(t *testing.T) {
	tests := []struct {
		fighter1      Fighter
		fighter2      Fighter
		firstAttacker string
		want          string
	}{
		{
			Fighter{"Lew", 10, 2},
			Fighter{"Harry", 5, 4},
			"Lew",
			"Lew",
		},
		{
			Fighter{"Lew", 10, 2},
			Fighter{"Harry", 5, 4},
			"Harry",
			"Harry",
		},
		{
			Fighter{"Harald", 20, 5},
			Fighter{"Harry", 5, 4},
			"Harry",
			"Harald",
		},
		{
			Fighter{"Harald", 20, 5},
			Fighter{"Harry", 5, 4},
			"Harald",
			"Harald",
		},
		{
			Fighter{"Jerry", 30, 3},
			Fighter{"Harald", 20, 5},
			"Jerry",
			"Harald",
		},
		{
			Fighter{"Jerry", 30, 3},
			Fighter{"Harald", 20, 5},
			"Harald",
			"Harald",
		},
	}
	for i, tt := range tests {
		t.Run(strconv.Itoa(i), func(t *testing.T) {
			if got := DeclareWinner(tt.fighter1, tt.fighter2, tt.firstAttacker); got != tt.want {
				t.Errorf("DeclareWinner() = %v, want %v", got, tt.want)
			}
		})
	}
}
