package projet

import (
	"fmt"
)

type Character struct {
	name         string
	money 		 int
	classe       string
	pvmax        int
	pv           int
	Mana 		 int
	Manamax		 int
	experience   int
	inventaire   map[string]int
	Resurrection bool
	Skill 		 []Skill
	AttaqueName string
	AttaqueDegats int
	SkillName     string
	SkillDegats   int
	Skillmana	  int
	Equipements map[string]string
}

func (p *Character) InventairePlein() bool {
    total := 0

    for _, quantite := range p.inventaire {
        total += quantite
    }

    return total >= inventaire_taillemax
}

func (p *Character) initcharacter(name string, classe string) {
	p.name = name
	p.money = 10000
	p.inventaire = make(map[string]int)
	p.classe = classe
	p.experience = 0
	p.Equipements = make(map[string]string)
	p.Manamax = 100
	p.Mana = 50
	switch classe {
	case "Guerrier" :
		p.pvmax = 200
		p.pv = p.pvmax / 2
		p.AttaqueName = "Coup d'épée"
		p.AttaqueDegats = 20
		p.SkillName = "Coup critique"
		p.SkillDegats = 35
		p.Skillmana = 25
	case "Sorcier" :
		p.pvmax = 100
		p.pv = p.pvmax / 2
		p.AttaqueName = "Coup de baton"
		p.AttaqueDegats = 10
		p.SkillName = "Boule de feu"
		p.SkillDegats = 70
		p.Skillmana = 50
	case "Assassin" :
		p.pvmax = 100
		p.pv = p.pvmax / 2
		p.AttaqueName = "Coup de dague"
		p.AttaqueDegats = 15
		p.SkillName = "Coups vicieux"
		p.SkillDegats = 45
		p.Skillmana = 35
	}
}

func (p *Character) displayinfo() {
}

func (p *Character) AccessInventory() {
	fmt.Println()
	for itemname, itemquantity := range p.inventaire {
		afficherTexte50("%s : %d\n", itemname, itemquantity)
	}
	fmt.Println()
	afficherTexte50("écus : %d\n", p.money)
	fmt.Println("============================inventaire============================")
	fmt.Println()
	afficherTexte20("Que voulez vous faire ?")
	afficherTexte20("P : Utiliser une potion de soin")
	afficherTexte20("M : Utiliser une potion de mana")
	afficherTexte20("E : Mettre mon équipement ")
	afficherTexte20("R : Retour")
	var inv string
		fmt.Scanln(&inv)
		if inv == "P" {
			p.takepotS()
			fmt.Println()
			fmt.Println()
			p.AccessInventory()
		}
		if inv == "M" {
			p.takepotM()
			fmt.Println()
			fmt.Println()
			p.AccessInventory()
		}
		if inv == "E" {
			if casque_equipe && p.inventaire[RessourceChapeaudegeant] == 1{
				casque_equipe = false
				p.inventaire[RessourceChapeaudegeant] -=1
			fmt.Println()
			fmt.Println("Vous équipez votre chapeau du géant ! (+30PV)")
			p.pvmax += 30
			affiche_casque = true 
			fmt.Println()
			fmt.Println()
			p.AccessInventory()
			}
			if botte_equipe && p.inventaire[RessourceBotteArcenciel] == 1 {
				p.inventaire[RessourceBotteArcenciel] -= 1
				botte_equipe = false 
				fmt.Println("Vous équipez vos bottes Arc-En-Ciel ! (+5PV / +20Mana)")
				p.pvmax += 5
				p.Manamax += 20
				affiche_botte = true 
				fmt.Println()
				fmt.Println()
				p.AccessInventory()
			}
			if armurecuir_equipe && p.inventaire[RessourceArmurecuir] == 1 {
				p.inventaire[RessourceArmurecuir] -= 1
				armurecuir_equipe = false
				fmt.Println("Vous équipez votre Armure en cuir ! (+15PV)")
				p.pvmax += 15
			affiche_armurecuir = true 
			fmt.Println()
			fmt.Println()
			p.AccessInventory()
			}
			if armureivoire_equipe && p.inventaire[RessourceArmureIvoire] == 1 {
				p.inventaire[RessourceArmureIvoire] -= 1
				armureivoire_equipe = false
				fmt.Println("Vous équipez votre Armure en  ivoire ! (+30PV / + 30Mana)")
				p.pvmax += 30
				p.Manamax += 30
			affiche_armureivoire = true 
			fmt.Println()
			fmt.Println()
			p.AccessInventory()
			}  
			if p.inventaire[RessourceArmureIvoire] == 0 && p.inventaire[RessourceArmurecuir] == 0 && p.inventaire[RessourceChapeaudegeant]==0 && p.inventaire[RessourceBotteArcenciel] == 0 {
				fmt.Println()
				fmt.Println("Vous n'avez rien d'autres à équiper !")
			fmt.Println()
			fmt.Println()
			p.AccessInventory()
			}
		}
}

func (p *Character) TakepotS() {
	potquantity, potcheck := p.inventaire[RessourcePotSoin]
	if !potcheck {
		fmt.Println("Vous n'avez pas de potion dans votre inventaire.")
		return
	}
	if potquantity <= 0 {
		fmt.Println("Vous n'avez plus de potion dans votre inventaire.")
		return
	}
	p.pv += 50
	if p.pv > p.pvmax {
		p.pv = p.pvmax
	}
	p.inventaire[RessourcePotSoin]--

	fmt.Printf("Vous avez utilisé une potion. Votre vie est maintenant de %d/%d.\n", p.pv, p.pvmax)
}
func (p *Character) TakepotM() {
	potquantity, potcheck := p.inventaire[RessourcePotSoin]
	if !potcheck {
		fmt.Println("Vous n'avez pas de potion dans votre inventaire.")
		return
	}
	if potquantity <= 0 {
		fmt.Println("Vous n'avez plus de potion dans votre inventaire.")
		return
	}
	p.Mana += 50
	if p.Mana > p.Manamax {
		p.Mana = p.Manamax
	}
	p.inventaire[RessourcePotMana]--

	fmt.Printf("Vous avez utilisé une potion. Votre mana est maintenant de %d/%d.\n", p.Mana, p.Manamax)
}
