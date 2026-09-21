package projet

import (
	"fmt"
)
type Skill struct {
	name string
	damage int
	ultime bool
}
type Character struct {
	name         string
	classe       string
	pvmax        int
	pv           int
	inventaire   map[string]int
	Resurrection bool
	Skill 		 []Skill
}
func (p *Character) initcharacter(name string, classe string) {
	p.name = name
	p.classe = classe
	switch classe {
	case "guerrier":
		p.pvmax = 200
		p.pv = p.pvmax / 2
	case "sorcier":	
	    p.pvmax = 100
		p.pv = p.pvmax / 2
	}
}

func (p *Character) displayinfo () {
}

func (p *Character) AccessInventory() {
	fmt.Println("=== Informations du personnage ===")
	for itemname, itemquantity := range p.inventaire {
		fmt.Printf("\t %s : %d\n", itemname, itemquantity)
	}
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
	c.inventaire["potion"]--

	fmt.Printf("Vous avez utilisé une potion. Votre vie est maintenant de %d/%d.\n", c.pv, c.pvmax)
}
