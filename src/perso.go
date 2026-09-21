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

func (c *Character) initcharacter(name string, classe string) {
	c.name = name
	c.classe = classe
	switch classe {
	case "guerrier":
		c.pvmax = 200
		c.pv = c.pvmax / 2
	case "sorcier":
		c.pvmax = 100
		c.pv = c.pvmax / 2
	}
}

func (c Character) displayinfo() {
}

func (c Character) accessIventory() {
	fmt.Println("=== Informations du personnage ===")
	for itemname, itemquantity := range c.inventaire {
		fmt.Printf("\t %s : %d\n", itemname, itemquantity)
	}
}
