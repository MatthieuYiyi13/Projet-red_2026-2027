package projet

import (
	"fmt"
	"math/rand"
)

const (
	experienceMinVictoire = 5
	experienceMaxVictoire = 15
	pvParVictoire         = 10
	degatsParVictoire     = 2
)

func (p *Character) GagnerCombat() {
	experienceGagnee := rand.Intn(experienceMaxVictoire-experienceMinVictoire+1) + experienceMinVictoire
	p.experience += experienceGagnee
	fmt.Printf("Vous gagnez %d points d'experience. Total : %d\n", experienceGagnee, p.experience)
	p.pvmax += pvParVictoire
	p.pv += pvParVictoire
	p.AttaqueDegats += degatsParVictoire
}
