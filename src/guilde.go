package projet

import (
	"fmt"
)

var coffre_guilde = make(map[string]int)
var premiere_fois = true

func (p *Character) Guilde() {

	for premiere_fois {
		afficherTexte20("Bienvenue dans la Guilde des Aventuriers d'Eldoria !")
		if premiere_fois {
			afficherTexte20("Ici, vous pouvez vous reposer autant que vous le voulez.")
			afficherTexte20("Si vous êtes trop chargé, vous pouvez également déposer vos objets en trop dans votre coffre de guilde.")
			afficherTexte20("Vous pouvez accéder à votre coffre de guilde depuis n'importe quelle guilde se trouvant dans une ville associée.")
			afficherTexte20("Bien entendu, vous pouvez récupérer vos objets à n'importe quel moment à condition de ne pas avoir l'inventaire plein.")
			premiere_fois = false
		}
		fmt.Println()
		afficherTexte20("Que voulez-vous faire ?")
		fmt.Println()
		afficherTexte20("[S] : Se reposer pour retrouver ses forces")
		afficherTexte20("[D] : Déposer un objet dans le coffre de guilde")
		afficherTexte20("[R] : Récupérer un objet dans le coffre de guilde")
		afficherTexte20("[O] : Voir les objets dans mon coffre de guilde")
		afficherTexte20("[E] : Retour")
		var choix string
		fmt.Scanln(&choix)
		switch choix {
		case "S", "s":
			p.pv = p.pvmax
			p.Mana = p.Manamax
			afficherTexte20("Vous prenez le temps de vous remettre de vos blessures.")
			stopSound()
			afficherTexte100(". . . . . . . . . . . . . . . . . . . .")
			afficherTexte20("Vous êtes complètement soigné !")
		case "D", "d":
			if len(p.inventaire) == 0 {
				afficherTexte20("Votre inventaire est vide.")
				continue
			}
			afficherTexte20("Voici les objets présents dans votre inventaire :")
			fmt.Println()
			for itemname, itemquantity := range p.inventaire {
				if itemquantity <= 0 {
					continue
				}
				afficherTexte50(
					"Voulez-vous ranger %s (%d) dans le coffre ? (Oui/Non)\n",
					itemname,
					itemquantity,
				)
				var depose string
				fmt.Scanln(&depose)
				switch depose {
				case "Oui", "oui", "OUI", "oUI", "OuI", "ouI":
					afficherTexte50(
						"Vous possédez %d %s.\n",
						itemquantity,
						itemname,
					)
					afficherTexte20("Combien voulez-vous en déposer ?")
					var nbr_depot int
					fmt.Scanln(&nbr_depot)
					if nbr_depot <= 0 {
						afficherTexte20("Vous devez déposer au moins 1 objet.")
						continue
					}
					if nbr_depot > itemquantity {
						afficherTexte50(
							"Vous ne pouvez pas déposer %d %s. Vous n'en possédez que %d.\n",
							nbr_depot,
							itemname,
							itemquantity,
						)
						continue
					}
					coffre_guilde[itemname] += nbr_depot
					p.inventaire[itemname] -= nbr_depot
					afficherTexte50(
						"Vous avez déposé %d %s dans le coffre.\n",
						nbr_depot,
						itemname,
					)
					if p.inventaire[itemname] <= 0 {
						delete(p.inventaire, itemname)
					}
				case "Non", "non", "NON", "NoN", "NOn", "nON", "noN":
					continue
				default:
					afficherTexte20("Réponse invalide.")
					continue
				}
			}
		case "R", "r":
			if len(coffre_guilde) == 0 {
				afficherTexte20("Votre coffre de guilde est vide.")
				continue
			}
			afficherTexte20("Voici les objets présents dans votre coffre :")
			fmt.Println()
			for itemname, itemquantity := range coffre_guilde {
				if itemquantity <= 0 {
					continue
				}
				afficherTexte50(
					"Voulez-vous prendre %s (%d) ? (Oui/Non)\n",
					itemname,
					itemquantity,
				)
				var prendre string
				fmt.Scanln(&prendre)
				switch prendre {
				case "Oui", "oui", "OUI", "oUI", "OuI", "ouI":
					if p.InventairePlein() {
						afficherTexte20(
							"Votre inventaire est plein ! Vous ne pouvez pas récupérer cet objet.",
						)
						continue
					}
					if coffre_guilde[itemname] <= 0 {
						afficherTexte50(
							"Vous n'avez pas de %s dans votre coffre de guilde.\n",
							itemname,
						)
						continue
					}
					coffre_guilde[itemname] -= 1
					p.inventaire[itemname] += 1
					afficherTexte50(
						"Vous avez récupéré 1 %s.\n",
						itemname,
					)
					afficherTexte50(
						"Il reste %d %s dans le coffre.\n",
						coffre_guilde[itemname],
						itemname,
					)
					if coffre_guilde[itemname] <= 0 {
						delete(coffre_guilde, itemname)
					}
				case "Non", "non", "NON", "NoN", "NOn", "nON", "noN":
					continue
				default:
					afficherTexte20("Réponse invalide.")
					continue
				}
			}
		case "O", "o":
			fmt.Println()
			afficherTexte20("========== COFFRE DE GUILDE ==========")
			fmt.Println()
			if len(coffre_guilde) == 0 {

				afficherTexte20("Votre coffre de guilde est vide.")
			} else {
				for itemname, itemquantity := range coffre_guilde {
					if itemquantity > 0 {
						fmt.Printf("- %s : %d\n", itemname, itemquantity)
					}
				}
			}
			fmt.Println()
			afficherTexte20("=======================================")
		case "E", "e":
			afficherTexte20("Vous quittez la Guilde.")
			return
		default:
			afficherTexte20("Choix invalide.")
		}
		fmt.Println()
	}
}
