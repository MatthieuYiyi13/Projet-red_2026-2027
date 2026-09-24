package projet

import "fmt"

var inventaire_taillemax int = 10
var sort_coupv bool = true
var sort_bdf bool = true
var sort_cc bool = true
var casque_equipe bool = true
var botte_equipe bool = true
var armureivoire_equipe bool = true
var armurecuir_equipe bool = true

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
	RessourcePotSoin 			 = "Potion de soin"
	RessourcePotMana			 = "Potion de mana"
	RessourcePotPoison			 = "Potion de poison"
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
	potquantity, potcheck := p.inventaire[RessourcePotSoin]
	if !potcheck {
		fmt.Println("Vous n'avez pas de potion de soin dans votre inventaire.")
		return
	}
	if potquantity <= 0 {
		fmt.Println("Vous n'avez plus de potion de soin dans votre inventaire.")
		return
	}
	p.pv += 50
	if p.pv > p.pvmax {
		p.pv = p.pvmax
	}
	p.inventaire[RessourcePotSoin]--

	afficherTexte50("Vous avez utilisé une potion de soin. Votre vie est maintenant de %d/%d.\n", p.pv, p.pvmax)
}

func (p *Character) takepotM() {
	potquantity, potcheck := p.inventaire[RessourcePotMana]
	if !potcheck {
		afficherTexte20("Vous n'avez pas de potion de mana dans votre inventaire.")
		return
	}
	if potquantity <= 0 {
		afficherTexte20("Vous n'avez plus de potion de mana dans votre inventaire.")
		return
	}
	p.Mana += 50
	if p.Mana > p.Manamax {
		p.Mana = p.Manamax
	}
	p.inventaire[RessourcePotMana]--

	afficherTexte50("Vous avez utilisé une potion de mana. Votre mana est maintenant de %s/%d.\n", p.Mana, p.Manamax)
}