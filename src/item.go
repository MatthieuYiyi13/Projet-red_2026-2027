package projet

import "fmt"

var inventaire_taillemax int = 10
var objet_marchand = map[string]int{
   "Tissu de spectre":            100,
	"Peau de géant":              150,
	"Fil d'araignée":             150, 
	"Dent de loup":               200,
	"Poil de mammouth":           250,
	"Griffe de lynx de fumée":    250,
	"Cendre de lynx de fumée":    400,
	"Fragment de glace":          500,
	"Fragment du Roi de la Nuit": 750,
	"Sabot de licorne":           999,
}

func (p *Character) takepotS() {
	potquantity, potcheck := p.inventaire["potion de vie"]
	if !potcheck {
		fmt.Println("Vous n'avez pas de potion de vie dans votre inventaire.")
		return
	}
	if potquantity <= 0 {
		fmt.Println("Vous n'avez plus de potion de vie dans votre inventaire.")
		return
	}
    p.pv += 50
	if p.pv > p.pvmax {
		p.pv = p.pvmax
	}
	p.inventaire["potion de vie"]--

	fmt.Printf("Vous avez utilisé une potion. Votre vie est maintenant de %d/%d.\n", p.pv, p.pvmax)
}
func (p *Character) takepotP(enemyPv *int) {
	 potquantity, potcheck := p.inventaire["potion de poison"]
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