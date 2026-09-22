package projet

import "fmt"

var chapeau_fabrique bool = true
var armurecuir_fabrique bool = true
var armureivoir_fabrique bool = true
var bottes_fabrique bool = true

func (p *Character) ForgeronMenu() {
	for {
		fmt.Println("=== Forgeron ===")
		fmt.Println("1. Chapeau de Géant")
		fmt.Println("2. Armure en cuir de mammouth")
		fmt.Println("3. Armure en ivoire")
		fmt.Println("4. Bottes Arc-En-Ciel")
		fmt.Println("5. Retour au menu principal")

		var choice int
		fmt.Scanln(&choice)

		switch choice {
		case 1:
			if p.inventaire[RessourcePeauDeGeant] >= 2 && p.inventaire[RessourceTissuDeSpectre] >= 2 && chapeau_fabrique {
				p.inventaire[RessourcePeauDeGeant] -= 2
				p.inventaire[RessourceTissuDeSpectre] -= 2
				chapeau_fabrique = false
				p.inventaire[RessourceChapeaudegeant]++
				fmt.Println("Vous obtenez le Chapeau de Géant ! ")
			} else if chapeau_fabrique {
				fmt.Println("Il te faut 2 peaux de Géant et 2 tissus de spectre pour fabriquer cela !")
			} else {
				fmt.Println()
				fmt.Println("Vous avez déjà fabriqué cette objet !")
				fmt.Println()
			}
		case 2:
			if p.inventaire[RessourcePoildemammouth] >= 2 && p.inventaire[RessourceFilDaraignee] >= 1 && armurecuir_fabrique {
				p.inventaire[RessourcePoildemammouth] -= 2
				p.inventaire[RessourceFilDaraignee]--
				armurecuir_fabrique = false
				p.inventaire[RessourceArmurecuir]++
				fmt.Println("Vous obtenez l'armure en cuir de mammouth !")
			} else if armurecuir_fabrique {
				fmt.Println("Il te faut 2 Poil de mammouth et 1 fil d'araignée pour fabriquer cela !")
			} else {
				fmt.Println()
				fmt.Println("Vous avez déjà fabriqué cette objet !")
				fmt.Println()
			}
		case 3:
			if p.inventaire[RessourceDentdeloup] >= 3 && p.inventaire[RessourceGriffeDelynxfumee] >= 3 && p.inventaire[RessourceFilDaraignee] >= 2 && armureivoir_fabrique {
				p.inventaire[RessourceDentdeloup] -= 3
				p.inventaire[RessourceGriffeDelynxfumee] -= 3
				p.inventaire[RessourceFilDaraignee] -= 2
				armureivoir_fabrique = false
				p.inventaire[RessourceArmureIvoire]++
				fmt.Println("Tu as fabriqué l'Armure en ivoire !")
			} else if armureivoir_fabrique {
				fmt.Println("Il te faut 3 Dents de loup, 3 Griffes de lynx fumée et 2 fils d'araignée pour fabriquer cela !")
			} else {
				fmt.Println()
				fmt.Println("Vous avez déjà fabriqué cette objet !")
				fmt.Println()
			}
		case 4:
			if p.inventaire[RessourceSabotDeLicorne] >= 2 && p.inventaire[RessourceCendreDelynxfumee] >= 2 && botte_equipe {
				p.inventaire[RessourceSabotDeLicorne] -= 2
				p.inventaire[RessourceCendreDelynxfumee] -= 2
				bottes_fabrique = false
				p.inventaire[RessourceBotteArcenciel]++
				fmt.Println("Tu as fabriqué les bottes Arc-En-Ciel !")
			} else if bottes_fabrique {
				fmt.Println("Il te faut 2 sabots de licorne et 2 cendres de lynx fumée pour fabriquer cela!")
			} else {
				fmt.Println()
				fmt.Println("Vous avez déjà fabriqué cette objet !")
				fmt.Println()
			}
		case 5:
			return
		default:
			fmt.Println("Commande invalide.")
		}
	}
}
