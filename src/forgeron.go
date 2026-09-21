package projet

import "fmt"

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
			if p.inventaire["Peau de géant"] >= 2 && p.inventaire["Tissu de spectre"] >= 2 {
				p.inventaire["Peau de géant"] -= 2
				p.inventaire["Tissu de spectre"] -= 2
				p.inventaire["Chapeau de Géant"]++
				fmt.Println("Vous obtenez le Chapeau de Géant ! ")
			} else {
				fmt.Println("Il te faut 2 peau de Géant et 2 tissus de spectre pour fabriquer cela !")
			}
		case 2:
			if p.inventaire["Poil de mammouth"] >= 2 && p.inventaire["Fil d'araignée"] >= 1 {
				p.inventaire["Poil de mammouth"] -= 2
				p.inventaire["Fil d'araignée"]--
				p.inventaire["Armure en cuir"]++
				fmt.Println("Vous obtenez l'armure en cuir de mammouth !")
			} else {
				fmt.Println("Il te faut 2 Poil de mammouth et 1 fil d'araignée pour fabriquer cela !")
			}
		case 3:
			if p.inventaire["Dent de loup"] >= 3 && p.inventaire["Griffe de lynx fumée"] >= 3 && p.inventaire["Fil d'araignée"] >= 2 {
				p.inventaire["Dent de loup"] -= 3
				p.inventaire["riffe de lynx fumée"] -= 3
            	p.inventaire["Fil d'araignée"] -= 2
				p.inventaire["Armure en ivoire"]++
				fmt.Println("Tu as fabriqué l'Armure en ivoire !")
			} else {
				fmt.Println("Il te faut 3 Dents de loup, 3 Griffes de lynx fumée et 2 fils d'araignée pour fabriquer cela !")
			}
            case 4:
			if p.inventaire["Sabot de licorne"] >= 2 && p.inventaire["Cendre de lynx fumée"] >= 2 {
				p.inventaire["Sabot de licorne"] -= 2
	 			p.inventaire["Cendre de lynx fumée"] -= 2
				p.inventaire["Bottes Arc-En-Ciel"]++
				fmt.Println("Tu as fabriqué les bottes Arc-En-Ciel !")
			} else {
				fmt.Println("Il te faut 2 sabots de licorne et 2 cendres de lynx fumée pour fabriquer cela!")
			}
		case 5:
			return
		default:
			fmt.Println("Choix invalide, choisis 1, 2, 3 ou 4.")
		}
	}
}
