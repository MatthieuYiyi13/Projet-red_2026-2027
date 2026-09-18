package projet

import (
	"fmt"
)
type Character struct {
	name string
	classe string
	pvmax int
	pv int
	inventaire map[string]int
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

func (c Character) displayinfo () {
}

func (c Character) accessIventory() {
	fmt.Println("=== Informations du personnage ===")
	for itemname, itemquantity := range c.inventaire {
		fmt.Printf("\t %s : %d\n", itemname, itemquantity)
	}
}

func (c *Character) takepot() {
	potquantity, potcheck := c.inventaire["potion"]
	if !potcheck {
		fmt.Println("Vous n'avez pas de potion dans votre inventaire.")
		return
	}
	if potquantity <= 0 {
		fmt.Println("Vous n'avez plus de potion dans votre inventaire.")
		return
	}
    c.pv += 50
	if c.pv > c.pvmax {
		c.pv = c.pvmax
	}
	c.inventaire["potion"]--
	// quoi faire quand quentite = 0 

	fmt.Printf("Vous avez utilisé une potion. Votre vie est maintenant de %d/%d.\n", c.pv, c.pvmax)
}
