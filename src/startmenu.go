package projet

import "fmt"

func (p *Character) StartMenu() {
<<<<<<< HEAD
	p.initcharacter("davy", "Sorcier")
=======
	vert := "\033[32m"
	reset := "\033[0m"

	fmt.Println(vert + `
██     ██ ███████ ██      ██████  ██████  ███    ███ ███████
██     ██ ██      ██     ██      ██    ██ ████  ████ ██
██  █  ██ █████   ██     ██      ██    ██ ██ ████ ██ █████
██ ███ ██ ██      ██     ██      ██    ██ ██  ██  ██ ██
 ███ ███  ███████ ███████ ██████  ██████  ██      ██ ███████
` + reset)
	p.initcharacter("davy", "guerrier")
>>>>>>> 0e06a701e638dc452e6c53b5d29d6f740e00d8df
	for true {
		fmt.Println("=== Menu principal ===")
		fmt.Println("1.  Afficher les informations du personnage")
		fmt.Println("2.  Accéder à l'inventaire")
		fmt.Println("3.  Marchand")
		fmt.Println("4.  Forgeron")
		fmt.Println("10. Quitter le jeu")
		var choice string
		fmt.Scanln(&choice)
		if choice == "1" {
			fmt.Println("=== stats du personnage ===")
		}
		if choice == "2" {
			fmt.Println("=== inventaire ===")
			p.AccessInventory()
		}
		if choice == "3" {
			fmt.Println("Bienvenue dans ma forge")
			p.Marchand()
		}
		if choice == "10" {
			fmt.Println("Au revoir !")
			break
		}
		if choice == "4" {
			p.ForgeronMenu()
		}
		if choice == "map" {
			p.AfficheMap()
		}
		if choice == "pos" {
			fmt.Printf("Vous êtes actuellement en %d, %d\n", x_position, y_position)
		}
		if choice == "combat" {
			p.Combat_start_premier()
		}
	}
}
