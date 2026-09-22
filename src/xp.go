package projet

import (
	"fmt"
	"math/rand"
)

const (
	niveauInitial         = 1
	experienceMinVictoire = 5
	experienceMaxVictoire = 15
	experienceParNiveau   = 6
	pvParNiveau           = 10
	degatsParNiveau       = 2
)

func (p *Character) Niveau() int {
	niveau := niveauInitial
	experienceRestante := p.experience
	for experienceRestante >= niveau*experienceParNiveau {
		experienceRestante -= niveau * experienceParNiveau
		niveau++
	}
	return niveau
}

func (p *Character) GagnerCombat() {
	ancienNiveau := p.Niveau()

	nombreValeursPossibles := experienceMaxVictoire - experienceMinVictoire + 1
	experienceAleatoire := rand.Intn(nombreValeursPossibles)
	experienceGagnee := experienceMinVictoire + experienceAleatoire

	p.experience += experienceGagnee
	fmt.Printf("Vous gagnez %d points d'experience. Total : %d\n", experienceGagnee, p.experience)

	nouveauNiveau := p.Niveau()
	niveauxGagnes := nouveauNiveau - ancienNiveau

	if niveauxGagnes > 0 {
		fmt.Printf("Vous gagnez %d niveau(x) !\n", niveauxGagnes)

		bonusPV := pvParNiveau * niveauxGagnes
		bonusDegats := degatsParNiveau * niveauxGagnes

		p.pvmax += bonusPV
		p.pv += bonusPV
		p.AttaqueDegats += bonusDegats
	}
}
