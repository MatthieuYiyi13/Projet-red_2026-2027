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
	experience   int
	inventaire   map[string]int
	Resurrection bool
	Skill 		 []Skill
	AttaqueName string
	AttaqueDegats int
	SkillName     string
	SkillDegats   int
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
	switch classe {
	case "Guerrier":
		p.pvmax = 200
		p.pv = p.pvmax / 2
		p.AttaqueName = "Coup d'épée"
		p.AttaqueDegats = 20
		p.SkillName = "Coup critique"
		p.SkillDegats = 35
	case "Sorcier":
		p.pvmax = 100
		p.pv = p.pvmax / 2
		p.AttaqueName = "Coup de baton"
		p.AttaqueDegats = 10
		p.SkillName = "Boule de feu"
		p.SkillDegats = 70
	case "Assassin":
		p.pvmax = 100
		p.pv = p.pvmax / 2
		p.AttaqueName = "Coup de dague"
		p.AttaqueDegats = 15
		p.SkillName = "Coups vicieux"
		p.SkillDegats = 45
	}
}

func (p *Character) displayinfo() {
}

func (p *Character) AccessInventory() {
	fmt.Println()
	for itemname, itemquantity := range p.inventaire {
		fmt.Printf("%s : %d\n", itemname, itemquantity)
		fmt.Println()
	}
	fmt.Printf("écus : %d\n", p.money)
	fmt.Println()
	fmt.Println()
	fmt.Println("Que voulez vous faire ?")
	fmt.Println("P : Utiliser une potion de soin ")
	fmt.Println("E : Mettre mon équipement ")
	fmt.Println("R : Retour")
	var inv string
		fmt.Scanln(&inv)
		if inv == "P" {
			p.Takepot()
			fmt.Println()
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
			}
			if botte_equipe && p.inventaire[RessourceBotteArcenciel] == 1 {
				p.inventaire[RessourceBotteArcenciel] -= 1
				botte_equipe = false 
				fmt.Println("Vous équipez vos bottes Arc-En-Ciel ! (+5PV)")
				p.pvmax += 5
				affiche_botte = true 
				fmt.Println()
			}
			if armurecuir_equipe && p.inventaire[RessourceArmurecuir] == 1 {
				p.inventaire[RessourceArmurecuir] -= 1
				armurecuir_equipe = false
				fmt.Println("Vous équipez votre Armure en cuir ! (+15PV)")
				p.pvmax += 15
			affiche_armurecuir = true 
			fmt.Println()
			}
			if armureivoire_equipe && p.inventaire[RessourceArmureIvoire] == 1 {
				p.inventaire[RessourceArmureIvoire] -= 1
				armureivoire_equipe = false
				fmt.Println("Vous équipez votre Armure en  ivoire ! (+30PV)")
				p.pvmax += 30
			affiche_armureivoire = true 
			fmt.Println()
			}  
			if p.inventaire[RessourceArmureIvoire] == 0 && p.inventaire[RessourceArmurecuir] == 0 && p.inventaire[RessourceChapeaudegeant]==0 && p.inventaire[RessourceBotteArcenciel] == 0 {
				fmt.Println()
				fmt.Println("Vous n'avez rien d'autres à équiper !")
			fmt.Println()
			}
		}
}

func (p *Character) Takepot() {
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
