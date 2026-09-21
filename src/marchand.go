package projet

import "fmt"

var objet_marchand = map[string]int{
   "Tissu de spectre":            10,
	"Peau de géant":              10,
	"Fil d'araignée":             10, 
	"Dent de loup":               10,
	"Poil de mammouth":           10,
	"Griffe de lynx de fumée":    10,
	"Cendre de lynx de fumée":    10,
	"Fragment de glace":          10,
	"Fragment du Roi de la Nuit": 10,
	"Sabot de licorne":           10,
}
 
func (p *Character) Acheter() {
	fmt.Println("Que voulez vous acheter ? ")
	fmt.Println("1 : Tissu de spectre: 10")
	fmt.Println("2 : Peau de géant: 10")
	fmt.Println("3 : Fil d'araignée: 10")
	fmt.Println("4 : Dent de loup: 10")
	fmt.Println("5 : Poil de mammouth: 10")
	fmt.Println("6 : Griffe de lynx de fumée: 10")
	fmt.Println("7 : Cendre de lynx de fumée: 10")
	fmt.Println("8 : Fragment de glace: 8")
	fmt.Println("9 : Fragment du Roi de la Nuit: 9")
	fmt.Println("10 : Sabot de licorne: 10")
	fmt.Println("11 : Retour au menu ")

var choix_objet int
	fmt.Scanln(&choix_objet)

	switch choix_objet {
	case 1:
		p.inventaire["tissu de spectre"] += 1 
		p.money -= 10
		fmt.Println("Vous avez acheté un tissu de spectre")
		fmt.Println()
		p.Acheter()
	case 2:
		p.inventaire["Peau de géant"] += 1 
		p.money -= 10
		fmt.Println("Vous avez acheté une peau de géant")
		fmt.Println()
		p.Acheter()
	case 3:
		p.inventaire["Fil d'araignéet"] += 1 
		p.money -= 10
		fmt.Println("Vous avez acheté un fil d'araignée")
		fmt.Println()
		p.Acheter()
	case 4:
		p.inventaire["Dent de loup"] += 1 
		p.money -= 10
		fmt.Println("Vous avez acheté une dent de loup")
		fmt.Println()
		p.Acheter()
	case 5:
		p.inventaire["Poil de mammouth"] += 1 
		p.money -= 10
		fmt.Println("Vous avez acheté du poil de mammouth")
		fmt.Println()
		p.Acheter()
	case 6:
		p.inventaire["Griffe de lynx de fumée"] += 1 
		p.money -= 10
		fmt.Println("Vous avez acheté une griffe de lynx de fumée")
		fmt.Println()
		p.Acheter()
	case 7:
		p.inventaire["Cendre de lynx de fumée"] += 1 
		p.money -= 10
		fmt.Println("Vous avez acheté de la cendre de lynx de fumée")
		fmt.Println()
		p.Acheter()
	case 8:
		p.inventaire["Fragment de glace"] += 1 
		p.money -= 10
		fmt.Println("Vous avez acheté un fragment de glace")
		fmt.Println()
		p.Acheter()
	case 9:
		p.inventaire["Fragment du Roi de la Nuit"] += 1 
		p.money -= 10
		fmt.Println("Vous avez acheté un fragment du Roi de la Nuit")
		fmt.Println()
		p.Acheter()
	case 10:
		p.inventaire["Sabot de licorne"] += 1 
		p.money -= 10
		fmt.Println("Vous avez acheté un sabot de licorne")
		fmt.Println()
		p.Acheter()
}
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
