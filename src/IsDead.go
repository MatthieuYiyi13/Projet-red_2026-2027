package projet

import "fmt"

func (p *Character) IsDead() {
	if p.pv <= 0 {
		if !p.Resurrection {
			fmt.Println("Vous êtes mort !")
			p.pv = p.pvmax / 2
			p.Resurrection = true
			fmt.Println("Angéline vous a ressucitée.")
		} else {
			gameOver()
		}
	}
}
