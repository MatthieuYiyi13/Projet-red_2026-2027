package projet

func Initgoblin() p*Character {
	goblin := &p*Character{
		name:          "Gobelin",
		pv:            30,
		pvmax:         30,
		experience:    10,
		AttaqueDegats: 5,
	}
	return goblin
}
