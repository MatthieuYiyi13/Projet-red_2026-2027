package projet

import "fmt"

var prix_objet = map[string]int{
    "objet_1": 10 ,
    "objet_2": 10 ,
    "objet_3": 10 , 
	"objet_4": 10 ,
	"objet_5": 10 ,
	"objet_6": 10 ,
	"objet_7": 10 ,
	"objet_8": 10 ,
	"objet_9": 10 ,
	"objet_10": 10 ,
}
 
func (p *Character) Acheter() {
	fmt.Println("Que voulez vous acheter ? ")
	fmt.Println("1 : objet_1: 10")
	fmt.Println("2 : objet_2: 10")
	fmt.Println("3 : objet_3: 10")
	fmt.Println("4 : objet_4: 10")
	fmt.Println("5 : objet_5: 10")
	fmt.Println("6 : objet_6: 10")
	fmt.Println("7 : objet_7: 10")
	fmt.Println("8 : objet_8: 8")
	fmt.Println("9 : objet_9: 9")
	fmt.Println("10 : objet_10: 10")

var choix_objet int
	fmt.Scanln(&choix_objet)

	switch choix_objet {
	case 1:
		p.inventaire["tissu de spectre"] += 1 
		fmt.Println("Vous avez acheté un tissu de spectre")
	case 2:
		p.Vendre()
	case 3:
		fmt.Println("Au revoir !")
	case 4:
		fmt.Println("Au revoir !")		
	case 5:
		fmt.Println("Au revoir !")
	case 6:
		fmt.Println("Au revoir !")
	case 7:
		fmt.Println("Au revoir !")
	case 8:
		fmt.Println("Au revoir !")
	case 9:
		fmt.Println("Au revoir !")
	case 10:
		fmt.Println("Au revoir !")
}
}

func (p *Character) Vendre() {
	fmt.Println("Vendre")
}

func (p *Character) Marchand() {
	fmt.Println("Bonjour jeune aventurier, je vois que tu as réussi à me trouver dans cette magnifique ville de Qarth !")
	fmt.Println()
	fmt.Println("Que souhaites-tu faire maintenant ?")
	fmt.Println()
	fmt.Println("1. Acheter")
	fmt.Println("2. Vendre")
	fmt.Println("3. Quitter le menu")

	var choiceM int
	fmt.Scanln(&choiceM)

	switch choiceM {
	case 1:
		p.Acheter()
	case 2:
		p.Vendre()
	case 3:
		fmt.Println("Au revoir !")
	default:
		fmt.Println("Choix invalide. Tu dois choisir 1, 2 ou 3.")
	}
}
