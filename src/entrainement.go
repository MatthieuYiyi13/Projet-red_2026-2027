package projet

import "fmt"

func (p *Character) TrainingFight() {
	if p.pv <= 0 {
		fmt.Println("Vous n'avez plus de PV pour vous entrainer.")
		return
	}

	gobelin := p.InitGobelin()
	tour := 1
	stopSound()
	PlaySoundAsyncCombatE()

	fmt.Printf("Un %s apparait pour le combat d'entrainement !\n", gobelin.name)

	for p.pv > 0 && gobelin.PV_actuelle > 0 {
		fmt.Printf("\n=== Tour %d ===\n", tour)
		fmt.Printf("Vos PV : %d/%d | PV du %s : %d/%d\n", p.pv, p.pvmax, gobelin.name, gobelin.PV_actuelle, gobelin.PV_max)
		fmt.Print("[A] Attaquer  [Q] Quitter l'entrainement : ")

		var choix string
		if _, err := fmt.Scanln(&choix); err != nil {
			fmt.Println("Entrainement interrompu.")
			return
		}

		switch choix {
		case "A", "a":
			gobelin.PV_actuelle -= p.AttaqueDegats
			if gobelin.PV_actuelle < 0 {
				gobelin.PV_actuelle = 0
			}
			fmt.Printf("Vous infligez %d degats au %s.\n", p.AttaqueDegats, gobelin.name)
		case "Q", "q":
			fmt.Println("Vous quittez l'entrainement.")
			stopSound()
			PlaySoundAsyncDebut()
			return
		default:
			fmt.Println("Choix invalide. Le tour ne change pas.")
			continue
		}

		if gobelin.PV_actuelle == 0 {
			fmt.Printf("Vous avez vaincu le %s en %d tour(s) !\n", gobelin.name, tour)
			p.GagnerCombat()
			stopSound()
			PlaySoundAsyncDebut()
			return
		}

		p.pv -= gobelin.attack
		fmt.Printf("Le %s vous inflige %d degats.\n", gobelin.name, gobelin.attack)
		if p.pv <= 0 {
			p.pv = 1
			fmt.Println("Vous perdez l'entrainement. Il vous reste 1 PV.")
			return
		}
		tour++
	}
}