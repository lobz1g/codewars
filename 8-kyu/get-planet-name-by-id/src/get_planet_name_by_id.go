package src

func GetPlanetName(ID int) string {
	var planet string
	switch ID {
	case 1:
		planet = "Mercury"
	case 2:
		planet = "Venus"
	case 3:
		planet = "Earth"
	case 4:
		planet = "Mars"
	case 5:
		planet = "Jupiter"
	case 6:
		planet = "Saturn"
	case 7:
		planet = "Uranus"
	case 8:
		planet = "Neptune"
	default:
		planet = "Pluto"
	}

	return planet
}
