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
            fmt.Println("fabrication chapeau.")
        case 2:
            fmt.Println("fabrication de la tunique.")
        case 3:
            fmt.Println("fabrication bottes.")
        case 4:
            return
        default:
            fmt.Println("Choix invalide, choisis 1, 2 ou 3 jeune aventurier.")
        }
    }
}	
func parcour_inventaire(p.inventaire []string, wanted string) bool {
    for _, item := range p.inventaire {
        if item == wanted {
            return true
        }
    }
    return false
}	
case 1:
    fmt.Println("Pour fabriquer ce chapeau, il te faut 1 plume de corbeau et 1 cuir de sanglier.")

    plume := parcour_inventaire(p, "plume de corbeau")
    cuir := parcour_inventaire(p, "cuir de sanglier")

    switch {
    case plume && cuir:
        fmt.Println("Tu as les deux matériaux : tu peux fabriquer le chapeau !")
    case plume:
        fmt.Println("Il te manque le cuir de sanglier.")
    case cuir:
        fmt.Println("Il te manque la plume de corbeau.")
    default:
        fmt.Println("Avec quoi veux-tu fabriquer ton chapeau ? Tu n'as rien dans ton sac, clochard.")
	}

case 2:
		fmt.Println("Pour fabriquer cette tunique, il te faut 2 fourrure de loup et 1 peau de troll.")
		
		fourrure := parcour_inventaire(p, "fourrure de loup")
		peau := parcour_inventaire(p, "peau de troll")

		switch {
		case fourrure && peau:
			fmt.Println("Tu as les deux matériaux : tu peux fabriquer la tunique !")
		case fourrure:
			fmt.Println("Il te manque la peau de troll.")
		case peau:
			fmt.Println("Il te manque la fourrure de loup.")
		default:
			fmt.Println("Avec quoi veux-tu fabriquer ta tunique ? Tu n'as rien dans ton sac, clochard.")
		}
case 3:
	fmt.Println("Pour fabriquer ces bottes, il te faut 1 cuir de sanglier et 1 fourrure de loup.")
	
	cuir := parcour_inventaire(p, "cuir de sanglier")
	fourrure := parcour_inventaire(p, "fourrure de loup")

	switch {
		case cuir && fourrure:
			fmt.Println("Tu as les deux matériaux : tu peux fabriquer les bottes !")
		case cuir:
			fmt.Println("Il te manque la fourrure de loup.")
		case fourrure:
			fmt.Println("Il te manque le cuir de sanglier.")
		default:
			fmt.Println("Avec quoi veux-tu fabriquer tes bottes ? Tu n'as rien dans ton sac, clochard.")
		}

func removeItem(p.inventaire, wanted string) []string {
    for i, item := range p.inventaire {
        if item == wanted {
            return append(inventory[:i], inventory[i+1:]...)
        }
    }
    return 
}

case plume && cuir:
    p.inventaire = removeItem(p.inventaire, "plume de corbeau")
    p.inventaire = removeItem(p.inventaire, "cuir de sanglier")
    p.inventaire = append(p.inventaire, "chapeau de l'aventurier")

    fmt.Println("Tu as fabriqué le chapeau de l'aventurier !")

case fourrure && peau:
	p.inventaire = removeItem(p.inventaire, "fourrure de loup")
	p.inventaire = removeItem(p.inventaire, "peau de troll")
	p.inventaire = append(p.inventaire, "tunique de l'aventurier")

case cuir && fourrure:
	p.inventaire = removeItem(p.inventaire, "cuir de sanglier")
	p.inventaire = removeItem(p.inventaire, "fourrure de loup")
	p.inventaire = append(p.inventaire, "bottes de l'aventurier")	