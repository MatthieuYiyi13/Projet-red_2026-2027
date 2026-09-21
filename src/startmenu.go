package projet

import "fmt"

<<<<<<< HEAD
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
			if choice == 4 
				p.Marchand()
		}
	}

=======
func (p *Character) StartMenu() {
	p.initcharacter("davy", "guerrier")
	for true {
		fmt.Println("=== Menu principal ===")
		fmt.Println("1. Afficher les informations du personnage")
		fmt.Println("2. Accéder à l'inventaire")
		fmt.Println("3. forgeron")
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
>>>>>>> 0d3d8302591b03a74139d529aef3f3b5e151075b
