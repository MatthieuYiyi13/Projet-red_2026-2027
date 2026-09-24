package projet

func (p *Character) IsDead() {
	if p.pv <= 0 {
		if !p.Resurrection {
			p.pv = p.pvmax / 2
			p.Resurrection = true
			afficherTexte50("Ange et Line vous ont ressucitée.")
			p.Affiche_ange()
			afficherTexte50("Vous avez ressucitée , vous n'avez plus le droit à l'erreur")
		} else {
			gameOver()
		}
	}
}