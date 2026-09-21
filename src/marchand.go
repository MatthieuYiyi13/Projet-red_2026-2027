package projet

import "fmt"

var potion_gratuite bool = true

func (p *Character) Acheter_Materiaux() {
	fmt.Println("Que voulez vous acheter ? ")
	fmt.Println("1 : Tissu de spectre: 100")
	fmt.Println("2 : Peau de géant: 150")
	fmt.Println("3 : Fil d'araignée: 150")
	fmt.Println("4 : Dent de loup: 200")
	fmt.Println("5 : Poil de mammouth: 250")
	fmt.Println("6 : Griffe de lynx de fumée: 250")
	fmt.Println("7 : Cendre de lynx de fumée: 400")
	fmt.Println("8 : Fragment de glace: 500")
	fmt.Println("9 : Fragment du Roi de la Nuit: 750")
	fmt.Println("10 : Sabot de licorne: 999")
	fmt.Println("11 : Retour au menu ")

	var choix_objet int
	fmt.Scanln(&choix_objet)

	switch choix_objet {
	case 1:
		if p.money >= 100 {
			p.inventaire["tissu de spectre"] += 1
			p.money -= 100
			fmt.Println("Vous avez acheté un tissu de spectre")
			fmt.Println()
			p.Acheter_Materiaux()
		} else {
			fmt.Println("Vous n'avez pas l'argent nécessaire pour acheter ça !")
			fmt.Println()
			p.Acheter_Materiaux()
		}
	case 2:
		if p.money >= 150 {
			p.inventaire["Peau de géant"] += 1
			p.money -= 150
			fmt.Println("Vous avez acheté une peau de géant")
			fmt.Println()
			p.Acheter_Materiaux()
		} else {
			fmt.Println("Vous n'avez pas l'argent nécessaire pour acheter ça !")
			fmt.Println()
			p.Acheter_Materiaux()
		}
	case 3:
		if p.money >= 150 {
			p.inventaire["Fil d'araignéet"] += 1
			p.money -= 150
			fmt.Println("Vous avez acheté un fil d'araignée")
			fmt.Println()
			p.Acheter_Materiaux()
		} else {
			fmt.Println("Vous n'avez pas l'argent nécessaire pour acheter ça !")
			fmt.Println()
			p.Acheter_Materiaux()
		}
	case 4:
		if p.money >= 200 {
			p.inventaire["Dent de loup"] += 1
			p.money -= 200
			fmt.Println("Vous avez acheté une dent de loup")
			fmt.Println()
			p.Acheter_Materiaux()
		} else {
			fmt.Println("Vous n'avez pas l'argent nécessaire pour acheter ça !")
			fmt.Println()
			p.Acheter_Materiaux()
		}
	case 5:
		if p.money >= 250 {
			p.inventaire["Poil de mammouth"] += 1
			p.money -= 250
			fmt.Println("Vous avez acheté du poil de mammouth")
			fmt.Println()
			p.Acheter_Materiaux()
		} else {
			fmt.Println("Vous n'avez pas l'argent nécessaire pour acheter ça !")
			fmt.Println()
			p.Acheter_Materiaux()
		}
	case 6:
		if p.money >= 250 {
			p.inventaire["Griffe de lynx de fumée"] += 1
			p.money -= 250
			fmt.Println("Vous avez acheté une griffe de lynx de fumée")
			fmt.Println()
			p.Acheter_Materiaux()
		} else {
			fmt.Println("Vous n'avez pas l'argent nécessaire pour acheter ça !")
			fmt.Println()
			p.Acheter_Materiaux()
		}
	case 7:
		if p.money >= 400 {
			p.inventaire["Cendre de lynx de fumée"] += 1
			p.money -= 400
			fmt.Println("Vous avez acheté de la cendre de lynx de fumée")
			fmt.Println()
			p.Acheter_Materiaux()
		} else {
			fmt.Println("Vous n'avez pas l'argent nécessaire pour acheter ça !")
			fmt.Println()
			p.Acheter_Materiaux()
		}
	case 8:
		if p.money >= 500 {
			p.inventaire["Fragment de glace"] += 1
			p.money -= 500
			fmt.Println("Vous avez acheté un fragment de glace")
			fmt.Println()
			p.Acheter_Materiaux()
		} else {
			fmt.Println("Vous n'avez pas l'argent nécessaire pour acheter ça !")
			fmt.Println()
			p.Acheter_Materiaux()
		}
	case 9:
		if p.money >= 750 {
			p.inventaire["Fragment du Roi de la Nuit"] += 1
			p.money -= 750
			fmt.Println("Vous avez acheté un fragment du Roi de la Nuit")
			fmt.Println()
			p.Acheter_Materiaux()
		} else {
			fmt.Println("Vous n'avez pas l'argent nécessaire pour acheter ça !")
			fmt.Println()
			p.Acheter_Materiaux()
		}
	case 10:
		if p.money >= 999 {
			p.inventaire["Sabot de licorne"] += 1
			p.money -= 999
			fmt.Println("Vous avez acheté un sabot de licorne")
			fmt.Println()
			p.Acheter_Materiaux()
		} else {
			fmt.Println("Vous n'avez pas l'argent nécessaire pour acheter ça !")
			fmt.Println()
			p.Acheter_Materiaux()
		}
	}
}

func (p *Character) Acheter_Utilitaire() {
	if potion_gratuite {
		fmt.Println("Que voulez vous acheter ? ")
		fmt.Println("1 : Potion de soin (la première gratuite): 0")
		fmt.Println("2 : Potion de poison: 10")
		fmt.Println("3 : Poche supplémentaire (+10 de stockage dans l'inventaire)")
		fmt.Println("4 : Retour")
		fmt.Println()

		var choix_utile int
		fmt.Scanln(&choix_utile)

		switch choix_utile {
		case 1:
			potion_gratuite = false
			fmt.Println("Vous avez obtenue la potion de soin gratuite")
			fmt.Println()
			p.inventaire["potion de soin"] += 1
			p.Acheter_Utilitaire()
		case 2:
			fmt.Println("Vous avez acheté une potion de poison")
			fmt.Println()
			p.inventaire["potion de poison"] += 1
			p.Acheter_Utilitaire()
		case 3:
			fmt.Println("Vous avez aggrandi votre inventaire !")
			fmt.Println()
			inventaire_taillemax += 10
			p.Acheter_Utilitaire()
		}
	} else {
		fmt.Println("Que voulez vous acheter ? ")
		fmt.Println("1 : Potion de soin : 10")
		fmt.Println("2 : Potion de poison: 10")
		fmt.Println("3 : Poche supplémentaire (+10 de stockage dans l'inventaire)")
		fmt.Println("4 : Retour")
		fmt.Println()

		var choix_utile int
		fmt.Scanln(&choix_utile)

		switch choix_utile {
		case 1:
			fmt.Println("Vous avez acheté une potion de soin")
			fmt.Println()
			p.inventaire["potion de soin"] += 1
			p.Acheter_Utilitaire()
		case 2:
			fmt.Println("Vous avez acheté une potion de poison")
			fmt.Println()
			p.inventaire["potion de poison"] += 1
			p.Acheter_Utilitaire()
		case 3:
			fmt.Println("Vous avez aggrandi votre inventaire !")
			fmt.Println()
			inventaire_taillemax += 10
			p.Acheter_Utilitaire()
		}
	}
}

func (p *Character) Acheter_Sorts() {
	fmt.Println("1 : Coup critique : 20")
	fmt.Println("2 : Boule de feu : 20")
	fmt.Println("3 : Coups vicieux : 20")

	var choix_sort int
	fmt.Scanln(&choix_sort)
	switch choix_sort {
	case 1:
		fmt.Println("Vous avez acheté le sort Coup critique !")
		fmt.Println()
		p.inventaire["Coup critique"] += 1
		p.Acheter_Sorts()
	case 2:
		fmt.Println("Vous avez acheté le sort Boule de feu !")
		fmt.Println()
		p.inventaire["Boule de feu"] += 1
		p.Acheter_Sorts()
	case 3:
		fmt.Println("Vous avez acheté le sort Coups vicieux")
		fmt.Println()
		p.inventaire["Coups vicieux"] += 1
		p.Acheter_Sorts()
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
		fmt.Println("Que souhaites-tu acheter ?")
		fmt.Println("1. Matériaux")
		fmt.Println("2. Utilitaires")
		fmt.Println("3. Sorts")
		var choix_achat int
		fmt.Scanln(&choix_achat)
		switch choix_achat {
		case 1:
			p.Acheter_Materiaux()
		case 2:
			p.Acheter_Utilitaire()
		case 3:
			p.Acheter_Sorts()
		}
	case 2:
		p.Vendre()
	case 3:
		fmt.Println("Au revoir !")
	}
}
