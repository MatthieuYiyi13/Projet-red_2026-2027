package projet

import "fmt"

func (p *Character) Acheter() {
	fmt.Println("Acheter")
}

func (p *Character) Vendre() {
	fmt.Println("Vendre")
}

func (p *Character) Marchand() {
	fmt.Println("Bonjour jeune aventurier, je vois que tu as réussi à me trouver dans cette magnifique ville de Qarth !")
	fmt.Println()
	fmt.Println("Que souhaites-tu faire maintenant ?")
	fmt.Println()
	fmt.Println("1. Acheter")
	fmt.Println("2. Vendre")
	fmt.Println("3. Quitter le menu")

	var choiceM int
	fmt.Scanln(&choiceM)

	switch choiceM {
	case 1:
		p.Acheter()
	case 2:
		p.Vendre()
	case 3:
		fmt.Println("Au revoir !")
	default:
		fmt.Println("Choix invalide. Tu dois choisir 1, 2 ou 3.")
	}
}
