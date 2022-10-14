package src

import (
	"math"
)

type Fighter struct {
	Name            string
	Health          int
	DamagePerAttack int
}

func DeclareWinner(fighter1 Fighter, fighter2 Fighter, firstAttacker string) string {
	f1 := math.Ceil(float64(fighter1.Health) / float64(fighter2.DamagePerAttack))
	f2 := math.Ceil(float64(fighter2.Health) / float64(fighter1.DamagePerAttack))

	if f1 > f2 {
		return fighter1.Name
	} else if f1 == f2 {
		return firstAttacker
	}
	return fighter2.Name
}
