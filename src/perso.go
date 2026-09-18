package main

import (
	"fmt"
)
type character struct {
	name string
	classe string
	pvmax int
	pv int
	inventaire map[string]int
}
func (c*character) initcharacter(name string, classe string) {
	c.name = name
	c.classe = classe
	switch classe {
	case "guerrier":
		c.pvmax = 200
		c.pv = c:pvmax / 2
	case "sorcier":	
	    c.pvmax = 100
		c.pv = c.pvmax / 2
	}
}

func (c.character) displayinfo () {
}

func (c character) accessIventory() {
	fmt.Println("=== Informations du personnage ===")
	for itemname, itemquantity := range c.inventaire {
		fmt.Printf("\t %s : %d\n", itemname, itemquantity)
	}
}

func (c *character) takepot() {
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

func main() {
	var player character
	player.initcharacter("Davy", "guerrier")
	for true {
		fmt.Println("=== Menu ===")
		fmt.Println("1. Afficher les informations du personnage")
		fmt.Println("2. Accéder à l'inventaire")
		fmt.Println("3. Quitter le jeu")
		var choice int
		fmt.Scanln(&choice)}

		switch choice {
			case 0:
}