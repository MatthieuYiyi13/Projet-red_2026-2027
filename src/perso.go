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
	Initiative	int
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
		p.pvmax = 140
		p.pv = p.pvmax / 2
		p.AttaqueName = "Coup d'épée"
		p.AttaqueDegats = 20
		p.SkillName = "Coup critique"
		p.SkillDegats = 80
		p.Skillmana = 25
		p.Initiative = 8
	case "Sorcier" :
		p.pvmax = 90
		p.pv = p.pvmax / 2
		p.AttaqueName = "Coup de baton"
		p.AttaqueDegats = 10
		p.SkillName = "Boule de feu"
		p.SkillDegats = 120
		p.Skillmana = 50
		p.Initiative = 4
	case "Assassin" :
		p.pvmax = 110
		p.pv = p.pvmax / 2
		p.AttaqueName = "Coup de dague"
		p.AttaqueDegats = 15
		p.SkillName = "Coups vicieux"
		p.SkillDegats = 100
		p.Skillmana = 35
		p.Initiative = 9
	}
}

func (p *Character) Displayinfo() {
fmt.Println("========================================================")
	afficherTexte50("Nom : %s",p.name)
	afficherTexte50("Classe : %s", p.classe)
	afficherTexte50("Pvs actuelle : %d / %d", p.pv,p.pvmax)
	afficherTexte50("Mana actuelle : %d/%d", p.Mana,p.Manamax)
	afficherTexte50("Niveau : %d", p.Niveau())
	afficherTexte50("Experience gagné: %d", p.experience)
	afficherTexte50("Degats de %s : %d", p.AttaqueName, p.AttaqueDegats)
	if Sort {
	afficherTexte50("Degats de %s : %d", p.SkillName,p.SkillDegats)
	}
	if p.Resurrection == false {
	afficherTexte50("1 Resurrection restante ")
	} else {
	afficherTexte50("Aucune Resurrection restante ")
	}
fmt.Println("========================================================")
}


