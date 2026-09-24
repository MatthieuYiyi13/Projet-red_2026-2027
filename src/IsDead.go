package projet

import "fmt"

func (p *Character) IsDead() {
	if p.pv <= 0 {
		if !p.Resurrection {
			p.pv = p.pvmax / 2
			p.Resurrection = true
			fmt.Println("Ange et Line vous ont ressucitée.")

		} else {
			gameOver()
		}
	}
}