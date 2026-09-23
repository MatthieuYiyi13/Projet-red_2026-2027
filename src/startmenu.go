package projet

import (
		"fmt"
		"time"
		"math/rand/v2"
)

var dans_ville bool = false 
var choix_classe bool = true
var choix_nom bool = true
var fin_du_jeu bool = true

type direction struct {
    Haut   string
    Bas    string
    Gauche string
    Droite string
}

var directionMap = direction{
    Haut:   "H",
    Bas:    "B",
    Gauche: "G",
    Droite: "D",
}

var touchemap string = "L"
var toucheInv string = "I"
var touchePersoInfo string = "P"
var toucheMarchand string = "M"
var ToucheForgeron string = "F"
var ToucheOption string = "O"
var ToucheQuitter string = "Q"
var ToucheGuilde string = "G"
var ToucheEntrainement string = "E"


func (p *Character) touchevalide(touche string)bool {
	if len(touche) != 1 {
		return false 
	}
	if touche == toucheInv || touche == toucheMarchand || touche == touchePersoInfo || touche == touchemap || touche == ToucheForgeron || touche==ToucheQuitter || touche ==ToucheOption || touche == directionMap.Haut || touche == directionMap.Bas || touche == directionMap.Gauche || touche == directionMap.Droite{
		return false
	}else {
	return true 
	} 
}




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
		nomV := true
		for _, r := range nom {
			if !((r >= 'a' && r <= 'z') || (r >= 'A' && r <= 'Z')) {
				nomV = false
				break
			}
			if len(nom) > 0 && nom[0] >= 'a' && nom[0] <= 'z' {
				nom = string(nom[0]-'a'+'A') + nom[1:]
			}
		}
		if !nomV {
			afficherTexte50("Nom invalide, veuillez réessayer.")
			fmt.Println()
			continue
		}
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
		afficherTexte100(". . . . . . . .")
		afficherTexte1("BAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAM")
		afficherTexte50("Une araignée vous tombe dessus , elle cherche le combat.")

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
		fmt.Printf("\t\t+------------------------------------------------+\n")
		fmt.Printf("\t\t|                                                |\n")
		fmt.Printf("\t\t|            Que voulez vous faire ?             |\n")
		fmt.Printf("\t\t|                                                |\n")
		fmt.Printf("\t\t|                 [A] Attaque                    |\n")
		fmt.Printf("\t\t|                                                |\n")
		fmt.Printf("\t\t+------------------------------------------------+\n")
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
 if dans_ville{
	fmt.Println()
	fmt.Printf("\t\t+-------------------------------------------------+\n")
	fmt.Printf("\t\t|               Menu Principal                    |\n")
	fmt.Printf("\t\t|                                                 |\n")
	fmt.Printf("\t\t|  [%s] Afficher les informations du personnage   |\n",touchePersoInfo)
	fmt.Printf("\t\t|  [%s] Accéder à l'inventaire                    |\n",toucheInv)
	fmt.Printf("\t\t|  [%s] Marchand                                  |\n",toucheMarchand)
	fmt.Printf("\t\t|  [%s] Forgeron                                  |\n",ToucheForgeron)
	fmt.Printf("\t\t|  [%s] Afficher la map                           |\n",touchemap)
	fmt.Printf("\t\t|  [%s] Guilde                                    |\n",ToucheGuilde)
	fmt.Printf("\t\t|  [%s] Entrainement                              |\n",ToucheEntrainement)
	fmt.Printf("\t\t|                                                 |\n")
	fmt.Printf("\t\t|                    [%s]↑                         |\n",directionMap.Haut)
	fmt.Printf("\t\t|         ←[%s]                    [%s]→            |\n",directionMap.Gauche,directionMap.Droite)
	fmt.Printf("\t\t|                    [%s]↓                         |\n",directionMap.Bas)
	fmt.Printf("\t\t|                                                 |\n")
	fmt.Printf("\t\t|  [%s] Options                                   |\n",ToucheOption)
	fmt.Printf("\t\t|  [%s] Quitter le jeu                            |\n",ToucheQuitter)
	fmt.Printf("\t\t|                                                 |\n")
	fmt.Printf("\t\t+-------------------------------------------------+\n")
		var choice string
		fmt.Scanln(&choice)
		switch choice {
		case touchePersoInfo : 
			fmt.Println("=== stats du personnage ===")

		case toucheInv : 
			fmt.Println("=== inventaire ===")
			p.AccessInventory()
	
		case toucheMarchand :
			fmt.Println("Bonjour jeune aventurier, je vois que tu as réussi à me trouver dans cette magnifique ville de Qarth !")
			p.Marchand()
		
		case ToucheForgeron :
			fmt.Println("Bienvenue dans ma forge")
			p.ForgeronMenu()
		case touchemap : 
			p.AfficheMapCo()
		case ToucheGuilde :

	case ToucheEntrainement : 
			p.TrainingFight()
	case directionMap.Haut :
		y_position ++
	case directionMap.Bas :
		y_position--
	case directionMap.Gauche :
		x_position--
	case directionMap.Droite :
		x_position++
	case ToucheOption :
		fmt.Println()
		fmt.Println("Touches actuelle :")
		fmt.Printf("Haut   : %s \n",directionMap.Haut)
		fmt.Printf("Bas    : %s \n",directionMap.Bas)
		fmt.Printf("Gauche : %s \n",directionMap.Gauche)
		fmt.Printf("Droite : %s \n",directionMap.Droite)
		fmt.Printf("Information du personnage : %s \n", touchePersoInfo)
		fmt.Printf("Inventaire  : %s \n", toucheInv)
		fmt.Printf("Map         : %s \n" ,touchemap)
		fmt.Printf("Options     : %s \n",ToucheOption)
		fmt.Printf("Quitter     : %s \n",ToucheQuitter)
		fmt.Printf("Marchand    : %s \n",toucheMarchand)
		fmt.Printf("Forgeron    : %s \n",ToucheForgeron)
		fmt.Printf("Entrainement: %s \n",ToucheEntrainement)
		fmt.Printf("Guilde      : %s \n",ToucheGuilde)
			case "M", "m":
				fmt.Println("Bonjour jeune aventurier, je vois que tu as réussi à me trouver dans cette magnifique ville de Qarth !")
				p.Marchand()
				fmt.Println()
				fmt.Println("Quelle touche voulez vous changer ?")
				var touche string
				fmt.Scanln(&touche)
				switch touche {
				case "Haut", "haut", "1":
					fmt.Println("Par quelle touche voulez vous la remplacer ?")
					var touche_replace string
					fmt.Scanln(&touche_replace)
					if p.touchevalide(touche_replace) {
						directionMap.Haut = touche_replace
					} else {
						fmt.Println("Changement impossible, touche déjà attribué.")
					}
				case "Bas", "bas", "2":
					fmt.Println("Par quelle touche voulez vous la remplacer ?")
					var touche_replace string
					fmt.Scanln(&touche_replace)
					if p.touchevalide(touche_replace) {
						directionMap.Bas = touche_replace
					} else {
						fmt.Println("Changement impossible, touche déjà attribué.")
					}
				case "Gauche", "gauche", "3":
					fmt.Println("Par quelle touche voulez vous la remplacer ?")
					var touche_replace string
					fmt.Scanln(&touche_replace)
					if p.touchevalide(touche_replace) {
						directionMap.Gauche = touche_replace
					} else {
						fmt.Println("Changement impossible, touche déjà attribué.")
					}
				case "Droite", "droite", "4":
					fmt.Println("Par quelle touche voulez vous la remplacer ?")
					var touche_replace string
					fmt.Scanln(&touche_replace)
					if p.touchevalide(touche_replace) {
						directionMap.Droite = touche_replace
					} else {
						fmt.Println("Changement impossible, touche déjà attribué.")
					}
				case "5", "Information du personnage", "Personnage", "information du personnage", "perso":
					fmt.Println("Par quelle touche voulez vous la remplacer ?")
					var touche_replace string
					fmt.Scanln(&touche_replace)
					if p.touchevalide(touche_replace) {
						touchePersoInfo = touche_replace
					} else {
						fmt.Println("Changement impossible, touche déjà attribué.")
					}
				case "6", "Inventaire", "inv", "inventaire":
					fmt.Println("Par quelle touche voulez vous la remplacer ?")
					var touche_replace string
					fmt.Scanln(&touche_replace)
					if p.touchevalide(touche_replace) {
						toucheInv = touche_replace
					} else {
						fmt.Println("Changement impossible, touche déjà attribué.")
					}
				case "7", "map", "Map":
					fmt.Println("Par quelle touche voulez vous la remplacer ?")
					var touche_replace string
					fmt.Scanln(&touche_replace)
					if p.touchevalide(touche_replace) {
						touchemap = touche_replace
					} else {
						fmt.Println("Changement impossible, touche déjà attribué.")
					}
				case "8", "Options", "Option", "option", "options":
					fmt.Println("Par quelle touche voulez vous la remplacer ?")
					var touche_replace string
					fmt.Scanln(&touche_replace)
					if p.touchevalide(touche_replace) {
						ToucheOption = touche_replace
					} else {
						fmt.Println("Changement impossible, touche déjà attribué.")
					}
				case "9", "Quitter", "quitter", "quit", "Quit":
					fmt.Println("Par quelle touche voulez vous la remplacer ?")
					var touche_replace string
					fmt.Scanln(&touche_replace)
					if p.touchevalide(touche_replace) {
						ToucheQuitter = touche_replace
					} else {
						fmt.Println("Changement impossible, touche déjà attribué.")
					}
				}
			case ToucheQuitter:
			fmt.Println("Au revoir !")
			fin_du_jeu = false	
			default:
				fmt.Println()
				fmt.Println("Veuillez saisir une touche valide.")
				fmt.Println()
			}
		fmt.Println()
		fmt.Println("Quelle touche voulez vous changer ?")
		var touche string
		fmt.Scanln(&touche)
		switch touche {
		case "Haut" , "haut" ,"1":
		fmt.Println("Par quelle touche voulez vous la remplacer ?")
		var touche_replace string
		fmt.Scanln(&touche_replace)
		if p.touchevalide(touche_replace) {
		directionMap.Haut = touche_replace
		} else {
			fmt.Println("Changement impossible.")
		}
		case "Bas" , "bas", "2" :
		fmt.Println("Par quelle touche voulez vous la remplacer ?")
		var touche_replace string
		fmt.Scanln(&touche_replace)
		if p.touchevalide(touche_replace) {
		directionMap.Bas = touche_replace
		} else {
			fmt.Println("Changement impossible.")
		}
		case "Gauche", "gauche","3":
		fmt.Println("Par quelle touche voulez vous la remplacer ?")
		var touche_replace string
		fmt.Scanln(&touche_replace)
		if p.touchevalide(touche_replace) {
		directionMap.Gauche = touche_replace
		} else {
			fmt.Println("Changement impossible.")
		}
		case "Droite", "droite", "4":
		fmt.Println("Par quelle touche voulez vous la remplacer ?")
		var touche_replace string
		fmt.Scanln(&touche_replace)
		if p.touchevalide(touche_replace) {
		directionMap.Droite = touche_replace
		} else {
			fmt.Println("Changement impossible.")
		}
		case "5","Information du personnage","Personnage","information du personnage","perso":
		fmt.Println("Par quelle touche voulez vous la remplacer ?")
		var touche_replace string
		fmt.Scanln(&touche_replace)
		if p.touchevalide(touche_replace) {
		touchePersoInfo = touche_replace
		} else {
			fmt.Println("Changement impossible.")
		}
		case "6", "Inventaire","inv","inventaire":
		fmt.Println("Par quelle touche voulez vous la remplacer ?")	
		var touche_replace string
		fmt.Scanln(&touche_replace)
		if p.touchevalide(touche_replace) {
		toucheInv = touche_replace
		} else {
			fmt.Println("Changement impossible.")
		}
		case "7", "map", "Map":
		fmt.Println("Par quelle touche voulez vous la remplacer ?")
		var touche_replace string
		fmt.Scanln(&touche_replace)
		if p.touchevalide(touche_replace) {
		touchemap = touche_replace
		} else {
			fmt.Println("Changement impossible.")
		}
		case "8" , "Options" , "Option", "option","options":
		fmt.Println("Par quelle touche voulez vous la remplacer ?")	
		var touche_replace string
		fmt.Scanln(&touche_replace)
		if p.touchevalide(touche_replace) {
		ToucheOption = touche_replace
		} else {
			fmt.Println("Changement impossible.")
		}
		case "9","Quitter","quitter","quit","Quit":
		fmt.Println("Par quelle touche voulez vous la remplacer ?")	
		var touche_replace string
		fmt.Scanln(&touche_replace)
		if p.touchevalide(touche_replace) {
		ToucheQuitter = touche_replace
		} else {
			fmt.Println("Changement impossible.")
		}
		case "10","Marchand","marchand":
		fmt.Println("Par quelle touche voulez vous la remplacer ?")	
		var touche_replace string
		fmt.Scanln(&touche_replace)
		if p.touchevalide(touche_replace) {
		toucheMarchand = touche_replace
		} else {
			fmt.Println("Changement impossible.")
		}
		case "11","Forgeron","forgeron","forge","Forge":
		fmt.Println("Par quelle touche voulez vous la remplacer ?")	
		var touche_replace string
		fmt.Scanln(&touche_replace)
		if p.touchevalide(touche_replace) {
		ToucheForgeron = touche_replace
		} else {
			fmt.Println("Changement impossible.")
		}
		case "12","Guilde","guilde","guild","Guild":
		fmt.Println("Par quelle touche voulez vous la remplacer ?")	
		var touche_replace string
		fmt.Scanln(&touche_replace)
		if p.touchevalide(touche_replace) {
		ToucheGuilde = touche_replace
		} else {
			fmt.Println("Changement impossible.")
		}
		case "13","Entrainement","entrainement","train","Train":
		fmt.Println("Par quelle touche voulez vous la remplacer ?")	
		var touche_replace string
		fmt.Scanln(&touche_replace)
		if p.touchevalide(touche_replace) {
		ToucheEntrainement = touche_replace
		} else {
			fmt.Println("Changement impossible.")
		}
	default : 
	fmt.Println()
	fmt.Println("Veuillez saisir une touche valide.")
	fmt.Println()
	continue
	}
	} else { 
		fmt.Println()
	fmt.Printf("\t\t+-------------------------------------------------+\n")
	fmt.Printf("\t\t|               Menu Principal                    |\n")
	fmt.Printf("\t\t|                                                 |\n")
	fmt.Printf("\t\t|  [%s] Afficher les informations du personnage    |\n",touchePersoInfo)
	fmt.Printf("\t\t|  [%s] Accéder à l'inventaire                     |\n",toucheInv)
	fmt.Printf("\t\t|  [%s] Afficher la map                            |\n",touchemap)
	fmt.Printf("\t\t|                                                 |\n")
	fmt.Printf("\t\t|                                                 |\n")
	fmt.Printf("\t\t|                    [%s]↑                         |\n",directionMap.Haut)
	fmt.Printf("\t\t|         ←[%s]                    [%s]→            |\n",directionMap.Gauche,directionMap.Droite)
	fmt.Printf("\t\t|                    [%s]↓                         |\n",directionMap.Bas)	
    fmt.Printf("\t\t|                                                 |\n")		
	fmt.Printf("\t\t|                                                 |\n")
	fmt.Printf("\t\t|  [%s] Options                                    |\n",ToucheOption)
	fmt.Printf("\t\t|  [%s] Quitter le jeu                             |\n",ToucheQuitter)
	fmt.Printf("\t\t|                                                 |\n")
	fmt.Printf("\t\t+-------------------------------------------------+\n")
	var choice string
	fmt.Scanln(&choice)
	switch choice {
		case touchePersoInfo : 
		fmt.Println("=== stats du personnage ===")

	case toucheInv : 
			fmt.Println("=== inventaire ===")
			p.AccessInventory()
	case touchemap : 
			p.AfficheMapCo()
	case directionMap.Haut :
		y_position ++
	case directionMap.Bas :
		y_position--
	case directionMap.Gauche :
		x_position--
	case directionMap.Droite :
		x_position++
	case ToucheQuitter:
		fmt.Println("Au revoir !")
		fin_du_jeu = false	
		case ToucheOption: 
		fmt.Println()
		fmt.Println("Touches actuelle :")
		fmt.Printf("Haut   : %s \n",directionMap.Haut)
		fmt.Printf("Bas    : %s \n",directionMap.Bas)
		fmt.Printf("Gauche : %s \n",directionMap.Gauche)
		fmt.Printf("Droite : %s \n",directionMap.Droite)
		fmt.Printf("Information du personnage : %s \n",touchePersoInfo)
		fmt.Printf("Inventaire : %s \n",toucheInv)
		fmt.Printf("Map        : %s \n",touchemap)
		fmt.Printf("Options    : %s \n",ToucheOption)
		fmt.Printf("Quitter    : %s  \n",ToucheQuitter)
		
	
		fmt.Println()
		fmt.Println("Quelle touche voulez vous changer ?")
		var touche string
		fmt.Scanln(&touche)
		switch touche {
		case "Haut" , "haut" ,"1":
		fmt.Println("Par quelle touche voulez vous la remplacer ?")
		var touche_replace string
		fmt.Scanln(&touche_replace)
		if p.touchevalide(touche_replace) {
		directionMap.Haut = touche_replace
		} else {
			fmt.Println("Changement impossible.")
		}
		case "Bas" , "bas", "2" :
		fmt.Println("Par quelle touche voulez vous la remplacer ?")
		var touche_replace string
		fmt.Scanln(&touche_replace)
		if p.touchevalide(touche_replace) {
		directionMap.Bas = touche_replace
		} else {
			fmt.Println("Changement impossible.")
		}
		case "Gauche", "gauche","3":
		fmt.Println("Par quelle touche voulez vous la remplacer ?")
		var touche_replace string
		fmt.Scanln(&touche_replace)
		if p.touchevalide(touche_replace) {
		directionMap.Gauche = touche_replace
		} else {
			fmt.Println("Changement impossible.")
		}
		case "Droite", "droite", "4":
		fmt.Println("Par quelle touche voulez vous la remplacer ?")
		var touche_replace string
		fmt.Scanln(&touche_replace)
		if p.touchevalide(touche_replace) {
		directionMap.Droite = touche_replace
		} else {
			fmt.Println("Changement impossible .")
		}
		case "5","Information du personnage","Personnage","information du personnage","perso":
		fmt.Println("Par quelle touche voulez vous la remplacer ?")
		var touche_replace string
		fmt.Scanln(&touche_replace)
		if p.touchevalide(touche_replace) {
		touchePersoInfo = touche_replace
		} else {
			fmt.Println("Changement impossible.")
		}
		case "6", "Inventaire","inv","inventaire":
		fmt.Println("Par quelle touche voulez vous la remplacer ?")	
		var touche_replace string
		fmt.Scanln(&touche_replace)
		if p.touchevalide(touche_replace) {
		toucheInv = touche_replace
		} else {
			fmt.Println("Changement impossible.")
		}
		case "7", "map", "Map":
		fmt.Println("Par quelle touche voulez vous la remplacer ?")
		var touche_replace string
		fmt.Scanln(&touche_replace)
		if p.touchevalide(touche_replace) {
		touchemap = touche_replace
		} else {
			fmt.Println("Changement impossible.")
		}
		case "8" , "Options" , "Option", "option","options":
		fmt.Println("Par quelle touche voulez vous la remplacer ?")	
		var touche_replace string
		fmt.Scanln(&touche_replace)
		if p.touchevalide(touche_replace) {
		ToucheOption = touche_replace
		} else {
			fmt.Println("Changement impossible.")
		}
		case "9","Quitter","quitter","quit","Quit":
		fmt.Println("Par quelle touche voulez vous la remplacer ?")	
		var touche_replace string
		fmt.Scanln(&touche_replace)
		if p.touchevalide(touche_replace) {
		ToucheQuitter = touche_replace
		} else {
			fmt.Println("Changement impossible.")
		}
		}

	case "combat" :
			p.Combat_start_premier()
	default : 
	fmt.Println()
	fmt.Println("Veuillez saisir une touche valide.")
	fmt.Println()
	}
	}
}
}