package projet

import "fmt"

func (p *Character) startMenu() { 
	var player Character
	player.initcharacter("Davy", "guerrier")
	for true {
		fmt.Println("=== Menu ===")
		fmt.Println("1. Afficher les informations du personnage")
		fmt.Println("2. Accéder à l'inventaire")
		fmt.Println("3. Quitter le jeu")
		var choice int
		fmt.Scanln(&choice)}
}
