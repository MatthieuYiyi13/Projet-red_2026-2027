package projet

import "fmt"

var inventaire_taillemax int = 10
var sort_bdf bool = true
var sort_cc bool = true
var casque_equipe bool = true
var botte_equipe bool = true
var armureivoire_equipe bool = true
var armurecuir_equipe bool = true
var affiche_casque bool = false
var affiche_armurecuir bool = false
var affiche_armureivoire bool = false
var affiche_botte bool = false

const (
	RessourceTissuDeSpectre      = "Tissu de spectre"
	RessourcePeauDeGeant         = "Peau de Géant"
	RessourceFilDaraignee        = "Fil d'araignée"
	RessourceDentdeloup          = "Dent de loup"
	RessourcePoildemammouth      = "Poil de mammouth"
	RessourceGriffeDelynxfumee   = "Griffe de lynx fumée"
	RessourceCendreDelynxfumee   = "Cendre de lynx fumée"
	RessourceFragmentdeglace     = "Fragment de glace"
	RessourceFragmentRoiDeLaNuit = "Fragment du Roi de la nuit"
	RessourceSabotDeLicorne      = "Sabot de licorne"
	RessourceChapeaudegeant      = "Chapeau de Géant"
	RessourceArmurecuir          = "Armure en cuir"
	RessourceArmureIvoire        = "Armure en ivoire"
	RessourceBotteArcenciel      = "Bottes Arc-En-Ciel"
)

var objet_marchand = map[string]int{
	RessourceTissuDeSpectre:      100,
	RessourcePeauDeGeant:         150,
	RessourceFilDaraignee:        150,
	RessourceDentdeloup:          200,
	RessourcePoildemammouth:      250,
	RessourceGriffeDelynxfumee:   250,
	RessourceCendreDelynxfumee:   400,
	RessourceFragmentdeglace:     500,
	RessourceFragmentRoiDeLaNuit: 750,
	RessourceSabotDeLicorne:      999,
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
	fmt.Printf("PV de l'ennemie après poison = %d\n", *enemyPv)
}
func (p *Character) Equiper(objet string) {
	if p.inventaire[objet] < 1 {
		fmt.Println("Equipement non possedé.")
		return
	}
	switch objet {
	case "Bottes arc en ciel":
		p.Equipements["bottes"] = objet
		p.pvmax += 5
		fmt.Println(objet, "Vous avez equipé les bottes en arc en ciel !")

	case "Armure en ivoire":
		p.Equipements["armure"] = objet
		p.pvmax += 30
		fmt.Println(objet, "Vous avez equipé l'armuyre en ivoire !")

	case "Armure en cuir":
		p.Equipements["armure"] = objet
		p.pvmax += 15
		fmt.Println(objet, "Vous avez equipé l'armure en cuir !")

	case "Chapeau de géant":
		p.Equipements["chapeau"] = objet
		p.pvmax += 8
		fmt.Println(objet, "Vous avez equipé le chapeau de géant !")
	default:
		fmt.Println("Cet objet ne peut pas être équipé.")
		return
	}
}
