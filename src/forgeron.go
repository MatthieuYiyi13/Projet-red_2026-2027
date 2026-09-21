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
	

	