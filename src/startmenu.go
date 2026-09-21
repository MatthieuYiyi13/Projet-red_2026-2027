	package projet

	import "fmt"

	func (p *Character)StartMenu() { 
		p.initcharacter("davy", "guerrier")
		for true {
			fmt.Println("=== Menu principal ===")
			fmt.Println("1. Afficher les informations du personnage")
			fmt.Println("2. Accéder à l'inventaire")
			fmt.Println("3. forgeron")
			fmt.Println("4. Quitter le jeu")
			var choice int
			fmt.Scanln(&choice)
			if choice == 1 {
				fmt.Println("=== stats du personnage ===")}
			if choice == 2 {
				fmt.Println("=== inventaire ===")
				p.AccessInventory() }	
			if choice == 3 {
				fmt.Println("Au revoir !")
				break}
			if choice == 4 {
				p.Marchand()
		}
	}
<<<<<<< HEAD
	}
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
=======
}
>>>>>>> de21e9e189f6dc42ed50b90f031e1048c07832be
