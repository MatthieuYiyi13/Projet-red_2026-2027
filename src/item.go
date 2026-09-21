package projet

import "fmt"

func (c *Character) takepotS() {
	potquantity, potcheck := c.inventaire["potion de vie"]
	if !potcheck {
		fmt.Println("Vous n'avez pas de potion de vie dans votre inventaire.")
		return
	}
	if potquantity <= 0 {
		fmt.Println("Vous n'avez plus de potion de vie dans votre inventaire.")
		return
	}
    c.pv += 50
	if c.pv > c.pvmax {
		c.pv = c.pvmax
	}
	c.inventaire["potion de vie"]--
	// quoi faire quand quantite = 0 

	fmt.Printf("Vous avez utilisé une potion. Votre vie est maintenant de %d/%d.\n", c.pv, c.pvmax)
}
func (c *Character) takepotP(enemyPv *int) {
	 potquantity, potcheck := c.inventaire["potion de poison"]
	 if !potcheck {
		fmt.Println("Vous n'avez pas de potion de poison dans votre inventaire")
		return
	 }
	 if potquantity <= 0 {
		fmt.Println("Vous n'avez plus de potion de poison dans votre inventaire")
		return
	 }
	 *enemyPv -= 10
	 if *enemyPv < 0 {
		*enemyPv = 0
	 }
	 fmt.Println("Potion de poison utilisée (-1)")
	 fmt.Printf("PV de l'ennemie apres poison = %d\n", *enemyPv)
}