package projet

import (
	"fmt"
	"math/rand/v2"
)

var zone_unlock int = 1
var count_super_monstre int = 1
var victoire bool = false
var combat_pas_finis bool = true

type Spectre struct {
	name        string
	attack      int
	PV_max      int
	PV_actuelle int
	Poison      bool
}
type Loup struct {
	name        string
	attack      int
	PV_max      int
	PV_actuelle int
	Poison      bool
}
type Géant struct {
	name        string
	attack      int
	PV_max      int
	PV_actuelle int
	Poison      bool
}
type Lynx_fumee struct {
	name        string
	attack      int
	PV_max      int
	PV_actuelle int
	Poison      bool
}
type Mammouth struct {
	name        string
	attack      int
	PV_max      int
	PV_actuelle int
	Poison      bool
}
type Araignee struct {
	name        string
	attack      int
	PV_max      int
	PV_actuelle int
	Poison      bool
}
type Roi_De_la_nuit struct {
	name        string
	attack      int
	PV_max      int
	PV_actuelle int
	Poison      bool
}
type Licorne struct {
	name        string
	attack      int
	PV_max      int
	PV_actuelle int
	Poison      bool
}
type White_walker struct {
	name        string
	attack      int
	PV_max      int
	PV_actuelle int
	Poison      bool
}
type Monstre struct {
	name        string
	attack      int
	PV_max      int
	PV_actuelle int
	Poison      bool
}

func (p *Character) choix_monstre() Monstre {

	n:= 0

	if zone_unlock == 1 {
	n = rand.IntN(3)}
	if zone_unlock == 2{
	n = rand.IntN(5)}
	if zone_unlock == 3{
	n = rand.IntN(7)}
	if zone_unlock == 4{
	n = rand.IntN(9)}

	if n == 0 {
		return Monstre{
			name:        "Spectre",
			attack:      10,
			PV_max:      75,
			PV_actuelle: 75,
			Poison:      false,
		}
	}

	if n == 1 {
		return Monstre{
			name:        "Géant",
			attack:      8,
			PV_max:      150,
			PV_actuelle: 150,
			Poison:      false,
		}
	}

	if n == 2 {
		return Monstre{
			name:        "Araignée",
			attack:      20,
			PV_max:      60,
			PV_actuelle: 60,
			Poison:      false,
		}
	}

	if n == 3 {
		return Monstre{
			name:        "Lynx fumée",
			attack:      25,
			PV_max:      200,
			PV_actuelle: 200,
			Poison:      false,
		}
	}

	if n == 4 {
		return Monstre{
			name:        "Loup",
			attack:      28,
			PV_max:      125,
			PV_actuelle: 125,
			Poison:      false,
		}
	}

	if n == 5 {
		return Monstre{
			name:        "Mammouth",
			attack:      12,
			PV_max:      350,
			PV_actuelle: 350,
			Poison:      false,
		}
	}

	if n == 6 {
		return Monstre{
			name:        "White walker",
			attack:      38,
			PV_max:      375,
			PV_actuelle: 375,
			Poison:      false,
		}
	}

	if n == 7 {
		return Monstre{
			name:        "Roi de la nuit",
			attack:      50,
			PV_max:      500,
			PV_actuelle: 500,
			Poison:      false,
		}
	}

	return Monstre{
		name:        "Licorne",
		attack:      125,
		PV_max:      1500,
		PV_actuelle: 1500,
		Poison:      false,
	}
}

func (p *Character) Usepot_poison(monstre *Monstre) {
	monstre.Poison = true
	fmt.Printf("%s est empoisonné !\n", monstre.name)
}

func (p *Character) Objet_utilitaire(monstre *Monstre) {

	fmt.Printf("\t+------------------------------------------------+\n")
	fmt.Printf("\t|                                                |\n")
	fmt.Printf("\t|            Que voulez vous faire ?             |\n")
	fmt.Printf("\t|                                                |\n")
	if p.inventaire[RessourcePotSoin] >= 1 {
		fmt.Printf("\t|  [S] Utiliser une potion de soin               |\n")
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
	case "S" , "s"  :
		if (p.inventaire[RessourcePotSoin] >= 1) {
		p.Takepot()
	} else {
		fmt.Println()
		fmt.Println("Vous n'avez de potion de soin à utiliser")
		fmt.Println()
	}
case "P" , "p" : 
	if (p.inventaire[RessourcePotPoison] >= 1) {
		p.Usepot_poison(monstre)
	} else {
		fmt.Println()
		fmt.Println("Vous n'avez de potion de poison à utiliser")
		fmt.Println()
	}
	default: fmt.Println("Commande invalide.")
}
}

func (p *Character) Fuir() {
	k := rand.IntN(2)
	if k == 0 {
		combat_pas_finis = false
		fmt.Println()
		fmt.Printf("Vous avez réussi à fuire le combat !")
		fmt.Println()
	} else {
		fmt.Println()
		fmt.Printf("Vous n'avez pas réussi à fuire le combat !")
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
	fmt.Printf("\t|     [C] Coup d'épée          [R] Retour        |\n")}
	if p.AttaqueName ==  "Coup de baton" {
	fmt.Printf("\t|     [C] Coup de baton        [R] Retour        |\n")}
	if p.AttaqueName ==  "Coups vicieux" {
	fmt.Printf("\t|     [C] Coups vicieux        [R] Retour        |\n")}
	fmt.Printf("\t|                                                |\n")
	fmt.Printf("\t|                                                |\n")
	if (p.SkillName == "Coup de dague") && (Sort) {
	fmt.Printf("\t|               [S] Coup de dague                |\n")}
	if (p.SkillName == "Boule de feu") && (Sort) {
	fmt.Printf("\t|               [S] Boule de feu                 |\n")}
	if (p.SkillName == "Coup critique") && (Sort) {
	fmt.Printf("\t|               [S] Coup critique                |\n")}
	fmt.Printf("\t|                                                |\n")
	fmt.Printf("\t|                                                |\n")
	fmt.Printf("\t+------------------------------------------------+\n")
	fmt.Println()
	var moove_attack string
	fmt.Scanln(&moove_attack)
	switch moove_attack{
	case "C" , "c" : {
		monstre.PV_actuelle -= p.AttaqueDegats
	}
	case "S" , "s" : {
		monstre.PV_actuelle -= p.SkillDegats
	}
	default: fmt.Println("Commande invalide.") 
	}
}

func (p *Character) Combat_start_premier() {
	victoire = false
	combat_pas_finis = true
	monstre := p.choix_monstre()
	fmt.Println()
	fmt.Printf("Vous rencontrez un %s sauvage !", monstre.name)
	fmt.Println()
	for combat_pas_finis {
		if monstre.Poison {
			fmt.Println()
			fmt.Printf("Le monstre ennemi %s perds %d PV", monstre.name, 10)
			fmt.Println()
			monstre.PV_actuelle -= 10
			fmt.Println()
			if monstre.PV_actuelle <= 0 {
				combat_pas_finis = false
				victoire = true
				continue
			}
		}
		fmt.Println()
		fmt.Printf("Les pv du monstre %s sont de %d/%d ", monstre.name , monstre.PV_actuelle, monstre.PV_max)
		fmt.Println()
		fmt.Println("Que voulez vous faire ?")
		fmt.Printf("\n")
		fmt.Printf("\t+------------------------------------------------+\n")
		fmt.Printf("\t|                                                |\n")
		fmt.Printf("\t|            Que voulez vous faire ?             |\n")
		fmt.Printf("\t|                                                |\n")
		fmt.Printf("\t|    [O] Objet       [F] Fuir       [A] Attaque  |\n")
		fmt.Printf("\t|                                                |\n")
		fmt.Printf("\t+------------------------------------------------+\n")
		fmt.Println()
		var moove string
		fmt.Scanln(&moove)
		switch moove { 
		case  "O" , "o"  :{
			p.Objet_utilitaire(&monstre)
		}
		case  "F" ,"f" : {
			p.Fuir()
			if combat_pas_finis == false {
				continue
			}
		}
		case "A" , "a" : {
			p.Attaque(&monstre)
		}
		default: fmt.Println("Commande invalide.") 
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
			fmt.Println("Le monstre vous attaque mais vous réussissez à l'esquiver !")
			fmt.Println()
		} else {
		fmt.Println("Le monstre ennemi est enervé , il vous charge")
		pv_perdu := monstre.attack
		if count_super_monstre%4 == 0 {
			pv_perdu = pv_perdu*2
		}
		fmt.Printf("Vous perdez %d Pvs\n",pv_perdu)
		p.pv -= monstre.attack
		fmt.Printf("Pvs actuelle : %d/%d",p.pv,p.pvmax)
		fmt.Println()
		count_super_monstre ++
	}
		if p.pv <= 0 {
			victoire = false 
			combat_pas_finis = false 
			continue 
		}
	}
	if victoire {
	fmt.Println()
	fmt.Println("Vous avez gagné le combat")
	fmt.Println()
	victoire = false
	combat_pas_finis = true
	} else {
	fmt.Println()
	fmt.Println("Vous avez perdu tout vos pvs , vous êtes mort !")
	fmt.Println()
	victoire = false 
	combat_pas_finis = true
	p.IsDead()
	}
}

func (p *Character) Combat_start_second() {}
