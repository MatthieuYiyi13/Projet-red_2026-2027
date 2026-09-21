package projet

import (
	"fmt"
)
type Character struct {
	name         string
	money 		 int
	classe       string
	pvmax        int
	pv           int
	experience   int
	inventaire   map[string]int
	Resurrection bool
	Skill 		 []Skill
	AttaqueName string
	AttaqueDegats int
	SkillName string
	SkillDegats int
}
func (p *Character) initcharacter(name string, classe string) {
	p.name = name
	p.money = 1000
	p.inventaire = make(map[string]int)
	p.classe = classe
	p.experience = 0
	switch classe {
	case "guerrier":
		p.pvmax = 200
		p.pv = p.pvmax / 2
		p.AttaqueName = "Coup d'épée"
		p.AttaqueDegats = 15
		p.SkillName = "Coup critique"
		p.SkillDegats = 35
	case "sorcier":	
	    p.pvmax = 100
		p.pv = p.pvmax / 2
		p.AttaqueName = "Coup de baton"
		p.AttaqueDegats = 10
		p.SkillName = "Boule de feu"
		p.SkillDegats = 50
	}
}

func (p *Character) displayinfo () {
}

func (p *Character) AccessInventory() {
	for itemname, itemquantity := range p.inventaire {
		fmt.Println()
		fmt.Printf("\t %s : %d\n", itemname, itemquantity)
		fmt.Println()
	}
	fmt.Printf("\t Argent : %d\n", p.money)
	fmt.Println()
}

func (p *Character) takepot() {
	potquantity, potcheck := p.inventaire["potion"]
	if !potcheck {
		fmt.Println("Vous n'avez pas de potion dans votre inventaire.")
		return
	}
	if potquantity <= 0 {
		fmt.Println("Vous n'avez plus de potion dans votre inventaire.")
		return
	}
    p.pv += 50
	if p.pv > p.pvmax {
		p.pv = p.pvmax
	}
	p.inventaire["potion"]--

	fmt.Printf("Vous avez utilisé une potion. Votre vie est maintenant de %d/%d.\n", p.pv, p.pvmax)
}
