package projet

import "fmt"

var potion_gratuite bool = true
var Sort bool = false

func (p *Character) Acheter_Materiaux() {
	fmt.Println("Que voulez vous acheter ?")
	fmt.Println("1 :  Tissu de spectre: 100 écus")
	fmt.Println("2 :  Peau de géant: 150 écus")
	fmt.Println("3 :  Fil d'araignée: 150 écus")
	fmt.Println("4 :  Dent de loup: 200 écus")
	fmt.Println("5 :  Poil de mammouth: 250 écus")
	fmt.Println("6 :  Griffe de lynx fumée: 250 écus")
	fmt.Println("7 :  Cendre de lynx fumée: 400 écus")
	fmt.Println("8 :  Fragment de glace: 500 écus")
	fmt.Println("9 :  Fragment du Roi de la Nuit: 750 écus")
	fmt.Println("10 : Sabot de licorne: 999 écus")
	fmt.Println("11 : Retour au menu ")

	var choix_objet int
	fmt.Scanln(&choix_objet)

	switch choix_objet {
	case 1:
		if p.InventairePlein() {
			fmt.Println("Votre inventaire est plein ! Vous ne pouvez pas acheter d'objet.")
			p.Marchand()
		} else if p.money >= 100 {
			p.inventaire[RessourceTissuDeSpectre] += 1
			p.money -= 100
			fmt.Println("Vous avez acheté un tissu de spectre pour 100 écus")
			fmt.Println()
			p.Acheter_Materiaux()
		} else {
			fmt.Println("Vous n'avez pas l'argent nécessaire pour acheter ça !")
			p.Acheter_Materiaux()
		}
	case 2:
		if p.InventairePlein() {
			fmt.Println("Votre inventaire est plein ! Vous ne pouvez pas acheter d'objet.")
			p.Marchand()
		} else if p.money >= 150 {
			p.inventaire[RessourcePeauDeGeant] += 1
			p.money -= 150
			fmt.Println("Vous avez acheté une peau de géant pour 150 écus")
			fmt.Println()
			p.Acheter_Materiaux()
		} else {
			fmt.Println("Vous n'avez pas l'argent nécessaire pour acheter ça !")
			fmt.Println()
			p.Acheter_Materiaux()
		}
	case 3:
		if p.InventairePlein() {
			fmt.Println("Votre inventaire est plein ! Vous ne pouvez pas acheter d'objet.")
			p.Marchand()
		} else if p.money >= 150 {
			p.inventaire[RessourceFilDaraignee] += 1
			p.money -= 150
			fmt.Println("Vous avez acheté un fil d'araignée pour 150 écus")
			fmt.Println()
			p.Acheter_Materiaux()
		} else {
			fmt.Println("Vous n'avez pas l'argent nécessaire pour acheter ça !")
			fmt.Println()
			p.Acheter_Materiaux()
		}
	case 4:
		if p.InventairePlein() {
			fmt.Println("Votre inventaire est plein ! Vous ne pouvez pas acheter d'objet.")
			p.Marchand()
		} else if p.money >= 200 {
			p.inventaire[RessourceDentdeloup] += 1
			p.money -= 200
			fmt.Println("Vous avez acheté une dent de loup pour 200 écus")
			fmt.Println()
			p.Acheter_Materiaux()
		} else {
			fmt.Println("Vous n'avez pas l'argent nécessaire pour acheter ça !")
			fmt.Println()
			p.Acheter_Materiaux()
		}
	case 5:
		if p.InventairePlein() {
			fmt.Println("Votre inventaire est plein ! Vous ne pouvez pas acheter d'objet.")
			p.Marchand()
		} else if p.money >= 250 {
			p.inventaire[RessourcePoildemammouth] += 1
			p.money -= 250
			fmt.Println("Vous avez acheté du poil de mammouth pour 250 écus")
			fmt.Println()
			p.Acheter_Materiaux()
		} else {
			fmt.Println("Vous n'avez pas l'argent nécessaire pour acheter ça !")
			fmt.Println()
			p.Acheter_Materiaux()
		}
	case 6:
		if p.InventairePlein() {
			fmt.Println("Votre inventaire est plein ! Vous ne pouvez pas acheter d'objet.")
			p.Marchand()
		} else if p.money >= 250 {
			p.inventaire[RessourceGriffeDelynxfumee] += 1
			p.money -= 250
			fmt.Println("Vous avez acheté une griffe de lynx fumée pour 250 écus")
			fmt.Println()
			p.Acheter_Materiaux()
		} else {
			fmt.Println("Vous n'avez pas l'argent nécessaire pour acheter ça !")
			fmt.Println()
			p.Acheter_Materiaux()
		}
	case 7:
		if p.InventairePlein() {
			fmt.Println("Votre inventaire est plein ! Vous ne pouvez pas acheter d'objet.")
			p.Marchand()
		} else if p.money >= 400 {
			p.inventaire[RessourceCendreDelynxfumee] += 1
			p.money -= 400
			fmt.Println("Vous avez acheté de la cendre de lynx fumée pour 400 écus")
			fmt.Println()
			p.Acheter_Materiaux()
		} else {
			fmt.Println("Vous n'avez pas l'argent nécessaire pour acheter ça !")
			fmt.Println()
			p.Acheter_Materiaux()
		}
	case 8:
		if p.InventairePlein() {
			fmt.Println("Votre inventaire est plein ! Vous ne pouvez pas acheter d'objet.")
			p.Marchand()
		} else if p.money >= 500 {
			p.inventaire[RessourceFragmentdeglace] += 1
			p.money -= 500
			fmt.Println("Vous avez acheté un fragment de glace pour 500 écus")
			fmt.Println()
			p.Acheter_Materiaux()
		} else {
			fmt.Println("Vous n'avez pas l'argent nécessaire pour acheter ça !")
			fmt.Println()
			p.Acheter_Materiaux()
		}
	case 9:
		if p.InventairePlein() {
			fmt.Println("Votre inventaire est plein ! Vous ne pouvez pas acheter d'objet.")
			p.Marchand()
		} else if p.money >= 750 {
			p.inventaire[RessourceFragmentRoiDeLaNuit] += 1
			p.money -= 750
			fmt.Println("Vous avez acheté un fragment du Roi de la Nuit pour 750 écus")
			fmt.Println()
			p.Acheter_Materiaux()
		} else {
			fmt.Println("Vous n'avez pas l'argent nécessaire pour acheter ça !")
			fmt.Println()
			p.Acheter_Materiaux()
		}
	case 10:
		if p.InventairePlein() {
			fmt.Println("Votre inventaire est plein ! Vous ne pouvez pas acheter d'objet.")
			p.Marchand()
		} else if p.money >= 999 {
			p.inventaire[RessourceSabotDeLicorne] += 1
			p.money -= 999
			fmt.Println("Vous avez acheté un sabot de licorne pour 999 écus")
			fmt.Println()
			p.Acheter_Materiaux()
		} else {
			fmt.Println("Vous n'avez pas l'argent nécessaire pour acheter ça !")
			fmt.Println()
			p.Acheter_Materiaux()
		}
	case 11 :
		p.Marchand()
	default : afficherTexte50("Veuillez saisir une touche valide")
	}
}

func (p *Character) Acheter_Utilitaire() {
	if potion_gratuite {
		fmt.Println()
		fmt.Println("Que voulez vous acheter ? ")
		fmt.Println("1 : Potion de soin (la première gratuite): 0 écus")
		fmt.Println("2 : Potion de mana: 50")
		fmt.Println("3 : Potion de poison: 150 écus")
		fmt.Println("4 : Poche supplémentaire (+10 de stockage dans l'inventaire 3 achats max) : 100 écus")
		fmt.Println("5 : Retour")
		fmt.Println()

		var choix_utile int
		fmt.Scanln(&choix_utile)

		switch choix_utile {
		case 1:
			if p.InventairePlein() {
				fmt.Println("Votre inventaire est plein ! Vous ne pouvez pas acheter d'objet.")
				p.Marchand()
			} else {
				potion_gratuite = false
				fmt.Println("Vous avez obtenue la potion de soin gratuite")
				fmt.Println()
				p.inventaire[RessourcePotSoin] += 1
				p.Acheter_Utilitaire()
			}
		case 2:
			if p.InventairePlein() {
				fmt.Println("Votre inventaire est plein ! Vous ne pouvez pas acheter d'objet.")
				p.Marchand()
			} else if p.money >= 50 {
				fmt.Println("Vous avez acheté une potion de mana pour 50 écus")
				fmt.Println()
				p.money -= 50
				p.inventaire[RessourcePotMana] += 1
				p.Acheter_Utilitaire()
			} else {
				fmt.Println()
				fmt.Println("Vous n'avez pas l'argent pour m'acheter cela !")
				fmt.Println()
				p.Acheter_Utilitaire()
			}
		case 3:
			if p.InventairePlein() {
				fmt.Println("Votre inventaire est plein ! Vous ne pouvez pas acheter d'objet.")
				p.Marchand()
			} else if p.money >= 150 {
				fmt.Println("Vous avez acheté une potion de poison pour 150 écus")
				fmt.Println()
				p.money -= 150
				p.inventaire[RessourcePotPoison] += 1
				p.Acheter_Utilitaire()
			} else {
				fmt.Println()
				fmt.Println("Vous n'avez pas l'argent pour m'acheter cela !")
				fmt.Println()
				p.Acheter_Utilitaire()
			}
		case 4:
			 if p.money >= 100  && inventaire_taillemax<40{
				fmt.Println("Vous avez aggrandi votre inventaire pour 100 écus!")
				fmt.Println()
				p.money -= 100
				inventaire_taillemax += 10
				p.Acheter_Utilitaire()
			} else {
				fmt.Println()
				fmt.Println("Désolé vous ne pouvez pas acheter ça pour le moment!")
				fmt.Println()
				p.Acheter_Utilitaire()
			}
		case 5 :
			p.Marchand()
		default : 
		afficherTexte50("Veuillez saisir une touche valide")
		}
	} else {
		afficherTexte50("Que voulez vous acheter ? ")
		afficherTexte50("1 : Potion de soin  :  50 écus")
		afficherTexte50("2 : Potion de mana : 50 écus")
		afficherTexte50("3 : Potion de poison: 150 écus")
		afficherTexte50("4 : Poche supplémentaire (+10 de stockage dans l'inventaire) : 100 écus")
		afficherTexte50("5 : Retour")
		fmt.Println()

		var choix_utile int
		fmt.Scanln(&choix_utile)

		switch choix_utile {
		case 1:
			if p.InventairePlein() {
				afficherTexte50("Votre inventaire est plein ! Vous ne pouvez pas acheter d'objet.")
				p.Marchand()
			} else if p.money >= 50 {
				afficherTexte50("Vous avez acheté une potion de soin pour 50 écus")
				fmt.Println()
				p.inventaire[RessourcePotSoin] += 1
				p.money -= 50
				p.Acheter_Utilitaire()
			} else {
				fmt.Println()
				afficherTexte50("Vous n'avez pas l'argent pour m'acheter cela !")
				fmt.Println()
				p.Acheter_Utilitaire()
			}
			case 2:
			if p.InventairePlein() {
				afficherTexte50("Votre inventaire est plein ! Vous ne pouvez pas acheter d'objet.")
				p.Marchand()
			} else if p.money >= 50 {
				afficherTexte50("Vous avez acheté une potion de mana pour 50 écus")
				fmt.Println()
				p.inventaire[RessourcePotMana] += 1
				p.money -= 50
				p.Acheter_Utilitaire()
			} else {
				fmt.Println()
				afficherTexte50("Vous n'avez pas l'argent pour m'acheter cela !")
				fmt.Println()
				p.Acheter_Utilitaire()
			}
		case 3:
			if p.InventairePlein() {
				afficherTexte50("Votre inventaire est plein ! Vous ne pouvez pas acheter d'objet.")
				p.Marchand()
			} else if p.money >= 150 {
				afficherTexte50("Vous avez acheté une potion de poison pour 150 écus")
				fmt.Println()
				p.money -= 150
				p.inventaire[RessourcePotPoison] += 1
				p.Acheter_Utilitaire()
			} else {
				fmt.Println()
				afficherTexte50("Vous n'avez pas l'argent pour m'acheter cela !")
				fmt.Println()
				p.Acheter_Utilitaire()
			}
		case 4:
			if p.money >= 100  && inventaire_taillemax<40{
				afficherTexte50("Vous avez aggrandi votre inventaire pour 100 écus!")
				fmt.Println()
				p.money -= 100
				inventaire_taillemax += 10
				p.Acheter_Utilitaire()
			} else {
				fmt.Println()
				afficherTexte50("Désolé vous ne pouvez pas acheter ça pour le moment!")
				fmt.Println()
				p.Acheter_Utilitaire()
			}
		case 5 :
			p.Marchand()
		default : 
		afficherTexte50("Veuillez saisir une touche valide")
		}
	}
}

func (p *Character) Acheter_Sorts() {

	if sort_bdf == false {
		fmt.Println()
		afficherTexte50("Désolé je n'ai plus rien à vous proposer !")
		fmt.Println()
		p.Marchand()
	}
	if sort_coupv == false  {
		fmt.Println()
		afficherTexte50("Désolé je n'ai plus rien à vous proposer !")
		fmt.Println()
		p.Marchand()
	}
	if sort_cc == false {
		fmt.Println()
		afficherTexte50("Désolé je n'ai plus rien à vous proposer !")
		fmt.Println()
		p.Marchand()
	}

	if sort_cc && (p.SkillName == "Coup critique") {
		afficherTexte50("C : Coup critique : 300 écus")
	}
	if sort_bdf && (p.SkillName == "Boule de feu") {
		afficherTexte50("B : Boule de feu : 300 écus")
	}
	if sort_coupv && (p.SkillName == "Coups vicieux"){
		afficherTexte50("V : Coup vicieux : 300 écus")
	}
	var choix_sort string
	fmt.Scanln(&choix_sort)
	switch choix_sort {
	case "C" , "c" :
		if sort_cc == false {
			p.Marchand()
		} else {
		afficherTexte50("Vous avez acheté le sort Coup critique pour 300 écus!")
		Sort = true
		fmt.Println()
		sort_cc = false
		p.inventaire["Coup critique"] += 1
		p.money -= 300
		p.Marchand() }
	case "B" , "b":
		if sort_bdf == false {
			p.Marchand()
		} else {
		afficherTexte50("Vous avez acheté le sort Boule de feu pour 300 écus !")
		sort_bdf = false
		fmt.Println()
		Sort = true
		p.inventaire["Boule de feu"] += 1
		p.money -= 300
		p.Marchand() }
		case "V","v" : 
			if sort_coupv == false {
			p.Marchand()
		} else {
		afficherTexte50("Vous avez acheté le sort Coups vicieux pour 300 écus !")
		sort_coupv = false
		fmt.Println()
		Sort = true
		p.inventaire["Coups vicieux"] += 1
		p.money -= 300
		p.Marchand() }
	}
}

func (p *Character) Vendre() {
	fmt.Println("Que voulez vous vendre ? ")
	fmt.Println("1 :  Tissu de spectre: 10 écus")
	fmt.Println("2 :  Peau de géant: 15 écus")
	fmt.Println("3 :  Fil d'araignée: 15 écus")
	fmt.Println("4 :  Dent de loup: 20 écus")
	fmt.Println("5 :  Poil de mammouth: 25 écus")
	fmt.Println("6 :  Griffe de lynx fumée: 25 écus")
	fmt.Println("7 :  Cendre de lynx fumée: 40 écus")
	fmt.Println("8 :  Fragment de glace: 50 écus")
	fmt.Println("9 :  Fragment du Roi de la Nuit: 75 écus")
	fmt.Println("10 : Sabot de licorne: 0 écus")
	fmt.Println("R : Retour au menu ")

	var choix_vente int
	fmt.Scanln(&choix_vente)

	switch choix_vente {
	case 1:
		if p.inventaire[RessourceTissuDeSpectre] >= 1 {
			p.inventaire[RessourceTissuDeSpectre] -= 1
			p.money += 10
			fmt.Println("Vous avez vendu un tissu de spectre pour 10 écus")
			fmt.Println()
			p.Acheter_Materiaux()
		} else {
			fmt.Println("Vous n'avez pas l'objet à vendre escroc!")
			fmt.Println()
			p.Vendre()
		}
	case 2:
		if p.inventaire[RessourcePeauDeGeant] >= 1 {
			p.inventaire[RessourcePeauDeGeant] -= 1
			p.money += 15
			fmt.Println("Vous avez vendu une peau de géant pour 15 écus")
			fmt.Println()
			p.Acheter_Materiaux()
		} else {
			fmt.Println("Vous n'avez pas l'objet à vendre escroc!")
			fmt.Println()
			p.Vendre()
		}
	case 3:
		if p.inventaire[RessourceFilDaraignee] >= 1 {
			p.inventaire[RessourceFilDaraignee] -= 1
			p.money += 15
			fmt.Println("Vous avez vendu un fil d'araignée pour 15 écus")
			fmt.Println()
			p.Vendre()
		} else {
			fmt.Println("Vous n'avez pas l'objet à vendre escroc!")
			fmt.Println()
			p.Vendre()
		}
	case 4:
		if p.inventaire[RessourceDentdeloup] >= 1 {
			p.inventaire[RessourceDentdeloup] -= 1
			p.money += 20
			fmt.Println("Vous avez vendu une dent de loup pour 20 écus")
			fmt.Println()
			p.Vendre()
		} else {
			fmt.Println("Vous n'avez pas l'objet à vendre escroc!")
			fmt.Println()
			p.Vendre()
		}
	case 5:
		if p.inventaire[RessourcePoildemammouth] >= 1 {
			p.inventaire[RessourcePoildemammouth] -= 1
			p.money += 25
			fmt.Println("Vous avez vendu un poil de mammouth pour 25 écus")
			fmt.Println()
			p.Vendre()
		} else {
			fmt.Println("Vous n'avez pas l'objet à vendre escroc!")
			fmt.Println()
			p.Vendre()
		}
	case 6:
		if p.inventaire[RessourceGriffeDelynxfumee] >= 1 {
			p.inventaire[RessourceGriffeDelynxfumee] -= 1
			p.money += 25
			fmt.Println("Vous avez vendu une griffe de lynx fumée pour 25 écus")
			fmt.Println()
			p.Vendre()
		} else {
			fmt.Println("Vous n'avez pas l'objet à vendre escroc!")
			fmt.Println()
			p.Vendre()
		}
	case 7:
		if p.inventaire[RessourceCendreDelynxfumee] >= 1 {
			p.inventaire[RessourceCendreDelynxfumee] -= 1
			p.money += 40
			fmt.Println("Vous avez vendu de la cendre de lynx fumée pour 40 écus")
			fmt.Println()
			p.Vendre()
		} else {
			fmt.Println("Vous n'avez pas l'objet à vendre escroc!")
			fmt.Println()
			p.Vendre()
		}
	case 8:
		if p.inventaire[RessourceFragmentdeglace] >= 1 {
			p.inventaire[RessourceFragmentdeglace] -= 1
			p.money += 50
			fmt.Println("Vous avez acheté un fragment de glace")
			fmt.Println()
			p.Vendre()
		} else {
			fmt.Println("Vous n'avez pas l'objet à vendre escroc!")
			fmt.Println()
			p.Vendre()
		}
	case 9:
		if p.inventaire[RessourceFragmentRoiDeLaNuit] >= 1 {
			p.inventaire[RessourceFragmentRoiDeLaNuit] -= 1
			p.money += 75
			fmt.Println("Vous avez vendu un fragment du Roi de la Nuit pour 75 écus")
			fmt.Println()
			p.Vendre()
		} else {
			fmt.Println("Vous n'avez pas l'objet à vendre escroc!")
			fmt.Println()
			p.Vendre()
		}
	case 10:
		if p.inventaire[RessourceSabotDeLicorne] >= 1 {
			p.inventaire[RessourceSabotDeLicorne] -= 1
			p.money += 0
			fmt.Println("Merci pour ce cadeau jeune aventurier !")
			fmt.Println()
			p.Vendre()
		} else {
			fmt.Println("Vous n'avez pas l'objet à vendre escroc!")
			fmt.Println()
			p.Vendre()
		}
	}
}

func (p *Character) Marchand() {
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
		fmt.Println()
		fmt.Println("Que souhaites-tu acheter ?")
		fmt.Println()
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

