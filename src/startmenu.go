package projet

import "fmt"

var choix_classe bool = true

func (p *Character) StartGame() {

	vert := "\033[32m"
	reset := "\033[0m"

	fmt.Println(vert + `
██╗    ██╗███████╗██╗      ██████╗ ██████╗ ███╗   ███╗███████╗
██║    ██║██╔════╝██║     ██╔════╝ ██╔══██╗████╗ ████║██╔════╝
██║ █╗ ██║█████╗  ██║     ██║     ██║  ██║██╔████╔██║█████╗
██║███╗██║██╔══╝  ██║     ██║     ██║  ██║██║╚██╔╝██║██╔══╝
╚███╔███╔╝███████╗███████╗╚██████╗ ██████╔╝██║ ╚═╝ ██║███████╗
 ╚══╝╚══╝ ╚══════╝╚══════╝ ╚═════╝ ╚═════╝ ╚═╝     ╚═╝╚══════╝
` + reset)
	fmt.Println()
	fmt.Printf("\t+------------------------------------------------+\n")
	fmt.Printf("\t|               Menu Principal                   |\n")
	fmt.Printf("\t|                                                |\n")
	fmt.Printf("\t|  [N] Nouvelle partie                           |\n")
	fmt.Printf("\t|                                                |\n")
	fmt.Printf("\t|  [Q] Quitter le jeu                            |\n")
	fmt.Printf("\t|                                                |\n")
	fmt.Printf("\t+------------------------------------------------+\n")
	var choice_jeu string
	fmt.Scanln(&choice_jeu)
	switch choice_jeu {
	case "N", "n":
		p.StartMenu()
	case "Q", "q":
		break
	default:
		fmt.Println()
		fmt.Println("Saisie invalide , veuillez recommencer.")
	}

}
func (p *Character) StartMenu() {

fmt.Println()
fmt.Println("Depuis des siècles, le royaume d'Eldoria vivait en paix.")
fmt.Println("Ses terres étaient divisées en plusieurs biomes, chacun abritant ses propres créatures, ses secrets et ses dangers.")
fmt.Println("Au Sud-Ouest s'étendait la Forêt d'Émeraude, une immense forêt où les arbres semblaient murmurer aux voyageurs.")
fmt.Println("Plus loin se trouvait le Désert des Cendres, une mer de sable brûlant où d'anciennes ruines étaient enfouies depuis des milliers d'années.")
fmt.Println("À l'est, les Montagnes de Glaces formaient une gigantesque frontière naturelle. On racontait que des monstres y dormaient sous la glace.")
fmt.Println("Au Nord-Est se trouvait le Volcan des Brumes, une région mystérieuse où peu de voyageurs osaient s'aventurer.")
fmt.Println("Depuis quelques années, des créatures étranges apparaissaient dans les différentes régions. Des villages disparaissaient. Des voyageurs racontaient avoir aperçu une immense ombre dans le ciel.")
fmt.Println("Puis, une nuit, les étoiles disparurent.")
fmt.Println("Une voix résonna dans tout le royaume")
fmt.Println()
fmt.Println("Le sceau est brisé... Celui qui portera le destin d'Eldoria devra choisir son chemin.")
fmt.Println("Vous vous réveillez au milieu de la fôret devant les portes de la cité de Qarth.")
fmt.Println("Devant vous se trouvent trois chemins.")
fmt.Println()
fmt.Println("Le Guerrier : ")
fmt.Println()
fmt.Println("Maître du combat rapproché, le Guerrier utilise sa force et sa résistance pour affronter ses ennemis directement.")
fmt.Println()
fmt.Println("Le Sorcier : ")
fmt.Println()
fmt.Println("Maître des arcanes, le Sorcier utilise la magie pour infliger de puissants dégâts et contrôler le champ de bataille.")
fmt.Println()
fmt.Println("L'assassin : ")
fmt.Println()
fmt.Println("Rapide et discret, l'assassin préfère la ruse, les attaques rapides et les coups dans l'ombre.")
for choix_classe {
	fmt.Println()
fmt.Println("Laquelle de ces 3 classes veux tu choisir ? ")
fmt.Println()
x  := "Temp"
var choice_classe string
fmt.Scanln(&choice_classe)
switch choice_classe {
case "Assassin","assassin" : p.initcharacter(x,"Assassin")
					    choix_classe = false
case "Guerrier","guerrier": p.initcharacter(x,"Guerrier")
						choix_classe = false
case "Sorcier","sorcier": p.initcharacter(x,"Sorcier")
						choix_classe = false
default : fmt.Println("Veuillez entrer une classe valide.")
}
}
	for true {
		stopSound()
		PlaySoundAsyncDebut()
		fmt.Println("=== Menu principal ===")
		fmt.Println("[1] Afficher les informations du personnage")
		fmt.Println("[2] Accéder à l'inventaire")
		fmt.Println("[3] Marchand")
		fmt.Println("[4] Forgeron")
		fmt.Println("[Q] Quitter le jeu")
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
			fmt.Println("Bonjour jeune aventurier, je vois que tu as réussi à me trouver dans cette magnifique ville de Qarth !")
			p.Marchand()
		}
		if choice == "4" {
			fmt.Println("Bienvenue dans ma forge")
			p.ForgeronMenu()
		}
		if choice == "map" {
			p.AfficheMap()
		}
			if choice == "Q" {
			fmt.Println("Au revoir !")
			break
		}
		if choice == "pos" {
			fmt.Printf("Vous êtes actuellement en %d, %d\n", x_position, y_position)
		}
		if choice == "combat" {
			p.Combat_start_premier()
		}
		if choice == "train" {
			p.TrainingFight()
		}
	}
}
