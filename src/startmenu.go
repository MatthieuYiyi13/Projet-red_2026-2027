package projet

import (
		"fmt"
		"time"
		"math/rand/v2"
)
var choix_classe bool = true
var choix_nom bool = true
var fin_du_jeu bool = true


func afficherTexte20(texte string) {
	for _, caractere := range texte {
		fmt.Print(string(caractere))
		time.Sleep(1 * time.Millisecond)
	}
	fmt.Println()
}

func afficherTexte1(format string, args ...interface{}) {
	texte := fmt.Sprintf(format, args...)

	for _, caractere := range texte {
		fmt.Print(string(caractere))
		time.Sleep(7 * time.Millisecond)
	}

	fmt.Println()
}

func afficherTexte50(format string, args ...interface{}) {
	texte := fmt.Sprintf(format, args...)

	for _, caractere := range texte {
		fmt.Print(string(caractere))
		time.Sleep(50 * time.Millisecond)
	}

	fmt.Println()
}

func afficherTexte100(format string, args ...interface{}) {
	texte := fmt.Sprintf(format, args...)

	for _, caractere := range texte {
		fmt.Print(string(caractere))
		time.Sleep(500 * time.Millisecond)
	}

	fmt.Println()
}




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
		 p.StartGame()
	}

}
func (p *Character) StartMenu() {

afficherTexte20("Depuis des siècles, le royaume d'Eldoria vivait en paix.")

afficherTexte20("Ses terres étaient divisées en plusieurs biomes, chacun abritant ses propres créatures, ses secrets et ses dangers.")

afficherTexte20("Au Sud-Ouest s'étendait la Forêt d'Émeraude, une immense forêt où les arbres semblaient murmurer aux voyageurs.")

afficherTexte20("Plus loin se trouvait le Désert des Cendres, une mer de sable brûlant où d'anciennes ruines étaient enfouies depuis des milliers d'années.")

afficherTexte20("À l'est, les Montagnes de Glaces formaient une gigantesque frontière naturelle. On racontait que des monstres y dormaient sous la glace.")

afficherTexte20("Au Nord-Est se trouvait le Volcan des Brumes, une région mystérieuse où peu de voyageurs osaient s'aventurer.")

afficherTexte20("Depuis quelques années, des créatures étranges apparaissaient dans les différentes régions. Des villages disparaissaient. Des voyageurs racontaient avoir aperçu une immense ombre dans le ciel.")

afficherTexte20("Puis, une nuit, les étoiles disparurent.")

afficherTexte20("Une voix résonna dans tout le royaume")

afficherTexte20("")

afficherTexte20("Le sceau est brisé... Celui qui portera le destin d'Eldoria devra choisir son chemin.")

afficherTexte20("Vous vous réveillez au milieu de la forêt devant les portes de la cité de Qarth.")

afficherTexte20("Devant vous se trouvent trois chemins.")

afficherTexte20("")

afficherTexte20("Le Guerrier :")

afficherTexte20("")

afficherTexte20("Maître du combat rapproché, le Guerrier utilise sa force et sa résistance pour affronter ses ennemis directement.")

afficherTexte20("")

afficherTexte20("Le Sorcier :")

afficherTexte20("")

afficherTexte20("Maître des arcanes, le Sorcier utilise la magie pour infliger de puissants dégâts et contrôler le champ de bataille.")

afficherTexte20("")

afficherTexte20("L'Assassin :")

afficherTexte20("")

afficherTexte20("Rapide et discret, l'Assassin préfère la ruse, les attaques rapides et les coups dans l'ombre.")
	for choix_classe {
	fmt.Println()
fmt.Println("Laquelle de ces 3 classes veux tu choisir ? ")
fmt.Println()
x  := "Temp"
var choice_classe string
fmt.Scanln(&choice_classe)
switch choice_classe {
case "3", "Assassin","assassin" : p.initcharacter(x,"Assassin")
					    choix_classe = false
						stopSound()
						PlaySoundAsyncDebut()
case "1" ,"Guerrier","guerrier": p.initcharacter(x,"Guerrier")
						choix_classe = false
						stopSound()
						PlaySoundAsyncDebut()
case "2", "Sorcier","sorcier": p.initcharacter(x,"Sorcier")
						choix_classe = false
						stopSound()
						PlaySoundAsyncDebut()
default : fmt.Println("Veuillez entrer une classe valide.")
}

}
	for choix_nom{
		fmt.Println()
		afficherTexte50("Très bien jeune %s, comment t'appelles tu ?", p.classe)
		fmt.Println()
		fmt.Println()
		var nom string
		fmt.Scanln(&nom)
		p.name = nom
		fmt.Println()
		afficherTexte50("Quel nom étrange...")
		fmt.Println()
		afficherTexte50("Bref j'espère que tu ne te perdras pas trop lors de ton aventure %s !", p.name)
		choix_nom = false 
		fmt.Println()
	}
		afficherTexte50("Vous vous réveillez dans une forêt à la fois silencieuse et sinistre.")
		if p.classe == "Guerrier" {
		afficherTexte50("Vous regardez autour de vous et n'y voyez qu'une épée un peu émoussée et un sac de pièces.")
		} else if p.classe == "Sorcier" {
		afficherTexte50("Vous regardez autour de vous et n'y voyez qu'un bâton un peu usé et un sac de pièces.")
		} else {
		afficherTexte50("Vous regardez autour de vous et n'y voyez qu'une dague un peu émoussée et un sac de pièces.")
		}
		afficherTexte100(". . . . . . . . . . .")
		afficherTexte1("BAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAM")
		afficherTexte50("Une araignée vous tombe dessus , elle cherche le combat")

	skiptuto = false 
	fuite = false
	victoire = false
	combat_pas_finis = true
	monstre := Araignee{
			name:        "Araignée",
			attack:      20,
			PV_max:      60,
			PV_actuelle: 60,
			Poison:      false,
	}
	fmt.Println()
	fmt.Printf("Vous rencontrez une %s sauvage !", monstre.name)
	fmt.Println()
	monstre.PV_actuelle -= 30
	for combat_pas_finis {
		skiptuto = false 
		if monstre.Poison {
			fmt.Println()
			fmt.Printf("Le monstre ennemi %s perds %d PV", monstre.name, 10)
			fmt.Println()
			monstre.PV_actuelle -= 10
			fmt.Println()
			if monstre.PV_actuelle <= 0 {
				combat_pas_finis = false
				victoire = true
				continue
			}
		}
		fmt.Println()
		fmt.Printf("Les pv du monstre %s sont de %d/%d ", monstre.name , monstre.PV_actuelle, monstre.PV_max)
		fmt.Println()
		fmt.Println("Que voulez vous faire ?")
		fmt.Printf("\n")
		fmt.Printf("\t+------------------------------------------------+\n")
		fmt.Printf("\t|                                                |\n")
		fmt.Printf("\t|            Que voulez vous faire ?             |\n")
		fmt.Printf("\t|                                                |\n")
		fmt.Printf("\t|                 [A] Attaque                    |\n")
		fmt.Printf("\t|                                                |\n")
		fmt.Printf("\t+------------------------------------------------+\n")
		fmt.Println()
		var moove string
		fmt.Scanln(&moove)
		switch moove { 
		case "A" , "a" : {
			p.AttaqueTuto((*Monstre)(&monstre))
			if skiptuto {continue
			}
		
		}
		default: fmt.Println("Commande invalide.") 
		continue
	}
		if monstre.PV_actuelle <= 0 {
			combat_pas_finis = false
			victoire = true
			continue
		}
		fmt.Println()
		l := rand.IntN(4) 
		if l == 0 {
			fmt.Println()
			afficherTexte20("Le monstre vous attaque mais vous réussissez à l'esquiver !")
			fmt.Println()
		} else {
		afficherTexte20("Le monstre ennemi est enervé , il vous charge")
		pv_perdu := monstre.attack
		if count_super_monstre%4 == 0 {
			pv_perdu = pv_perdu*2
		}
		fmt.Printf("Vous perdez %d Pvs\n",pv_perdu)
		p.pv -= monstre.attack
		fmt.Printf("Pvs actuelle : %d/%d",p.pv,p.pvmax)
		fmt.Println()
		count_super_monstre ++
	}
		if p.pv <= 0 {
			victoire = false 
			combat_pas_finis = false 
			continue 
		}
	}
	if victoire && fuite == false {
	fmt.Println()
	afficherTexte20("Vous avez gagné votre premier combat")
	fmt.Println()
	p.GagnerCombat()
	victoire = false
	combat_pas_finis = true
	} else if victoire && fuite{
		fmt.Println()
	fmt.Println()
	} else {
	fmt.Println()
	fmt.Println("Vous avez perdu tout vos pvs , vous êtes mort !")
	fmt.Println()
	victoire = false 
	combat_pas_finis = true
	p.IsDead()
}
p.pv = p.pvmax
	for fin_du_jeu {
	fmt.Println()
	fmt.Printf("\t+------------------------------------------------+\n")
	fmt.Printf("\t|               Menu Principal                   |\n")
	fmt.Printf("\t|                                                |\n")
	fmt.Printf("\t|  [P] Afficher les informations du personnage   |\n")
	fmt.Printf("\t|  [I] Accéder à l'inventaire                    |\n")
	fmt.Printf("\t|  [M] Marchand                                  |\n")
	fmt.Printf("\t|  [F] Forgeron                                  |\n")
	fmt.Printf("\t|  [Map] Afficher la map                         |\n")
	fmt.Printf("\t|                                                |\n")
	fmt.Printf("\t|                                                |\n")
	fmt.Printf("\t|  [O] Options                                   |\n")
	fmt.Printf("\t|  [Q] Quitter le jeu                            |\n")
	fmt.Printf("\t|                                                |\n")
	fmt.Printf("\t+------------------------------------------------+\n")
		var choice string
		fmt.Scanln(&choice)
		switch choice {
		case "P" , "p" : 
			fmt.Println("=== stats du personnage ===")

		case "I", "i" : 
			fmt.Println("=== inventaire ===")
			p.AccessInventory()
	
		case "M","m" :
			fmt.Println("Bonjour jeune aventurier, je vois que tu as réussi à me trouver dans cette magnifique ville de Qarth !")
			p.Marchand()
		
		case "F","f" :
			fmt.Println("Bienvenue dans ma forge")
			p.ForgeronMenu()
		case "Map","map" : 
			p.AfficheMapAscii()
	case "Q" , "q":
			fmt.Println("Au revoir !")
			fin_du_jeu = false
	case "pos" :
			fmt.Printf("Vous êtes actuellement en %d, %d\n", x_position, y_position)
	case "co" : 
			p.AfficheMapCo()
	case "combat" :
			p.Combat_start_premier()
	case "train" : 
			p.TrainingFight()
	default : 
	fmt.Println()
	fmt.Println("Veuillez saisir une touche valide.")
	fmt.Println()
	}
	
	}
}
