package projet

import (
	"fmt"
	"math/rand/v2"
)

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
	n := rand.IntN(9)

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
			name:        "Mammouth",
			attack:      12,
			PV_max:      350,
			PV_actuelle: 350,
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
			name:        "Lynx fumée",
			attack:      25,
			PV_max:      200,
			PV_actuelle: 200,
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
	if (moove_objet == "S" || moove_objet == "s") && (p.inventaire[RessourcePotSoin] >= 1) {
		p.Takepot()
	}
	if (moove_objet == "P" || moove_objet == "p") && (p.inventaire[RessourcePotPoison] >= 1) {
		p.Usepot_poison(monstre)
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
	if moove_attack == "C" || moove_attack == "c" {
		monstre.PV_actuelle -= p.AttaqueDegats
	}
	if moove_attack == "S" || moove_attack == "s" {
		monstre.PV_actuelle -= p.SkillDegats
	}

}

func (p *Character) Combat_start_premier() {
	victoire = false
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
			fmt.Printf("Les pv du monstre sont de %d/%d ", monstre.PV_actuelle, monstre.PV_max)
			fmt.Println()
			if monstre.PV_actuelle <= 0 {
				combat_pas_finis = false
				victoire = true
				continue
			}
		}
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
		if moove == "O" || moove == "o" {
			p.Objet_utilitaire(&monstre)
		}
		if moove == "F" || moove == "f" {
			p.Fuir()
			if combat_pas_finis == false {
				continue
			}
		}
		if moove == "A" || moove == "a" {
			p.Attaque(&monstre)
		}
		if monstre.PV_actuelle <= 0 {
			combat_pas_finis = false
			continue
		}
	}
}

func (p *Character) Combat_start_second() {}
