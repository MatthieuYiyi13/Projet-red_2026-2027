package projet

import "fmt"

func (p *Character) StartMenu() {
	p.initcharacter("davy", "guerrier")
	for true {
		fmt.Println("=== Menu principal ===")
		fmt.Println("1. Afficher les informations du personnage")
		fmt.Println("2. Accéder à l'inventaire")
		fmt.Println("3. Forgeron")
		fmt.Println("4. Marchand")
		fmt.Println("10. Quitter le jeu")
		var choice int
		fmt.Scanln(&choice)
		if choice == 1 {
			fmt.Println("=== stats du personnage ===")
		}
		if choice == 2 {
			fmt.Println("=== inventaire ===")
			p.AccessInventory()
		}
		if choice == 10 {
			fmt.Println("Au revoir !")
			break
		}
		if choice == 4 {
			p.Marchand()
		}
	}
}

