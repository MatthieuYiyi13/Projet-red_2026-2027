package projet

import (
	"fmt"
	"math/rand/v2"
)

var skiptuto = false
var skip bool = false
var monstreinitbool bool = true
var info_debut_combat bool = true
var zone_unlock int = 1
var count_super_monstre int = 1
var victoire bool = false
var combat_pas_finis bool = true
var fuite bool = false

type Spectre struct {
	name        string
	attack      int
	PV_max      int
	PV_actuelle int
	Poison      bool
	Initiative  int
}
type Loup struct {
	name        string
	attack      int
	PV_max      int
	PV_actuelle int
	Poison      bool
	Initiative  int
}
type Géant struct {
	name        string
	attack      int
	PV_max      int
	PV_actuelle int
	Poison      bool
	Initiative  int
}
type Lynx_fumee struct {
	name        string
	attack      int
	PV_max      int
	PV_actuelle int
	Poison      bool
	Initiative  int
}
type Mammouth struct {
	name        string
	attack      int
	PV_max      int
	PV_actuelle int
	Poison      bool
	Initiative  int
}
type Araignee struct {
	name        string
	attack      int
	PV_max      int
	PV_actuelle int
	Poison      bool
	Initiative  int
}
type Roi_De_la_nuit struct {
	name        string
	attack      int
	PV_max      int
	PV_actuelle int
	Poison      bool
	Initiative  int
}
type Licorne struct {
	name        string
	attack      int
	PV_max      int
	PV_actuelle int
	Poison      bool
	Initiative  int
}
type White_walker struct {
	name        string
	attack      int
	PV_max      int
	PV_actuelle int
	Poison      bool
	Initiative  int
}
type Monstre struct {
	name        string
	attack      int
	PV_max      int
	PV_actuelle int
	Poison      bool
	Initiative  int
}


func (p *Character) choix_monstre() Monstre {

	n := 0

	if zone_unlock == 1 {
		n = rand.IntN(3)
	}
	if zone_unlock == 2 {
		n = rand.IntN(5)
	}
	if zone_unlock == 3 {
		n = rand.IntN(7)
	}
	if zone_unlock == 4 {
		n = rand.IntN(9)
	}

	if n == 0 {
		return Monstre{
			name:        "Spectre",
			attack:      10,
			PV_max:      75,
			PV_actuelle: 75,
			Poison:      false,
			Initiative:  2,
		}
	}

	if n == 1 {
		return Monstre{
			name:        "Géant",
			attack:      8,
			PV_max:      150,
			PV_actuelle: 150,
			Poison:      false,
			Initiative:  7,
		}
	}

	if n == 2 {
		return Monstre{
			name:        "Araignée",
			attack:      20,
			PV_max:      60,
			PV_actuelle: 60,
			Poison:      false,
			Initiative:  1,
		}
	}

	if n == 3 {
		return Monstre{
			name:        "Lynx fumée",
			attack:      25,
			PV_max:      200,
			PV_actuelle: 200,
			Poison:      false,
			Initiative:  2,
		}
	}

	if n == 4 {
		return Monstre{
			name:        "Loup",
			attack:      28,
			PV_max:      125,
			PV_actuelle: 125,
			Poison:      false,
			Initiative:  3,
		}
	}

	if n == 5 {
		return Monstre{
			name:        "Mammouth",
			attack:      12,
			PV_max:      350,
			PV_actuelle: 350,
			Poison:      false,
			Initiative:  7,
		}
	}

	if n == 6 {
		return Monstre{
			name:        "White walker",
			attack:      38,
			PV_max:      375,
			PV_actuelle: 375,
			Poison:      false,
			Initiative:  5,
		}
	}

	if n == 7 {
		return Monstre{
			name:        "Roi de la nuit",
			attack:      50,
			PV_max:      500,
			PV_actuelle: 500,
			Poison:      false,
			Initiative:  9,
		}
	}

	return Monstre{
		name:        "Licorne",
		attack:      125,
		PV_max:      1500,
		PV_actuelle: 1500,
		Poison:      false,
		Initiative:  10,
	}
}

func (p *Character) Loot(monstre Monstre){
	if p.InventairePlein() {
		afficherTexte1("Votre inventaire est plein ! Vous ne pouvez pas récuperer d'objets.")
		return
	}	
	if monstre.name == "Spectre" {
		n := rand.IntN(2)	
		if n == 1 {
		p.inventaire[RessourceTissuDeSpectre] += 1
		afficherTexte50("Vous récuperez 1 %s sur le corps du monstre",RessourceTissuDeSpectre)
	}
}
	if monstre.name == "Géant" {
		n := rand.IntN(2)
		if n == 1 {
		p.inventaire[RessourcePeauDeGeant] += 1
		afficherTexte50("Vous récuperez 1 %s sur le corps du monstre",RessourcePeauDeGeant)
		}
	}
	if monstre.name == "Araignée" {
	n := rand.IntN(2)
		if n == 1 {
		p.inventaire[RessourceFilDaraignee] += 1
		afficherTexte50("Vous récuperez 1 %s sur le corps du monstre",RessourceFilDaraignee)
		}
	}
	if monstre.name == "Lynx fumée" {
	n := rand.IntN(2)
		if n == 1 {
		p.inventaire[RessourceGriffeDelynxfumee] += 1
		afficherTexte50("Vous récuperez 1 %s sur le corps du monstre",RessourceGriffeDelynxfumee)
	} else {
		p.inventaire[RessourceCendreDelynxfumee] += 1
		afficherTexte50("Vous récuperez 1 %s sur le corps du monstre",RessourceCendreDelynxfumee)
	}
}
	if monstre.name == "Loup" {
	n := rand.IntN(2)
		if n == 1 {
		p.inventaire[RessourceDentdeloup] += 1
		afficherTexte50("Vous récuperez 1 %s sur le corps du monstre",RessourceDentdeloup)
		}
	}
	if monstre.name == "Mammouth" {
	n := rand.IntN(2)
		if n == 1 {
		p.inventaire[RessourcePoildemammouth] += 1
		afficherTexte50("Vous récuperez 1 %s sur le corps du monstre",RessourcePoildemammouth)
		}
	}
	if monstre.name == "White walker" {
			n := rand.IntN(2)
		if n == 1 {
		p.inventaire[RessourceFragmentdeglace] += 1
		afficherTexte50("Vous récuperez 1 %s sur le corps du monstre",RessourceFragmentdeglace)
		}
	}
	if monstre.name == "Roi de la nuit" {
	n := rand.IntN(2)
		if n == 1 {
		p.inventaire[RessourceFragmentRoiDeLaNuit] += 1
		afficherTexte50("Vous récuperez 1 %s sur le corps du monstre",RessourceFragmentRoiDeLaNuit)
		}
	}
	if monstre.name == "Licorne" {
	n := rand.IntN(2)
		if n == 1 {
		p.inventaire[RessourceSabotDeLicorne] += 1
		afficherTexte50("Vous récuperez 1 %s sur le corps du monstre",RessourceSabotDeLicorne)
		}
	}
	g := rand.IntN(150)
	p.money += g
	afficherTexte50("Vous gagnez %d écus",g)
}
	

func (p *Character) Usepot_poison(monstre *Monstre) {
	monstre.Poison = true
	afficherTexte50("%s est empoisonné !\n", monstre.name)
}

func (p *Character) Objet_utilitaire(monstre *Monstre) {

	fmt.Printf("\t+------------------------------------------------+\n")
	fmt.Printf("\t|                                                |\n")
	fmt.Printf("\t|            Que voulez vous faire ?             |\n")
	fmt.Printf("\t|                                                |\n")
	if p.inventaire[RessourcePotSoin] >= 1 {
	fmt.Printf("\t|  [S] Utiliser une potion de soin               |\n")
	}
	if p.inventaire[RessourcePotMana] >= 1 {
	fmt.Printf("\t|  [M] Utiliser une potion de mana               |\n")
	}
	fmt.Printf("\t|                                                |\n")
	if p.inventaire[RessourcePotPoison] >= 1 {
	fmt.Printf("\t|  [P] Utiliser une potion de poison             |\n")
	}
	fmt.Printf("\t|                                                |\n")
	fmt.Printf("\t|  [R] Retour                                    |\n")
	fmt.Printf("\t|                                                |\n")
	fmt.Printf("\t+------------------------------------------------+\n")
	var moove_objet string
	fmt.Scanln(&moove_objet)
	switch moove_objet {
	case "R", "r":
		skip = true
	case "S", "s":
		if p.inventaire[RessourcePotSoin] >= 1 {
			p.takepotS()
		} else {
			fmt.Println()
			afficherTexte50("Vous n'avez pas de potion de soin à utiliser")
			fmt.Println()
			skip = true
		}
	case "M", "m":
		if p.inventaire[RessourcePotMana] >= 1 {
			p.takepotM()
		} else {
			fmt.Println()
			afficherTexte50("Vous n'avez pas de potion de mana à utiliser")
			fmt.Println()
			skip = true
		}
	case "P", "p":
		if p.inventaire[RessourcePotPoison] >= 1 {
			p.Usepot_poison(monstre)
		} else {
			fmt.Println()
			afficherTexte50("Vous n'avez de potion de poison à utiliser")
			fmt.Println()
			skip = true
		}
	default:
		afficherTexte1("Commande invalide.")
	}
}

func (p *Character) Fuir() {
	k := rand.IntN(3)
	if (k == 0)  || (k ==1) {
		combat_pas_finis = false
		victoire = true
		fuite = true
		fmt.Println()
		afficherTexte20("Vous avez réussi à fuire le combat !")
		fmt.Println()
	} else {
		fmt.Println()
		afficherTexte20("Vous n'avez pas réussi à fuire le combat !")
		fmt.Println()
		fmt.Println()
	}
}
func (p *Character) Attaque(monstre *Monstre) {
	fmt.Printf("\n")
	fmt.Printf("\t+------------------------------------------------+\n")
	fmt.Printf("\t|           Quelle attaque utiliser ?            |\n")
	fmt.Printf("\t|                                                |\n")
	fmt.Printf("\t|                                                |\n")
	if p.AttaqueName == "Coup d'épée" {
	fmt.Printf("\t|     [C] Coup d'épée          [R] Retour        |\n")
	}
	if p.AttaqueName == "Coup de baton" {
	fmt.Printf("\t|     [C] Coup de baton        [R] Retour        |\n")
	}
	if p.AttaqueName == "Coup de dague" {
	fmt.Printf("\t|     [C] Coup de dague       [R] Retour         |\n")
	}
	fmt.Printf("\t|                                                |\n")
	fmt.Printf("\t|                                                |\n")
	if (p.SkillName == "Coups vicieux") && (Sort) {
	fmt.Printf("\t|          [S] Coups vicieux  (35 Mana)          |\n")
	}
	if (p.SkillName == "Boule de feu") && (Sort) {
	fmt.Printf("\t|          [S] Boule de feu   (50 Mana)          |\n")
	}
	if (p.SkillName == "Coup critique") && (Sort) {
	fmt.Printf("\t|          [S] Coup critique  (25 Mana)          |\n")
	}
	fmt.Printf("\t|                                                |\n")
	fmt.Printf("\t|                                                |\n")
	fmt.Printf("\t+------------------------------------------------+\n")
	fmt.Println()
	var moove_attack string
	fmt.Scanln(&moove_attack)
	switch moove_attack {
	case "R", "r":
		skip = true
	case "C", "c":
		{
			afficherTexte50("Vous attaquez le monstre il perd %d pvs",p.AttaqueDegats)

			monstre.PV_actuelle -= p.AttaqueDegats
		}
	case "S", "s":
		if !Sort {
			afficherTexte20("Vous ne pouvez pas utiliser de sort.")
			return
		}

		if p.Mana < p.Skillmana {
			afficherTexte20("Vous n'avez pas assez de mana pour utiliser ce sort.")
			return
		}

		monstre.PV_actuelle -= p.SkillDegats
		p.Mana -= p.Skillmana

		afficherTexte50(
			"Vous utilisez %s et infligez %d dégâts !\n",
			p.SkillName,
			p.SkillDegats,
		)
		afficherTexte50("Mana actuelle %d\n", p.Mana)
	default:
		afficherTexte1("Commande invalide.")
		skip = true
	}
}

func (p *Character) AttaqueTuto(monstre *Monstre) {
	fmt.Printf("\n")
	fmt.Printf("\t\t+------------------------------------------------+\n")
	fmt.Printf("\t\t|           Quelle attaque utiliser ?            |\n")
	fmt.Printf("\t\t|                                                |\n")
	fmt.Printf("\t\t|                                                |\n")
	if p.AttaqueName == "Coup d'épée" {
	fmt.Printf("\t\t|     [C] Coup d'épée                            |\n")
	}
	if p.AttaqueName == "Coup de baton" {
	fmt.Printf("\t\t|     [C] Coup de baton                          |\n")
	}
	if p.AttaqueName == "Coup de dague" {
	fmt.Printf("\t\t|     [C] Coup de dague                          |\n")
	}
	fmt.Printf("\t\t|                                                |\n")
	fmt.Printf("\t\t|                                                |\n")
	if (p.SkillName == "Coups vicieux") && (Sort) {
		fmt.Printf("\t\t|               [S] Coups vicieux  (25 Mana) |\n")
	}
	if (p.SkillName == "Boule de feu") && (Sort) {
		fmt.Printf("\t\t|               [S] Boule de feu   (50 Mana) |\n")
	}
	if (p.SkillName == "Coup critique") && (Sort) {
		fmt.Printf("\t\t|               [S] Coup critique  (35 Mana) |\n")
	}
	fmt.Printf("\t\t|                                                |\n")
	fmt.Printf("\t\t|                                                |\n")
	fmt.Printf("\t\t+------------------------------------------------+\n")
	fmt.Println()
	var moove_attack string
	fmt.Scanln(&moove_attack)
	switch moove_attack {
	case "C", "c":
		{
			afficherTexte50("Vous attaquez le monstre il perd %d pvs",p.AttaqueDegats)
			monstre.PV_actuelle -= p.AttaqueDegats
		}
	default:
		afficherTexte50("Commande invalide.")
		skiptuto = true
	}
}


func (p *Character) Combat_start_premier() {
	stopSound()
	PlaySoundAsyncCombat1()
	fuite = false
	victoire = false
	combat_pas_finis = true
	monstre := p.choix_monstre()
	fmt.Println()
	afficherTexte50("Le monstre sauvage %s vous attaque !", monstre.name)
	fmt.Println()
	info_debut_combat = false

	if p.Initiative < monstre.Initiative {
		fmt.Println()
		fmt.Println("Le monstre prends l'initiative !")
		l := rand.IntN(4)
		if l == 0 {
			afficherTexte20("Le monstre vous attaque mais vous réussissez à l'esquiver !")
		} else {
			afficherTexte20("Le monstre ennemi est énervé, il vous charge")
			pv_perdu := monstre.attack
			if count_super_monstre%4 == 0 {
				pv_perdu = pv_perdu * 2
			}
			afficherTexte50("Vous perdez %d Pvs\n", pv_perdu)
			p.pv -= monstre.attack
			afficherTexte50("Pvs actuelle : %d/%d", p.pv, p.pvmax)
			fmt.Println()
			count_super_monstre++
		}
		if p.pv <= 0 {
			victoire = false
			combat_pas_finis = false
			afficherTexte50("Vous avez perdu tout vos pvs , vous êtes mort !")
			fmt.Println()
			p.IsDead()
			return
		}
	}
	for combat_pas_finis {
		if monstre.Poison && skip == false{
			fmt.Println()
			afficherTexte50("Le monstre ennemi %s perd %d PV", monstre.name, 10)
			fmt.Println()
			monstre.PV_actuelle -= 10
			fmt.Println()
			if monstre.PV_actuelle <= 0 {
				combat_pas_finis = false
				victoire = true
				continue
			}
		}
		skip = false
		fmt.Println()
		afficherTexte50("Les pv du monstre %s sont de %d/%d ", monstre.name, monstre.PV_actuelle, monstre.PV_max)
		fmt.Println()
		fmt.Printf("\n")
		fmt.Printf("\t+------------------------------------------------+\n")
		fmt.Printf("\t|                                                |\n")
		fmt.Printf("\t|           Que voulez vous faire ?              |\n")
		fmt.Printf("\t|                                                |\n")
		fmt.Printf("\t|    [O] Objet       [F] Fuir       [A] Attaque  |\n")
		fmt.Printf("\t|                                                |\n")
		fmt.Printf("\t+------------------------------------------------+\n")
		fmt.Println()
		var moove string
		fmt.Scanln(&moove)
		switch moove {
		case "O", "o":
			p.Objet_utilitaire(&monstre)
			if skip == true {
				continue
			}
		case "F", "f":
			p.Fuir()
			if combat_pas_finis == false {
				continue
			}
		case "A", "a":
			p.Attaque(&monstre)
			if skip == true {
				continue
			}
		default:
			skip = true
			continue
		}
		if monstre.PV_actuelle <= 0 {
			combat_pas_finis = false
			victoire = true
			continue
		}
		fmt.Println()
		l := rand.IntN(4)
		if l == 0 {
			fmt.Println()
			afficherTexte50("Le monstre vous attaque mais vous réussissez à l'esquiver !")
			fmt.Println()
		} else {
			afficherTexte50("Le monstre ennemi est enervé , il vous charge")
			pv_perdu := monstre.attack
			if count_super_monstre%4 == 0 {
				pv_perdu = pv_perdu * 2
			}
			afficherTexte50("Vous perdez %d Pvs\n", pv_perdu)
			p.pv -= monstre.attack
			afficherTexte50("Pvs actuelle : %d/%d", p.pv, p.pvmax)
			fmt.Println()
			count_super_monstre++
		}
		if p.pv <= 0 {
			victoire = false
			combat_pas_finis = false
			continue
		}
	}
	if victoire && fuite == false {
		fmt.Println()
		afficherTexte50("Vous avez gagné le combat")
		p.Loot(monstre)
		fmt.Println()
		p.GagnerCombat()
		victoire = false
		stopSound()
		combat_pas_finis = true
	} else if victoire && fuite {
		fmt.Println()
		fmt.Println()
		stopSound()
	} else {
		fmt.Println()
		afficherTexte50("Vous avez perdu tout vos pvs , vous êtes mort !")
		fmt.Println()
		victoire = false
		combat_pas_finis = true
		stopSound()
		p.IsDead()
	}
}

