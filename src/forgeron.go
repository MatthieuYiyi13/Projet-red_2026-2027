<<<<<<< HEAD
package projet
=======
package projet

import "fmt"

func (p *Character) forgeronMenu() {
	for {
		fmt.Println("=== Forgeron ===")
		fmt.Println("1. Chapeau de l'aventurier")
		fmt.Println("2. Tunique de l'aventurier")
		fmt.Println("3. Bottes de l'aventurier")
		fmt.Println("4. Retour au menu principal")

		var choice int
		fmt.Scanln(&choice)

		switch choice {
		case 1:
			if p.inventaire["plume de corbeau"] >= 1 && p.inventaire["cuir de sanglier"] >= 1 {
				p.inventaire["plume de corbeau"]--
				p.inventaire["cuir de sanglier"]--
				p.inventaire["chapeau de l'aventurier"]++
				fmt.Println("Tu as fabriqué le chapeau de l'aventurier !")
			} else {
				fmt.Println("Il te faut 1 plume de corbeau et 1 cuir de sanglier.")
			}
		case 2:
			if p.inventaire["fourrure de loup"] >= 2 && p.inventaire["peau de troll"] >= 1 {
				p.inventaire["fourrure de loup"] -= 2
				p.inventaire["peau de troll"]--
				p.inventaire["tunique de l'aventurier"]++
				fmt.Println("Tu as fabriqué la tunique de l'aventurier !")
			} else {
				fmt.Println("Il te faut 2 fourrures de loup et 1 peau de troll.")
			}
		case 3:
			if p.inventaire["cuir de sanglier"] >= 1 && p.inventaire["fourrure de loup"] >= 1 {
				p.inventaire["cuir de sanglier"]--
				p.inventaire["fourrure de loup"]--
				p.inventaire["bottes de l'aventurier"]++
				fmt.Println("Tu as fabriqué les bottes de l'aventurier !")
			} else {
				fmt.Println("Il te faut 1 cuir de sanglier et 1 fourrure de loup.")
			}
		case 4:
			return
		default:
			fmt.Println("Choix invalide, choisis 1, 2, 3 ou 4.")
		}
	}
}
>>>>>>> 6b384c7b3be9a85f1bd3eecad4918905b4d830b9
