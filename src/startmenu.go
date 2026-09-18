package projet

import "fmt"

func (p *Character) startMenu() { 
	var player Character
	player.initcharacter("davy", "guerrier")
	for true {
		fmt.Println("=== Menu principal ===")
		fmt.Println("1. Afficher les informations du personnage")
		fmt.Println("2. Accéder à l'inventaire")
		fmt.Println("3. Quitter le jeu")

		var choice int

		fmt.Scanln(&choice)}
		if choice == 1 {
			fmt.Println("=== stats du personnage ===")}
		if choice == 2 {
			fmt.Println("=== inventaire ===")}	
		if choice == 3 {
			fmt.Println("Au revoir !")
			break
		}	

}
