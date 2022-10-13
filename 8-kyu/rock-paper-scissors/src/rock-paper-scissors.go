package src

func Rps(p1, p2 string) string {
	if p1 == p2 {
		return "Draw!"
	}

	rules := map[string]map[string]string{
		"rock":     {"paper": "2", "scissors": "1"},
		"paper":    {"scissors": "2", "rock": "1"},
		"scissors": {"rock": "2", "paper": "1"},
	}

	return "Player " + rules[p1][p2] + " won!"
}
