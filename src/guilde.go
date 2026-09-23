package projet 

import (
	"fmt"
)

var premiere_fois bool = true

func (p *Character) Guilde(){
		afficherTexte50("Bienvenue dans la Guilde des Aventurier de Eldoria !")
		if premiere_fois {
			afficherTexte50("Ici, vous pouvez vous reposer autant que vous le voulez.")
			afficherTexte50("Si vous êtes trop chargé, vous pouvez également déposer vos objets en trop dans votre coffre de guilde.")
			afficherTexte50("Vous pouvez accéder à votre coffre de guilde depuis n'importe quelle guilde se trouvant dans une ville associé.")
			afficherTexte50("Bien entendu, vous pouvez récupérer vos objets à n'importe quel moment à condition de ne pas avoir l'inventaire plein.")
		}
		fmt.Println()
		afficherTexte50("Que voulez vous faire?")
		afficherTexte50("Se reposer pour retrouver ses forces")
		afficherTexte50("Déposer un objet dans le coffre de Guilde")
		afficherTexte50("Récuperer un objet dans le coffre de Guilde")
		var choix string
		fmt.Scanln(&choix)
		switch choix {

		}
}