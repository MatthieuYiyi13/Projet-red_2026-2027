package projet

type Gobelin struct {
	name        string
	attack      int
	PV_max      int
	PV_actuelle int
}

func (p*Character) InitGobelin() Monstre {
	return Monstre{
			name:        "Gobelin",
			attack:      5,
			PV_max:      40,
			PV_actuelle: 40,
		}
}
